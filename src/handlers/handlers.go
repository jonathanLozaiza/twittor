package handlers

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"twitter/src/middleware"
	"twitter/src/routes"

	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

func Manejadores() {
	router := mux.NewRouter()
	fmt.Println("baby")
	//routes
	router.HandleFunc("/registro", middleware.ChequeoBD(routes.Registro)).Methods("POST")
	router.HandleFunc("/login", middleware.ChequeoBD(routes.Login)).Methods("POST")
	router.HandleFunc("/verperfil", middleware.ChequeoBD(middleware.ValidoJWT(routes.VerPerfil))).Methods("GET")
	router.HandleFunc("/modificarperfil", middleware.ChequeoBD(middleware.ValidoJWT(routes.ModificarPerfil))).Methods("PUT")

	PORT := os.Getenv("PORT")
	if PORT == "" {
		PORT = "8080"
	}
	handler := cors.AllowAll().Handler(router)
	log.Fatal(http.ListenAndServe(":"+PORT, handler))
}
