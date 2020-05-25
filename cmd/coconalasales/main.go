package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/ShintaNakama/csv_parse/internal/coconala"
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

	datas, err := coconala.Output(file)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
	}

	for _, data := range datas {
		fmt.Println(data.SalesDate)
		fmt.Println(data.SalesAmount)
	}
}
