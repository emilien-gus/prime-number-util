package main

import (
	"flag"
	"fmt"
	"strconv"
	"strings"
)

var timeout int
var fileName string

type Range struct {
	Start uint32
	End   uint32
}

type RangeSlice []Range

func (rs *RangeSlice) String() string {
	var parts []string
	for _, r := range *rs {
		parts = append(parts, fmt.Sprintf("%d:%d", r.Start, r.End))
	}
	return strings.Join(parts, ", ")
}

func (rs *RangeSlice) Set(value string) error {
	parts := strings.Split(value, ":")
	if len(parts) != 2 {
		return fmt.Errorf("wrong range format: %s (expect start:end)", value)
	}
	start64, err := strconv.ParseUint(parts[0], 10, 32)
	if err != nil {
		return fmt.Errorf("parse error start: %w", err)
	}
	end64, err := strconv.ParseUint(parts[1], 10, 32)
	if err != nil {
		return fmt.Errorf("parse error end: %w", err)
	}

	start := uint32(start64)
	end := uint32(end64)
	if start > end {
		return fmt.Errorf("start end be less then start %s", value)
	}
	*rs = append(*rs, Range{Start: start, End: end})
	return nil
}

func main() {
	var ranges RangeSlice
	flag.StringVar(&fileName, "file", "output.txt", "output file name (standart: output.txt)")
	flag.IntVar(&timeout, "timeout", 0, "timeout in seconds")
	flag.Var(&ranges, "range", "Диапазон чисел в формате start:end. Values limited with uint32 type")
	flag.Parse()

}

// func processRange()
