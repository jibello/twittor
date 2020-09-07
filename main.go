package main

import (
	"log"

	"github.com/jibello/twittor73/bd"
	"github.com/jibello/twittor73/handlers"
)

func main() {
	if bd.ChequeoConnection() == 0 {
		log.Fatal("Sin conexión a la BD")
		return
	}
	handlers.Manejadores()
}
