package signatures

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"crypto/sha256"
	"encoding/base64"
	"io"
	"os"
)

type AESSignature struct {
	privateKey []byte
	publicKey  []byte
	Signature
}

// NewAESSignature returns a new AESSignature
func NewAESSignature() *AESSignature {
	return &AESSignature{
		privateKey: nil,
		publicKey:  nil,
	}
}

// Generate generates a new private and public key for AESSignature
func (as *AESSignature) Generate() error {
	privateKey := make([]byte, 32)
	if _, err := rand.Read(privateKey); err != nil {
		return err
	}

	publicKey := make([]byte, 32)
	if _, err := rand.Read(publicKey); err != nil {
		return err
	}

	as.privateKey = privateKey
	as.publicKey = publicKey

	return nil
}

// GetPrivateKey returns the private key of AESSignature
func (as *AESSignature) GetPrivateKey() interface{} {
	return as.privateKey
}

// GetPublicKey returns the public key of AESSignature
func (as *AESSignature) GetPublicKey() interface{} {
	return as.publicKey
}

// GetPublicKeyString returns the base64-encoded public key as a string
func (as *AESSignature) GetPublicKeyString() string {
	return base64.StdEncoding.EncodeToString(as.publicKey)
}

// SetPrivateKey sets the private key of AESSignature
func (as *AESSignature) SetPrivateKey(privateKey interface{}) error {
	key, ok := privateKey.([]byte)
	if !ok {
		return ErrInvalidKeyType
	}
	as.privateKey = key
	return nil
}

// SetPublicKey sets the public key of AESSignature
func (as *AESSignature) SetPublicKey(publicKey interface{}) error {
	key, ok := publicKey.([]byte)
	if !ok {
		return ErrInvalidKeyType
	}
	as.publicKey = key
	return nil
}

// Save saves the private and public keys to files in the specified folder path
func (as *AESSignature) Save(folderPath string) error {
	if err := writeFile(folderPath+"/private_key.txt", as.privateKey); err != nil {
		return err
	}
	if err := writeFile(folderPath+"/public_key.txt", as.publicKey); err != nil {
		return err
	}
	return nil
}

// Load loads the private and public keys from files in the specified folder path
func (as *AESSignature) Load(folderPath string) error {
	privateKey, err := readFile(folderPath + "/private_key.txt")
	if err != nil {
		return err
	}

	publicKey, err := readFile(folderPath + "/public_key.txt")
	if err != nil {
		return err
	}

	as.privateKey = privateKey
	as.publicKey = publicKey

	return nil
}

// Encrypt encrypts the data using AES 256
func (as *AESSignature) Encrypt(data []byte) (*Output, error) {
	block, err := aes.NewCipher(as.publicKey)
	if err != nil {
		return nil, err
	}

	ciphertext := make([]byte, aes.BlockSize+len(data))
	iv := ciphertext[:aes.BlockSize]
	if _, err := io.ReadFull(rand.Reader, iv); err != nil {
		return nil, err
	}

	mode := cipher.NewCFBEncrypter(block, iv)
	mode.XORKeyStream(ciphertext[aes.BlockSize:], data)

	hashSum := sha256.Sum256(data)

	return &Output{
		hash:    sha256.New(),
		hashSum: hashSum,
		content: ciphertext,
	}, nil
}

// Decrypt decrypts the data using AES 256
func (as *AESSignature) Decrypt(data []byte) (*Output, error) {
	block, err := aes.NewCipher(as.privateKey)
	if err != nil {
		return nil, err
	}

	if len(data) < aes.BlockSize {
		return nil, ErrCipherTextShort
	}

	iv := data[:aes.BlockSize]
	data = data[aes.BlockSize:]

	output := make([]byte, len(data))

	mode := cipher.NewCFBDecrypter(block, iv)
	mode.XORKeyStream(output, data)

	hashSum := sha256.Sum256(output)

	return &Output{
		hash:    sha256.New(),
		hashSum: hashSum,
		content: output,
	}, nil
}

// writeFile writes data to a file
func writeFile(filename string, data []byte) error {
	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer file.Close()
	_, err = file.Write(data)
	return err
}

// readFile reads data from a file
func readFile(filename string) ([]byte, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()
	return io.ReadAll(file)
}
