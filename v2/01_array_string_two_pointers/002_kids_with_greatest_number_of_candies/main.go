package main

// Kids With the Greatest Number of Candies - LeetCode 1431
// Сложность: Easy
// https://leetcode.com/problems/kids-with-the-greatest-number-of-candies/
// Даны целочисленный массив candies (количество конфет у каждого ребенка)
// и целое число extraCandies.
// Вернуть массив булевых значений, где true означает, что ребенок
// может иметь наибольшее количество конфет после добавления extraCandies.

func main() {
	// Реализуйте решение здесь
}

// kidsWithCandies возвращает массив, показывающий может ли каждый ребенок
// иметь максимальное количество конфет
func kidsWithCandies(candies []int, extraCandies int) []bool {
	max := 0
	for i := 0; i < len(candies); i++ {
		if candies[i] > max {
			max = candies[i]
		}
	}

	result := make([]bool, len(candies))
	for i := 0; i < len(candies); i++ {
		if (candies[i] + extraCandies) >= max {
			result[i] = true
		}
	}

	return result
}
