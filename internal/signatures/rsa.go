package signatures

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
	"hash"
)

type RSASignature struct {
	privateKey *rsa.PrivateKey
	publicKey  *rsa.PublicKey
	hash       hash.Hash
	hashSum    [32]byte
	Signature
}

func NewRSASignature() *RSASignature {
	return &RSASignature{
		privateKey: nil,
		publicKey:  nil,
		hash:       nil,
		hashSum:    [32]byte{},
	}
}

func (s *RSASignature) Generate() error {
	s.hash = sha256.New()
	privateKey, err := rsa.GenerateKey(rand.Reader, 2048)
	if err != nil {
		return err
	}
	s.privateKey = privateKey
	s.publicKey = &privateKey.PublicKey
	return nil
}

func (s *RSASignature) GetPrivateKey() interface{} {
	return s.privateKey
}

func (s *RSASignature) GetPublicKey() interface{} {
	return s.publicKey
}

func (s *RSASignature) GetPublicKeyString() string {
	pubkey := s.publicKey
	return string(pubkey.N.Bytes())
}

func (s *RSASignature) SetPrivateKey(privateKey interface{}) error {
	pk, ok := privateKey.(*rsa.PrivateKey)
	if !ok {
		return ErrInvalidKeyType
	}
	s.privateKey = pk
	s.publicKey = &pk.PublicKey
	return nil
}

func (s *RSASignature) SetPublicKey(publicKey interface{}) error {
	pk, ok := publicKey.(*rsa.PublicKey)
	if !ok {
		return ErrInvalidKeyType
	}
	s.publicKey = pk
	return nil
}

func (s *RSASignature) Encrypt(data []byte) ([]byte, error) {
	s.hashSum = sha256.Sum256(data)
	return rsa.EncryptOAEP(
		s.hash,
		rand.Reader,
		s.publicKey,
		data,
		[]byte(""),
	)
}

func (s *RSASignature) Decrypt(data []byte) ([]byte, error) {
	return rsa.DecryptOAEP(s.hash, rand.Reader, s.privateKey, data, []byte(""))
}

func (s *RSASignature) GetHashSum() [32]byte {
	return s.hashSum
}

func (s *RSASignature) GetHash() hash.Hash {
	return s.hash
}
