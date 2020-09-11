package bd

import (
	"context"
	"fmt"
	"time"

	"github.com/jibello/twittor73/models"
	"go.mongodb.org/mongo-driver/bson"
)

/*ChequeoYaExisteUsuario recibe un email y comprueba si existe*/
func ChequeoYaExisteUsuario(email string) (models.Usuario, bool, string) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	db := MongoCN.Database("Twittor73")
	col := db.Collection("usuarios")

	condicion := bson.M{"email": email}

	var resultado models.Usuario

	fmt.Println("Resultado: ", resultado)

	err := col.FindOne(ctx, condicion).Decode(&resultado)
	fmt.Println("err", err)
	ID := resultado.ID.Hex()

	if err != nil {
		return resultado, false, ID
	}
	return resultado, true, ID
}
