package coconala_test

import (
	"bytes"
	"encoding/csv"
	"io"
	"testing"

	"github.com/ShintaNakama/csv_parse/internal/coconala"
)

const sampleData = `売上確定日,トークルームNo,決済No,サービス名/見積り・仕事の相談名,購入者ID,購入者名,内訳,売上金額
2019/12/15,4707685,5372134,ショップカード制作の件,777257,ccnl0703,基本料金,"36,500"
2019/12/15,4707685,5504060,ショップカード制作の件,777257,ccnl0703,おひねり(追加),"2,190"
2019/11/8,4596458,5407343,シンプルなロゴデザインを制作しませんか,105,ココナラ太郎,基本料金,"21,750"
2019/11/8,4596458,5319576,シンプルなロゴデザインを制作しませんか,105,ココナラ太郎,オプション支払い,"2,175"
2019/11/8,4596458,5276122,シンプルなロゴデザインを制作しませんか,105,ココナラ太郎,おひねり(追加),725
`

func createExpectedData(sample io.Reader) ([][]string, error) {
	r := csv.NewReader(sample)

	records, err := r.ReadAll()
	if err != nil {
		return nil, err
	}

	return records, nil
}

func TestParse(t *testing.T) {
	expected, err := createExpectedData(bytes.NewBufferString(sampleData))
	if err != nil {
		t.Fatal(err)
	}

	coconalaSalesData, err := coconala.Parse(bytes.NewBufferString(sampleData))
	if err != nil {
		t.Fatal(err)
	}

	for i, d := range coconalaSalesData {
		if d.SalesDate != expected[i+1][0] {
			t.Errorf("SalesDate not matched expected: %s actual: %s", expected[i+1][0], d.SalesDate)
		}

		if d.TolkRoomNo != expected[i+1][1] {
			t.Errorf("TolkRoomNo not matched expected: %s actual: %s", expected[i+1][1], d.TolkRoomNo)
		}

		if d.PaymentNo != expected[i+1][2] {
			t.Errorf("PaymentNo not matched expected: %s actual: %s", expected[i+1][2], d.PaymentNo)
		}

		if d.ServiceName != expected[i+1][3] {
			t.Errorf("ServiceName not matched expected: %s actual: %s", expected[i+1][3], d.ServiceName)
		}

		if d.BuyerID != expected[i+1][4] {
			t.Errorf("BuyerID not matched expected: %s actual: %s", expected[i+1][4], d.BuyerID)
		}

		if d.BuyerName != expected[i+1][5] {
			t.Errorf("BuyerName not matched expected: %s actual: %s", expected[i+1][5], d.BuyerName)
		}

		if d.SalesBreakdown != expected[i+1][6] {
			t.Errorf("SalesBreakdown not matched expected: %s actual: %s", expected[i+1][6], d.SalesBreakdown)
		}

		if d.SalesAmount != expected[i+1][7] {
			t.Errorf("SalesAmount not matched expected: %s actual: %s", expected[i+1][7], d.SalesAmount)
		}
	}
}
