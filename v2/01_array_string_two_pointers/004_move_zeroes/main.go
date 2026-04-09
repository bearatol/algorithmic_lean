package main

// Move Zeroes - LeetCode 283
// Сложность: Easy
// https://leetcode.com/problems/move-zeroes/
// Дан массив nums. Переместить все нули в конец массива,
// сохраняя относительный порядок ненулевых элементов.

func main() {
	// Реализуйте решение здесь
}

// moveZeroes перемещает все нули в конец массива
func moveZeroes(nums []int) {
	idx := 0
	for i := 0; i < len(nums); i++ {
		if nums[i] == 0 {
			continue
		}
		if nums[i] != 0 {
			nums[idx], nums[i] = nums[i], nums[idx]
			idx++
		}
	}
}
