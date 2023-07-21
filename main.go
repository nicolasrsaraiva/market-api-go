package main

import (
	"net/http"

	"github.com/nicolasrsaraiva/market-api-go/routes"
)

func main() {
	routes.HandleProductRequests()
	http.ListenAndServe("127.0.0.1:8080", nil)
}
