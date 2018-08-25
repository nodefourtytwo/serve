package main

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/astroflow/astroflow-go"
	"github.com/astroflow/astroflow-go/log"
	"github.com/spf13/cobra"
)

var (
	servePort          = "8080"
	serveDirectory     = "."
	serveAddress       = "0.0.0.0"
	serveLoggingFormat = "console"
)

var availableLoggingFormats = map[string]astroflow.Formatter{
	"console": astroflow.ConsoleFormatter{},
	"json":    astroflow.JSONFormatter{},
}

func getAllAvailableLoggingFormats() []string {
	ret := make([]string, len(availableLoggingFormats))
	i := 0
	for f := range availableLoggingFormats {
		ret[i] = f
		i += 1
	}
	return ret
}

func init() {
	rootCmd.PersistentFlags().StringVarP(&servePort, "port", "p", "8080", "port to listen to")
	rootCmd.PersistentFlags().StringVarP(&serveAddress, "address", "a", "0.0.0.0", "address to listen to")
	rootCmd.PersistentFlags().StringVarP(&serveDirectory, "directory", "d", ".", "directory to serve")
	rootCmd.PersistentFlags().StringVarP(&serveLoggingFormat, "format", "f", "console", fmt.Sprintf("logging format. Valid values are: [%s]", strings.Join(getAllAvailableLoggingFormats(), ", ")))
}

func getFormatter(format string) (astroflow.Formatter, error) {
	for f, formatter := range availableLoggingFormats {
		if format == f {
			return formatter, nil
		}
	}

	return nil, fmt.Errorf("format: %s is not available", format)
}

var rootCmd = &cobra.Command{
	Use:   "serve",
	Short: "The simplest directory serving web server",
	Long:  "The simplest directory serving web server. See https://github.com/z0mbie42/serve",
	Run: func(cmd *cobra.Command, args []string) {
		var formatter astroflow.Formatter
		var err error

		formatter, err = getFormatter(serveLoggingFormat)
		if err != nil {
			log.Fatal(err.Error())
		}

		address := fmt.Sprintf("%s:%s", serveAddress, servePort)

		log.Config(
			astroflow.SetFormatter(formatter),
		)

		http.Handle("/", http.FileServer(http.Dir(serveDirectory)))

		middleware := astroflow.HTTPHandler(log.With())
		log.Info("Listening on " + address)
		err = http.ListenAndServe(address, middleware(http.DefaultServeMux))
		if err != nil {
			log.Fatal(err.Error())
		}
	},
}
