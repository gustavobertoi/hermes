package signatures

var (
	RSA = "RSA"
)

func IsValidAlgorithm(algorithm string) bool {
	switch algorithm {
	case RSA:
		return true
	default:
		return false
	}
}
