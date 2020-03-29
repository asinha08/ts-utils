package password

import "golang.org/x/crypto/bcrypt"

func EncryptPassword(password string, cost int) string {
	pass := []byte(password)
	hashedPassword, err := bcrypt.GenerateFromPassword(pass, cost)
	if err != nil {
		panic(err)
	}
	return string(hashedPassword)
}

func DoPasswordMatch(hashedPassword, password string) (status bool) {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
	return err == nil
}
