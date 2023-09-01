package list

// O(n*log(n))
func ThirdGreatest(slc []int) int {
	length := len(slc)

	if length == 1 {
		return slc[0]
	}

	ordSlc := MergeSort(slc, func(a, b int) bool { return a > b })

	if length == 2 {
		return slc[1]
	}

	return ordSlc[2]
}

// O(n*log(n))
func ListPartition(l1, l2 []int) (bool, []int, []int) {

	len1, len2 := len(l1), len(l2)
	listFull := make([]int, 0)

	if len1 != len2 || len1 == 0 || len2 == 0 {
		return false, nil, nil
	}

	// O(n)
	listFull = append(listFull, l1...)
	// O(n)
	listFull = append(listFull, l2...)
	// O(n*log(n))
	listFull = MergeSort(listFull, func(a, b int) bool { return a < b })

	firstHalf := listFull[:len1]
	secHalf := listFull[len1 : len1+len2]

	return firstHalf[len1-1] < secHalf[0], firstHalf, secHalf
}

// O(n)
func Permutations(l1, l2 []int) bool {
	values := make(map[int]int)

	// O(n)
	for _, v := range l2 {
		values[v]++
	}

	// O(n)
	for _, v := range l1 {
		if values[v] == 0 {
			return false
		}
		values[v]--
	}

	// O(n)
	for _, v := range values {
		if v != 0 {
			return false
		}
	}

	return true
}
