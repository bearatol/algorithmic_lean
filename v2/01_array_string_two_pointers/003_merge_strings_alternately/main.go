package main

// Merge Strings Alternately - LeetCode 1768
// Сложность: Easy
// https://leetcode.com/problems/merge-strings-alternately/
// Даны две строки word1 и word2.
// Вернуть строку, где символы из word1 и word2 чередуются.
// Если одна строка длиннее, добавить остальные символы в конец.

func main() {
	// Реализуйте решение здесь
}

// mergeAlternately сливает две строки, чередуя символы
func mergeAlternately(word1, word2 string) string {
	if len(word1) == 0 {
		return word2
	}
	if len(word2) == 0 {
		return word1
	}

	var result []rune
	w1 := []rune(word1)
	w2 := []rune(word2)
	for k := range len(w1) + len(w2) {
		if len(w1) > k {
			result = append(result, w1[k])
		}
		if len(w2) > k {
			result = append(result, w2[k])
		}
	}
	return string(result)
}
