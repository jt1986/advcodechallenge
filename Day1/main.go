package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	fmt.Println("Advent code challenge day 1")
	f, err := os.Open("./Day1/frequency.txt")
	if err != nil {
		fmt.Println(err)
	}
	defer f.Close()

	var freq []int

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		i, err := strconv.Atoi(scanner.Text())
		if err != nil {
			fmt.Println(err)
		}
		freq = append(freq, i)
	}

	currentFreq := 0

	r := resultingFrequency(freq, currentFreq)
	fmt.Println("resulting freq", r)

	b := dupFreq(freq, currentFreq)
	fmt.Printf("Frequency %d has reached twice\n", b)
}

func resultingFrequency(freq []int, currentFreq int) int {
	var changeFreq int
	var rf int
	for j := range freq {
		changeFreq = freq[j]
		rf = currentFreq + changeFreq
		currentFreq = rf
	}
	return rf
}

func dupFreq(f []int, currentFreq int) int {

	dupFreqMap := make(map[int]bool)
	dupFreqMap[currentFreq] = true
	for {
		for i := 0; i < len(f); i++ {
			currentFreq += f[i]

			if _, ok := dupFreqMap[currentFreq]; ok {
				return currentFreq
			}
			dupFreqMap[currentFreq] = true
		}
	}
}
