package config

import (
	"errors"

	"github.com/gustavobertoi/hermes/internal/signatures"
)

type PersonalSignature struct {
	Algorithm string `json:"algorithm"`
	KeysPath  string `json:"keys_path"`
	signature *signatures.Signature
}

func (ps *PersonalSignature) GetSignature() (*signatures.Signature, error) {
	if ps.signature != nil {
		return ps.signature, nil
	}
	err := ps.loadSignature()
	if err != nil {
		return nil, err
	}
	return ps.signature, nil
}

func (ps *PersonalSignature) loadSignature() error {
	signature := signatures.NewSignature(ps.Algorithm)
	if signature == nil {
		return errors.New("signature algorithm not supported")
	}
	err := signature.Load(ps.KeysPath)
	if err != nil {
		return err
	}
	ps.signature = &signature
	return nil
}
