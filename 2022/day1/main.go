package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	f, err := os.Open("input.txt")

	defer func() {
		if err := f.Close(); err != nil {
			fmt.Println("Error while closing the file", err)
		}
	}()

	if err != nil {
		fmt.Println("Could not open input.txt file", err)
		return
	}

	scanner := bufio.NewScanner(f)
	temp := 0
	var maxHeap = [3]int{0, 0, 0}

	for scanner.Scan() {
		line := scanner.Text()

		if line == "" {
			addToMaxHeap(temp, &maxHeap)
			temp = 0
			continue
		}

		i, err := strconv.Atoi(line)
		if err != nil {
			panic(err)
		}

		temp += i
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error while scanning the file")
		return
	}
	finalSum := 0
	for _, ele := range maxHeap {
		finalSum += ele
	}

	fmt.Println("Done!! Answer is", maxHeap)
	fmt.Println("Added answer", finalSum)
}

func addToMaxHeap(t int, heap *[3]int) {
	for i, ele := range heap {
		if t > ele {
			// Shift the elements downwards to make room for t
			copy(heap[i+1:], heap[i:])
			heap[i] = t
			break
		}
	}
}
