package main

import (
	"encoding/json"
	"fmt"
	"runtime"

	"github.com/astroflow/astroflow-go/log"
	"github.com/spf13/cobra"
)

// const set at build time
const (
	UTCBuildTime = "undefined"
	GitCommit    = "undefined"
	OS           = runtime.GOOS
	Arch         = runtime.GOARCH
	GoVersion    = "undefined"
)

const (
	Version = "0.2.0"
)

var versionFormat string

func init() {
	rootCmd.AddCommand(versionCmd)
	versionCmd.PersistentFlags().StringVarP(&versionFormat, "format", "f", "console", "Output format. Valid values are [console, json]")
}

type versionJSON struct {
	Version      string `json:"version"`
	GitCommit    string `json:"git_commit"`
	UTCBuildTime string `json:"utc_build_time"`
	OS           string `json:"os"`
	Architecture string `json:"architecture"`
	GoVersion    string `json:"go_version"`
}

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Display the version and build information",
	Long:  "Display the version and build information",
	Run: func(cmd *cobra.Command, args []string) {
		var err error
		var output []byte

		switch versionFormat {
		case "console":
			fmt.Printf("Version        : %s\n", Version)
			fmt.Printf("Git commit     : %s\n", GitCommit)
			fmt.Printf("UTC build time : %s\n", UTCBuildTime)
			fmt.Printf("OS             : %s\n", OS)
			fmt.Printf("Architecture   : %s\n", Arch)
			fmt.Printf("Go version     : %s\n", GoVersion)
		case "json":
			jsonVersion := versionJSON{
				Version:      Version,
				GitCommit:    GitCommit,
				UTCBuildTime: UTCBuildTime,
				OS:           OS,
				Architecture: Arch,
				GoVersion:    GoVersion,
			}
			output, err = json.Marshal(&jsonVersion)
			if err == nil {
				fmt.Println(string(output))
			}
		default:
			err = fmt.Errorf("Error: %s is not a valid output format", versionFormat)

		}
		if err != nil {
			log.Fatal(err.Error())
		}

	},
}
