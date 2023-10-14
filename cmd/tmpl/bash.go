package main

import (
	"fmt"

	msg "github.com/jvzantvoort/tmpl/messages"
	tmpl "github.com/jvzantvoort/tmpl/templates"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

// bashCmd represents the bash command
var bashCmd = &cobra.Command{
	Use:   "bash",
	Short: "templates for bash",
	Long:  msg.GetLong("bash/root"),
	Run:   handleBashCmd,
}

func handleBashCmd(cmd *cobra.Command, args []string) {

	if verbose {
		log.SetLevel(log.DebugLevel)
	}

	log.Debugf("%s: start", cmd.Use)
	defer log.Debugf("%s: end", cmd.Use)

	list, _ := cmd.Flags().GetBool("list")

	if list {
		/// tmpl := NewTemplate("bash")
		files, err := tmpl.ListTemplates("bash")
		if err != nil {
			log.Errorf("listing error: %s", err)
		}
		for _, i := range files {
			fmt.Printf("- %s\n", i)
		}
		return
	}
	tmpl := NewTemplate("bash")
	for _, arg := range args {
		fmt.Printf("%s\n", tmpl.Template(arg))
	}
}

func init() {
	rootCmd.AddCommand(bashCmd)
	bashCmd.Flags().BoolP("list", "l", false, "List files")
}
