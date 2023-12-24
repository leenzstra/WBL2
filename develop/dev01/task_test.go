package main

import (
	"testing"
	"time"
)

func TestGetTimeOk(t *testing.T) {
	got, err := GetTime(defaultNtpAddr)
	expected := time.Now()
	maxOffsetSec := float64(5)

	if err != nil {
		t.Error("TestGetTimeOK", err)
	}

	offset := got.Sub(expected).Abs()
	if offset.Seconds() > maxOffsetSec {
		t.Error("TestGetTimeOK big offset", offset)
	}
}

func TestGetTimeError(t *testing.T) {
	_, err := GetTime("random addr")

	if err == nil {
		t.Error("TestGetTimeOK", err)
	}
}