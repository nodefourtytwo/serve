package main

import (
	"github.com/astroflow/astroflow-go/log"
)

func main() {
	if err := rootCmd.Execute(); err != nil {
		log.Fatal(err.Error())
	}
}
