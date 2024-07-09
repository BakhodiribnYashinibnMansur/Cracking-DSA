package searching

// LinearSearch searches for a value in a slice of integers using linear search algorithm
func LinearSearch(slice []int, value int) int {
	for i, v := range slice {
		if v == value {
			return i
		}
	}
	return -1
}
