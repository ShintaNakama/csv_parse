package cmd

import (
	"fmt"
	"io"
	"log"
	"os"

	"github.com/ShintaNakama/csv_parse/internal/coconala"
)

func parse(args string, stdout io.Writer) {
	csv := args

	file, err := os.Open(csv)
	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()

	salesData, err := coconala.Parse(file)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
	}

	for _, d := range salesData {
		fmt.Fprintln(stdout, d.SalesDate)
		fmt.Fprintln(stdout, d.SalesAmount)
	}
}
