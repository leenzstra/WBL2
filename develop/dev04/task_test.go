package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFindGroups(t *testing.T) {
	var tests = []struct {
		name string
		data     []string
		expected map[string][]string
	}{
		{
			"2 группы",
			[]string{"пятак", "слиток", "ПЯТКА", "пятка", "тяпка", "листок", "СтоЛИК", "листок"},
			map[string][]string{"пятак": {"пятак", "пятка", "тяпка"}, "слиток": {"листок", "слиток", "столик"}},
		},
		{
			"1 группа, 1 удаляется",
			[]string{"игорь","собака","басока"},
			map[string][]string{"собака": {"басока", "собака"}},
		},
		{
			"0 групп",
			[]string{},
			map[string][]string{},
		},
		{
			"0 групп",
			[]string{"игорь","собака"},
			map[string][]string{},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			finder := NewAnagramFinder(tt.data)
			assert.Equal(t, tt.expected, finder.FindGroups())
		})
	}
}
