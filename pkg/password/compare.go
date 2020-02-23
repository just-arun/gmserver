package password

import "golang.org/x/crypto/bcrypt"

// CompareHash function compare hashpassword and entered password
func CompareHash(hash []byte, pwd []byte) error {
	err := bcrypt.CompareHashAndPassword(hash, pwd)
	return err
}
