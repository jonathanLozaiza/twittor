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
	router.HandleFunc("/tweet", middleware.ChequeoBD(middleware.ValidoJWT(routes.GraboTweet))).Methods("POST")
	router.HandleFunc("/leoTeewts", middleware.ChequeoBD(middleware.ValidoJWT(routes.LeoTeewts))).Methods("GET")
	router.HandleFunc("/eliminarTweet", middleware.ChequeoBD(middleware.ValidoJWT(routes.EliminarTweet))).Methods("DELETE")

	router.HandleFunc("/subirAvatar", middleware.ChequeoBD(middleware.ValidoJWT(routes.SubirAvatar))).Methods("POST")
	router.HandleFunc("/obtenerAvatar", middleware.ChequeoBD(middleware.ValidoJWT(routes.ObtenerAvatar))).Methods("GET")
	router.HandleFunc("/subirBanner", middleware.ChequeoBD(middleware.ValidoJWT(routes.SubirBanner))).Methods("POST")
	router.HandleFunc("/obtenerBanner", middleware.ChequeoBD(middleware.ValidoJWT(routes.ObtenerBanner))).Methods("GET")

	router.HandleFunc("/altarelacion", middleware.ChequeoBD(middleware.ValidoJWT(routes.AltaRelacion))).Methods("POST")
	router.HandleFunc("/bajarelacion", middleware.ChequeoBD(middleware.ValidoJWT(routes.BajaRelacion))).Methods("DELETE")
	router.HandleFunc("/consultarelacion", middleware.ChequeoBD(middleware.ValidoJWT(routes.ConsultaRelacion))).Methods("GET")

	router.HandleFunc("/listaUsuarios", middleware.ChequeoBD(middleware.ValidoJWT(routes.ListaUsuarios))).Methods("GET")
	router.HandleFunc("/leoTweetsSeguidores", middleware.ChequeoBD(middleware.ValidoJWT(routes.LeoTweetsSeguidores))).Methods("GET")

	PORT := os.Getenv("PORT")
	if PORT == "" {
		PORT = "8080"
	}
	handler := cors.AllowAll().Handler(router)
	log.Fatal(http.ListenAndServe(":"+PORT, handler))
}
