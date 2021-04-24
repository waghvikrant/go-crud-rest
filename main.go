package main

import (
	"net/http"
	"sample-backend/server"

	logger "github.com/sirupsen/logrus"
)

func main() {
	logger.Fatal(http.ListenAndServe(":8080", server.NewRouter()))
}
