package main

import "testing"

func TestSearchInsert(t *testing.T){
	test_data := []struct{target, index int}{
		{5,2},
		{2,1},
		{7,4},
		{0,0},
	}
	test_arr := []int{1,3,5,6}
	for _, test := range test_data{
		if actual := searchInsert(test_arr,test.target) ; actual!= test.index{
			t.Errorf("Expect to %d, but got %d",test.index, actual)
		}
	}
}
