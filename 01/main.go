package main

import (
	"bufio"
	"context"
	"flag"
	"fmt"
	"io"
	"log"
	"os"
	"os/signal"
	"strconv"
	"strings"
	"sync/atomic"
	"syscall"
)

var (
	input   = flag.String("i", "", "input file (stdin if empty)")
	output  = flag.String("o", "", "output file (stdout if empty)")
	verbose = flag.Bool("v", false, "verbose logging")
)

func openInput(path string) (io.Reader, func(), error) {
	if path == "" {
		return os.Stdin, func() {}, nil
	}
	f, err := os.Open(path)
	return f, func() { f.Close() }, err
}

func openOutput(path string) (io.Writer, func(), error) {
	if path == "" {
		return os.Stdout, func() {}, nil
	}
	f, err := os.Create(path)
	return f, func() { f.Close() }, err
}

func run(ctx context.Context) error {
	// Setup input reader
	reader, closeIn, err := openInput(*input)
	if err != nil {
		return err
	}
	defer closeIn()

	// Setup output writer
	writer, closeOut, err := openOutput(*output)
	if err != nil {
		return err
	}
	defer closeOut()

	// Process lines
	scanner := bufio.NewScanner(reader)
	bufWriter := bufio.NewWriter(writer)
	defer bufWriter.Flush()

	lineCount := 0
	var solCount1 int64 = 0
	var solCount2 int64 = 0
	var state int64 = 50
	for scanner.Scan() {
		select {
		case <-ctx.Done():
			return ctx.Err()
		default:
		}

		s1Delta, s2Delta := processLine(scanner.Text(), &state)

		solCount1 += s1Delta
		solCount2 += s2Delta

		lineCount++

		if *verbose && lineCount%1000 == 0 {
			log.Printf("Processed %d lines", lineCount)
		}
	}

	if *verbose {
		log.Printf("Complete: %d lines processed", lineCount)
	}
	fmt.Printf("solution1 %v\n", solCount1)
	fmt.Printf("solution2 %v\n", solCount2)
	return scanner.Err()
}

func main() {
	flag.Parse()

	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	defer cancel()

	if err := run(ctx); err != nil {
		log.Fatal(err)
	}
}

func processLine(line string, state *int64) (int64, int64) {
	str := strings.TrimSpace(strings.ToUpper(line))

	delta, err := strconv.ParseInt(str[1:], 0, 64)
	if err != nil {
		log.Fatalf("err conv %v\n", err)
	}
	div, q := delta/100, delta%100

	if str[0] == 'L' {
		delta = 100 - q // offset delta as it is a left turn
	}

	var passedZero = div
	if str[0] == 'L' {

		if *state != 0 && (*state-q) <= 0 {
			passedZero++
			fmt.Printf("Passed zero L-turn! %v\n", passedZero)
		}
	} else {
		if *state != 0 && (*state+q) >= 100 {
			passedZero++
		}
	}
	atomic.StoreInt64(state, (*state+delta)%100)
	if *state == 0 {
		return 1, passedZero
	}
	return 0, passedZero
}
