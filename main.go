package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	if err := run(); err != nil {
		log.Fatalf("an error occurred: %v", err)
	}
}

func run() error {
	in, err := readInput()
	if err != nil {
		return fmt.Errorf("failed to read the input: %w", err)
	}
	log.Println(in)
	result := CalculatePathEnergy(in)
	resultStrs := make([]string, 0, len(result))
	for _, r := range result {
		resultStrs = append(resultStrs, strconv.Itoa(r))
	}
	fmt.Println(strings.Join(resultStrs, " "))
	return nil
}

func readInput() (*Input, error) {
	reader := bufio.NewReader(os.Stdin)
	readBytes, err := io.ReadAll(reader)
	if err != nil {
		return nil, fmt.Errorf("failed to read the input: %w", err)
	}
	readStr := string(readBytes)
	parts := strings.Split(readStr, "\n")
	const expectedPartsLen = 2
	if len(parts) < expectedPartsLen {
		return nil, fmt.Errorf("expected %d parts, got %d: %v", expectedPartsLen, len(parts), parts)
	}
	nStr := parts[0]
	if err != nil {
		return nil, fmt.Errorf("failed to read the number of input: %w", err)
	}
	n, err := strconv.ParseInt(nStr, 10, 0)
	if err != nil {
		return nil, fmt.Errorf("failed to parse the N parameter: %w", err)
	}
	shortcutsStr := parts[1]
	if err != nil {
		return nil, fmt.Errorf("failed to read the shortcuts string: %w", err)
	}
	shortcutsNums := strings.Split(shortcutsStr, " ")
	if len(shortcutsNums) != int(n) {
		return nil, fmt.Errorf("expected to get %d shortcuts, got %d: %v", n, len(shortcutsNums), shortcutsNums)
	}
	shortcuts := make([]int, len(shortcutsNums))
	for i, s := range shortcutsNums {
		sn, err := strconv.ParseInt(s, 10, 0)
		if err != nil {
			return nil, fmt.Errorf("failed to parse shortcut %s: %w", s, err)
		}
		shortcuts[i] = int(sn)
	}
	return &Input{
		Shortcuts: shortcuts,
	}, nil
}
