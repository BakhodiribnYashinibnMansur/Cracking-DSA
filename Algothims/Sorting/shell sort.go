package sorting

func ShellSort(arr []int) []int {
	gap := len(arr) / 2
	for gap > 0 {
		for i := gap; i < len(arr); i++ {
			temp := arr[i]
			j := i
			for j >= gap && arr[j-gap] > temp {
				arr[j] = arr[j-gap]
				j -= gap
			}
			arr[j] = temp
		}
		gap = gap / 2
	}
	return arr
}
