// Package census simulates a system used to collect census data.
package census
/*
import (
	"fmt"
)
*/
// Resident represents a resident in this city.
type Resident struct {
	Name    string
	Age     int
	Address map[string]string
}

// NewResident registers a new resident in this city.
func NewResident(name string, age int, address map[string]string) *Resident {
	var newResident Resident
	newResident.Name = name
	newResident.Age = age
	newResident.Address = address
	return &newResident 
	panic("Please implement NewResident.")
}

// HasRequiredInfo determines if a given resident has all of the required information.
func (r *Resident) HasRequiredInfo() bool {
	if r.Name == "" {
		return false
	}

	if r.Address == nil {
		return false
	}

	if value, ok := r.Address["street"]; ok {
		if value == ""  {
			return false
		} else {
		return true
	}
	} else {
		return false
	}


	return true
	panic("Please implement HasRequiredInfo.")
}

// Delete deletes a resident's information.
func (r *Resident) Delete() {
	r.Name = ""
	r.Age = 0
	r.Address = nil
	return
	panic("Please implement Delete.")
}

// Count counts all residents that have provided the required information.
func Count(residents []*Resident) int {
	number := 0

	for _, val := range residents {
		if val.Name == "" {
			continue
		}
		if val.Age == 0 {
			continue
		}

		if val.Address == nil {
			continue
		}
		number ++
	}
	return number
	panic("Please implement Count.")
}
