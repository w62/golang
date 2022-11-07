package tree

import (
	"fmt"
	"errors"
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
    return append(slice[:s], slice[s+1:]...)
}

func duplicate (slice []Record) bool {
	if len(slice) == 1 {
		return false
	}
	for i := 0; i < len(slice); i ++ {
		for j := i+1; j < len(slice) ; j ++ {
			if slice[i].ID == slice[j].ID &&
			 slice[i].Parent == slice[j].Parent {
				 fmt.Println("i=", i, "j=",j)
				 fmt.Println("slice[i]", slice[i], "slice[j]", slice[j])
				return true
			}
		}
	}
	return false
}
func Build(records []Record) (*Node, error) {

	var rootNode Node
	var rootRecord int

	if len(records) <=0 {
		return nil, nil
	}

	fmt.Println("There are", len(records), " records.")
	fmt.Println("The entire data set is", records)
	if duplicate(records) {
		return nil, errors.New("Duplicate records")
	}


	numberOfRoots := 0

	for i,r := range records {
		if r.ID == r.Parent {
			rootRecord = i
			fmt.Println("in loop r=",r, " RootRecord", rootRecord)
			numberOfRoots ++
		}
	}

	if numberOfRoots != 1 {
		return nil, errors.New("not one and only one root")
	}
			fmt.Println("RootRecord", rootRecord)
	fmt.Println("root node is ", rootRecord, "value", records[rootRecord])


	rootNode.ID = records[rootRecord].ID
	records = remove(records, rootRecord)

	if len(records) > 0 {

	for i := 0; i < len(records) ; i ++ {

	}
} else {
}

	return &rootNode, nil
	panic("Please implement the Build function")
}
