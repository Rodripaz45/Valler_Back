package routers

import (
	"encoding/json"
	"net/http"
	"strconv"
	"github.com/Rodripaz45/twitter/bd"
)

/* ListaUsuarios leo la lista de usuarios */
func ListaUsuarios(w http.ResponseWriter, r *http.Request){

	typeUser := r.URL.Query().Get("type") 
	page := r.URL.Query().Get("page")
	search := r.URL.Query().Get("search")

	pagTemp, err := strconv.Atoi(page)
	if err != nil{
		http.Error(w, "Debe enviar el formato pagina con un entero mayor a 0", http.StatusBadRequest)
		return
	}

	pag := int64(pagTemp)

	result, status := bd.LeoUsuariosTodos(IDUsuario, pag, search, typeUser)
	if status == false {
		http.Error(w, "Error al leer el usuario", http.StatusBadRequest)
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(result)
}