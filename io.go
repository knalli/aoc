package aoc

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"log"
	"os"
	"strings"
)

/**
Read a file and returning the lines as array (without newlines)
*/
func ReadFileToArray(path string) ([]string, error) {
	f, err := os.Open(path)
	if err != nil {
		return nil, errors.New(fmt.Sprintf("error opening file: %v", err))
	}
	//noinspection GoUnhandledErrorResult
	defer f.Close()

	r := bufio.NewReader(f)
	var lines []string
	for {
		line, err := r.ReadString('\n')
		if line != "" {
			if line[len(line)-1] == '\n' {
				line = line[0 : len(line)-1] // remove newline (last char)
			}
			lines = append(lines, line)
		}
		if err != nil {
			if err == io.EOF {
				break
			}

			log.Fatalf("read file line error: %v", err)
			return nil, err
		}
	}
	return lines, nil
}

/**
Read a file and returns its content as one string
*/
func ReadFileToString(path string) (string, error) {
	if lines, err := ReadFileToArray(path); err != nil {
		return "", err
	} else {
		result := strings.Join(lines, "\n")
		return result, nil
	}
}

func ReadFileAsIntArray(file string) ([]int, error) {
	if str, err := ReadFileToString(file); err != nil {
		return nil, err
	} else {
		arr := strings.Split(strings.TrimSpace(str), ",")
		return ParseStringToIntArray(arr), nil
	}
}

func ParseStringToIntArray(lines []string) []int {
	res := make([]int, len(lines))
	for i, line := range lines {
		res[i] = ParseInt(line)
	}
	return res
}
