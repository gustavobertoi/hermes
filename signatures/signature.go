package signatures

func NewSignature(algorithm string) Signature {
	if algorithm == RSA {
		return NewRSASignature()
	}
	return nil
}
