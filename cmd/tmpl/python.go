package main

import (
	"fmt"

	msg "github.com/jvzantvoort/tmpl/messages"
	tmpl "github.com/jvzantvoort/tmpl/templates"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

// pythonCmd represents the python command
var pythonCmd = &cobra.Command{
	Use:   "python",
	Short: msg.GetUsage("python"),
	Long:  msg.GetLong("python"),
	Run:   handlePythonCmd,
}

func handlePythonCmd(cmd *cobra.Command, args []string) {

	if verbose {
		log.SetLevel(log.DebugLevel)
	}

	log.Debugf("%s: start", cmd.Use)
	defer log.Debugf("%s: end", cmd.Use)

	list, _ := cmd.Flags().GetBool("list")

	if len(args) == 0 {
		list = true
	}

	if list {
		files, err := tmpl.ListTemplates("python")
		if err != nil {
			log.Errorf("listing error: %s", err)
		}
		for _, i := range files {
			fmt.Printf("- %s\n", i)
		}
		return
	}
	tmpl := NewTemplate("python")
	for _, arg := range args {
		fmt.Printf("%s\n", tmpl.Template(arg))
	}
}

func init() {
	rootCmd.AddCommand(pythonCmd)
	pythonCmd.Flags().BoolP("list", "l", false, "List files")
}
