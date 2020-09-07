package routers

import (
	"encoding/json"
	"net/http"

	"github.com/jibello/twittor73/bd"
	"github.com/jibello/twittor73/models"
)

/*Registro (EndPoint) Crea en la BD el usuario*/
func Registro(w http.ResponseWriter, r *http.Request) {
	var t models.Usuario

	//El objeto Body (stream) solo se puede leer una vez, luego se destruye.
	err := json.NewDecoder(r.Body).Decode(&t)
	if err != nil {
		http.Error(w, "Error en los datos recibidos. "+err.Error(), 400)
		return
	}

	if len(t.Email) == 0 {
		http.Error(w, "El email del usuario es requerido", 400)
		return
	}

	if len(t.Password) < 6 {
		http.Error(w, "El password debe contener, al menos, 6 caracteres.", 400)
		return
	}

	_, encontrado, _ := bd.ChequeoYaExisteUsuario(t.Email)

	if encontrado == true {
		http.Error(w, "El usuario "+t.Email+" ya existe.", 400)
		return
	}

	_, status, err := bd.InsertoRegistro(t)
	if err != nil {
		http.Error(w, "Hubo un problema en el registro del usuario."+err.Error(), 400)
		return
	}

	if status == false {
		http.Error(w, "No se ha logrado insertar el registro del usuario.", 400)
		return
	}

	w.WriteHeader(http.StatusCreated)

}
