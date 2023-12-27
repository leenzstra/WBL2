package main

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestOrChannel(t *testing.T) {
	t.Run("нет каналов", func(t *testing.T) {
		maxDelta := 200 * time.Millisecond
		s := time.Now()

		<- or()

		assert.LessOrEqual(t, time.Since(s).Milliseconds(), maxDelta)
	})
	t.Run("1 канал", func(t *testing.T) {
		maxDelta := 200 * time.Millisecond
		expectedTime := 1 * time.Second
		s := time.Now()

		<- or(
			sig(1 * time.Second),
		)

		assert.LessOrEqual(t, time.Since(s).Milliseconds(), expectedTime + maxDelta)
	})

	t.Run("N каналов", func(t *testing.T) {
		maxDelta := 200 * time.Millisecond
		expectedTime := 2 * time.Second
		s := time.Now()

		<- or(
			sig(1 * time.Hour),
			sig(2 * time.Second),
			sig(2 * time.Second),
			sig(2 * time.Second),
			sig(1 * time.Minute),
		)

		assert.LessOrEqual(t, time.Since(s).Milliseconds(), expectedTime + maxDelta)
	})
}