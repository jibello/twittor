package bd

import "golang.org/x/crypto/bcrypt"

/*EncriptarPassword permite encriptar el password*/
func EncriptarPassword(pass string) (string, error) {
	costo := 8 //Pasadas al Password ^ 2
	bytes, err := bcrypt.GenerateFromPassword([]byte(pass), costo)
	return string(bytes), err
}
