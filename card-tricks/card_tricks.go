package cards

// GetItem retrieves an item from a slice at given position. The second return value indicates whether
// the given index exists in the slice or not.
func GetItem(slice []int, index int) (int, bool) {
	if index >= len(slice) || index < 0 {
		return 0, false
	} else {
		return slice[index], true
	}
}

// SetItem writes an item to a slice at given position overwriting an existing value.
// If the index is out of range the value needs to be appended.
func SetItem(slice []int, index, value int) []int {
	if index >= len(slice) || index < 0 {
		slice = append(slice, value)
		return slice
	} else {
		slice[index] = value
		return slice
	}
}

// PrefilledSlice creates a slice of given length and prefills it with the given value.
func PrefilledSlice(value, length int) []int {
	if length <= 0 {
		return []int{}
	}
	return fillOfValue(make([]int, length), value, 0)
}

func fillOfValue(slice []int, value, idx int) []int {
	if idx == len(slice) {
		return slice
	}
	slice[idx] = value
	idx++
	return fillOfValue(slice, value, idx)
}

// RemoveItem removes an item from a slice by modifying the existing slice.
func RemoveItem(slice []int, index int) []int {
	if index >= len(slice) || index < 0 {
		return slice
	} else {
		return fillOfSlice(slice[:index], slice[index+1:])
	}
}

func fillOfSlice(slice1, slice2 []int) []int {
	if len(slice2) == 0 {
		return slice1
	}
	return fillOfSlice(append(slice1, slice2[0]), slice2[1:])
}
