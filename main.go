package main

import (
	"net/http"

	"github.com/waghvikrant/go-crud-rest/server"

	logger "github.com/sirupsen/logrus"
)

func main() {
	logger.Fatal(http.ListenAndServe(":8080", server.NewRouter()))
}
