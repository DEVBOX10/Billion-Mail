package public

import (
	"billionmail-core/internal/consts"
	"context"
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"crypto/x509/pkix"
	"encoding/pem"
	"github.com/gogf/gf/v2/frame/g"
	"io/fs"
	"math/big"
	random "math/rand"
	"net"
	"os"
	"path/filepath"
	"time"
)

type selfSignedCertProvider struct {
	certFile       string
	privateKeyFile string
}

// SelfSignedCert creates a new self-signed certificate provider
func SelfSignedCert() *selfSignedCertProvider {
	return &selfSignedCertProvider{
		certFile:       AbsPath(filepath.Join(consts.SSL_PATH, "cert.pem")),
		privateKeyFile: AbsPath(filepath.Join(consts.SSL_PATH, "key.pem")),
	}
}

// Generate creates self-signed certificate if not exists
func (sp *selfSignedCertProvider) Generate() {
	// Only generate when SSL certificate not configured
	_, certFileExists := os.Stat(sp.certFile)
	_, privateKeyFileExists := os.Stat(sp.privateKeyFile)

	if certFileExists == nil && privateKeyFileExists == nil {
		return
	}

	g.Log().Debug(context.Background(), "No SSL certificate configured, generating self-signed certificate...")

	if err := sp.createOffline(); err != nil {
		return
	}

	g.Log().Debug(context.Background(), "Self-signed certificate generated successfully")
}

// createOffline creates self-signed certificate offline
func (sp *selfSignedCertProvider) createOffline() error {
	// Ensure ssl directory exists
	err := os.MkdirAll(filepath.Dir(sp.certFile), 0755)

	if err != nil {
		g.Log().Error(context.Background(), "Failed to generate self-signed certificate:", err)
		return err
	}

	csr, key, err := sp.createCertWithDynamic()

	if err != nil {
		g.Log().Error(context.Background(), "Failed to generate self-signed certificate:", err)
		return err
	}

	err = os.WriteFile(sp.certFile, csr, fs.ModePerm)

	if err != nil {
		g.Log().Error(context.Background(), "Failed to generate self-signed certificate:", err)
		return err
	}

	err = os.WriteFile(sp.privateKeyFile, key, fs.ModePerm)

	if err != nil {
		g.Log().Error(context.Background(), "Failed to generate self-signed certificate:", err)
		return err
	}

	return nil
}

// createCertWithRoot generates certificate using existing CA
func (sp *selfSignedCertProvider) createCertWithRoot(rootCa, rootKey []byte) ([]byte, []byte, error) {
	// Load CA certificate
	ca, caKey, err := sp.loadRootCertificate(rootCa, rootKey)

	if err != nil {
		g.Log().Error(context.Background(), "Failed to load CA certificate:", err)
		return []byte{}, []byte{}, err
	}

	// Generate certificate template
	templateCsr := sp.generateCertTemplate()

	// Issue certificate
	csr, key, err := sp.generateCert(ca, templateCsr, caKey)

	if err != nil {
		g.Log().Error(context.Background(), "Failed to issue certificate:", err)
		return []byte{}, []byte{}, err
	}

	return csr, key, nil
}

// createCertWithDynamic dynamically generates CA certificate and issues certificate
func (sp *selfSignedCertProvider) createCertWithDynamic() ([]byte, []byte, error) {
	templateCa := sp.generateCACertTemplate()

	caCsr, caKey, err := sp.generateCACert(templateCa)

	if err != nil {
		return []byte{}, []byte{}, err
	}

	return sp.createCertWithRoot(caCsr, caKey)
}

// loadRootCertificate loads CA certificate and private key
func (sp *selfSignedCertProvider) loadRootCertificate(rootCa, rootKey []byte) (*x509.Certificate, *rsa.PrivateKey, error) {
	// Parse root certificate
	caBlock, _ := pem.Decode(rootCa)
	rootCert, err := x509.ParseCertificate(caBlock.Bytes)

	if err != nil {
		return nil, nil, err
	}

	// Parse private key
	keyBlock, _ := pem.Decode(rootKey)
	praKey, err := x509.ParsePKCS1PrivateKey(keyBlock.Bytes)
	if err != nil {
		return nil, nil, err
	}

	return rootCert, praKey, nil
}

