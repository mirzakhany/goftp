package commands

import (
	"fmt"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(versionCmd)
}

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print the version number of gofmt",
	Long:  `All software has versions. This is goFtp's`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Fast in place Ftp server v0.1")
	},
}
