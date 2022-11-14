package tree

import (
	"errors"
	"fmt"
	"sort"
)

type Record struct {
	ID     int
	Parent int
	// feel free to add fields as you see fit
}

type Node struct {
	ID       int
	Children []*Node
	// feel free to add fields as you see fit
}

func remove(slice []Record, s int) []Record {
	if len(slice) <= 1 {
		return nil
	}
	return append(slice[:s], slice[s+1:]...)
}

func duplicate(slice []Record) bool {
	if len(slice) == 1 {
		return false
	}
	for i := 0; i < len(slice); i++ {
		for j := i + 1; j < len(slice); j++ {
			if slice[i].ID == slice[j].ID &&
				slice[i].Parent == slice[j].Parent {
				fmt.Println("i=", i, "j=", j)
				fmt.Println("slice[i]", slice[i], "slice[j]", slice[j])
				return true
			}
		}
	}
	return false
}
func Build(records []Record) (*Node, error) {

	var root *Node
	var next *Node
	var rootRecord int

	//check for 0 or zero records
	if len(records) <= 0 {
		return nil, nil
	}

	// Print out the initial data for verification
	fmt.Println("There are", len(records), " records.")
	fmt.Println("The entire data set is", records)
	if duplicate(records) {
		return nil, errors.New("Duplicate records")
	}

	// number of root checking
	// There should be one and only one root, no more no less

	numberOfRoots := 0

	for i, r := range records {
		fmt.Printf("r.ID=%d r.Parent=%d\n", r.ID, r.Parent)
		if r.ID == r.Parent {
			rootRecord = i
			fmt.Println("in loop r=", r, " RootRecord", rootRecord)
			numberOfRoots++
		}

		if r.ID != 0 && r.Parent != 0 {
			if r.Parent > r.ID {
				return nil, errors.New("ID larger than Parent")
			}
		}
	}

	if numberOfRoots != 1 {
		return nil, errors.New("not one and only one root")
	}
	// end of number of root checking

	fmt.Println("RootRecord", rootRecord)
	fmt.Println("root node is ", rootRecord, "value", records[rootRecord])

	n := Node{
		ID:       records[rootRecord].ID,
		Children: nil,
	}

	records = remove(records, rootRecord)

	records = sortrecord(records)
	// If the root is not the only node, then construct the tree

	if records != nil {
		kids := children(records, 0)

		n.Children = kids

		for len(records) > 0 {

			for i, val := range kids {
				next = val
				fmt.Printf("records=%v\n", records)
				fmt.Printf("kids=%v records=%v\n", n.Children, records)
				next.Children = children(records, next.ID)
				fmt.Printf("after records=%v\n", records)
				fmt.Printf("Build i=%d next=%v next.ID=%v next.Children=%v \n", i, next, next.ID, next.Children)
			}
		}

	}

	root = &n

	return root, nil
	panic("Please implement the Build function")
}

func children(list []Record, pid int) []*Node {

	var child []*Node
	var new_list []Record

	for i, _ := range list {

		if list[i].Parent == pid {
			n := Node{
				ID:       list[i].ID,
				Children: nil,
			}
			child = append(child, &n)
		} else {
			new_list = append(new_list, list[i])
		}

		fmt.Println("inside children i=", i)
	}
	list = new_list
	fmt.Println("child=", child)
	fmt.Println("new_list=", new_list)
	return child
}

func sortrecord(list []Record) []Record {
	sort.Slice(list, func(i, j int) bool {
		if list[i].Parent == list[j].Parent {
			return list[i].ID < list[j].ID
		}
		return list[i].Parent < list[j].Parent
	})

	return list
}
