package cmd_test

import (
	"bytes"
	"io/ioutil"
	"testing"

	"github.com/ShintaNakama/csv_parse/cmd"
)

func createExpectedData(path string) ([]byte, error) {
	expectedData, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, err
	}

	return expectedData, nil
}

func TestParse(t *testing.T) {
	stdin := bytes.NewBufferString("../testdata/sample.csv")
	stdout := new(bytes.Buffer)

	cmd.Parse(stdin.String(), stdout)

	expected, err := createExpectedData("../testdata/expected.txt")
	if err != nil {
		t.Fatal(err)
	}

	if !bytes.Equal(expected, stdout.Bytes()) {
		t.Fatal("not matched")
	}
}
