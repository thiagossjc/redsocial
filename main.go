package main

import (
	"log"
	"redsocial/db"
	"redsocial/handlers"
)

func main() {
	if db.CheckConnection() == 0 {
		log.Fatal("Sin conexión a la banco de dados")
		return
	}
	handlers.Manejadores()
}
