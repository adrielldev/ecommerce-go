package main

import (
	"net/http"
	"text/template"

	"github.com/adrielldev/ecommerce-go/routes"
)

var temp = template.Must(template.ParseGlob("templates/*.html"))

func main() {
	routes.CarregaRotas()
	http.ListenAndServe("127.0.0.1:8000", nil)
}
