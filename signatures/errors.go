package signatures

import "errors"

var (
	ErrSignatureNotFound = errors.New("signature not found")
	ErrInvalidKeyType    = errors.New("invalid key type")
	ErrDecodingPemBlock  = errors.New("error decoding pem block")
	ErrUnknownPemType    = errors.New("unknown pem type")
	ErrNullPrivateKey    = errors.New("null private key")
	ErrNullPublicKey     = errors.New("null public key")
	ErrCipherTextShort   = errors.New("ciphertext too short")
)
