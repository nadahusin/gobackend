package libs

import "golang.org/x/crypto/bcrypt"

func HashingPassword(pass string) (string, error) {
	hashPassword, err := bcrypt.GenerateFromPassword([]byte(pass), bcrypt.DefaultCost)

	if err != nil {
		return "", err
	}
	return string(hashPassword), nil
}

func CheckPassword(hassPassword, passwordDB string) bool {

	err := bcrypt.CompareHashAndPassword([]byte(hassPassword), []byte(passwordDB))

	if err != nil {
		return false
	}
	return true
}
