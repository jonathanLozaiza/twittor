package routes

import (
	"encoding/json"
	"net/http"
	"time"
	"twitter/src/db"
	"twitter/src/jwt"
	"twitter/src/models"
)

func Login(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("content-type", "application/json")

	var t models.Usuario

	err := json.NewDecoder(r.Body).Decode(&t)
	if err != nil {
		http.Error(w, "Usuario y/o contraseña invalida"+err.Error(), 400)
		return
	}
	if len(t.Email) == 0 {
		http.Error(w, "El email del usuario es requerido ", 400)
		return
	}
	documento, existe := db.IntentoLogin(t.Email, t.Password)
	if existe == false {
		http.Error(w, "Usuario y/o password invalidos", 400)
		return
	}
	jwtKey, err := jwt.GeneroJWT(documento)
	if err != nil {
		http.Error(w, "Ocurrio un error al intentar generar el token correspondiente "+err.Error(), 400)
		return
	}
	resp := models.RespuestaLogin{
		Token: jwtKey,
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(resp)

	expirationTime := time.Now().Add(24 * time.Hour)
	http.SetCookie(w, &http.Cookie{
		Name:    "token",
		Value:   jwtKey,
		Expires: expirationTime,
	})
}
