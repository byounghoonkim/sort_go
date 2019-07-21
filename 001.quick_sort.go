package main

func quickSort(nums []int) []int {

	if len(nums) <= 1 {
		return nums
	}

	p := 0
	l := 1
	r := len(nums) -1

	for l <= r {
		if nums[l] <= nums[p]{
			l++
			continue
		}

		if nums[r] > nums[p] {
			r--
			continue
		}

		nums[r], nums[l] = nums[l], nums[r]
		l++
		r--
	}

	nums[p], nums[r] = nums[r], nums[p]

	quickSort(nums[:r])
	quickSort(nums[r+1:])

	return nums
}