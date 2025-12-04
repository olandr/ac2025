package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/charmbracelet/log"
)

var input, output = flag.String("i", "", "input"), flag.String("o", "", "output")
var logger = log.New(os.Stderr)

func main() {
	flag.Parse()
	if level, err := log.ParseLevel(os.Getenv("LOG_LEVEL")); err == nil {
		logger.SetLevel(level)
	}
	in, out := os.Stdin, os.Stdout
	if *input != "" {
		f, err := os.Open(*input)
		if err != nil {
			logger.Fatal("Open failed", "error", err)
		}
		defer f.Close()
		in = f
	}
	if *output != "" {
		f, err := os.Create(*output)
		if err != nil {
			logger.Fatal("Create failed", "error", err)
		}
		defer f.Close()
		out = f
	}
	result1 := process(in, out, 2)
	fmt.Println(result1)

	// // Reset file to beginning for second pass
	// if *input != "" {
	// 	in.Seek(0, 0)
	// }

	result2 := process(in, out, 12)
	fmt.Println(result2)
}
func _int64(v string) int64 {
	if r, err := strconv.ParseInt(v, 0, 64); err == nil {
		return r
	}
	return -1
}
func findMax(digits string, start, end int) (int, string) {
	//fmt.Printf("start %v end %v\n", start, end)

	var max = "-1"
	var ix = -1
	for i := start; i < end; i++ {
		if _int64(max) < _int64(string(digits[i])) {
			max = string(digits[i])
			ix = i
		}
	}
	return ix, max
}

func process(in, out *os.File, length_of_battery_toggles int) int64 {
	var count int64
	s, w := bufio.NewScanner(in), bufio.NewWriter(out)
	s.Buffer(make([]byte, 64*1024), 1024*1024)
	defer w.Flush()
	for s.Scan() {
		line := s.Text()
		var sb strings.Builder
		var ix = 0
		var max string
		for digitsSelected := range length_of_battery_toggles {
			endIdx := len(line) - (length_of_battery_toggles - digitsSelected - 1)
			ix, max = findMax(line, ix, endIdx)
			ix++
			sb.Write([]byte(max))
		}
		delta1, _ := strconv.ParseInt(sb.String(), 0, 64)
		count += delta1
	}

	return count
}
