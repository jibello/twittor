package bd

import (
	"github.com/jibello/twittor73/models"
	"golang.org/x/crypto/bcrypt"
)

/*IntentoLogin login del usuario*/
func IntentoLogin(email string, password string) (models.Usuario, bool) {

	usu, encontrado, _ := ChequeoYaExisteUsuario(email)
	if encontrado == false {
		return usu, false
	}
	passwordBytes := []byte(password)
	passwordBD := []byte(usu.Password)

	err := bcrypt.CompareHashAndPassword(passwordBD, passwordBytes)
	if err != nil {
		return usu, false
	}
	return usu, true
}
