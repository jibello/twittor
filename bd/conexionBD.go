package bd

import (
	"context"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

/*MongoCN es el objeto de la conexión a la BD*/
var MongoCN = ConectarBD()
var clientOptions = options.Client().ApplyURI("mongodb+srv://jibello:Tubular73@twitter.ccko0.mongodb.net/Twittor73?retryWrites=true&w=majority")

/*ConectarBD es la función de conexión a la BD */
func ConectarBD() *mongo.Client {
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		log.Fatal(err.Error())
		return client
	}

	err = client.Ping(context.TODO(), nil)
	if err != nil {
		log.Fatal(err.Error())
		return client
	}
	log.Println("Conexión Exitosa con la BD")
	return client
}

/*ChequeoConnection es el Ping a la BD*/
func ChequeoConnection() int {
	err := MongoCN.Ping(context.TODO(), nil)
	if err != nil {
		return 0
	}
	return 1
}
