package output

import "github.com/ShintaNakama/csv_parse/internal/coconala"

type OutputData struct {
	SalesDate   string
	SalesAmount string
}

func NewOutputDatas(salesDatas []*coconala.SalesData) []*OutputData {
	output := []*OutputData{}

	for _, data := range salesDatas {
		o := OutputData{
			SalesDate:   data.SalesDate,
			SalesAmount: data.SalesAmount,
		}
		output = append(output, &o)
	}

	return output
}

//func parse(file *os.File) ([]*SalesData, error) {
//	coconalaSalesDatas := []*SalesData{}
//
//	if err := gocsv.UnmarshalFile(file, &coconalaSalesDatas); err != nil {
//		return coconalaSalesDatas, err
//	}
//
//	return coconalaSalesDatas, nil
//}

//func convert(coconalaDatas []*SalesData, output []*OutputData) []*OutputData {
//
//	for _, data := range coconalaDatas {
//		o := OutputData{
//			SalesDate:   data.SalesDate,
//			SalesAmount: data.SalesAmount,
//		}
//		output = append(output, &o)
//	}
//
//	return output
//}
