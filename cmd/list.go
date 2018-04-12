package cmd

import (
	"errors"

	"github.com/cmsd2/zattr/pkg/zattr"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(listCmd)
}

var listCmd = &cobra.Command{
	Use:   "list [archive file]",
	Short: "List the contents of an archive",
	Run:   ListArchiveFiles,
	Args: func(cmd *cobra.Command, args []string) error {
		if len(args) < 1 {
			return errors.New("no archive file specified")
		}

		return nil
	},
}

// ListArchiveFiles lists the contents of an archive
func ListArchiveFiles(cmd *cobra.Command, args []string) {
	file := zattr.OpenZip(args[0])
	defer file.Close()

	zattr.TransformZip(file, zattr.PrintFile)
}
