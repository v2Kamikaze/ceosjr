package list

// O(n*log(n))
func MergeSort(slc []int, f func(int, int) bool) []int {

	if len(slc) < 2 {
		return slc
	}
	mid := (len(slc)) / 2
	return Merge(MergeSort(slc[:mid], f), MergeSort(slc[mid:], f), f)
}

func Merge(left, right []int, f func(int, int) bool) []int {
	size, i, j := len(left)+len(right), 0, 0
	slc := make([]int, size)

	for k := 0; k < size; k++ {
		if i > len(left)-1 && j <= len(right)-1 {
			slc[k] = right[j]
			j++
		} else if j > len(right)-1 && i <= len(left)-1 {
			slc[k] = left[i]
			i++
		} else if f(left[i], right[j]) {
			slc[k] = left[i]
			i++
		} else {
			slc[k] = right[j]
			j++
		}
	}
	return slc
}
