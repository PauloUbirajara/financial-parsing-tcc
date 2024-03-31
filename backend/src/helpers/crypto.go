package helpers

import "crypto/sha512"

func StringToSHA512Sum(value string) string {
	hash := sha512.New()
	hash.Write([]byte(value))
	sha512sum := hash.Sum(nil)
	return string(sha512sum)
}
