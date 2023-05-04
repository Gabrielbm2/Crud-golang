package main

import (
	"net/http"

	"gocrud/routes"
)

func main() {
	routes.CarregaRotas()
	http.ListenAndServe(":8000", nil)
}
