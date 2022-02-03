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
	"strconv"
)

const FILENAME_LOCAL = ".tridos.json"

//
type Item struct {
	Text     string
	Priority int
	position int
	Done     bool
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
	for i, _ := range items {
		items[i].position = i + 1
	}
	return items, nil
}

//

func (i *Item) SetPriority(pri int) {
	switch pri {
	case 1:
		i.Priority = 1
	case 3:
		i.Priority = 3
	default:
		i.Priority = 2
	}
}

//
func (i *Item) PrettyP() string {
	if i.Priority == 1 {
		return "(1)"
	}
	if i.Priority == 3 {
		return "(3)"
	}
	return " "
}

func (i *Item) PrettyDone() string {
	if i.Done {
		return "X"
	} else {
		return ""
	}
}

func (i *Item) Label() string {
	return strconv.Itoa(i.position) + "."
}

//ByPri implements sort.Interface for []Item based on the priority & position field.
type ByPri []Item

//
func (s ByPri) Len() int      { return len(s) }
func (s ByPri) Swap(i, j int) { s[i], s[j] = s[j], s[i] }
func (s ByPri) Less(i, j int) bool {
	if s[i].Done != s[j].Done {
		return s[i].Done
	}
	if s[i].Priority == s[j].Priority {
		return s[i].position < s[j].position
	} else {
		return s[i].Priority < s[j].Priority
	}
}
