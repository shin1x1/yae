package cmd

import (
	"log"
	"os"

	"github.com/shin1x1/yamlfmt/handlers"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "yamlfmt [YAML file]",
	Short: "yamlfmt is a YAML formatter",
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		filepath := args[0]

		handlers.NewYamlFormatter(filepath).Run(os.Stdout)
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		log.Fatal(err)
	}
}
