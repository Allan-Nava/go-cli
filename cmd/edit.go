/*
* Copyright © 2022 Allan Nava <>
* Created 02/02/2022
* Updated 03/02/2022
*
 */
package cmd

import (
	"fmt"
	"log"

	"github.com/Allan-Nava/go-cli/todo"
	"github.com/spf13/cobra"
)

var (
	position int
	//todoName string
)

// editCmd represents the edit command
var editCmd = &cobra.Command{
	Use:   "edit",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: editRun,
}

func init() {
	rootCmd.AddCommand(editCmd)
	// Here you will define your flags and configuration settings.
	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// editCmd.PersistentFlags().String("foo", "", "A help for foo")
	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// editCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	//
	editCmd.Flags().IntVarP(&position, "position ", "i", 0, "Position: 1,2,3,4..etc")
	//editCmd.Flags().StringP("", "n", "", "Edit the name")
}

//
func editRun(cmd *cobra.Command, args []string) {
	fmt.Println("add edit function")
	//
	items, errorRead := todo.ReadItems(dataFile)
	if errorRead != nil {
		fmt.Errorf("editRun error: %v", errorRead)
	}
	//fmt.Printf("position:%v \n items: %v\n", position, items)
	if len(items) < position {
		log.Print("position item doesn't exists")
	}
	//
	item := &items[position]
	fmt.Printf(" item: %v\n", item)
	//
	if len(args) > 0 {
		fmt.Printf("args %v", args)
		name := args[0]
		item.Text = name
	}

	fmt.Printf(" items: %v\n", items)
	err := todo.SaveItems(dataFile, items)
	if err != nil {
		fmt.Errorf("%v", err)
	}
	fmt.Printf("%#v \n ", items)
}

//
