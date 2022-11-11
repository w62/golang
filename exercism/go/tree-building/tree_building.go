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
			if r.ID > r.Parent {
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

	// If the root is not the only node, then construct the tree

	if records != nil {
		kids := children(records, 0)

		n.Children = kids
		fmt.Printf("kids=%v records=%v\n", n.Children, records)
	}

	root = &n

	/*


		var tree map[int]int

		tree = make(map[int]int)



		for i := 0; i <len(records); i++ {
			fmt.Printf("%d ID=%d Parent=%d\n",i, records[i].ID, records[i].Parent)
			tree [records[i].Parent]++
		}

		fmt.Println("tree map[int]int=", tree)

		for val, _ := range  tree {

		fmt.Printf("Parent %d has %d childs.\n",val, tree[val])
	}

		records = sortrecord(records)
		fmt.Println("After removing root and sort, there are", len(records), " records. Data ", records)


		for len(records) > 0 {

			for i := 0; i < len(records); i++ {
				if records[i].Parent == (*next).ID{
					n := Node{
						ID:       records[i].ID,
						Children: nil,
					}
					fmt.Println("n=",n)

					kids := children(records,(*next).ID)
					(*next).Children = append((*next).Children, &n)
					fmt.Println("number kids=", len(kids), " kids=", kids)

					if kids == nil {

					next = (*next).Children[0]
				}
			fmt.Println("before records=", records)
					records = remove(records, i)
			fmt.Println("after records=", records)

				}
			}
		}*/
	return root, nil
	panic("Please implement the Build function")
}

func children(list []Record, pid int) []*Node {

	var child []*Node
	j := len(list)

	for i := 0; i < len(list); i++ {

		if list[i].Parent == pid {
			n := Node{
				ID:       list[i].ID,
				Children: nil,
			}
			child = append(child, &n)
			list = remove(list, i)
		}
	}
	fmt.Println("child=", child)
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
