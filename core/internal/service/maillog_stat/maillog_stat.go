package maillog_stat

import (
	"billionmail-core/internal/service/public"
	"context"
	"fmt"
	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/glog"
	"os"
	"regexp"
	"strconv"
	"strings"
	"time"
)

// Define constants
var (
	defaultMaillogPath                    = public.AbsPath("../logs/postfix/mail.log")
	defaultPostfixMainConf                = public.AbsPath("../conf/postfix/main.cf")
	defaultLatestLogTimeFile              = "data/last_maillog_time"
	defaultLatestAggregateMaillogTimeFile = "data/last_aggregate_maillog_time"
)

// MailRecord basic mail record structure
type MailRecord struct {
	PostfixMessageID string `json:"postfix_message_id"`
	LogTimeMillis    int64  `json:"log_time_millis"`
}

// MailSendRecord mail sending record
type MailSendRecord struct {
	MailRecord
	Recipient    string  `json:"recipient"`
	MailProvider string  `json:"mail_provider"`
	Status       string  `json:"status"`
	Delay        float64 `json:"delay"`
	Delays       string  `json:"delays"`
	Dsn          string  `json:"dsn"`
	Relay        string  `json:"relay"`
	Description  string  `json:"description"`
}

// MailReceiveRecord mail receiving record
type MailReceiveRecord struct {
	MailRecord
	Recipient   string  `json:"recipient"`
	Status      string  `json:"status"`
	Delay       float64 `json:"delay"`
	Delays      string  `json:"delays"`
	Dsn         string  `json:"dsn"`
	Relay       string  `json:"relay"`
	Description string  `json:"description"`
}

// MailMessageID mail MessageID record
type MailMessageID struct {
	MailRecord
	MessageID string `json:"message_id"`
}

// MailSender mail sender record
type MailSender struct {
	MailRecord
	Sender string `json:"sender"`
	Size   int64  `json:"size"`
}

// MailRemoved mail removed record
type MailRemoved struct {
	MailRecord
}

// MailDeferredRecord mail deferred record
type MailDeferredRecord struct {
	MailRecord
	Delay       float64 `json:"delay"`
	Delays      string  `json:"delays"`
	Dsn         string  `json:"dsn"`
	Relay       string  `json:"relay"`
	Description string  `json:"description"`
}

// MaillogStat mail log statistics analyzer
type MaillogStat struct {
	maillogPath         string
	startTime           int64
	endTime             int64
	doSummary           bool
	ignoreRelays        map[string]struct{}
	ignoreMailAddresses map[string]struct{}
	currentYear         int
	delivered           int
	bounced             int
	deferred            int
	deferredTotal       int
	bounceDetails       map[string]map[string]struct{}
	deferralDetails     map[string]map[string]struct{}
	mailHostname        string
	monthMap            map[string]int
	statusPattern       *regexp.Regexp
	recipientPattern    *regexp.Regexp
	delayPattern        *regexp.Regexp
	delaysPattern       *regexp.Regexp
	dsnPattern          *regexp.Regexp
	relayPattern        *regexp.Regexp
	descriptionPattern  *regexp.Regexp
	messageIDPattern    *regexp.Regexp
	mailRemovedPattern  *regexp.Regexp
	mailSenderPattern   *regexp.Regexp
}

// NewMaillogStat creates a new mail log statistics analyzer
func NewMaillogStat(maillogPath string, startTime, endTime int64, doSummary bool) *MaillogStat {
	ms := &MaillogStat{
		maillogPath:         maillogPath,
		startTime:           startTime,
		endTime:             endTime,
		doSummary:           doSummary,
		ignoreRelays:        make(map[string]struct{}),
		ignoreMailAddresses: make(map[string]struct{}),
		monthMap: map[string]int{
			"Jan": 1, "Feb": 2, "Mar": 3, "Apr": 4,
			"May": 5, "Jun": 6, "Jul": 7, "Aug": 8,
			"Sep": 9, "Oct": 10, "Nov": 11, "Dec": 12,
		},
		bounceDetails:   make(map[string]map[string]struct{}),
		deferralDetails: make(map[string]map[string]struct{}),
	}

	// Compile regular expressions
	ms.statusPattern = regexp.MustCompile(`status=(\S+) `)
	ms.recipientPattern = regexp.MustCompile(`to=<([^>]+)>`)
	ms.delayPattern = regexp.MustCompile(`delay=(\d+(?:\.\d+)?),`)
	ms.delaysPattern = regexp.MustCompile(`delays=(\d+(?:\.\d+)?(?:/\d+(?:\.\d+)?){3}),`)
	ms.dsnPattern = regexp.MustCompile(`dsn=([^,]+),`)
	ms.relayPattern = regexp.MustCompile(`relay=([^,]+),`)
	ms.descriptionPattern = regexp.MustCompile(`\((.*?)\)$`)
	ms.messageIDPattern = regexp.MustCompile(`postfix/[^\[]+\[\d+]: *([^:]+):`)
	ms.mailRemovedPattern = regexp.MustCompile(`postfix/qmgr\[\d+]: *([^:]+): *removed$`)
	ms.mailSenderPattern = regexp.MustCompile(`postfix/qmgr\[\d+]: *([^:]+): *from=<([^>]+)>, +size=(\d+),`)

	ms.resetData()
	ms.initIgnoreRelays()
	ms.initIgnoreMailAddresses()

	return ms
}

