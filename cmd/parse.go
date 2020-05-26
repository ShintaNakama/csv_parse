package cmd

import (
	"fmt"
	"log"
	"os"

	"github.com/ShintaNakama/csv_parse/internal/coconala"
)

func parse(args string) {
	csv := args

	file, err := os.Open(csv)
	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()

	datas, err := coconala.Output(file)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
	}

	for _, data := range datas {
		fmt.Println(data.SalesDate)
		fmt.Println(data.SalesAmount)
	}
}
