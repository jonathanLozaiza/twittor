package db

import (
	"context"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// MongoCN es el objeto de coneccion a la base de datos
var MongoCN = ConectarDB()
var clientOptions = options.Client().ApplyURI("mongodb+srv://jonathan:Isabela-10@cluster0.r8g89.mongodb.net/?retryWrites=true&w=majority")

func ConectarDB() *mongo.Client {
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
	log.Println("Conexion Exitosa con la DB")
	return client
}

func CheckConnection() int {
	err := MongoCN.Ping(context.TODO(), nil)
	if err != nil {
		return 0
	}
	return 1
}