// resetData resets statistical data
func (ms *MaillogStat) resetData() {
	ms.bounceDetails = make(map[string]map[string]struct{})
	ms.deferralDetails = make(map[string]map[string]struct{})

	if fi, err := os.Stat(ms.maillogPath); err == nil {
		ms.currentYear = fi.ModTime().Year()
	}

	ms.delivered = 0
	ms.bounced = 0
	ms.deferred = 0
	ms.deferredTotal = 0
}

// initIgnoreRelays initializes ignored relays
func (ms *MaillogStat) initIgnoreRelays() {
	if _, err := os.Stat(defaultPostfixMainConf); err != nil {
		return
	}

	content, err := os.ReadFile(defaultPostfixMainConf)
	if err != nil {
		return
	}

	lines := strings.Split(string(content), "\n")
	var myhostname, virtualTransport string

	for _, line := range lines {
		line = strings.TrimSpace(line)
		if line == "" {
			continue
		}

		if myhostname != "" && virtualTransport != "" {
			break
		}

		if myhostname == "" {
			if matches := regexp.MustCompile(`myhostname *= *([^\r\n]+)`).FindStringSubmatch(line); matches != nil {
				myhostname = matches[1]
				continue
			}
		}

		if virtualTransport == "" {
			if matches := regexp.MustCompile(`virtual_transport *= *([^\r\n]+)`).FindStringSubmatch(line); matches != nil {
				virtualTransport = matches[1]
				continue
			}
		}
	}

	if myhostname != "" && virtualTransport != "" {
		ms.mailHostname = myhostname
		relay := fmt.Sprintf("%s[%s]", myhostname, strings.Split(virtualTransport, ":")[len(strings.Split(virtualTransport, ":"))-1])
		ms.ignoreRelays[relay] = struct{}{}
	}
}

// initIgnoreMailAddresses initializes ignored mail addresses
func (ms *MaillogStat) initIgnoreMailAddresses() {
	ms.ignoreMailAddresses["root@localhost"] = struct{}{}
	if ms.mailHostname != "" {
		ms.ignoreMailAddresses[fmt.Sprintf("root@%s", ms.mailHostname)] = struct{}{}
	}
}

// getMailProvider gets mail service provider
func (ms *MaillogStat) getMailProvider(email string) string {
	hostname := strings.Split(email, "@")[1]

	switch {
	case strings.Contains(hostname, "gmail.com") || strings.Contains(hostname, "googlemail.com"):
		return "google"
	case strings.Contains(hostname, "outlook.com") || strings.Contains(hostname, "hotmail.com") ||
		strings.Contains(hostname, "live.com") || strings.Contains(hostname, "msn.com") ||
		strings.HasPrefix(hostname, "outlook."):
		return "outlook"
	case strings.Contains(hostname, "yahoo.com") || strings.Contains(hostname, "ymail.com") ||
		strings.Contains(hostname, "rocketmail.com") || strings.HasPrefix(hostname, "yahoo."):
		return "yahoo"
	case strings.Contains(hostname, "icloud.com") || strings.Contains(hostname, "me.com") ||
		strings.Contains(hostname, "mac.com") || strings.Contains(hostname, "apple.com"):
		return "apple"
	case strings.Contains(hostname, "protonmail.com") || strings.Contains(hostname, "proton.me") ||
		strings.Contains(hostname, "pm.me"):
		return "proton"
	case strings.Contains(hostname, "zoho.com") || strings.Contains(hostname, "zohomail.com") ||
		strings.Contains(hostname, "zohocorp.com") || strings.Contains(hostname, "zmail.com") ||
		strings.HasPrefix(hostname, "zoho."):
		return "zoho"
	case strings.Contains(hostname, "kindle.com") || strings.Contains(hostname, "amazon.com") ||
		strings.Contains(hostname, "awsapps.com"):
		return "amazon"
	default:
		return "other"
	}
}

