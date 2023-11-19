package main

import (
	"net/http"

	"github.com/gobuffalo/logger"
)

func main() {

	log := logger.NewLogger("info")

	log.Info("Server Starting up...")

	serve(setupServer(log), log)
}

func serve(svr *http.Server, log logger.Logger) {
	// Start server
	err := svr.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}
}
