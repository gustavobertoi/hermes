package signatures

type Signature interface {
	Generate() error
	GetPrivateKey() interface{}
	GetPublicKey() interface{}
	GetPublicKeyString() string
	SetPrivateKey(privateKey interface{}) error
	SetPublicKey(publicKey interface{}) error
	Save(folderPath string) error
	Load(folderPath string) error
	Encrypt(data []byte) (*Output, error)
	Decrypt(output *Output) ([]byte, error)
}
