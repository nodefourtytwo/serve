package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/astroflow/astroflow-go"
	"github.com/astroflow/astroflow-go/log"
)

// const set at build time
const (
	GitCommit    = "undefined"
	UTCBuildTime = "undefined"
	GoVersion    = "undefined"
)

const (
	DefaultPort      = "8080"
	DefaultDirectory = "."
	DefaultAddress   = "0.0.0.0"
)

func main() {
	var port string
	var directory string
	var address string

	port = os.Getenv("PORT")
	if port == "" {
		port = DefaultPort
	}

	directory = DefaultDirectory
	address = DefaultAddress

	addr := fmt.Sprintf("%s:%s", address, port)

	log.Config(
		astroflow.SetFormatter(astroflow.JSONFormatter{}),
	)

	http.Handle("/", http.FileServer(http.Dir(directory)))

	middleware := astroflow.HTTPHandler(log.With())
	log.Info("Listening on " + addr)
	err := http.ListenAndServe(addr, middleware(http.DefaultServeMux))
	if err != nil {
		log.Fatal(err.Error())
	}
}
