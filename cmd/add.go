/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"github.com/spf13/cobra"
	"github.com/toffernator/todo-cli/task"
)

// addCmd represents the add command
var addCmd = &cobra.Command{
	Use:   "add",
	Short: "add a new TODO",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		isUrgent, _ := cmd.Flags().GetBool("IsUrgent")
		isImportant, _ := cmd.Flags().GetBool("IsImportant")

		t := task.Task{Name: args[0], IsUrgent: isUrgent, IsImportant: isImportant}
		task.Add(t)
	},
}

func init() {
	rootCmd.AddCommand(addCmd)

	// Here you will define your flags and configuration settings.
	addCmd.PersistentFlags().BoolP("IsUrgent", "u", false, "Marks the tasks 'Urgent' according to the Eisenhower matrix")
	addCmd.PersistentFlags().BoolP("IsImportant", "i", false, "Marks the tasks 'Important' according to the Eisenhower matrix")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// addCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
