package main

import (
	"fmt"
)

func main() {
	// scanner := bufio.NewScanner(os.Stdin)
	// scanner.Scan()
	// text := scanner.Text()

	text := `
	She sells sea shells on the sea shore;
	The shells that she sells are sea shells I'm sure.
	So if she sells sea shells on the sea shore,
	I'm sure that the shells are sea shore shells.
	`

	words := make(map[string]struct{})
	var wordWindow string
	for _, v := range text {
		if v == '.' ||
			v == ',' ||
			v == ';' ||
			v == '?' ||
			v == '!' ||
			v == ':' {
			continue
		}
		if v == ' ' || v == '\n' {
			if wordWindow == "" {
				continue
			}
			words[wordWindow] = struct{}{}
			wordWindow = ""
			continue
		}
		wordWindow += string(v)
	}

	i := 0
	for k := range words {
		i++
		fmt.Println(i, "-", k)
	}
	fmt.Println(len(words))
}
