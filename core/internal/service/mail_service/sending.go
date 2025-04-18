package mail_service

import (
	"billionmail-core/internal/service/mail_boxes"
	"billionmail-core/internal/service/public"
	"context"
	"crypto/tls"
	"encoding/hex"
	"errors"
	"fmt"
	"github.com/gogf/gf/util/grand"
	"github.com/gogf/gf/v2/frame/g"
	"net/smtp"
	"strings"
	"sync"
	"time"
)

type Message struct {
	Title   string
	Content string
	Headers map[string]string
}

// MailTitle get title of email
func (m Message) MailTitle() string {
	return m.Title
}

// MailText get content of email
func (m Message) MailText() string {
	return m.Content
}

// MailHeader get headers of email
func (m Message) MailHeader() string {
	headers := ""
	for key, value := range m.Headers {
		headers += fmt.Sprintf("%s: %s\r\n", key, value)
	}
	return headers
}

func NewMessage(title string, content string) Message {
	return Message{
		Title:   title,
		Content: content,
		Headers: make(map[string]string),
	}
}

type EmailSender struct {
	Email     string       `json:"email" v:"required|email"` // email of sender
	Host      string       `json:"host" v:"required|domain"` // SMTP server address
	Port      string       `json:"port" v:"required"`        // SMTP server port
	Password  string       `json:"password" v:"required"`    // SMTP password
	client    *smtp.Client // persistent SMTP client connection
	mutex     sync.Mutex   // mutex for thread safety
	connected bool         // connection status
}

func NewEmailSender() *EmailSender {
	e := &EmailSender{
		Host:      "localhost",
		Port:      "25",
		connected: false,
	}

	if public.IsRunningInContainer() {
		e.Host = "postfix"
	}

	return e
}

func NewEmailSenderWithLocal(email string) (es *EmailSender, err error) {
	es = NewEmailSender()

	es.Email = email
	es.Password, err = mail_boxes.PasswordByEmail(context.Background(), email)

	if err != nil {
		return
	}

	if !es.IsConfigured() {
		err = errors.New("Email Sender not configured")
		return
	}

	return
}

// Close closes the SMTP connection
func (e *EmailSender) Close() {
	_ = e.Disconnect()
}

// Connect establishes an SMTP connection
func (e *EmailSender) Connect() error {
	e.mutex.Lock()
	defer e.mutex.Unlock()

	if e.connected && e.client != nil {
		// Connection already exists
		return nil
	}

	var err error

	if e.isSecure() {
		err = e.connectWithSSL()
	} else {
		err = e.connectPlain()
	}

	if err != nil {
		e.connected = false
		e.client = nil
		return err
	}

	e.connected = true
	return nil
}

// connectWithSSL establishes a secure SMTP connection
func (e *EmailSender) connectWithSSL() error {
	conn, err := tls.Dial("tcp", e.Host+":"+e.Port, nil)
	if err != nil {
		return fmt.Errorf("TLS dial: %w", err)
	}

	client, err := smtp.NewClient(conn, e.Host)
	if err != nil {
		conn.Close()
		return fmt.Errorf("new SMTP client: %w", err)
	}

	auth := smtp.PlainAuth("", e.Email, e.Password, e.Host)
	if err = client.Auth(auth); err != nil {
		client.Close()
		return fmt.Errorf("SMTP auth: %w", err)
	}

	e.client = client
	return nil
}

// connectPlain establishes a plain SMTP connection
func (e *EmailSender) connectPlain() error {
	client, err := smtp.Dial(e.Host + ":" + e.Port)
	if err != nil {
		return fmt.Errorf("SMTP dial: %w", err)
	}

	// Check if STARTTLS is needed
	if e.Port == "587" {
		if err = client.StartTLS(nil); err != nil {
			client.Close()
			return fmt.Errorf("SMTP STARTTLS: %w", err)
		}
	}

	auth := smtp.PlainAuth("", e.Email, e.Password, e.Host)
	if err = client.Auth(auth); err != nil {
		client.Close()
		return fmt.Errorf("SMTP auth: %w", err)
	}

	e.client = client
	return nil
}

// Disconnect closes the SMTP connection
func (e *EmailSender) Disconnect() error {
	e.mutex.Lock()
	defer e.mutex.Unlock()

	if !e.connected || e.client == nil {
		return nil
	}

	err := e.client.Quit()
	if err != nil {
		// Force close connection if Quit fails
		e.client.Close()
		g.Log().Warning(context.Background(), "Error closing SMTP connection: ", err)
	}

	e.connected = false
	e.client = nil
	return err
}

