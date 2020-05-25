package coconala

import (
	"os"

	"github.com/gocarina/gocsv"
)

type OutputData struct {
	SalesDate   string
	SalesAmount string
}

func Output(file *os.File) ([]*OutputData, error) {
	output := []*OutputData{}
	datas, err := parse(file)

	if err != nil {
		return output, err
	}

	output = convert(datas, output)

	return output, nil
}

func parse(file *os.File) ([]*SalesData, error) {
	coconalaSalesDatas := []*SalesData{}

	if err := gocsv.UnmarshalFile(file, &coconalaSalesDatas); err != nil {
		return coconalaSalesDatas, err
	}

	return coconalaSalesDatas, nil
}

func convert(coconalaDatas []*SalesData, output []*OutputData) []*OutputData {
	for _, data := range coconalaDatas {
		o := OutputData{
			SalesDate:   data.SalesDate,
			SalesAmount: data.SalesAmount,
		}
		output = append(output, &o)
	}

	return output
}
