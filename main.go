package main

import (
	"net/http"

	"github.com/armedcor/learn_GoLang/controllers"
)

func main() {
	controllers.RegisterController()
	http.ListenAndServe(":3000", nil)
}
