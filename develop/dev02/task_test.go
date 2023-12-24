package main

import "testing"

func TestUnpack(t *testing.T) {
	tests := []struct {
		input    string
		expected string
		err      bool
	}{
		{"a4bc2d5e", "aaaabccddddde", false},
		{"d1e5", "deeeee", false},
		{"abcd", "abcd", false},
		{"45", "", true},
		{"45a", "", true},
		{"", "", false},
		{`qwe\4\5\`, "qwe45", false},
		{`qwe\45`, "qwe44444", false},
		{`qwe\\5`, `qwe\\\\\`, false},
		{`abc2\`, `abcc`, false},
	}

	for _, test := range tests {
		t.Run(test.input, func(t *testing.T) {
			actual, err := StringUnpack(test.input)
			if actual != test.expected {
				t.Errorf("input: %s, actual: %s, expected: %s", test.input, actual, test.expected)
			}

			if (err == nil) == test.err {
				t.Errorf("error expected: %t, error actual: %s", test.err, err.Error())
			}
		})
	}
}
