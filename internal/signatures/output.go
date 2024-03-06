package signatures

import (
	"encoding/base64"
	"encoding/hex"
	"hash"
	"os"
	"path"
)

type Output struct {
	hash    hash.Hash
	hashSum [32]byte
	content []byte
}

func NewOutput(hash hash.Hash, hashSum [32]byte, content []byte) *Output {
	return &Output{
		hash:    hash,
		hashSum: hashSum,
		content: content,
	}
}

func (o *Output) Hash() hash.Hash {
	return o.hash
}

func (o *Output) HashSum() [32]byte {
	return o.hashSum
}

func (o *Output) Content() []byte {
	return o.content
}

func (o *Output) SaveContentToFile(folderPath string, fileName string) error {
	encodedText := base64.StdEncoding.EncodeToString(o.content)
	filePath := path.Join(folderPath, fileName)
	file, err := os.Create(filePath)
	if err != nil {
		return err
	}
	defer file.Close()
	_, err = file.WriteString(encodedText)
	if err != nil {
		return err
	}
	return nil
}

func (o *Output) SaveHashSum(folderPath string, fileName string) error {
	hashHexString := hex.EncodeToString(o.hashSum[:])
	filePath := path.Join(folderPath, fileName)
	file, err := os.Create(filePath)
	if err != nil {
		return err
	}
	defer file.Close()
	_, err = file.WriteString(hashHexString)
	if err != nil {
		return err
	}
	return nil
}
