package cmd

import (
	"fmt"

	"github.com/shin1x1/yae/version"
	"github.com/spf13/cobra"
)

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Show version",
	Long:  "Show version",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println(version.Text())
	},
}

func init() {
	rootCmd.AddCommand(versionCmd)
}
