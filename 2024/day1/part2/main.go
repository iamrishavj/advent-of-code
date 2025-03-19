package main

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"sort"
	"strconv"
	"strings"
	"sync"
)

func main() {
	path := filepath.Join("..", "input.txt")
	f, err := os.Open(path)

	defer func() {
		if err := f.Close(); err != nil {
			fmt.Println("Error while closing the file", err)
		}
	}()

	if err != nil {
		fmt.Println("Could not open input.txt file", err)
		return
	}

	var list1 []int
	var list2 []int

	scanner := bufio.NewScanner(f)

	for scanner.Scan() {
		line := scanner.Text()

		numbers := strings.Split(line, "   ")
		number1, _ := strconv.Atoi(numbers[0])
		number2, _ := strconv.Atoi(numbers[1])
		list1 = append(list1, number1)
		list2 = append(list2, number2)
	}

	var wg sync.WaitGroup
	wg.Add(2)
	go func() {
		defer wg.Done()
		sort.Ints(list1)
	}()
	go func() {
		defer wg.Done()
		sort.Ints(list2)
	}()

	wg.Wait()

	var totalDistance int

	for _, value := range list1 {
		totalDistance += value * findCount(list2, value)
	}

	fmt.Println("Done!! Answer is", totalDistance)
}

func Abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func findCount(list []int, value int) int {
	count := 0

	for _, val := range list {
		if val > value {
			break
		}
		if val == value {
			count += 1
		}
	}

	return count
}
