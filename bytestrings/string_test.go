package bytestrings

import "testing"

func TestSearchString(t *testing.T) {
	var searchTests = []struct {
		name string
	}{
		{"base-case"},
	}
	for _, tt := range searchTests {
		t.Run(tt.name, func(t *testing.T) {
			SearchString()
		})
	}
}

func TestModifyString(t *testing.T) {
	var modifyTests = []struct {
		name string
	}{
		{"base-case"},
	}
	for _, tt := range modifyTests {
		t.Run(tt.name, func(t *testing.T) {
			ModifyString()
		})
	}
}

func TestStringReader(t *testing.T) {
	var stringReaderTests = []struct {
		name string
	}{
		{"base-case"},
	}
	for _, tt := range stringReaderTests {
		t.Run(tt.name, func(t *testing.T) {
			StringReader()
		})
	}
}
