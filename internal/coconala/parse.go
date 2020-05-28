package coconala

import (
	"os"

	"github.com/gocarina/gocsv"
)

func Parse(file *os.File) ([]*SalesData, error) {
	var coconalaSalesDatas []*SalesData

	if err := gocsv.UnmarshalFile(file, &coconalaSalesDatas); err != nil {
		return nil, err
	}

	return coconalaSalesDatas, nil
}