// parseLogTimeMillis parses log time (millisecond timestamp)
func (ms *MaillogStat) parseLogTimeMillis(line string) int64 {
	line = strings.TrimSpace(line)

	// Try to parse ISO 8601 format
	if len(line) >= 4 && isDigit(line[:4]) {
		t, err := time.Parse(time.RFC3339, strings.Split(line, " ")[0])
		if err == nil {
			return t.UnixMilli()
		}
	}

	parts := strings.Fields(line)
	if len(parts) < 4 {
		return -1
	}

	month, ok := ms.monthMap[parts[0]]
	if !ok {
		return -1
	}

	day, err := strconv.Atoi(parts[1])
	if err != nil {
		return -1
	}

	timeParts := strings.Split(parts[2], ":")
	if len(timeParts) != 3 {
		return -1
	}

	hour, err := strconv.Atoi(timeParts[0])
	if err != nil {
		return -1
	}

	minute, err := strconv.Atoi(timeParts[1])
	if err != nil {
		return -1
	}

	second, err := strconv.Atoi(timeParts[2])
	if err != nil {
		return -1
	}

	t := time.Date(ms.currentYear, time.Month(month), day, hour, minute, second, 0, time.Local)
	return t.UnixMilli()
}

// isDigit checks if string consists of digits only
func isDigit(s string) bool {
	for _, c := range s {
		if c < '0' || c > '9' {
			return false
		}
	}
	return true
}

// Analysis analyzes mail logs
func (ms *MaillogStat) Analysis(ctx context.Context) (<-chan interface{}, error) {
	ms.resetData()

	if _, err := os.Stat(ms.maillogPath); err != nil {
		return nil, err
	}

	recordChan := make(chan interface{}, 1000)
	go func() {
		defer close(recordChan)

		err := public.ReadEachReverse(ms.maillogPath, func(line string, cnt int) bool {
			select {
			case <-ctx.Done():
				return false
			default:
				record := ms.analyzeLine(line)
				if record != nil {
					if r, ok := record.(*MailSendRecord); ok && r.Status == "deferred" {
						recordChan <- MailDeferredRecord{
							MailRecord: MailRecord{
								PostfixMessageID: r.PostfixMessageID,
								LogTimeMillis:    r.LogTimeMillis,
							},
							Delay:       r.Delay,
							Delays:      r.Delays,
							Dsn:         r.Dsn,
							Relay:       r.Relay,
							Description: r.Description,
						}
					}

					recordChan <- record
				}
			}
			return true
		})

		if err != nil {
			fmt.Printf("Error reading maillog: %v\n", err)
		}
	}()

	return recordChan, nil
}

// analyzeLine analyzes single log line
func (ms *MaillogStat) analyzeLine(line string) interface{} {
	line = strings.TrimSpace(line)
	logTime := ms.parseLogTimeMillis(line)

	if logTime < 1 {
		return nil
	}

	if ms.endTime > 0 && logTime > ms.endTime {
		return nil
	}

	if ms.startTime > 0 && logTime < ms.startTime {
		return nil
	}

	// Analyze different types of logs
	if strings.Contains(line, "postfix/lmtp[") {
		return ms.analyzeReceiveMail(line, logTime)
	}

	if strings.Contains(line, "postfix/qmgr[") {
		if strings.Contains(line, "from=<") {
			return ms.analyzeSender(line, logTime)
		}
		if strings.Contains(line, "removed") {
			return ms.analyzeRemoved(line, logTime)
		}
	}

	if strings.Contains(line, "postfix/cleanup[") && strings.Contains(line, "message-id=<") {
		return ms.analyzeMessageID(line, logTime)
	}

	if strings.Contains(line, "to=<") && strings.Contains(line, "status=") && strings.Contains(line, "dsn=") {
		return ms.analyzeSendMail(line, logTime)
	}

	return nil
}

