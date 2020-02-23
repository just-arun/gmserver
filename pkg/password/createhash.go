package password

import "golang.org/x/crypto/bcrypt"

// CreateHash function create hash from entered password
func CreateHash(pwd []byte) (string, error) {
	hash, err := bcrypt.GenerateFromPassword(pwd, bcrypt.DefaultCost)
	if err != nil {
		return string(hash), err
	}
	return string(hash), nil
}
