package searching

func BinarySearch(slice []int, value int) int {
	low := 0
	high := len(slice) - 1
	for low <= high {
		mid := (low + high) / 2
		if slice[mid] == value {
			return mid
		} else if slice[mid] < value {
			low = mid + 1
		} else {
			high = mid - 1
		}
	}
	return -1
}
