package cmd

import (
	"log"
	"os"

	"github.com/shin1x1/yae/handlers"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "yae [YAML file]",
	Short: "yae is a YAML aliases expander",
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
