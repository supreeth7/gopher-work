/*
Copyright © 2023 Supreeth Basabattini <supreeth.b@outlook.com>
*/

package cmd

import (
	"fmt"
	"strconv"

	"github.com/spf13/cobra"
)

// doneCmd represents the done command
var doneCmd = &cobra.Command{
	Use:   "done",
	Short: "Marks a task as completed",
	Run: func(cmd *cobra.Command, args []string) {
		var taskIds []int

		for _, arg := range args {
			id, err := strconv.Atoi(arg)
			if err != nil {
				panic(err)
			}

			taskIds = append(taskIds, id)
		}

		fmt.Printf("Tasks %d have been marked as done", taskIds)
	},
}

func init() {
	rootCmd.AddCommand(doneCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// doneCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// doneCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
