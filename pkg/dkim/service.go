package dkim

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"log"
)

func GenerateRSA(domain string, selector string, bitSize int) (string, string, error) {
	random := rand.Reader

	key, err := rsa.GenerateKey(random, bitSize)
	if err != nil {
		log.Println(err.Error())
		return "", "", err
	}

	privateKeyBlock := &pem.Block{
		Type:  "PRIVATE KEY",
		Bytes: x509.MarshalPKCS1PrivateKey(key),
	}

	publicKeyBlock := &pem.Block{
		Type:  "PUBLIC KEY",
		Bytes: x509.MarshalPKCS1PublicKey(&key.PublicKey),
	}

	publicKeyBlock.Headers = map[string]string{
		"v": "DKIM1",
		"k": "rsa",
		"s": selector,
		"d": domain,
	}

	privateKeyPEM := pem.EncodeToMemory(privateKeyBlock)
	publicKeyPEM := pem.EncodeToMemory(&pem.Block{Type: "PUBLIC KEY", Bytes: publicKeyBlock.Bytes})

	return string(privateKeyPEM), string(publicKeyPEM), nil
}
