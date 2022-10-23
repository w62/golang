package gross

import (
	"fmt"
)

// Units stores the Gross Store unit measurements.
func Units() map[string]int {
	return map[string]int {"quarter_of_a_dozen":3,
	"half_of_a_dozen":6,
	"dozen":12,
	"small_gross":120,
	"gross":144,
	"great_gross":1728,
}
	panic("Please implement the Units() function")
}

// NewBill creates a new bill.
func NewBill() map[string]int {
	return map[string]int {}
	panic("Please implement the NewBill() function")
}

// AddItem adds an item to customer bill.
func AddItem(bill, units map[string]int, item, unit string) bool {
	fmt.Println("item=", item, " unit=", unit)
	if _, ok := units [unit]; ok {

		if val, found := bill[item]; found {

		bill [item] = units [unit] + val
	} else {

		bill [item] = units [unit]
	}
		return true 
	} else {
	fmt.Println("not found item=", item, " unit=", unit)
		return false
	}
	panic("Please implement the AddItem() function")
}

// RemoveItem removes an item from customer bill.
func RemoveItem(bill, units map[string]int, item, unit string) bool {

	value := 0

	if val, ok := units[unit]; ok {
		value = val
	} else {
		fmt.Println("invalid unit", unit)
		return false
	} 

	if _, ok := bill[item]; ok {
		diff := bill[item] - value

		if diff < 0 {
			return false
		} else if diff == 0 {
			delete (bill, item)
			return true
		} else {
			bill[item] = diff
			return true
		}
		
	} else {
		fmt.Println(item, "is not in the bill")
		return false
	}

	

	panic("Please implement the RemoveItem() function")
}

// GetItem returns the quantity of an item that the customer has in his/her bill.
func GetItem(bill map[string]int, item string) (int, bool) {
	if val, ok := bill[item]; ok {
		return val, true
	} else {
		return 0, false 
	}
	panic("Please implement the GetItem() function")
}
