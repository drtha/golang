package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"os"
	"github.com/drtha/golang/webcli"
)

var rootCmd = &cobra.Command {
	Use:	"golang",
	Short:	"Different commands to learn go",
	Long: 	"Different commands to learn go",
}

var cmdLs = &cobra.Command{
	Use:   "ls",
	Short: "List articles",
	Long:  `Show the first 5 articles from heise.de`,
	Run: func(cmd *cobra.Command, args []string) {
		webcli.List()
	},
}


var cmdDescribe = &cobra.Command{
	Use:   "describe [id]",
	Short: "Show details for an article",
	Long:  `Show details for an article`,
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		webcli.Describe(args[0])
	},
}


func Exec(){
	rootCmd.AddCommand(cmdLs)
	rootCmd.AddCommand(cmdDescribe)
	err := rootCmd.Execute()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
