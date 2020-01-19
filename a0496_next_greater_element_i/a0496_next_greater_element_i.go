package a0496_next_greater_element_i

func nextGreaterElement(nums1 []int, nums2 []int) []int {
	stack := make([]int, len(nums2))
	m := map[int]int{}
	for i, v := range nums2 {
		for len(stack) > 0 && nums2[stack[len(stack)-1]] < v {
			m[nums2[stack[len(stack)-1]]] = v
			stack = stack[0 : len(stack)-1]
		}
		stack = append(stack, i)
	}

	for len(stack) > 0 {
		m[nums2[stack[len(stack)-1]]] = -1
		stack = stack[0 : len(stack)-1]
	}

	res := make([]int, len(nums1))
	for i, v := range nums1 {
		res[i] = m[v]
	}
	return res
}