// analyzeReceiveMail analyzes mail receiving logs
func (ms *MaillogStat) analyzeReceiveMail(line string, logTime int64) interface{} {
	matches := ms.messageIDPattern.FindStringSubmatch(line)
	if matches == nil {
		return nil
	}

	recipientMatches := ms.recipientPattern.FindStringSubmatch(line)
	if recipientMatches == nil {
		return nil
	}

	statusMatches := ms.statusPattern.FindStringSubmatch(line)
	if statusMatches == nil {
		return nil
	}

	record := &MailReceiveRecord{
		MailRecord: MailRecord{
			PostfixMessageID: matches[1],
			LogTimeMillis:    logTime,
		},
		Recipient: recipientMatches[1],
		Status:    statusMatches[1],
	}

	if _, ok := ms.ignoreMailAddresses[record.Recipient]; ok {
		return nil
	}

	if relayMatches := ms.relayPattern.FindStringSubmatch(line); relayMatches != nil {
		record.Relay = relayMatches[1]
	}

	if delayMatches := ms.delayPattern.FindStringSubmatch(line); delayMatches != nil {
		record.Delay, _ = strconv.ParseFloat(delayMatches[1], 64)
	}

	if delaysMatches := ms.delaysPattern.FindStringSubmatch(line); delaysMatches != nil {
		record.Delays = delaysMatches[1]
	}

	if dsnMatches := ms.dsnPattern.FindStringSubmatch(line); dsnMatches != nil {
		record.Dsn = dsnMatches[1]
	}

	if descMatches := ms.descriptionPattern.FindStringSubmatch(line); descMatches != nil {
		record.Description = descMatches[1]
	}

	return record
}

// analyzeSendMail analyzes mail sending logs
func (ms *MaillogStat) analyzeSendMail(line string, logTime int64) interface{} {
	matches := ms.messageIDPattern.FindStringSubmatch(line)
	if matches == nil {
		return nil
	}

	recipientMatches := ms.recipientPattern.FindStringSubmatch(line)
	if recipientMatches == nil {
		return nil
	}

	statusMatches := ms.statusPattern.FindStringSubmatch(line)
	if statusMatches == nil {
		return nil
	}

	record := &MailSendRecord{
		MailRecord: MailRecord{
			PostfixMessageID: matches[1],
			LogTimeMillis:    logTime,
		},
		Recipient: recipientMatches[1],
		Status:    statusMatches[1],
	}

	if _, ok := ms.ignoreMailAddresses[record.Recipient]; ok {
		return nil
	}

	if relayMatches := ms.relayPattern.FindStringSubmatch(line); relayMatches != nil {
		relay := relayMatches[1]
		if _, ok := ms.ignoreRelays[relay]; ok {
			return nil
		}
		record.Relay = relay
	}

	if delayMatches := ms.delayPattern.FindStringSubmatch(line); delayMatches != nil {
		record.Delay, _ = strconv.ParseFloat(delayMatches[1], 64)
	}

	if delaysMatches := ms.delaysPattern.FindStringSubmatch(line); delaysMatches != nil {
		record.Delays = delaysMatches[1]
	}

	if dsnMatches := ms.dsnPattern.FindStringSubmatch(line); dsnMatches != nil {
		record.Dsn = dsnMatches[1]
	}

	if descMatches := ms.descriptionPattern.FindStringSubmatch(line); descMatches != nil {
		record.Description = descMatches[1]
	}

	record.MailProvider = ms.getMailProvider(record.Recipient)

	switch record.Status {
	case "sent":
		ms.delivered++
	case "bounced":
		ms.bounced++
		if _, ok := ms.bounceDetails[record.Dsn]; !ok {
			ms.bounceDetails[record.Dsn] = make(map[string]struct{})
		}
		ms.bounceDetails[record.Dsn][record.Description] = struct{}{}
	case "deferred":
		ms.deferred++
		ms.deferredTotal++
		if _, ok := ms.deferralDetails[record.Dsn]; !ok {
			ms.deferralDetails[record.Dsn] = make(map[string]struct{})
		}
		ms.deferralDetails[record.Dsn][record.Description] = struct{}{}
	}

	return record
}

// analyzeMessageID analyzes MessageID logs
func (ms *MaillogStat) analyzeMessageID(line string, logTime int64) interface{} {
	matches := ms.messageIDPattern.FindStringSubmatch(line)
	if matches == nil {
		return nil
	}

	messageIDMatches := regexp.MustCompile(`message-id=<([^>]+)>`).FindStringSubmatch(line)
	if messageIDMatches == nil {
		return nil
	}

	return &MailMessageID{
		MailRecord: MailRecord{
			PostfixMessageID: matches[1],
			LogTimeMillis:    logTime,
		},
		MessageID: messageIDMatches[1],
	}
}

// analyzeRemoved analyzes mail removed logs
func (ms *MaillogStat) analyzeRemoved(line string, logTime int64) interface{} {
	matches := ms.mailRemovedPattern.FindStringSubmatch(line)
	if matches == nil {
		return nil
	}

	return &MailRemoved{
		MailRecord: MailRecord{
			PostfixMessageID: matches[1],
			LogTimeMillis:    logTime,
		},
	}
}

