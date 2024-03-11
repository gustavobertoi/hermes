package signatures

type Key interface{}

type Signature interface {
	Generate() error
	Save(folderPath string) error
	Load(folderPath string) error
	Encrypt(data []byte) ([]byte, error)
	Decrypt(data []byte) ([]byte, error)
}
