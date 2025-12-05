package main

import (
	"bufio"
	"flag"
	"os"

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
	process(in, out)
}

func _checkaround(x int, y int, space Grid, xmax, ymax int) int {
	adj := 0
	for dy := -1; dy <= 1; dy++ {
		for dx := -1; dx <= 1; dx++ {
			if dx == dy && dx == 0 {
				continue
			}
			if y+dy < 0 || x+dx < 0 || y+dy >= ymax || x+dx >= xmax {
				continue
			}
			switch space[y+dy][x+dx] {
			case '@':
				adj++
			}
		}
	}
	if adj < 4 {
		return 1
	}
	return 0
}

type Grid [][]rune

func process(in, out *os.File) {
	s, w := bufio.NewScanner(in), bufio.NewWriter(out)
	s.Buffer(make([]byte, 64*1024), 1024*1024)
	defer w.Flush()
	var space Grid = make(Grid, 140)
	xmax, y := 0, 0
	for s.Scan() {
		line := s.Text()
		xmax = len(line)
		for x, cell := range line {
			if space[y] == nil {
				space[y] = make([]rune, 200)
			}
			space[y][x] = cell
		}

		y++
	}

	ymax := y
	log.Info("solution 1", "count", logic(xmax, ymax, space, false))
	log.Info("solution 1", "count", logic(xmax, ymax, space, true))
}

func logic(xmax, ymax int, space Grid, inplace_edit bool) int {
	count := 0
	removed := false
	for y := range ymax {
		for x := range space[y] {
			cell := space[y][x]
			var v int
			switch cell {
			case '@':
				v = _checkaround(x, y, space, xmax, ymax)
				count += v
			case '.':
			}
			if v == 1 && inplace_edit {
				space[y][x] = 'r'
				removed = true
			} else {
				//fmt.Printf("%v", string(cell))
			}
		}
		//fmt.Println()
	}
	if removed && inplace_edit {
		return count + logic(xmax, ymax, space, true)
	}
	return count
}
