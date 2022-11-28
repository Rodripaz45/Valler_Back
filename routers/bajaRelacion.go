package routers

import (
	"net/http"

	"github.com/Rodripaz45/twitter/bd"
	"github.com/Rodripaz45/twitter/models"
)

func BajaRelacion(w http.ResponseWriter, r *http.Request){
	ID := r.URL.Query().Get("id")
	
	var t models.Relacion
	t.UsuarioID = IDUsuario
	t.UsuarioRelacionID = ID

	status, err := bd.BorroRelacion(t)
	if err != nil {
		http.Error(w, "Ocurrio un error al borrar relacion "+err.Error(), http.StatusBadRequest)
		return
	}
	if status == false {
		http.Error(w, "No se ha logrado borrar la relacion "+err.Error(), http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-type", "application/json")
	w.WriteHeader(http.StatusCreated)
}