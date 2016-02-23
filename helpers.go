package main

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"errors"
	"fmt"
	"os"
	"time"
)

func checkKeyFile() bool {
	if _, err := os.Stat(KEY); err == nil {
		return true
	}
	return false
}

func generateKeyFile() error {
	privatekey, err := rsa.GenerateKey(rand.Reader, 2048)

	if err != nil {
		fmt.Printf("Error: %v", err)
		return err
	}

	pemfile, _ := os.Create(KEY)
	var pemkey = &pem.Block{
		Type:  "RSA PRIVATE KEY",
		Bytes: x509.MarshalPKCS1PrivateKey(privatekey)}
	pem.Encode(pemfile, pemkey)
	pemfile.Close()
	if checkKeyFile() {
		return nil
	}
	err = errors.New("Couldn't create key file!!!")
	return err
}

// Return true if current date is higher input date
func expired(exp int64) bool {
	now := time.Now().Unix()
	if now > exp {
		return true
	}
	return false
}