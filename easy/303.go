package main

type NumArray struct {
	pre_sum []int
}

func Constructor(nums []int) NumArray {
	preSum := make([]int, len(nums)+1)
	for i, v := range nums {
		preSum[i+1] = preSum[i] + v
	}
	return NumArray{preSum}
}

func (this *NumArray) SumRange(i int, j int) int {
	return this.pre_sum[j+1] - this.pre_sum[i]
}
