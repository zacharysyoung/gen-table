/*
Table prints a simple CSV table of rows X cols size, plus a header. The rows
and cols args must both be positive integers.

"Usage of table: [-h] rows cols
*/
package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"strings"
)

var (
	rows, cols int

	w   *csv.Writer
	row []string

	randCharsFlag = flag.Int("randChars", 0, "instead of R1C1 format for field data, create n different random chars, repeated randomLens")
	randLensFlag  = flag.Int("randLens", 0, "instead of R1C1 format for field data, create n different random lengths of randomChars")
)

func main() {
	const usage = `[-h] [-randChars -randLens] rows cols

Table prints a simple CSV table of rows X cols size, plus a header. The rows
and cols args must both be positive integers.`

	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, "Usage of table: %s\n", usage)
		flag.PrintDefaults()
		os.Exit(2)
	}
	flag.Parse()

	var (
		args = flag.Args()
		err  error
	)
	if len(args) != 2 {
		flag.Usage()
	}
	if rows, err = strconv.Atoi(args[0]); err != nil || rows < 1 {
		fmt.Fprintln(os.Stderr, "error: rows must be a positive integer")
		os.Exit(2)
	}
	if cols, err = strconv.Atoi(args[1]); err != nil || cols < 1 {
		fmt.Fprintln(os.Stderr, "error: cols must be a positive integer")
		os.Exit(2)
	}
	if (*randCharsFlag > 0 || *randLensFlag > 0) && !(*randCharsFlag > 0 && *randLensFlag > 0) {
		fmt.Fprintln(os.Stderr, "error: randChars and randLens must be set together")
		os.Exit(2)
	}
	if *randCharsFlag < 0 {
		fmt.Fprintln(os.Stderr, "error: random chars must be a positive integer")
		os.Exit(2)
	}
	if *randLensFlag < 0 {
		fmt.Fprintln(os.Stderr, "error: random lengths must be a positive integer")
		os.Exit(2)
	}

	w = csv.NewWriter(os.Stdout)
	row = make([]string, cols)

	for i := range cols {
		row[i] = fmt.Sprintf("Col_%d", i+1)
	}
	w.Write(row)

	if *randLensFlag > 0 {
		writeRand()
	} else {
		writeR1C1()
	}

	w.Flush()
}

func writeR1C1() {
	for i := range rows {
		for j := range cols {
			row[j] = fmt.Sprintf("r%d_c%d", i+1, j+1)
		}
		w.Write(row)
	}
}

func writeRand() {
	var randChars []string
	for i := range *randCharsFlag {
		randChars = append(randChars, string('a'+i))
	}

	var randLens []int
	for i := range *randLensFlag {
		randLens = append(randLens, i+1)
	}

	var (
		len  int
		char string
	)
	for range rows {
		for j := range cols {
			char = randChars[rand.Intn(*randCharsFlag)]
			len = randLens[rand.Intn(*randLensFlag)]
			row[j] = strings.Repeat(char, len)
		}
		w.Write(row)
	}

}
