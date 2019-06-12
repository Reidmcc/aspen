package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// addBotCmd represents the addBot command
var addBotCmd = &cobra.Command{
	Use:   "addBot",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("addBot called")
	},
}

func init() {
	botGroupCmd.AddCommand(addBotCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// addBotCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// addBotCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
