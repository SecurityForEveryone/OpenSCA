package cmd

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/SecurityForEveryone/software-composition-scanner/scanner"
	"github.com/SecurityForEveryone/software-composition-scanner/scanner/version"
	"github.com/spf13/cobra"
)

var outputFormat string

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "show the version",
	RunE:  printVersion,
}

func init() {
	versionCmd.Flags().StringVarP(&outputFormat, "output", "o", "text", "format to show version information (available=[text, json])")

	rootCmd.AddCommand(versionCmd)
}

func printVersion(_ *cobra.Command, _ []string) error {
	versionInfo := version.FromBuild()
	switch outputFormat {
	case "text":
		fmt.Println("Application:         ", scanner.ApplicationName)
		fmt.Println("Version:             ", versionInfo.Version)
		fmt.Println("BuildDate:           ", versionInfo.BuildDate)
		fmt.Println("GitCommit:           ", versionInfo.GitCommit)
		fmt.Println("GitTreeState:        ", versionInfo.GitTreeState)
		fmt.Println("Platform:            ", versionInfo.Platform)
		fmt.Println("GoVersion:           ", versionInfo.GoVersion)
		fmt.Println("Compiler:            ", versionInfo.Compiler)
	case "json":

		enc := json.NewEncoder(os.Stdout)
		enc.SetEscapeHTML(false)
		enc.SetIndent("", " ")
		err := enc.Encode(&struct {
			version.Version
			Application   string `json:"application"`
			SchemaVersion int    `json:"supportedDbSchema"`
		}{
			Version:     versionInfo,
			Application: scanner.ApplicationName,
		})
		if err != nil {
			return fmt.Errorf("failed to show version information: %+v", err)
		}
	default:
		return fmt.Errorf("unsupported output format: %s", outputFormat)
	}
	return nil
}
