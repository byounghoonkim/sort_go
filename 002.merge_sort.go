package main

func mergeSort(nums []int) []int {
	if len(nums) == 1 {
		return nums
	}

	left := mergeSort(nums[:len(nums)/2])
	right := mergeSort(nums[len(nums)/2:])
	ret := []int{}
	i, j := 0, 0
	for i < len(left) || j < len(right) {

		if i >= len(left) {
			ret = append(ret, right[j])
			j++
		} else if j >= len(right) {
			ret = append(ret, left[i])
			i++
		} else {
			if left[i] > right[j] {
				ret = append(ret, right[j])
				j++
			} else {
				ret = append(ret, left[i])
				i++
			}
		}
	}

	return ret
}
