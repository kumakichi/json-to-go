package json-to-go

import "testing"

func TestParse(t *testing.T) {
	parser := NewParser()
	t.Log(parser.Parse(testData, "nodeName", true))
}

var (
	testData = `{
  "name": {"first": "Tom", "last": "Anderson"},
  "age":37,
  "score": 92.0,
  "children": ["Sara","Alex","Jack"],
  "friends": [
    {"first": "James", "last": "Murphy"},
    {"first": "Roger", "last": "Craig"}
  ]
}`
)
