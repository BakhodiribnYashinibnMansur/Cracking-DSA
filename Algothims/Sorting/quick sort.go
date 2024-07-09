package sorting

func QuickSort(arr []int) []int {
	if len(arr) < 2 {
		return arr
	}
	pivot := arr[0]
	lowPart := make([]int, 0)
	highPart := make([]int, 0)
	for i := 1; i < len(arr); i++ {
		if arr[i] < pivot {
			lowPart = append(lowPart, arr[i])
		} else {
			highPart = append(highPart, arr[i])
		}
	}
	lowPart = QuickSort(lowPart)
	highPart = QuickSort(highPart)
	return append(append(lowPart, pivot), highPart...)
}
