package main

import (
	"log"

	"github.com/Rodripaz45/twitter/bd"
	"github.com/Rodripaz45/twitter/handlers"
)

func main() {
	if bd.ChequeoConnection() == 0 {
		log.Fatal("Sin conexion a la BD")
		return
	}
	handlers.Manejadores()
}
