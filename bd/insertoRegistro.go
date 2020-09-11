package bd

import (
	"context"
	"fmt"
	"time"

	"github.com/jibello/twittor73/models"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

/*InsertoRegistro es la parada final con la BD para insertar
los datos del usuario*/
func InsertoRegistro(u models.Usuario) (string, bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	db := MongoCN.Database("twittor73")
	col := db.Collection("usuarios")

	u.Password, _ = EncriptarPassword(u.Password)

	result, err := col.InsertOne(ctx, u)

	if err != nil {
		return "", false, err
	}

	ObjID, _ := result.InsertedID.(primitive.ObjectID)
	fmt.Println("ObjID", ObjID)

	return ObjID.String(), true, nil
}
