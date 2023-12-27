package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRunCut(t *testing.T) {
	var tests = []struct {
		name string
		input     string
		expected string
		err bool

		fields string
		delimeter string
		separated bool
	}{
		{
			name: "первое поле",
			input: "asd1 asd2 asd3",
			expected: "asd1", 
			err: false,
			fields: "1",
			delimeter: " ",
			separated: false,
		},
		{
			name: "нет полей",
			input: "asd1 asd2 asd3",
			expected: "asd1 asd2 asd3", 
			err: false,
			fields: "",
			delimeter: " ",
			separated: false,
		},
		{
			name: "поля выходят за границу",
			input: "asd1 asd2 asd3",
			expected: "", 
			err: true,
			fields: "0-1",
			delimeter: " ",
			separated: false,
		},
		{
			name: "ошибка в записи диапазона",
			input: "asd1 asd2 asd3",
			expected: "", 
			err: true,
			fields: "1-2-3",
			delimeter: " ",
			separated: false,
		},
		{
			name: "обратный диапазон",
			input: "asd1 asd2 asd3",
			expected: "", 
			err: true,
			fields: "2-1",
			delimeter: " ",
			separated: false,
		},
		{
			name: "несколько полей",
			input: "asd1 asd2 asd3",
			expected: "asd1 asd3", 
			err: false,
			fields: "1,3",
			delimeter: " ",
			separated: false,
		},
		{
			name: "разделитель ','",
			input: "asd1,asd2,asd3",
			expected: "asd1,asd3", 
			err: false,
			fields: "1,3",
			delimeter: ",",
			separated: false,
		},
		{
			name: "разные разделители ",
			input: "asd1 asd2,asd3",
			expected: "asd2,asd3", 
			err: false,
			fields: "2",
			delimeter: " ",
			separated: false,
		},
		{
			name: "диапазон полей",
			input: "asd1 asd2 asd3",
			expected: "asd1 asd2 asd3", 
			err: false,
			fields: "1-3",
			delimeter: " ",
			separated: false,
		},
		{
			name: "порядок полей",
			input: "asd1 asd2 asd3",
			expected: "asd2 asd1 asd3", 
			err: false,
			fields: "2,1,3",
			delimeter: " ",
			separated: false,
		},
		{
			name: "параметр -s",
			input: "asd1 asd2 asd3",
			expected: "", 
			err: false,
			fields: "2,1,3",
			delimeter: ",",
			separated: true,
		},
		{
			name: "ошибка в полях",
			input: "asd1 asd2 asd3",
			expected: "", 
			err: true,
			fields: "2a-3",
			delimeter: " ",
			separated: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r, err := runCut(tt.input, tt.fields, tt.delimeter, tt.separated)
			assert.Equal(t, tt.expected, r)
			assert.Equal(t, tt.err, err != nil)
		})
	}
}

func TestParseFieldRangeSyntaxErr(t *testing.T) {
	rng := "1_3"
	rngSep := "-"

	r, err := parseFieldRange(rng, rngSep)

	assert.Error(t, err)
	assert.Nil(t, r)
}