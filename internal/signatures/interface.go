package signatures

import "hash"

type Signature interface {
	Generate() error
	GetPrivateKey() interface{}
	GetPublicKey() interface{}
	GetPublicKeyString() string
	SetPrivateKey(privateKey interface{}) error
	SetPublicKey(publicKey interface{}) error
	GetHashSum() [32]byte
	GetHash() hash.Hash
	Encrypt(data []byte) ([]byte, error)
	Decrypt(data []byte) ([]byte, error)
}
