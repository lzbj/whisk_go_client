package whisk

import (
	"fmt"
	"os"
	"testing"
)

func TestPropfile(t *testing.T) {
	pwd, err := os.Getwd()
	if err != nil {
		fmt.Fprintf(os.Stderr, "error get cwd %v \n", err)
	}
	pwd += "/test_data"
	var res = Propfile(pwd)
	var expected = fmt.Sprintf("%s/whisk.properties", pwd)
	if res != expected {
		t.Error(`"res" != "data/whisk.property"`)
	}
}

func TestPropsIfAvailable(t *testing.T) {
	expected := map[string]string{
		"A1":  "value1",
		"A2":  "value2",
		"A-3": "value3",
	}
	pwd, err := os.Getwd()
	if err != nil {
		fmt.Fprintf(os.Stderr, "error get cwd %v \n", err)
	}
	pwd += "/test_data"
	var res = Propfile(pwd)

	var actual = ImportPropsIfAvailable(res)
	if !Map_equal(actual, expected) {
		t.Error(`"results" != "expected"`)
	}
}
