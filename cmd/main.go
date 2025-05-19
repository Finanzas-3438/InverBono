package main

import (
	"log"
	"net/http"

	"github.com/Finanzas-3438/InverBono.git/pkg/web"
)

func main() {
	web.Router()

	log.Println("Servidor iniciado en http://localhost:8080")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal(err)
	}
}
