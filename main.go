package main

import (
	"log"

	"github.com/thiagossjc/redsocial/db"
	"github.com/thiagossjc/redsocial/handlers"
)

func main() {
	if db.CheckConnection() == 0 {
		log.Fatal("Sin conexión a la banco de dados")
		return
	}
	handlers.Manejadores()
}
