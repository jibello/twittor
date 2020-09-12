package routers

import (
	"errors"
	"strings"

	jwt "github.com/dgrijalva/jwt.go"
	"github.com/jibello/twittor73/bd"
	"github.com/jibello/twittor73/models"
)

/*Email valor del Email usado en todos los EndPoints*/
var Email string

/*IDUsuario es el ID devuel del modelo que se usará aen todos los EndPoints*/
var IDUsuario string

/*ProcesoToken Proceso token para extraer sus valores*/
func ProcesoToken(tk string) (*models.Claim, bool, string, error) {
	miClave := []byte("MastersdelDesarrollo_grupodeFacebook")
	claims := &models.Claim{}
	splitToken := strings.Split(tk, "Bearer")
	if len(splitToken) != 2 {
		return claims, false, string(""), errors.New("formato de token invalido")
	}
	tk = strings.TrimSpace(splitToken[1])
	tkn, err := jwt.ParseWithClaims(tk, claims, func(token *jwt.Token) (interface{}, error) {
		return miClave, nil
	})

	if err == nil {
		_, encontrado, _ := bd.ChequeoYaExisteUsuario(claims.Email)
		if encontrado == true {
			Email = claims.Email
			IDUsuario = claims.ID.Hex()
		}
		return claims, encontrado, ID, nil
	}

	if !tkn.Valid {
		return claims, false, string(""), erros.New("formato de token invalido")
	}
	return claims, false, string(""), err
}
