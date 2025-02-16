package main

import (
	"bufio"
	"fmt"
	"os"
)

func reverse(s rune) rune {
	switch s {
	case '0':
		return '1'
	case '1':
		return '0'
	}
	return 0
}

func solve(s string, outputChannel chan int) {
	expected_char := '0'
	result := 0
	for _, c := range s {
		if c != expected_char {
			result++
			expected_char = reverse(expected_char)
		}
	}
	outputChannel <- result
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	var t int
	fmt.Fscan(reader, &t)
	var outputChannels []chan int = make([]chan int, t)

	for i := 0; i < t; i++ {
		var n int
		var s string
		fmt.Fscan(reader, &n)
		fmt.Fscan(reader, &s)
		outputChannels[i] = make(chan int, 1)
		go solve(s, outputChannels[i])
	}

	for i := 0; i < t; i++ {
		fmt.Fprintln(writer, <-outputChannels[i])
	}
}
