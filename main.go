package main

import (
	"log"
	"twitter/src/db"
	"twitter/src/handlers"
)

func main() {
	if db.CheckConnection() == 0 {
		log.Fatal("Sin conexion a la BD")
		return
	}
	handlers.Manejadores()
}
