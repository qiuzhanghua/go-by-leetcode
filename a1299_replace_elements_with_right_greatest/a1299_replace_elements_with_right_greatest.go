package a1299_replace_elements_with_right_greatest

func replaceElements(arr []int) []int {
	n := len(arr)
	if n == 0 {
		return arr
	}
	max := arr[n-1]
	for i := n - 2; i >= 0; i-- {
		t := arr[i]
		arr[i] = max
		if max < t {
			max = t
		}
	}
	arr[n-1] = -1
	return arr
}
