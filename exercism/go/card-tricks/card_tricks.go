package cards


// FavoriteCards returns a slice with the cards 2, 6 and 9 in that order.
func FavoriteCards() []int {
	return [] int {2, 6, 9}
	panic("Please implement the FavoriteCards function")
}

// GetItem retrieves an item from a slice at given position.
// If the index is out of range, we want it to return -1.
func GetItem(slice []int, index int) int {
	if len(slice) <=0 {
		return -1 
	}

	if index < 0 {
		return -1 
	}
	if index > len(slice) - 1 {
		return -1 
	} else {
		return slice[index]
	}
	panic("Please implement the GetItem function")
}

// SetItem writes an item to a slice at given position overwriting an existing value.
// If the index is out of range the value needs to be appended.
func SetItem(slice []int, index, value int) []int {
	if index < 0 || index >= len(slice) {
		slice = append(slice, value)
	} else 

	{
		slice[index] = value
	}

	return slice 
	panic("Please implement the SetItem function")
}

// PrependItems adds an arbitrary number of values at the front of a slice.
func PrependItems(slice []int, values ...int) []int {
	var return_value  [] int
	if len(values) <= 0 {
		return_value = slice
	} else {
		for i :=0; i < len(values); i++ {
			return_value = append(return_value, values[i])
		}
			return_value = append(return_value, slice...)
	}
		return return_value
	panic("Please implement the PrependItems function")
}

// RemoveItem removes an item from a slice by modifying the existing slice.
func RemoveItem(slice []int, index int) []int {
	var result  [] int
	if index < 0 || index >= len(slice) {
		result = slice
	} else if index == 0{
		result = slice[1:]
	} else if index >= 1{
		result = append(slice[0:index],slice[index+1:]...)
	} else {
		result = slice[:index]
	}


	return result

	panic("Please implement the RemoveItem function")
}