// generateCACertTemplate generates CA certificate template
func (sp *selfSignedCertProvider) generateCACertTemplate() *x509.Certificate {
	commonName := "Billion-Mail"
	rd := random.New(random.NewSource(time.Now().UnixNano()))
	cer := &x509.Certificate{
		SerialNumber: big.NewInt(rd.Int63()), // Certificate serial number
		Subject: pkix.Name{
			Country:            []string{"HK"},
			Organization:       []string{"aaPanel"},
			OrganizationalUnit: []string{"aaPanel"},
			Province:           []string{"HongKong"},
			CommonName:         commonName,
			Locality:           []string{"HongKong"},
		},
		NotBefore:             time.Now(),                                                                 // Validity start time
		NotAfter:              time.Now().AddDate(1, 0, 0),                                                // Validity end time
		BasicConstraintsValid: true,                                                                       // Basic constraints
		IsCA:                  true,                                                                       // Is CA certificate
		ExtKeyUsage:           []x509.ExtKeyUsage{x509.ExtKeyUsageClientAuth, x509.ExtKeyUsageServerAuth}, // Certificate purposes
		KeyUsage:              x509.KeyUsageDigitalSignature | x509.KeyUsageCertSign,
		EmailAddresses:        []string{"billionmail@aapanel.com"},
	}

	return cer
}

// generateCertTemplate generates certificate template
func (sp *selfSignedCertProvider) generateCertTemplate() *x509.Certificate {
	// Get server IP
	serverIp, localIp, err := GetServerIPAndLocalIP()

	if err != nil {
		g.Log().Error(context.Background(), "Failed to get server IP:", err)
		return nil
	}

	commonName := "Billion-Mail"
	rd := random.New(random.NewSource(time.Now().UnixNano()))
	cer := &x509.Certificate{
		SerialNumber: big.NewInt(rd.Int63()), // Certificate serial number
		Subject: pkix.Name{
			Country:            []string{"HK"},
			Organization:       []string{"aaPanel"},
			OrganizationalUnit: []string{"aaPanel"},
			Province:           []string{"HongKong"},
			CommonName:         commonName,
			Locality:           []string{"HongKong"},
		},
		NotBefore:             time.Now(),                                                                 // Validity start time
		NotAfter:              time.Now().AddDate(1, 0, 0),                                                // Validity end time
		BasicConstraintsValid: true,                                                                       // Basic constraints
		IsCA:                  false,                                                                      // Is CA certificate
		ExtKeyUsage:           []x509.ExtKeyUsage{x509.ExtKeyUsageClientAuth, x509.ExtKeyUsageServerAuth}, // Certificate purposes
		KeyUsage:              x509.KeyUsageDigitalSignature | x509.KeyUsageDataEncipherment,
		EmailAddresses:        []string{"billionmail@aapanel.com"},
		IPAddresses:           []net.IP{net.ParseIP(serverIp), net.ParseIP(localIp)}, // Supported IP addresses
		DNSNames:              []string{},                                            // Supported domains
	}

	return cer
}

// generateCACert issues CA certificate
func (sp *selfSignedCertProvider) generateCACert(templateCert *x509.Certificate) ([]byte, []byte, error) {
	// Generate key pair
	priKey, err := rsa.GenerateKey(rand.Reader, 4096)
	if err != nil {
		return nil, nil, err
	}

	ca, err := x509.CreateCertificate(rand.Reader, templateCert, templateCert, &priKey.PublicKey, priKey)
	if err != nil {
		return nil, nil, err
	}

	// Encode certificate and private key
	caPem := &pem.Block{
		Type:  "CERTIFICATE",
		Bytes: ca,
	}
	ca = pem.EncodeToMemory(caPem)
	buf := x509.MarshalPKCS1PrivateKey(priKey)
	keyPem := &pem.Block{
		Type:  "PRIVATE KEY",
		Bytes: buf,
	}
	key := pem.EncodeToMemory(keyPem)

	return ca, key, nil
}

// generateCert issues certificate using CA
func (sp *selfSignedCertProvider) generateCert(rootCert, templateCert *x509.Certificate, rootKey *rsa.PrivateKey) ([]byte, []byte, error) {
	// Generate key pair
	priKey, err := rsa.GenerateKey(rand.Reader, 1024)
	if err != nil {
		return nil, nil, err
	}

	ca, err := x509.CreateCertificate(rand.Reader, templateCert, rootCert, &priKey.PublicKey, rootKey)
	if err != nil {
		return nil, nil, err
	}

	// Encode certificate and private key
	caPem := &pem.Block{
		Type:  "CERTIFICATE",
		Bytes: ca,
	}
	ca = pem.EncodeToMemory(caPem)
	buf := x509.MarshalPKCS1PrivateKey(priKey)
	keyPem := &pem.Block{
		Type:  "PRIVATE KEY",
		Bytes: buf,
	}
	key := pem.EncodeToMemory(keyPem)

	return ca, key, nil
}
