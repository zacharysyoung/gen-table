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
	"os"
	"strconv"
)

func main() {
	const usage = `[-h] rows cols

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

		rows, cols int
		err        error
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

	w := csv.NewWriter(os.Stdout)
	row := make([]string, cols)

	for i := range cols {
		row[i] = fmt.Sprintf("Col_%d", i+1)
	}
	w.Write(row)

	for i := range rows {
		for j := range cols {
			row[j] = fmt.Sprintf("r%d_c%d", i+1, j+1)
		}
		w.Write(row)
	}

	w.Flush()
}
