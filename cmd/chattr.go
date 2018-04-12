package cmd

import (
	"errors"

	"github.com/cmsd2/zattr/pkg/zattr"
	"github.com/gobwas/glob"
	"github.com/spf13/cobra"
)

var Dest string
var Pattern string
var Mode string

func init() {
	chattrCmd.Flags().StringVarP(&Pattern, "pattern", "p", "*", "Shell glob pattern to match against file paths in the zip archive")
	chattrCmd.Flags().StringVarP(&Dest, "output", "o", "", "Path to destination zip file")
	chattrCmd.Flags().StringVarP(&Mode, "mode", "m", "", "Mode to set on matching files")
	rootCmd.AddCommand(chattrCmd)
}

var chattrCmd = &cobra.Command{
	Use:   "chattr [archive file]",
	Short: "Change the attributes of files in an archive",
	Run:   ChattrArchiveFiles,
	Args: func(cmd *cobra.Command, args []string) error {
		if len(args) < 1 {
			return errors.New("no archive file specified")
		}

		if Mode == "" {
			return errors.New("no mode specified")
		}

		if Dest == "" {
			return errors.New("no destination specified")
		}

		return nil
	},
}

// ListArchiveFiles lists the contents of an archive
func ChattrArchiveFiles(cmd *cobra.Command, args []string) {
	compiledPattern := glob.MustCompile(Pattern)
	filter := zattr.AttrChanger(compiledPattern, Mode)

	zattr.CopyZipPath(args[0], Dest, filter)
}
