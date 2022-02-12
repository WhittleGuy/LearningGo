package main

func Sum(nums []int) (sum int) {
	sum = 0
	for _, number := range nums {
		sum += number
	}
	return
}

// func SumAll(numsToSum ...[]int) (sums []int) {
// 	lengthOfNums := len(numsToSum)
// 	sums = make([]int, lengthOfNums)
// 	for i, numbers := range numsToSum {
// 		sums[i] = Sum(numbers)
// 	}
// 	return
// }
func SumAll(numsToSum ...[]int) (sums []int) {
	for _, numbers := range numsToSum {
		sums = append(sums, Sum(numbers))
	}
	return
}

func SumTails(numsToSum ...[]int) (sums []int) {
	for _, nums := range numsToSum {
		if len(nums) == 0 {
			sums = append(sums, 0)
		} else {
			tail := nums[1:]
			sums = append(sums, Sum(tail))
		}
	}
	return
}
