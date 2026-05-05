package main

import (
	"flag"
	"fmt"
	"os"
	"time"

	"github.com/user/logslice/internal/filereader"
	"github.com/user/logslice/internal/logfilter"
	"github.com/user/logslice/internal/output"
	"github.com/user/logslice/internal/processor"
	"github.com/user/logslice/internal/timeparser"
	"github.com/user/logslice/internal/timerange"
)

func main() {
	var (
		pattern    = flag.String("pattern", "", "regex pattern to filter lines")
		start      = flag.String("start", "", "start time (e.g. 2006-01-02T15:04:05)")
		end        = flag.String("end", "", "end time (e.g. 2006-01-02T15:04:05)")
		maxLines   = flag.Int("n", 0, "maximum number of lines to output (0 = unlimited)")
		lineNums   = flag.Bool("line-numbers", false, "prefix output with line numbers")
		prefix     = flag.String("prefix", "", "prefix string for each output line")
	)
	flag.Parse()

	args := flag.Args()
	if len(args) == 0 {
		fmt.Fprintln(os.Stderr, "usage: logslice [flags] <file> [file...]")
		flag.PrintDefaults()
		os.Exit(1)
	}

	tp := timeparser.New(time.UTC)

	var tr *timerange.TimeRange
	if *start != "" || *end != "" {
		s, err := tp.Parse(*start)
		if err != nil {
			fmt.Fprintf(os.Stderr, "invalid start time: %v\n", err)
			os.Exit(1)
		}
		e, err := tp.Parse(*end)
		if err != nil {
			fmt.Fprintf(os.Stderr, "invalid end time: %v\n", err)
			os.Exit(1)
		}
		tr, err = timerange.New(s, e)
		if err != nil {
			fmt.Fprintf(os.Stderr, "invalid time range: %v\n", err)
			os.Exit(1)
		}
	}

	f, err := logfilter.New(*pattern, tr)
	if err != nil {
		fmt.Fprintf(os.Stderr, "invalid filter: %v\n", err)
		os.Exit(1)
	}

	out := output.New(os.Stdout, *lineNums, *prefix)
	proc := processor.New(f, out, *maxLines)

	for _, path := range args {
		if err := processFile(path, proc); err != nil {
			fmt.Fprintf(os.Stderr, "error processing %q: %v\n", path, err)
			os.Exit(1)
		}
	}
}

func processFile(path string, proc *processor.Processor) error {
	r, err := filereader.New(path)
	if err != nil {
		return err
	}
	defer r.Close()
	return proc.Process(r)
}
