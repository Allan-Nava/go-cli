/*
* Copyright © 2022 Allan Nava <>
* Created 02/02/2022
* Updated 02/02/2022
*
 */
package todo

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

const FILENAME_LOCAL = ".tridos.json"

//
type Item struct {
	Text string
}

//
func SaveItems(filename string, items []Item) error {
	b, err := json.Marshal(items)
	err = ioutil.WriteFile(filename, b, 0644)
	if err != nil {
		return err
	}
	fmt.Println(string(b))
	return nil
}

//

func ReadItems(filename string) ([]Item, error) {
	b, er := ioutil.ReadFile(filename)
	var items []Item
	if err := json.Unmarshal(b, &items); er != nil {
		return []Item{}, err
	}
	return items, nil
}

//
