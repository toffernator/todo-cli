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

// doCmd represents the do command
var doCmd = &cobra.Command{
	Use:   "do",
	Short: "List all tasks that classify as do",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		tasks := task.List()
		tasks = task.FilterImportant(tasks)
		tasks = task.FilterUrgent(tasks)
		prettyPrint := display.Table(tasks)
		fmt.Println(prettyPrint)
	},
}

func init() {
	listCmd.AddCommand(doCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// doCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// doCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
