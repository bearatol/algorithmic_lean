package main

import "fmt"

// String Compression - LeetCode 443
// Сложность: Medium
// https://leetcode.com/problems/string-compression/
// Сжать строку: aabcccccaaa -> a2b1c5a3

func main() {
	fmt.Printf("%+v", getNumsList(1402))
}

func compress(chars []byte) int {
	if len(chars) == 1 {
		return 1
	}

	var (
		read byte = chars[0]
		num  int  = 1
		idx  int  = 0
	)
	for i := 1; i < len(chars); i++ {
		if read == chars[i] {
			num++
			continue
		}

		chars[idx] = read
		idx++
		read = chars[i]
		if num > 1 {
			for _, v := range getNumsList(num) {
				chars[idx] = v
				idx++
			}
		}
		num = 1
		continue
	}

	chars[idx] = read
	idx++
	if num > 1 {
		for _, v := range getNumsList(num) {
			chars[idx] = v
			idx++
		}
	}

	return idx
}

func getNumsList(whole int) []byte {
	res := make([]byte, 0, 4)
	var reminder int
	for whole != 0 {
		reminder = whole % 10
		res = append(res, byte(reminder)+48)
		whole /= 10
	}

	for i, j := 0, len(res)-1; i < j; i, j = i+1, j-1 {
		res[i], res[j] = res[j], res[i]
	}
	return res
}
