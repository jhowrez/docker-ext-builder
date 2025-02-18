/*
Copyright © 2025 Jhonatas Rezende <jhonatas.rezende@alltis.com.br>
*/
package cmd

import (
	"os"

	"github.com/spf13/cobra"

	"github.com/jhowrez/docker-ext-builder/cmd/exporter"
)

var exportCfg = exporter.ExportOptions{}

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "docker-ext-builder",
	Short: "Build and export content using dockerfiles",
	Long: `Build and export content using dockerfile to a tar file locally. 
	Example: 
		docker-ext-builder -f ./examples/Dockerfile.test -o ./test.tar -c `,
	Run: func(cmd *cobra.Command, args []string) {
		exporter.RunExporter(exportCfg)
	},
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	rootCmd.Flags().StringVarP(&exportCfg.Dockerfile, "file", "f", "Dockerfile", "dockerfile to be used for build")
	rootCmd.Flags().StringVarP(&exportCfg.OutputFilename, "out", "o", "out.tar", "output container content .tar filename")
	rootCmd.Flags().StringVarP(&exportCfg.ExportPath, "path", "p", "/opt/out", "container output path to be exported")
}
