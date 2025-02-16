package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func calcUniqueElements(array []int) map[int]bool {
	uniqueElements := make(map[int]bool)
	for _, element := range array {
		if _, ok := uniqueElements[element]; ok {
			uniqueElements[element] = false
		} else {
			uniqueElements[element] = true
		}
	}
	return uniqueElements
}

func stringToIntArray(s string) []int {
	var intArray []int
	stringArray := strings.Fields(s)
	for _, stringValue := range stringArray {
		intValue, _ := strconv.Atoi(stringValue)
		intArray = append(intArray, intValue)
	}
	return intArray
}

// Should return (range start, range end, true). (0, 0, false) if no such range exists.
func calcMaxConsecutiveElementsRange(array []int) (int, int, bool) {

	uniqueElements := calcUniqueElements(array)

	bestStart, bestEnd := 0, 0

	start := 0
	maxLength := 0
	currentLength := 0
	for i, element := range array {
		if uniqueElements[element] {
			currentLength++
			if currentLength > maxLength {
				maxLength = currentLength
				bestStart = start
				bestEnd = i
			}
		} else {
			currentLength = 0
			start = i + 1
		}
	}
	if maxLength == 0 {
		return 0, 0, false
	}
	return bestStart, bestEnd, true
}

func solve(s string) string {
	array := stringToIntArray(s)
	start, end, ok := calcMaxConsecutiveElementsRange(array)
	if !ok {
		return "0"
	}
	return fmt.Sprintf("%d %d", start+1, end+1)
}

func readInt(reader *bufio.Reader) int {
	s, _ := reader.ReadString('\n')
	s = strings.TrimSpace(s)
	n, _ := strconv.Atoi(s)
	return n
}

func syncSolve(f func(string) string) {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	var t int = readInt(reader)

	for i := 0; i < t; i++ {
		readInt(reader)
		s, _ := reader.ReadString('\n')
		result := f(s)
		fmt.Fprintln(writer, result)
	}
}

func asyncSolve(f func(string) string) {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)

	var t int = readInt(reader)

	tasks := make([]chan string, t)

	for i := 0; i < t; i++ {
		tasks[i] = make(chan string, 1)
		readInt(reader)
		s, _ := reader.ReadString('\n')
		go func(i int) {
			result := f(s)
			tasks[i] <- result
		}(i)
	}

	for i := 0; i < t; i++ {
		result := <-tasks[i]
		fmt.Fprintf(writer, "%v\n", result)
	}

	writer.Flush()
}

func main() {
	asyncSolve(solve)
}
