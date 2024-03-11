package signatures

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"errors"
	"os"
	"path"

	"github.com/gustavobertoi/hermes/pkg"
)

const fileName = "key.txt"

type AESSignature struct {
	key []byte
	Signature
}

// NewAESSignature returns a new AESSignature
func NewAESSignature() *AESSignature {
	return &AESSignature{
		key: nil,
	}
}

func (a *AESSignature) Generate() error {
	// 32 bytes = AES 256
	key := make([]byte, 32)
	if _, err := rand.Read(key); err != nil {
		panic(err.Error())
	}
	a.key = key
	return nil
}

func (a *AESSignature) Save(folderPath string) error {
	return pkg.WriteToFile(folderPath, fileName, a.key)
}

func (a *AESSignature) Load(folderPath string) error {
	key, err := os.ReadFile(path.Join(folderPath, fileName))
	if err != nil {
		return err
	}
	a.key = key
	return nil
}

func (a *AESSignature) Encrypt(data []byte) ([]byte, error) {
	block, err := aes.NewCipher(a.key)
	if err != nil {
		return nil, err
	}
	output := make([]byte, aes.BlockSize+len(data))
	iv := output[:aes.BlockSize]
	if _, err := rand.Read(iv); err != nil {
		return nil, err
	}
	mode := cipher.NewCFBEncrypter(block, iv)
	mode.XORKeyStream(output[aes.BlockSize:], data)
	return output, nil
}

func (a *AESSignature) Decrypt(data []byte) ([]byte, error) {
	block, err := aes.NewCipher(a.key)
	if err != nil {
		return nil, err
	}
	if len(data) < aes.BlockSize {
		return nil, errors.New("ciphertext too short")
	}
	iv := data[:aes.BlockSize]
	data = data[aes.BlockSize:]
	mode := cipher.NewCFBDecrypter(block, iv)
	mode.XORKeyStream(data, data)
	return data, nil
}