// analyzeSender analyzes sender logs
func (ms *MaillogStat) analyzeSender(line string, logTime int64) interface{} {
	matches := ms.mailSenderPattern.FindStringSubmatch(line)
	if matches == nil {
		return nil
	}

	size, err := strconv.ParseInt(matches[3], 10, 64)
	if err != nil {
		return nil
	}

	return &MailSender{
		MailRecord: MailRecord{
			PostfixMessageID: matches[1],
			LogTimeMillis:    logTime,
		},
		Sender: matches[2],
		Size:   size,
	}
}

// AnalysisAndSaveToDatabase analyzes mail logs and saves to database
func (ms *MaillogStat) AnalysisAndSaveToDatabase(ctx context.Context) error {
	recordChan, err := ms.Analysis(ctx)
	if err != nil {
		return err
	}

	sendRecords := make([]*MailSendRecord, 256)
	receiveRecords := make([]*MailReceiveRecord, 256)
	messageIDRecords := make([]*MailMessageID, 256)
	senderRecords := make([]*MailSender, 256)
	removedRecords := make([]*MailRemoved, 256)
	deferredRecords := make([]*MailDeferredRecord, 256)

	for record := range recordChan {
		switch r := record.(type) {
		case *MailSendRecord:
			sendRecords = append(sendRecords, r)
		case *MailReceiveRecord:
			receiveRecords = append(receiveRecords, r)
		case *MailSender:
			senderRecords = append(senderRecords, r)
		case *MailRemoved:
			removedRecords = append(removedRecords, r)
		case *MailMessageID:
			messageIDRecords = append(messageIDRecords, r)
		case *MailDeferredRecord:
			deferredRecords = append(deferredRecords, r)
		default:
			continue
		}
	}

	if len(sendRecords) > 0 {
		_, err = g.DB().Model("mailstat_send_mails").OnDuplicate(g.Map{
			"status":          gdb.Raw("case when excluded.log_time_millis > log_time_millis then excluded.status else status end"),
			"delay":           gdb.Raw("case when excluded.log_time_millis > log_time_millis then excluded.delay else delay end"),
			"delays":          gdb.Raw("case when excluded.log_time_millis > log_time_millis then excluded.delays else delays end"),
			"dsn":             gdb.Raw("case when excluded.log_time_millis > log_time_millis then excluded.dsn else dsn end"),
			"relay":           gdb.Raw("case when excluded.log_time_millis > log_time_millis then excluded.relay else relay end"),
			"description":     gdb.Raw("case when excluded.log_time_millis > log_time_millis then excluded.description else description end"),
			"log_time_millis": gdb.Raw("case when excluded.log_time_millis > log_time_millis then excluded.log_time_millis else log_time_millis end"),
		}).Insert(sendRecords)
		if err != nil {
			glog.Error(context.Background(), err.Error())
		}
	}

	if len(receiveRecords) > 0 {
		_, err = g.DB().Model("mailstat_receive_mails").OnDuplicate(g.Map{
			"status":          gdb.Raw("case when excluded.log_time_millis > log_time_millis then excluded.status else status end"),
			"delay":           gdb.Raw("case when excluded.log_time_millis > log_time_millis then excluded.delay else delay end"),
			"delays":          gdb.Raw("case when excluded.log_time_millis > log_time_millis then excluded.delays else delays end"),
			"dsn":             gdb.Raw("case when excluded.log_time_millis > log_time_millis then excluded.dsn else dsn end"),
			"relay":           gdb.Raw("case when excluded.log_time_millis > log_time_millis then excluded.relay else relay end"),
			"description":     gdb.Raw("case when excluded.log_time_millis > log_time_millis then excluded.description else description end"),
			"log_time_millis": gdb.Raw("case when excluded.log_time_millis > log_time_millis then excluded.log_time_millis else log_time_millis end"),
		}).Insert(receiveRecords)
		if err != nil {
			glog.Error(context.Background(), err.Error())
		}
	}

	if len(senderRecords) > 0 {
		_, err = g.DB().Model("mailstat_senders").Insert(senderRecords)
		if err != nil {
			glog.Error(context.Background(), err.Error())
		}
	}

	if len(removedRecords) > 0 {
		_, err = g.DB().Model("mailstat_removed").Insert(removedRecords)
		if err != nil {
			glog.Error(context.Background(), err.Error())
		}
	}

	if len(messageIDRecords) > 0 {
		_, err = g.DB().Model("mailstat_message_ids").Insert(messageIDRecords)
		if err != nil {
			glog.Error(context.Background(), err.Error())
		}
	}

	if len(deferredRecords) > 0 {
		_, err = g.DB().Model("mailstat_deferred").Insert(deferredRecords)
		if err != nil {
			glog.Error(context.Background(), err.Error())
		}
	}

	return nil
}
