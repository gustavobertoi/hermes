package signatures

var (
	RSA = "RSA"
	AES = "AES"
)

func IsValidAlgorithm(algorithm string) bool {
	switch algorithm {
	case RSA:
		return true
	case AES:
		return true
	default:
		return false
	}
}
