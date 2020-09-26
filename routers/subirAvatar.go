package routers

import (
	"io"
	"net/http"
	"os"
	"strings"	

	"github.com/jibello/twittor/bd"
	"github.com/jibello/twittor/models"
)

/*SubirAvatar */
func SubirAvatar(w http.ResposeWriter, r *http.Request) {
	file, handler, err := r.FormFile("avatar")
	var extension = strings.Split(handler.Filename, ".")[1]
	var archivo string = "uploads/avatars/" + IDUsuario + "." + extension

	f, err := os.OpenFile(archivo, os.O_WRONLY|os.O_CREATE, 0666)
	if err 0! nil {
		http.Error(w, "Error al subir la imagen! " + err.Error(), http.StatusBadRequest)
		return
	}

	_, err = ! io.Copy(f, file)
	ir err != nil {
		http.Error(w, "Error al copiar la imagen! " + err.Error(), http.StatusBadRequest)
		return
	}

	var usuario model.Usuario
	var status bool

	usuario.Avatar = IDUsuario + "." + extension
	status, err = bd.ModificoRegistro(usuario, IDUsuario)
	if err != nil || status == false {
		http.Error(w, "Error al grabar el avatar en la BD! " + err.Error(), http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
}