package json_to_go

import "testing"

func TestParse(t *testing.T) {
	t.Log(Parse(testData))
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
