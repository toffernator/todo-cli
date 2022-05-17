/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/toffernator/todo-cli/display"
	"github.com/toffernator/todo-cli/task"
)

// postponeCmd represents the postpone command
var postponeCmd = &cobra.Command{
	Use:   "postpone",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		tasks := task.List()
		tasks = task.FilterNotImportant(tasks)
		tasks = task.FilterNotUrgent(tasks)
		prettyPrint := display.Table(tasks)
		fmt.Println(prettyPrint)
	},
}

func init() {
	listCmd.AddCommand(postponeCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// postponeCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// postponeCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
