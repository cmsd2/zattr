package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "zattr",
	Short: "Manipulate archive files",
	Run:   ZattrCommand,
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func ZattrCommand(cmd *cobra.Command, args []string) {
	// Do Stuff Here
	fmt.Println("not much...")
}
