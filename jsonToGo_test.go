package json_to_go

import "testing"

func TestParse(t *testing.T) {
	result, err := Parse(testData)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(result)
}

var (
	// https://api.github.com/repos/kumakichi/json-to-go

	// https://github.com/mholt/json-to-go/issues/71
	testData = `[{"holder1": {"type1": [["value1", "value2"], ["value1", "value2"]]}}]`
)
