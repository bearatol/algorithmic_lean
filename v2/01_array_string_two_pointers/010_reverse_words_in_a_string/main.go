package main

// Reverse Words in a String - LeetCode 151
// Сложность: Medium
// https://leetcode.com/problems/reverse-words-in-a-string/
// Дана строка s. Переверните порядок слов в строке.
// Слова определяются как последовательность непробельных символов.
// Строка может содержать лишние пробелы, их нужно удалить.

func main() {
	// Реализуйте решение здесь
}

// reverseWords переворачивает порядок слов в строке
func reverseWords(s string) string {
	sChar := []rune(s)
	sCharList := make([]string, 0)
	sString := ""
	for i := 0; i < len(sChar); i++ {
		if sChar[i] == ' ' {
			if len(sString) != 0 {
				sCharList = append(sCharList, sString)
				sString = ""
			}
			continue
		}
		sString += string(sChar[i])
	}

	if len(sString) != 0 {
		sCharList = append(sCharList, sString)
	}

	if len(sCharList) == 0 {
		return ""
	}

	sString = sCharList[len(sCharList)-1]
	for i := len(sCharList) - 2; i >= 0; i-- {
		sString += " " + sCharList[i]
	}

	return sString
}
