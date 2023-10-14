package main

import (
	"fmt"

	"github.com/jvzantvoort/tmpl/messages"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

// HeaderBashCmd represents the header command
var HeaderBashCmd = &cobra.Command{
	Use:   "header",
	Short: "Create a new project",
	Long:  messages.GetLong("bash/header"),
	Run:   handleHeaderBashCmd,
}

func handleHeaderBashCmd(cmd *cobra.Command, args []string) {

	if verbose {
		log.SetLevel(log.DebugLevel)
	}

	log.Debugf("%s: start", cmd.Use)
	defer log.Debugf("%s: end", cmd.Use)

	tmpl := NewTemplate("bash")
	// name, _ := cmd.Flags().GetString("name")
	fmt.Printf("%s\n", tmpl.Template("header"))
}

func init() {
	// Header
	bashCmd.AddCommand(HeaderBashCmd)
	HeaderBashCmd.Flags().StringP("name", "n", "", "project name")
	// HeaderBashCmd.MarkFlagRequired("name")
}
