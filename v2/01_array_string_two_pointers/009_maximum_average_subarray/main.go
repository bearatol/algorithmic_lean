package main

// Maximum Average Subarray I - LeetCode 643
// Сложность: Easy
// https://leetcode.com/problems/maximum-average-subarray-i/
// Дан массив nums и целое число k.
// Вернуть максимальную среднюю значение подмассива длины k.

func main() {
	// Реализуйте решение здесь
}

// findMaxAverage возвращает максимальную среднюю подмассива длины k
func findMaxAverage(nums []int, k int) float64 {
	window := 0
	for i := 0; i < k; i++ {
		window += nums[i]
	}

	res := window
	for i := k; i < len(nums); i++ {
		window -= nums[i-k]
		window += nums[i]

		if window > res {
			res = window
		}
	}

	return float64(res) / float64(k)
}
