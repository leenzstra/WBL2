package main

import (
	"bytes"
	"testing"

	"github.com/stretchr/testify/assert"
)

var testLines = []string{
	"Lorem ipsum dolor sit amet, consectetur adipiscing elit",
	"Sed do eiusmod tempor incididunt ut labore et dolore magna aliqua.",
	"Arcu felis bibendum ut tristique et egestas.",
	"A arcu cursus vitae congue.",
	"Quis viverra nibh cras pulvinar mattis nunc sed blandit.",
	"Mattis rhoncus urna neque viverra justo nec ultrices dui sapien.",
	"Rhoncus urna neque viverra justo nec ultrices dui.",
	"Tortor posuere ac ut consequat semper viverra nam libero justo.",
	"Maecenas ultricies mi eget mauris pharetra et.",
	"Quam pellentesque nec nam aliquam sem et tortor consequat.",
	"Aliquam sem et tortor consequat id porta nibh venenatis cras.",
	"Quam vulputate dignissim suspendisse in est ante in.",
	"Risus sed vulputate odio ut enim blandit volutpat maecenas volutpat.",
	"Tellus mauris a diam maecenas congue.",
	"Dis parturient montes nascetur ridiculus.",
	"Fermentum leo vel orci porta non pulvinar neque.",
	"Ipsum consequat nisl vel pretium lectus.",
	"Nisl tincidunt eget nullam non nisi est sit amet facilisis.",
	"Leo urna molestie at elementum eu facilisis sed odio morbi.",
	"Urna neque viverra justo nec ultrices dui.",
}

func TestFixedFlag(t *testing.T) {
	// command := `go run . -n -c -C 1 -F "Tellus mauris a diam maecenas congue." .\input.txt`
	var (
		lineNum    = true
		count      = true
		context    = 1
		before     = 0
		after      = 0
		fixed      = true
		invert     = false
		ignoreCase = false
		pattern    = "Tellus mauris a diam maecenas congue."

		out = bytes.NewBufferString("")

		expected = `13. Risus sed vulputate odio ut enim blandit volutpat maecenas volutpat.
14. Tellus mauris a diam maecenas congue.
15. Dis parturient montes nascetur ridiculus.
3`
	)

	run(testLines, pattern, ignoreCase, invert, fixed, count, lineNum, before, after, context, out)

	assert.Equal(t, expected, out.String())
}

func TestDefaultWithOffsetsFlag(t *testing.T) {
	// command := `go run . -n -B 1 -A 2 congue .\input.txt`
	var (
		lineNum    = true
		count      = false
		context    = 0
		before     = 1
		after      = 2
		fixed      = false
		invert     = false
		ignoreCase = false
		pattern    = "congue."

		out = bytes.NewBufferString("")

		expected = `3. Arcu felis bibendum ut tristique et egestas.
4. A arcu cursus vitae congue.
5. Quis viverra nibh cras pulvinar mattis nunc sed blandit.
6. Mattis rhoncus urna neque viverra justo nec ultrices dui sapien.
13. Risus sed vulputate odio ut enim blandit volutpat maecenas volutpat.
14. Tellus mauris a diam maecenas congue.
15. Dis parturient montes nascetur ridiculus.
16. Fermentum leo vel orci porta non pulvinar neque.
`
	)

	run(testLines, pattern, ignoreCase, invert, fixed, count, lineNum, before, after, context, out)

	assert.Equal(t, expected, out.String())
}

func TestIgnoreCaseFlag(t *testing.T) {
	// command := `go run . -n -i "urna" .\input.txt`
	var (
		lineNum    = true
		count      = false
		context    = 0
		before     = 0
		after      = 0
		fixed      = false
		invert     = false
		ignoreCase = true
		pattern    = "urna"

		out = bytes.NewBufferString("")

		expected = `6. Mattis rhoncus urna neque viverra justo nec ultrices dui sapien.
7. Rhoncus urna neque viverra justo nec ultrices dui.
19. Leo urna molestie at elementum eu facilisis sed odio morbi.
20. Urna neque viverra justo nec ultrices dui.
`
	)

	run(testLines, pattern, ignoreCase, invert, fixed, count, lineNum, before, after, context, out)

	assert.Equal(t, expected, out.String())
}

func TestInvertFlag(t *testing.T) {
	// command := `go run . -i -v -n "leo" .\input.txt `
	var (
		lineNum    = true
		count      = false
		context    = 0
		before     = 0
		after      = 0
		fixed      = false
		invert     = true
		ignoreCase = true
		pattern    = "leo"

		out = bytes.NewBufferString("")

		expected = `1. Lorem ipsum dolor sit amet, consectetur adipiscing elit
2. Sed do eiusmod tempor incididunt ut labore et dolore magna aliqua.
3. Arcu felis bibendum ut tristique et egestas.
4. A arcu cursus vitae congue.
5. Quis viverra nibh cras pulvinar mattis nunc sed blandit.
6. Mattis rhoncus urna neque viverra justo nec ultrices dui sapien.
7. Rhoncus urna neque viverra justo nec ultrices dui.
8. Tortor posuere ac ut consequat semper viverra nam libero justo.
9. Maecenas ultricies mi eget mauris pharetra et.
10. Quam pellentesque nec nam aliquam sem et tortor consequat.
11. Aliquam sem et tortor consequat id porta nibh venenatis cras.
12. Quam vulputate dignissim suspendisse in est ante in.
13. Risus sed vulputate odio ut enim blandit volutpat maecenas volutpat.
14. Tellus mauris a diam maecenas congue.
15. Dis parturient montes nascetur ridiculus.
17. Ipsum consequat nisl vel pretium lectus.
18. Nisl tincidunt eget nullam non nisi est sit amet facilisis.
20. Urna neque viverra justo nec ultrices dui.
`
	)

	run(testLines, pattern, ignoreCase, invert, fixed, count, lineNum, before, after, context, out)

	assert.Equal(t, expected, out.String())
}