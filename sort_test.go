package main

import "testing"
import "fmt"

func copyData(data[]int) []int {
	ret := make([]int, len(data))
	copy(ret, data)
	return ret
}

func testSortFunc(f func([]int)[]int) error{
	data := []int{4, 3, 2, 1}
	result := f(copyData(data))

	for i:= 0 ; i< len(result) -1; i++ {
		if result[i] > result[i+1] {
			return fmt.Errorf("failed to sort : %v", result)
		}
	}

	return nil
}

func TestMergeSort(t *testing.T)  {
	if err := testSortFunc(mergeSort); err != nil {
		t.Error(err)
	}
}

func TestQuickSort(t *testing.T)  {
	if err := testSortFunc(quickSort); err != nil {
		t.Error(err)
	}
}