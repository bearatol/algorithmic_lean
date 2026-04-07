package main

// Can Place Flowers - LeetCode 605
// Сложность: Easy
// https://leetcode.com/problems/can-place-flowers/
// Есть цветочная клумба с грядками. 1 = занято, 0 = пусто.
// Можно посадить цветок, если нет соседних занятых грядок.
// Дано целое число n - количество цветов для посадки.
// Вернуть true, если можно посадить n цветов.

func main() {
	// Реализуйте решение здесь
}

// canPlaceFlowers проверяет можно ли посадить n цветов
func canPlaceFlowers(flowerbed []int, n int) bool {
	for i := 0; i < len(flowerbed); i++ {
		if flowerbed[i] == 1 {
			i += 1
			continue
		}
		if (i == 0 || flowerbed[i-1] == 0) &&
			(i == len(flowerbed)-1 || flowerbed[i+1] == 0) {
			n--
			i += 1
		}
	}
	return n < 1
}
