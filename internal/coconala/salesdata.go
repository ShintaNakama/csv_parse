package coconala

import (
	"os"

	"github.com/gocarina/gocsv"
)

type SalesData struct {
	SalesDate      string `csv:"売上確定日"`
	TolkRoomNo     string `csv:"トークルームNo"`
	ServiceName    string `csv:"サービス名/見積り・仕事の相談名"`
	BuyerID        string `csv:"購入者ID"`
	BuyerName      string `csv:"購入者名"`
	SalesBreakdown string `csv:"内訳"`
	SalesAmount    string `csv:"売上金額"`
}

func Parse(file *os.File) ([]*SalesData, error) {
	coconalaSalesDatas := []*SalesData{}

	if err := gocsv.UnmarshalFile(file, &coconalaSalesDatas); err != nil {
		return coconalaSalesDatas, err
	}

	return coconalaSalesDatas, nil
}
