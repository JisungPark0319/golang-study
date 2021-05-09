package main

import (
	"net/http"
	"restapi/myapp"
)

func main() {
	http.ListenAndServe(":3000", myapp.NewHandler())
}
