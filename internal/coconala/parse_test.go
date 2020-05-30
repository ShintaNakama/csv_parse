package coconala_test

import (
	"encoding/csv"
	"os"
	"testing"

	"github.com/ShintaNakama/csv_parse/internal/coconala"
)

func createExpectedData(path string) ([][]string, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	r := csv.NewReader(file)

	records, err := r.ReadAll()
	if err != nil {
		return nil, err
	}

	return records, nil
}

func TestParse(t *testing.T) {
	expected, err := createExpectedData("../../testdata/sample.csv")
	if err != nil {
		t.Fatal(err)
	}

	sampleData, err := os.Open("../../testdata/sample.csv")
	if err != nil {
		t.Fatal(err)
	}
	defer sampleData.Close()

	coconalaSalesData, err := coconala.Parse(sampleData)
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
