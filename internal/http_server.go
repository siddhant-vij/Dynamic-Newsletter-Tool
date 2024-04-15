package internal

import (
	"net/http"
)

func StartHttpServer() error {
	return http.ListenAndServe(":8080", http.FileServer(http.Dir("public")))
}
