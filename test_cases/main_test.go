package main

import "testing"

func TestConnection(t *testing.T) {
	if connectPostgresDB(err) {
		t.Error("Expected connection to work")
	}
}

func TestTableCalculate(t *testing.T) {
	var tests = []struct {
		input    int
		expected int
	}{
		{2, 4},
		{-1, 1},
		{0, 2},
		{99999, 100001},
	}
	for _, test := range tests {
		if output := Calculate(test.input); output != test.expected {
			t.Error("Test failed: {} inputted, {} expected, recieved: {}", test.input, test.expected, output)
		}
	}
}
