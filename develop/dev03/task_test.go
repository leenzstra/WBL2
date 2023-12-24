package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

const sep = " "

func TestNumericSort(t *testing.T) {
	data := [][]string{
		splitRow("1 2 3 4 5 6", sep),
		splitRow("2 8 13 -1 7 4", sep),
		splitRow("-12 -125 623 -53 3 85", sep),
		splitRow("0 64 2 2 -4 -7", sep),
		splitRow("8 4 2 6 2 1", sep),
	}
	expected := [][]string{
		{"0","64","2","2","-4","-7"},
		{"8","4","2","6","2","1"},
		{"1","2","3","4","5","6"},
		{"2","8","13","-1","7","4"},
		{"-12","-125","623","-53","3","85"},
	}
	col := 2

	table := NewDataTable(data)
	table.SetSortStrategy(NumericSortStrategy{t: table, col: col})
	table.Sort()

	assert.Equal(t, expected, table.Data())
}

func TestStringSort(t *testing.T) {
	data := [][]string{
		splitRow("abs 2 3 4 5 6", sep),
		splitRow("hrh 8 13 -1 7 4", sep),
		splitRow("zTgh -125 623 -53 3 85", sep),
		splitRow("ahbs 64 2 2 -4 -7", sep),
		splitRow("rhhawd 4 2 6 2 1", sep),
	}
	expected := [][]string{
		{"abs","2","3","4","5","6"},
		{"ahbs","64","2","2","-4","-7"},
		{"hrh","8","13","-1","7","4"},
		{"rhhawd","4","2","6","2","1"},
		{"zTgh","-125","623","-53","3","85"},
	}
	col := 0

	table := NewDataTable(data)
	table.SetSortStrategy(StringSortStrategy{t: table, col: col})
	table.Sort()

	assert.Equal(t, expected, table.Data())
}

func TestMonthSort(t *testing.T) {
	data := [][]string{
		splitRow("January 2 3 4 5 6", sep),
		splitRow("September 8 13 -1 7 4", sep),
		splitRow("March -125 623 -53 3 85", sep),
		splitRow("June 64 2 2 -4 -7", sep),
		splitRow("February 4 2 6 2 1", sep),
	}
	expected := [][]string{
		{"January","2","3","4","5","6"},
		{"February","4","2","6","2","1"},
		{"March","-125","623","-53","3","85"},
		{"June","64","2","2","-4","-7"},
		{"September","8","13","-1","7","4"},
	}
	col := 0

	table := NewDataTable(data)
	table.SetSortStrategy(MonthSortStrategy{t: table, col: col})
	table.Sort()

	assert.Equal(t, expected, table.Data())
}

func TestNumberSufSort(t *testing.T) {
	data := [][]string{
		splitRow("1b 2 3 4 5 6", sep),
		splitRow("2aa 8 13 -1 7 4", sep),
		splitRow("2cnr -125 623 -53 3 85", sep),
		splitRow("1z 64 2 2 -4 -7", sep),
		splitRow("-5gr 4 2 6 2 1", sep),
	}
	expected := [][]string{
		{"-5gr","4","2","6","2","1"},
		{"1b","2","3","4","5","6"},
		{"1z","64","2","2","-4","-7"},
		{"2aa","8","13","-1","7","4"},
		{"2cnr","-125","623","-53","3","85"},
	}
	col := 0

	table := NewDataTable(data)
	table.SetSortStrategy(NumberSufSortStrategy{t: table, col: col})
	table.Sort()

	assert.Equal(t, expected, table.Data())
}

func TestUnique(t *testing.T) {
	data := [][]string{
		splitRow("1 8 13 -1 7 4", sep),
		splitRow("1 8 13 -1 7 4", sep),
		splitRow("2cnr -125 623 -53 3 85", sep),
		splitRow("1z 64 2 2 -4 -7", sep),
		splitRow("-5gr 4 2 6 2 1", sep),
	}
	expected := [][]string{
		{"1","8","13","-1","7","4"},
		{"2cnr","-125","623","-53","3","85"},
		{"1z","64","2","2","-4","-7"},
		{"-5gr","4","2","6","2","1"},
	}

	table := NewDataTable(data)
	table.Unique()

	assert.Equal(t, expected, table.Data())
}

func TestReverseNumericSort(t *testing.T) {
	data := [][]string{
		splitRow("1 2 3 4 5 6", sep),
		splitRow("2 8 13 -1 7 4", sep),
		splitRow("-12 -125 623 -53 3 85", sep),
		splitRow("0 64 2 2 -4 -7", sep),
		splitRow("8 4 2 6 2 1", sep),
	}
	expected := [][]string{
		{"-12","-125","623","-53","3","85"},
		{"2","8","13","-1","7","4"},
		{"1","2","3","4","5","6"},
		{"0","64","2","2","-4","-7"},
		{"8","4","2","6","2","1"},
	}
	col := 2

	table := NewDataTable(data)
	table.SetSortStrategy(NumericSortStrategy{t: table, col: col})
	table.ReverseSort()

	assert.Equal(t, expected, table.Data())
}

func TestIsSorted(t *testing.T) {
	data := [][]string{
		splitRow("1 2 3 4 5 6", sep),
		splitRow("2 8 13 -1 7 4", sep),
		splitRow("-12 -125 623 -53 3 85", sep),
		splitRow("0 64 2 2 -4 -7", sep),
		splitRow("8 4 2 6 2 1", sep),
	}

	col := 2

	table := NewDataTable(data)
	table.SetSortStrategy(NumericSortStrategy{t: table, col: col})
	table.Sort()

	assert.True(t, table.IsSorted())
}

func TestIsNotSorted(t *testing.T) {
	data := [][]string{
		splitRow("1 2 3 4 5 6", sep),
		splitRow("2 8 13 -1 7 4", sep),
		splitRow("-12 -125 623 -53 3 85", sep),
		splitRow("0 64 2 2 -4 -7", sep),
		splitRow("8 4 2 6 2 1", sep),
	}

	col := 2

	table := NewDataTable(data)
	table.SetSortStrategy(NumericSortStrategy{t: table, col: col})

	assert.False(t, table.IsSorted())
}