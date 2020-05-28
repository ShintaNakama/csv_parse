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

	//file, err := os.Open(path)
	//if err != nil {
	//	return nil, err
	//}

	//defer file.Close()
	//reader := bufio.NewReader(file)
	//var expectData []string
	//for {
	//	line, _, err := reader.ReadLine()
	//	if err == io.EOF {
	//		break
	//	}
	//	if err != nil {
	//		return nil, err
	//	}
	//	expectData = append(expectData, string(line))
	//}

	//return expectData, nil
}

func TestParse(t *testing.T) {
	stdin := bytes.NewBufferString("../testdata/sample.csv")
	stdout := new(bytes.Buffer)

	cmd.Parse(stdin.String(), stdout)

	expected, err := createExpectedData("../testdata/expected.txt")
	if err != nil {
		t.Fatal(err)
	}

	if bytes.Compare(expected, stdout.Bytes()) != 0 {
		t.Fatal("not matched")
	}
	//scanner := bufio.NewScanner(stdout)
	//var i int
	//for scanner.Scan() {
	//	line := scanner.Text()
	//	if line != expected[i] {
	//		t.Errorf("not matched expected: %s result: %s",expected[i], line)
	//	}
	//  i++
	//}
}
