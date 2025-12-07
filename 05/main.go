package main

import (
	"bufio"
	"flag"
	"fmt"
	"maps"
	"os"
	"slices"
	"strings"

	"github.com/charmbracelet/log"
	"github.com/olandr/ac2025/utils"
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

func comp(ing int64, fresh map[int64]int64) int64 {
	for start, end := range maps.All(fresh) {
		if start <= ing && ing <= end {
			return 1
		}
	}
	return 0
}

func reduceOverlap(sets []*MaxSet) []*MaxSet {
	//ret := make([]*MaxSet, 0)
	for ci, check := range sets {
		for _, v := range sets[ci+1:] {
			cm, _ := check.Max()
			//log.Info(" checking", "cm", cm, "v.start", v.start)
			if v.start <= cm {
				vm, _ := v.Max()
				//log.Info("  found startlap!", "vm", vm, "cm", cm)
				if vm <= cm {
					//log.Info("   max of this range is contained in check! skip!", "v.Max()", vm, "vmi", vmi)
				} else {
					//log.Info("   vm is bigger! expanding check range!", "check.all", check.all, "with", vm)
					check.all = append(check.all, vm)
					//log.Info("   new range", "check.all", check.all)
					check.end = vm
				}
				//v.all = slices.Delete(v.all, vi, vi+1)
				v.skip = true
			}
		}
	}

	return withoutEmpty(sets)
}

type MaxSet struct {
	start int64
	end   int64
	skip  bool
	all   []int64
}

func (m *MaxSet) Max() (int64, int) {
	var ret int64
	var reti int
	for i, v := range m.all {
		if ret < v {
			ret = v
			reti = i
		}
	}
	m.end = ret
	return ret, reti
}

func (m *MaxSet) String() string {
	var sb strings.Builder
	sb.Write([]byte(fmt.Sprintf("{%v", m.start)))
	sb.Write([]byte(":"))
	for i, v := range m.all {
		sb.Write([]byte(fmt.Sprintf("%v", v)))
		if i < len(m.all)-1 {
			sb.Write([]byte(" "))
		}
	}
	sb.Write([]byte(fmt.Sprintf("}(max=%v", m.end)))
	sb.Write([]byte(fmt.Sprintf(" %v", m.skip)))
	sb.Write([]byte(")"))

	return sb.String()
}

func (m *MaxSet) Append(v int64) {
	m.all = append(m.all, v)
}

func withoutEmpty(sets []*MaxSet) []*MaxSet {
	ret := make([]*MaxSet, 0)
	for _, v := range sets {
		if !v.skip {
			ret = append(ret, v)
		}
	}
	return ret
}

//func (m *MaxSet) String() {
//	return fmt.Sprintf("")
//}

func sortMaxSet(set map[int64]*MaxSet) []*MaxSet {
	ret := make([]*MaxSet, 0)
	for v := range maps.Values(set) {
		ret = append(ret, v)
	}
	slices.SortFunc(ret, func(a, b *MaxSet) int {
		return int(a.start) - int(b.start)
	})
	return ret
}

func process(in, out *os.File) {
	s, w := bufio.NewScanner(in), bufio.NewWriter(out)
	s.Buffer(make([]byte, 64*1024), 1024*1024)
	defer w.Flush()
	fresh1 := make(map[int64]int64)
	fresh := make(map[int64]*MaxSet)

	for s.Scan() {
		line := s.Text()
		if line == "" {
			break
		}
		split := strings.Split(line, "-")
		start, end := utils.Int64(split[0]), utils.Int64(split[1])
		if _, ok := fresh[start]; !ok {
			fresh[start] = &MaxSet{start: start}
		}
		if cm, _ := fresh[start].Max(); cm < end {
			fresh[start].Append(end)
		}
		if fresh1[start] < end {
			fresh1[start] = end
		}
	}
	all := make([]int64, 0)
	for s.Scan() {
		all = append(all, utils.Int64(s.Text()))
	}
	count1 := 0
	seen := make(map[int64]bool)
	for _, f := range all {
		for start, end := range fresh1 {
			if !seen[f] && start <= f && f <= end {
				count1++
				seen[f] = true
			}
		}
	}

	freshSorted := sortMaxSet(fresh)
	var count2 int64

	freshSorted = reduceOverlap(freshSorted)

	for _, mx := range freshSorted {
		mx.Max()
		count2 += (mx.end - mx.start) + 1
	}

	log.Info("solution 1", "count", count1)
	log.Info("solution 2", "count", count2, "PASS", 343329651880509 == count2, "diff", 343329651880509-count2)
}
