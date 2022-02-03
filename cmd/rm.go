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

	"github.com/spf13/cobra"

	"github.com/Allan-Nava/go-cli/todo"
)

// rmCmd represents the rm command
var rmCmd = &cobra.Command{
	Use:   "rm",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: rmRun,
}

func init() {
	rootCmd.AddCommand(rmCmd)
	// Here you will define your flags and configuration settings.
	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// rmCmd.PersistentFlags().String("foo", "", "A help for foo")
	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// rmCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	//
	rmCmd.Flags().IntVarP(&position, "position ", "i", 0, "Position: 1,2,3,4..etc")
	//

}

//
func rmRun(cmd *cobra.Command, args []string) {
	fmt.Println("rm called")
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
	//del item
	fmt.Printf(" item: %v\n", item)
	//
}

//
