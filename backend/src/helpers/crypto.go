package helpers

import (
	"crypto/sha512"
	"fmt"
)

func StringToSHA512Sum(value string) string {
	hash := sha512.New()
	hash.Write([]byte(value))
	sha512sum := hash.Sum(nil)
	return fmt.Sprintf("%x", sha512sum)
}
