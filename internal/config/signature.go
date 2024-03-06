package config

import "github.com/gustavobertoi/hermes/internal/signatures"

type PersonalSignature struct {
	PrivateKeyPath string
	PublicKeyPath  string
}

func NewPersonalSignature() *PersonalSignature {
	return &PersonalSignature{
		Signature: signatures.NewSignature(signatures.RSA),
	}
}
