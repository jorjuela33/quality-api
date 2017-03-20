package controller

import (
	"fmt"
	"net/http"
)

func Index(responseWriter http.ResponseWriter, request *http.Request) {
	fmt.Printf("Hello World")
}
