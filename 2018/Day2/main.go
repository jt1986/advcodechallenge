package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func boxCheckSum() int {
	bl := boxListReader()
	twiceString := make(map[string]int)
	threeString := make(map[string]int)
	for i := range bl {
		// fmt.Println(bl[i])

		b := []byte(bl[i])
		for j := 0; j < len(b); j++ {
			currentString := bl[i]

			if strings.Count(bl[i], string(b[j])) == 2 {
				if _, ok := twiceString[currentString]; !ok {

					twiceString[currentString] = strings.Count(currentString, string(b[j]))
				}
				// fmt.Println(bl[i], string(b[j]), 2)
			}

		}
	}

	for i := range bl {
		// fmt.Println(bl[i])
		b := []byte(bl[i])
		for j := 0; j < len(b); j++ {
			currentString := bl[i]

			if strings.Count(bl[i], string(b[j])) == 3 {
				// fmt.Println(bl[i], string(b[j]), 3)
				if _, ok := threeString[currentString]; !ok {

					threeString[currentString] = strings.Count(currentString, string(b[j]))
				}
			}

		}
	}

	return len(twiceString) * len(threeString)
}

func boxListReader() []string {
	f, err := os.Open("./Day2/boxlist.txt")
	if err != nil {
		fmt.Println("File reading has errored out", err)
	}
	defer f.Close()
	// var boxIds []string
	fs := bufio.NewScanner(f)
	// for fs.Scan() {
	// 	boxIds = append(boxIds, fs.Text())
	// }

	strings := []string{}

	for fs.Scan() {
		x := fs.Text()

		for _, str := range strings {
			if pos := compareStrings(str, x); pos >= 0 {
				fmt.Println("Found the two strings:", str, x)
				fmt.Println("They differ for character at position", pos)
				fmt.Println("The common part is:", str[:pos]+str[pos+1:])
				return nil
			}
		}

		strings = append(strings, x)
	}
	fmt.Println(strings)
	return strings
}

func compareStrings(s string, t string) int {
	if len(s) != len(t) {
		return -1
	}

	diffs := 0
	posDiff := 0

	for i := 0; i < len(s); i++ {
		if s[i] != t[i] {
			diffs++
			posDiff = i
		}
	}

	if diffs == 1 {
		return posDiff
	}

	return -1
}
func main() {
	fmt.Println("Day2 Inventory Management System")
	fmt.Println("the checksum for the list of box ids", boxCheckSum())
	boxListReader()

	// fmt.Println("letters are common between the two correct box IDs", bl)
}
