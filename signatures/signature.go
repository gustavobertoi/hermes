package signatures

func NewSignature(algorithm string) Signature {
	if algorithm == RSA {
		return NewRSASignature()
	}
	if algorithm == AES {
		return NewAESSignature()
	}
	return nil
}
