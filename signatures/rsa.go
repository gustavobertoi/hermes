package signatures

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
	"crypto/x509"
	"encoding/pem"
	"io"
	"os"
	"path"
)

type RSASignature struct {
	privateKey *rsa.PrivateKey
	publicKey  *rsa.PublicKey
	Signature
}

func NewRSASignature() *RSASignature {
	return &RSASignature{
		privateKey: nil,
		publicKey:  nil,
	}
}

func (s *RSASignature) Generate() error {
	privateKey, err := rsa.GenerateKey(rand.Reader, 4096)
	if err != nil {
		return err
	}
	s.privateKey = privateKey
	s.publicKey = &privateKey.PublicKey
	return nil
}

func (s *RSASignature) Encrypt(data []byte) ([]byte, error) {
	data, err := rsa.EncryptOAEP(
		sha256.New(),
		rand.Reader,
		s.publicKey,
		data,
		[]byte(""),
	)
	if err != nil {
		return nil, err
	}
	return data, nil
}

func (s *RSASignature) Decrypt(data []byte) ([]byte, error) {
	data, err := rsa.DecryptOAEP(sha256.New(), rand.Reader, s.privateKey, data, []byte(""))
	if err != nil {
		return nil, err
	}
	return data, nil
}

func (s *RSASignature) Save(folderPath string) error {
	privateKeyBytes := x509.MarshalPKCS1PrivateKey(s.privateKey)
	publicKeyBytes, err := x509.MarshalPKIXPublicKey(s.publicKey)
	if err != nil {
		return err
	}

	privateKeyPath := path.Join(folderPath, "private_key.pem")
	err = saveRsaKeyToFile(privateKeyPath, privateKeyBytes, "RSA PRIVATE KEY")
	if err != nil {
		return err
	}

	publicKeyPath := path.Join(folderPath, "public_key.pem")
	err = saveRsaKeyToFile(publicKeyPath, publicKeyBytes, "RSA PUBLIC KEY")
	if err != nil {
		return err
	}

	return nil
}

func (s *RSASignature) Load(folderPath string) error {
	privateKeyPath := path.Join(folderPath, "private_key.pem")
	privateKey, publicKey, err := readRsaKeyFromFile(privateKeyPath)
	if err != nil {
		return err
	}
	if privateKey == nil {
		return ErrNullPrivateKey
	}
	if publicKey == nil {
		return ErrNullPublicKey
	}
	s.privateKey = privateKey
	s.publicKey = publicKey
	return nil
}

func saveRsaKeyToFile(filename string, keyBytes []byte, keyType string) error {
	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer file.Close()
	err = pem.Encode(file, &pem.Block{Type: keyType, Bytes: keyBytes})
	if err != nil {
		return err
	}
	return nil
}

func readRsaKeyFromFile(filePath string) (*rsa.PrivateKey, *rsa.PublicKey, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, nil, err
	}
	defer file.Close()
	keyBytes, err := io.ReadAll(file)
	if err != nil {
		return nil, nil, err
	}
	block, _ := pem.Decode(keyBytes)
	if block == nil {
		return nil, nil, ErrDecodingPemBlock
	}
	if block.Type == "RSA PRIVATE KEY" {
		privateKey, err := x509.ParsePKCS1PrivateKey(block.Bytes)
		if err != nil {
			return nil, nil, err
		}
		return privateKey, &privateKey.PublicKey, nil
	}
	if block.Type == "RSA PUBLIC KEY" {
		publicKey, err := x509.ParsePKIXPublicKey(block.Bytes)
		if err != nil {
			return nil, nil, err
		}
		return nil, publicKey.(*rsa.PublicKey), nil
	}
	return nil, nil, ErrUnknownPemType
}
