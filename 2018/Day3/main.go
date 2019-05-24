package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
)

func Index(vs []int, n int) int {
	for i, v := range vs {
		if v == n {
			return i
		}
	}
	return -1
}

func main() {
	f, err := os.Open("input.txt")
	if err != nil {
		fmt.Println("File reading has errored out", err)
	}
	defer f.Close()

	fs := bufio.NewScanner(f)

	var fb [1000][1000]int

	intactClaims := []int{}
	total := 0
	isIntact := true

	r, _ := regexp.Compile("#([0-9]+) @ ([0-9]+),([0-9]+): ([0-9]+)x([0-9]+)")
	for fs.Scan() {
		v := fs.Text()
		id, _ := strconv.Atoi(r.FindStringSubmatch(v)[1])
		s, _ := strconv.Atoi(r.FindStringSubmatch(v)[2])
		e, _ := strconv.Atoi(r.FindStringSubmatch(v)[3])
		w, _ := strconv.Atoi(r.FindStringSubmatch(v)[4])
		l, _ := strconv.Atoi(r.FindStringSubmatch(v)[5])
		for i := s; i < s+w; i++ {
			for j := e; j < e+l; j++ {

				if fb[i][j] > 0 {
					// the current claim is certainly not intact
					isIntact = false
					// not the first visit of this square, mark the claim as not intact
					index := Index(intactClaims, fb[i][j])
					if index >= 0 {
						intactClaims = append(intactClaims[:index], intactClaims[index+1:]...)
					}

				}
				fb[i][j] = id
			}

		}
		if isIntact {
			intactClaims = append(intactClaims, id)
		}

	}

	fmt.Println("square inches of fabric which overlap", total)
	fmt.Println("ID of the only claim that doesn't overlap", intactClaims, isIntact)
}
