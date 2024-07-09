package sorting

func merge(a []int, b []int) []int {
	r := make([]int, 0)
	i := 0
	j := 0
	for i < len(a) && j < len(b) {
		if a[i] <= b[j] {
			r = append(r, a[i])
			i++
		} else {
			r = append(r, b[j])
			j++
		}
	}
	for i < len(a) {
		r = append(r, a[i])
		i++
	}
	for j < len(b) {
		r = append(r, b[j])
		j++
	}
	return r
}

func MergeSort(arr []int) []int {
	if len(arr) < 2 {
		return arr
	}
	var middle = len(arr) / 2
	var a = MergeSort(arr[:middle])
	var b = MergeSort(arr[middle:])
	return merge(a, b)
}
