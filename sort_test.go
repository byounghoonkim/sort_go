package main

import "testing"
import "fmt"
import "math/rand"

func testSortFunc(f func([]int)[]int, n int) error{

	data := make([]int, n)
	for i:=0; i< len(data);i++{
		data[i] = rand.Intn(n)
	}

	result := f(data)

	for i:= 0 ; i< len(result) -1; i++ {
		if result[i] > result[i+1] {
			return fmt.Errorf("failed to sort : %v", result)
		}
	}

	return nil
}

func TestMergeSort(t *testing.T)  {
	if err := testSortFunc(mergeSort, 100); err != nil {
		t.Error(err)
	}
}

func TestQuickSort(t *testing.T)  {
	if err := testSortFunc(quickSort, 100); err != nil {
		t.Error(err)
	}
}