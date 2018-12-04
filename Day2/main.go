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
	var boxIds []string
	fs := bufio.NewScanner(f)
	for fs.Scan() {
		boxIds = append(boxIds, fs.Text())
	}

	return boxIds
}
func main() {
	fmt.Println("Day2 Inventory Management System")
	fmt.Println("the checksum for the list of box ids", boxCheckSum())
	fmt.Println("letters are common between the two correct box IDs")
}
