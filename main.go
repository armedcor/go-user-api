package main

import (
	"net/http"

	"github.com/armedcor/go-user-api/controllers"
)

func main() {
	controllers.RegisterController()
	http.ListenAndServe(":3000", nil)
}
