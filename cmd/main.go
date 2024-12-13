package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var mainCmd = &cobra.Command{
	Use:   "main",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("main called")
	},
}

func init() {
	rootCmd.AddCommand(mainCmd)
}
