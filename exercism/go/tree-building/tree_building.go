package tree
import (
	"fmt"
)
const rootID = 0
type Record struct {
	ID, Parent int
}
type Node struct {
	ID       int
	Children []*Node
}
// Build converts an unordered slice of Records into a hierarchical tree of Nodes,
// after validating that the tree is not a graph and that the records have
// a contiguous range of IDs.
func Build(records []Record) (*Node, error) {
	if len(records) == 0 {
		return nil, nil
	}
	positions := make([]int, len(records))
	for i, r := range records {
		if r.ID < rootID || r.ID >= len(records) {
			return nil, fmt.Errorf("out of bounds record id %d", r.ID)
		}
		positions[r.ID] = i
		fmt.Printf("i=%d positions[%d]=%d r.ID=%d r.Parent=%d \n", i, r.ID, positions[r.ID], r.ID, r.Parent)
	}
	nodes := make([]Node, len(records))
	fmt.Println("After make []Node ")
	for i := range positions {
		r := records[positions[i]]
		fmt.Printf("i=%d r.ID=%d r.Parent=%d  ",i,  r.ID, r.Parent)
		if r.ID != i {
		fmt.Printf("r.ID != i - r.ID=%d i=%d \n",  r.ID, i)
			return nil, fmt.Errorf("non-contiguous node %d (want %d)", r.ID, i)
		}
		validParentForChild := (r.ID > r.Parent) || (r.ID == rootID && r.Parent == rootID)
		if !validParentForChild {
			return nil, fmt.Errorf("node %d has impossible parent %d", r.ID, r.Parent)
		}
		fmt.Printf("node[i].ID=%d ", nodes[i].ID)
		nodes[i].ID = i
		fmt.Printf(" After nodes[i].ID = i i=%d r.ID=%d r.Parent=%d  \n",i,  r.ID, r.Parent)
		if i != rootID {
			p := &nodes[r.Parent]
			p.Children = append(p.Children, &nodes[i])
		}
	}
	return &nodes[0], nil
}
