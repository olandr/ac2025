package main

import (
	"bufio"
	"flag"
	"fmt"
	"iter"
	"log"
	"os"
	"slices"
	"strconv"
	"strings"
	"sync"
	"sync/atomic"
)

var input, output = flag.String("i", "", "input"), flag.String("o", "", "output")

func main() {
	flag.Parse()
	in, out := os.Stdin, os.Stdout
	if *input != "" {
		if f, err := os.Open(*input); err == nil {
			defer f.Close()
			in = f
		} else {
			log.Fatal(err)
		}
	}
	if *output != "" {
		if f, err := os.Create(*output); err == nil {
			defer f.Close()
			out = f
		} else {
			log.Fatal(err)
		}
	}
	process(in, out)
}

func genRange(mi, mx string) iter.Seq2[int64, string] {
	min, _ := strconv.ParseInt(mi, 0, 64)
	max, _ := strconv.ParseInt(mx, 0, 64)
	return func(yield func(int64, string) bool) {
		for k := min; k <= max; k++ {
			if !yield(k, fmt.Sprintf("%v", k)) {
				return
			}
		}
	}
}

func _recsplit(id string, size int) bool {

	//for c := range slices.Chunk([]byte(id), size) {
	//	fmt.Printf("c: %s ", c)
	//}
	//fmt.Println()
	chunks := slices.Collect(slices.Chunk([]byte(id), size))
	ret := true
	prev := chunks[0]
	for _, chunk := range chunks {
		//fmt.Printf("  prev %s\n", prev)
		//fmt.Printf("  chunk %s\n", chunk)
		ret = ret && slices.Equal(prev, chunk)
		//fmt.Printf("   slices.Equal(prev, chunk) %v\n", slices.Equal(prev, chunk))
		//fmt.Printf("  ret %v\n", ret)

		if !ret {
			break
		}
	}
	return ret
}

func rule(id string) (bool, bool) {
	////fmt.Printf("running on ID: %v\n", id)
	ret1 := true
	for window := 1; window <= len(id)/2; window++ {
		ret1 = id[0:window] == id[window:]
	}
	for window := 1; window <= len(id)/2; window++ {
		//fmt.Printf(" window %v\n", window)
		if _recsplit(id, window) {
			return ret1, true
		}
	}
	return ret1, false
}

func process(in, out *os.File) {
	seen := sync.Map{}
	s, w := bufio.NewScanner(in), bufio.NewWriter(out)
	s.Buffer(make([]byte, 64*1024), 1024*1024)
	count1 := new(atomic.Int64)
	count2 := new(atomic.Int64)
	//var wg sync.WaitGroup
	defer w.Flush()
	for s.Scan() {
		ids := strings.SplitSeq(s.Text(), ",")
		for _idv := range ids {
			//wg.Go(func() {
			_v := strings.Split(_idv, "-")
			for numID, ID := range genRange(_v[0], _v[1]) {
				seen.Store(ID, true)
				ret1, ret2 := rule(ID)
				if ret1 {
					//fmt.Printf("ID is invalid %v\n", numID)
					count1.Add(numID)
				}
				if ret2 {
					//fmt.Printf("ID is invalid %v\n", numID)
					count2.Add(numID)
				}
			}
			//	})
		}
	}
	//wg.Wait()
	fmt.Fprintln(w, "soulution 1:", count1.Load())
	fmt.Fprintln(w, "soulution 2:", count2.Load())
}
