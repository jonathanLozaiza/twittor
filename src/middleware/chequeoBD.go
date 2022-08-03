package middleware

import (
	"net/http"
	"twitter/src/db"
)

func ChequeoBD(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if db.CheckConnection() == 0 {
			http.Error(w, "Conexion perdida con la base de datos", 500)
			return
		}
		next.ServeHTTP(w, r)
	}
}
