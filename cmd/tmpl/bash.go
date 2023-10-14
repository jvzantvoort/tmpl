package main

import (
	"github.com/jvzantvoort/tmpl/messages"
	"github.com/spf13/cobra"
)

// bashCmd represents the bash command
var bashCmd = &cobra.Command{
	Use:   "bash",
	Short: "templates for bash",
	Long:  messages.GetLong("bash/root"),
}

func init() {
	rootCmd.AddCommand(bashCmd)

}
