package middlew

import (
	"net/http"
)

/*ValidoJWT permite validar el JWT que nos viene en la petición.*/
func ValidoJWT(next http.HandleFunc) {
	return func(w http.ResponseWriter, r *http.Request) {
		_, _, _, err := routers.ProcesoToken(r.Header.Get("Authorization"))
		if err != nil {
			http.Error(w, "Error en el Token!"+err.Error(), http.StatusBadRequest)
			return
		}
		next.ServeHTTP(w, r)
	}
}