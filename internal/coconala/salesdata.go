package coconala

type SalesData struct {
	SalesDate      string `csv:"売上確定日"`
	TolkRoomNo     string `csv:"トークルームNo"`
	PaymentNo      string `csv:"決済No"`
	ServiceName    string `csv:"サービス名/見積り・仕事の相談名"`
	BuyerID        string `csv:"購入者ID"`
	BuyerName      string `csv:"購入者名"`
	SalesBreakdown string `csv:"内訳"`
	SalesAmount    string `csv:"売上金額"`
}
