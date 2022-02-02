/*
* Copyright © 2022 Allan Nava <>
* Created 02/02/2022
* Updated 02/02/2022
*
 */
package cmd

import (
	"fmt"

	"github.com/Allan-Nava/go-cli/todo"
	"github.com/spf13/cobra"
)

//
// addCmd represents the add command
var addCmd = &cobra.Command{
	Use:   "add",
	Short: "A new todo",
	Long:  `Add will create a new todo item to the list`,
	Run:   addRun,
}

//
func init() {
	rootCmd.AddCommand(addCmd)
	// Here you will define your flags and configuration settings.
	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// addCmd.PersistentFlags().String("foo", "", "A help for foo")
	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// addCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

//
func addRun(cmd *cobra.Command, args []string) {
	fmt.Println("add called ")
	//items := []todo.Item{}
	items, errorRead := todo.ReadItems(todo.FILENAME_LOCAL)
	if errorRead != nil {
		fmt.Errorf("%v", errorRead)
	}
	for _, x := range args {
		fmt.Println(x)
		items = append(items, todo.Item{Text: x})
	}
	//
	err := todo.SaveItems(todo.FILENAME_LOCAL, items)
	if err != nil {
		fmt.Errorf("%v", err)
	}
	//
	fmt.Printf("%#v \n ", items)
}

//
