package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/ShintaNakama/csv_parse/internal/coconala"
	"github.com/ShintaNakama/csv_parse/internal/output"
)

func passArgs() string {
	flag.Parse()
	args := strings.Join(flag.Args(), "")

	return args
}

func main() {
	args := passArgs()

	file, err := os.Open(args)
	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()

	salesDatas, err := coconala.Parse(file)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
	}

	datas := output.NewOutputDatas(salesDatas)

	for _, data := range datas {
		fmt.Println(data.SalesDate)
		fmt.Println(data.SalesAmount)
	}
}
