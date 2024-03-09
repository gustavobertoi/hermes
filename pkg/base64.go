package pkg

import "encoding/base64"

func IsBase64(s string) bool {
	_, err := base64.StdEncoding.DecodeString(s)
	return err == nil
}

func EncodeToBase64(s string) string {
	return base64.StdEncoding.EncodeToString([]byte(s))
}

func DecodeFromBase64(s string) ([]byte, error) {
	decoded, err := base64.StdEncoding.DecodeString(s)
	return decoded, err
}
