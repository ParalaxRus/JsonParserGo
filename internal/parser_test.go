package parser

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestParse(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected string
	}{
		{"empty object", "{}", "{}"},
		{"empty array", "[]", "[]"},
		{"array of numbers", "[ 1, 2,3 ]", "[1,2,3]"},
		{"array of strings", "[\"a\",\"b\",	\"c\"]", "[\"a\",\"b\",\"c\"]"},
		{"array of bools", "[  true,   false, 	true 	 ]", "[true,false,true]"},
		{
			"custom object",
			"{ \"name\": \"Jane Doe\",  \"age\": 30 ,  \"balance\": -50.2}",
			"{\"age\":30,\"balance\":-50.2,\"name\":\"Jane Doe\"}",
		},
		{
			"custom object with nested array",
			"{ \"name\": \"Jane Doe\",  \"age\": 30 ,  \"balance\": -50.2, \"pay\" : [-100.3, 50,  10000 ]}",
			"{\"age\":30,\"balance\":-50.2,\"name\":\"Jane Doe\",\"pay\":[-100.3,50,10000]}",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			parser := NewParser()
			node := parser.Parse(test.input)
			require.Equal(t, test.expected, node.String(true))
		})
	}
}