// GenerateMessageID generates a unique Message-ID for email
func (e *EmailSender) GenerateMessageID() string {
	randomBytes := grand.B(16)
	randomID := hex.EncodeToString(randomBytes)
	timestampMillis := time.Now().UnixMilli()

	domain := strings.Split(e.Email, "@")
	domainPart := e.Host
	if len(domain) > 1 {
		domainPart = domain[1]
	}

	return fmt.Sprintf("<%d.%s@%s>", timestampMillis, randomID, domainPart)
}

// Send sends an email to specified recipients
func (e *EmailSender) Send(message Message, recipients []string) error {
	if len(recipients) == 0 {
		return fmt.Errorf("no recipients specified")
	}

	e.mutex.Lock()
	defer e.mutex.Unlock()

	// Make sure we have a connection
	if !e.connected || e.client == nil {
		e.mutex.Unlock()
		if err := e.Connect(); err != nil {
			return fmt.Errorf("failed to connect: %w", err)
		}
		e.mutex.Lock()
	}

	// Try to send the message, with reconnect on failure
	err := e.doSend(message, recipients)
	if err != nil {
		// Connection might be stale, try to reconnect once
		g.Log().Debug(context.Background(), "SMTP send failed, attempting reconnection")

		// Clean up
		e.client.Close()
		e.connected = false
		e.client = nil

		// Try to reconnect
		e.mutex.Unlock()
		if err := e.Connect(); err != nil {
			return fmt.Errorf("reconnect failed: %w", err)
		}
		e.mutex.Lock()

		// Try sending again after reconnect
		return e.doSend(message, recipients)
	}

	return nil
}

// doSend performs the actual message sending
func (e *EmailSender) doSend(message Message, recipients []string) error {
	// Add default headers if not present
	if message.Headers == nil {
		message.Headers = make(map[string]string)
	}

	// Add Message-ID if not already set
	if _, exists := message.Headers["Message-ID"]; !exists {
		message.Headers["Message-ID"] = e.GenerateMessageID()
	}

	// Add Date if not already set
	if _, exists := message.Headers["Date"]; !exists {
		message.Headers["Date"] = time.Now().Format(time.RFC1123Z)
	}

	// Default Content-Type if not already set
	if _, exists := message.Headers["Content-Type"]; !exists {
		message.Headers["Content-Type"] = "text/html; charset=utf-8"
	}

	// Build email message with headers
	headerString := fmt.Sprintf("From: %s\r\n", e.Email) +
		fmt.Sprintf("To: %s\r\n", strings.Join(recipients, ",")) +
		fmt.Sprintf("Subject: %s\r\n", message.MailTitle()) +
		message.MailHeader() +
		"\r\n" +
		message.MailText() +
		"\r\n"

	msg := []byte(headerString)

	// Reset the connection state
	if err := e.client.Reset(); err != nil {
		return fmt.Errorf("SMTP reset: %w", err)
	}

	// Set the sender
	if err := e.client.Mail(e.Email); err != nil {
		return fmt.Errorf("SMTP mail: %w", err)
	}

	// Set the recipients
	for _, to := range recipients {
		if err := e.client.Rcpt(to); err != nil {
			return fmt.Errorf("SMTP rcpt: %w", err)
		}
	}

	// Get a writer for the message body
	w, err := e.client.Data()
	if err != nil {
		return fmt.Errorf("SMTP data: %w", err)
	}

	// Write the message
	if _, err = w.Write(msg); err != nil {
		return fmt.Errorf("SMTP write: %w", err)
	}

	// Close the writer
	if err = w.Close(); err != nil {
		return fmt.Errorf("SMTP close writer: %w", err)
	}

	return nil
}

// isSecure check if TLS connection is needed
func (e *EmailSender) isSecure() bool {
	// usually,
	// port 465 is used for SMTP with SSL (implicit TLS)
	// port 587 is used for SMTP with STARTTLS (explicit TLS)
	if e.Port == "465" {
		return true
	} else if e.Port == "587" {
		// We will use STARTTLS if the port is 587
		return false
	}
	return false
}

// IsConfigured check if Email notification is configured
func (e *EmailSender) IsConfigured() bool {
	e.Email = strings.TrimSpace(e.Email)
	e.Host = strings.TrimSpace(e.Host)
	e.Port = strings.TrimSpace(e.Port)
	e.Password = strings.TrimSpace(e.Password)

	if e.Email == "" || e.Host == "" || e.Port == "" || e.Password == "" {
		return false
	}

	if !public.IsPort(e.Port) {
		return false
	}

	if !public.IsHost(e.Host) {
		return false
	}

	if !public.IsEmail(e.Email) {
		return false
	}

	return true
}
