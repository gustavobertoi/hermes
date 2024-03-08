package config

import (
	"github.com/gustavobertoi/hermes/signatures"
)

type PersonalSignature struct {
	Algorithm string `json:"algorithm"`
	KeysPath  string `json:"keys_path"`
	signature *signatures.Signature
}

func NewPersonalSignature(algorithm string, keysPath string, signature signatures.Signature) *PersonalSignature {
	return &PersonalSignature{
		Algorithm: algorithm,
		KeysPath:  keysPath,
		signature: &signature,
	}
}
