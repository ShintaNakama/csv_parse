package coconala

import (
	"io"

	"github.com/gocarina/gocsv"
)

func Parse(file io.Reader) ([]*SalesData, error) {
	var coconalaSalesDatas []*SalesData

	if err := gocsv.Unmarshal(file, &coconalaSalesDatas); err != nil {
		return nil, err
	}

	return coconalaSalesDatas, nil
}
