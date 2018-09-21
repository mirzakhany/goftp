package cmd

import (
	"fmt"

	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "goftp",
	Short: "goFtp is a fast in place ftp server",
	Long: `A fast Ftp server to share a directory,
with a memory from python -m SimpleHTTPServer
But this is a multi thread and very fast
Complete documentation is available at http://github.com/mirzakhany/goftp`,

	Run: func(cmd *cobra.Command, args []string) {
		cmd.Help()
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
