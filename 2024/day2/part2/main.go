package main

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"strconv"
	"strings"
)

func main() {
	path := filepath.Join("..", "input.txt")

	f, err := os.Open(path)

	defer func() {
		if err := f.Close(); err != nil {
			fmt.Println("Error closing file")
		}
	}()

	if err != nil {
		fmt.Println("Error opening file " + path)
		return
	}

	scanner := bufio.NewScanner(f)

	var totalSafeLevels int
	for scanner.Scan() {
		line := scanner.Text()
		level := strings.Split(line, " ")

		if isSafe(level) {
			totalSafeLevels += 1
		}
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error while scanning the file")
	}

	fmt.Println("Total Safe levels", totalSafeLevels)
}

func isSafe(level []string) bool {
	trend := 0
	for i := 0; i < len(level)-1; i++ {
		curr, _ := strconv.Atoi(level[i])
		next, _ := strconv.Atoi(level[i+1])

		if next > curr {
			if trend == -1 || next-curr > 3 {
				return false
			}
			trend = 1
		} else if next < curr {
			if trend == 1 || next-curr < -3 {
				return false
			}
			trend = -1
		} else {
			return false
		}
	}
	return true
}
