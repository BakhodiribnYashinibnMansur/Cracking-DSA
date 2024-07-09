package searching

func interpolationSearch(slice []int, value int) int {
	low := 0
	high := len(slice) - 1
	for low <= high && value >= slice[low] && value <= slice[high] {
		mid := low + ((high-low)*(value-slice[low]))/(slice[high]-slice[low])
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
