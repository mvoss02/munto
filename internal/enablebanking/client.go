package enablebanking

import (
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"fmt"
	"os"
)

// API Respone
type Env string

const (
	EnvSandbox    Env = "SANDBOX"
	EnvProduction Env = "PRODUCTION"
)

type Application struct {
	Name         string   `json:"name"`
	Description  string   `json:"description"`
	Kid          string   `json:"kid"`
	Environment  Env      `json:"environment"`
	RedirectURLs []string `json:"redirect_urls"`
	Active       bool     `json:"active"`
	Countries    []string `json:"countries"`
	Services     []string `json:"services"`
}

func LoadKey(path string) (*rsa.PrivateKey, error) {
	raw, err := os.ReadFile(path)
	if err != nil {
		return nil, fmt.Errorf("read %s: %w", path, err)
	}

	block, _ := pem.Decode(raw)
	if block == nil {
		return nil, fmt.Errorf("no PEM block in %s", path)
	}

	key, err := x509.ParsePKCS8PrivateKey(block.Bytes)
	if err != nil {
		return nil, fmt.Errorf("parse key: %w", err)
	}

	keyParsed, ok := key.(*rsa.PrivateKey)
	if !ok {
		return nil, fmt.Errorf("expected rsa.PrivateKey, got %T", key)
	}
	return keyParsed, nil
}
