package main

import (
	"bytes"
	"net/url"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestParseHtml(t *testing.T) {
	tests := []struct {
		name string
		html string
		err  bool
	}{
		{
			"parse ok",
			`<html><body><h1>Asd</h1><a href="/other-page">A</a></body></html>`,
			false,
		},
		{
			"parse random string",
			`<random output<`,
			false,
		},
	}

	wget := NewWget(&url.URL{}, "", 1)

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			doc, err := wget.parseHtml(bytes.NewReader([]byte(tt.html)))
			require.Equal(t, tt.err, err != nil, err)
			assert.NotNil(t, doc)
		})
	}
}
func TestRunSync(t *testing.T) {
	depth := 1
	dir := "./tmp"

	tests := []struct {
		name string
		url  string
		err  bool
	}{
		{
			"run sync ok",
			`https://go.dev/doc`,
			false,
		},
		{
			"run sync err",
			`wrong url`,
			true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			url, err := url.Parse(tt.url)
			require.NoError(t, err)

			wget := NewWget(url, dir, depth)

			err = wget.RunSync()
			if !tt.err {
				assert.DirExists(t, dir)
			}
			assert.Equal(t, tt.err, err != nil)

			os.Remove(dir)
		})
	}
}

func TestRunAsync(t *testing.T) {
	depth := 1
	dir := "./tmp"

	tests := []struct {
		name string
		url  string
		err  bool
	}{
		{
			"run async ok",
			`https://go.dev/doc`,
			false,
		},
		{
			"run async err",
			`wrong url`,
			true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			os.Remove(dir)

			url, err := url.Parse(tt.url)
			require.NoError(t, err)

			wget := NewWget(url, dir, depth)

			err = wget.RunAsync()
			assert.DirExists(t, dir)
			assert.Equal(t, tt.err, err != nil)
		})
	}
}
