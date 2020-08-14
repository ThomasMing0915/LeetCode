package main

func intersection(nums1 []int, nums2 []int) []int {
	m1 := make(map[int]struct{}, len(nums1))
	m2 := make(map[int]struct{}, len(nums2))

	for _, n1 := range nums1 {
		m1[n1] = struct{}{}
	}

	for _, n2 := range nums2 {
		m2[n2] = struct{}{}
	}

	res := make([]int, 0, min(len(nums1), len(nums2)))

	for k := range m1 {
		if _, ok := m2[k]; ok {
			res = append(res, k)
		}
	}

	return res
}

func min(a, b int) int {
	if a <= b {
		return a
	}
	return b
}
