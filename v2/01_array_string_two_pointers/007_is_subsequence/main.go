package main

// Is Subsequence - LeetCode 392
// Сложность: Easy
// https://leetcode.com/problems/is-subsequence/
// Даны две строки s и t.
// Вернуть true, если s является подпоследовательностью t.
// Подпоследовательность - это последовательность символов,
// которая появляется в том же относительном порядке.

func main() {
	// Реализуйте решение здесь
}

// isSubsequence проверяет является ли s подпоследовательностью t
func isSubsequence(s, t string) bool {
	if len(s) == 0 {
		return true
	}
	if len(t) == 0 {
		return false
	}

	sIdx := 0
	for i := 0; i < len(t); i++ {
		if t[i] == s[sIdx] {
			sIdx++
		}

		if len(s) == sIdx {
			return true
		}
	}
	return false
}
