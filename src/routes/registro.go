package routes

import (
	"encoding/json"
	"net/http"
	"twitter/src/db"
	"twitter/src/models"
)

// resgistro es la funcion para crear en la base de datos el usuario
func Registro(w http.ResponseWriter, r *http.Request) {
	var t models.Usuario
	err := json.NewDecoder(r.Body).Decode(&t)
	if err != nil {
		http.Error(w, "Error en los datos recibidos "+err.Error(), 400)
		return
	}
	if len(t.Email) == 0 {
		http.Error(w, "El email de usuario es requerido", 400)
		return
	}
	if len(t.Password) < 6 {
		http.Error(w, "Debe especificar una password de al menos 6 caracteres", 400)
		return
	}
	_, encontrado, _ := db.CheckUser(t.Email)
	if encontrado {
		http.Error(w, "Ya existe un usuario registrado con ese email", 400)
		return
	}
	_, status, err := db.InsertRegister(t)
	if err != nil {
		http.Error(w, "Ocurrio un error al intentar realizar un registro de usuario"+err.Error(), 400)
		return
	}
	if status == false {
		http.Error(w, "No se a logrado insertar un registro de usuario", 400)
		return
	}

	w.WriteHeader(http.StatusCreated)
}
