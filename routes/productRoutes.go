package routes

import (
	"fmt"
	"net/http"

	"github.com/nicolasrsaraiva/market-api-go/controllers"
)

func HandleProductRequests() {
	fmt.Println("Product routes initialized")
	http.HandleFunc("/api/product", controllers.CreateProductController)
	http.HandleFunc("/api/products", controllers.ReadAllProductsController)
}
