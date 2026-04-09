package main

// Find Pivot Index - LeetCode 724
// Сложность: Easy
// https://leetcode.com/problems/find-pivot-index/
// Дан массив nums. Вернуть индекс pivot, где сумма элементов слева
// равна сумме элементов справа. Если такого индекса нет, вернуть -1.

func main() {
	// Реализуйте решение здесь
}

// pivotIndex ищет индекс, где суммы слева и справа равны
func pivotIndex(nums []int) int {
	if len(nums) == 1 {
		return 0
	}

	sumLeft := 0
	sumRight := 0
	for i := 1; i < len(nums); i++ {
		sumRight += nums[i]
	}
	if sumLeft == sumRight {
		return 0
	}

	for i := 1; i < len(nums); i++ {
		sumLeft += nums[i-1]
		sumRight -= nums[i]
		if sumLeft == sumRight {
			return i
		}
	}

	return -1
}
