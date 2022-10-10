package array

func Sum(nums []int) int {
	count := 0
	for _, num := range nums {
		count += num
	}
	return count
}

func SumAll(num ...[]int) []int {
	length := len(num)
	sums := make([]int, length)

	for i, nums := range num {
		sums[i] = Sum(nums)
	}
	return sums
}
