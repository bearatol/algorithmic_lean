package main

// Maximum Number of Vowels in a Substring of Given Length - LeetCode 1456
// Сложность: Medium
// https://leetcode.com/problems/maximum-number-of-vowels-in-a-substring-of-given-length/
// Дана строка s и целое число k.
// Вернуть максимальное количество гласных в любой подстроке длины k.

func main() {
	// Реализуйте решение здесь
}

// maxVowels возвращает максимальное количество гласных в подстроке длины k
func maxVowels(s string, k int) int {
	maxWindow := 0
	for i := range k {
		if checkVowels(s[i]) {
			maxWindow++
		}
	}

	resSum := maxWindow
	for i := k; i < len(s); i++ {
		if checkVowels(s[i]) {
			maxWindow++
		}
		if checkVowels(s[i-k]) {
			maxWindow--
		}
		if maxWindow > resSum {
			resSum = maxWindow
		}
	}

	return resSum
}

func checkVowels(symbol byte) bool {
	return symbol == 97 ||
		symbol == 101 ||
		symbol == 105 ||
		symbol == 111 ||
		symbol == 117
}
