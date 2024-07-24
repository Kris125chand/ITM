package main

import (
	"testing"
)

func TestMask(t *testing.T) {
	tests := []struct {
		name   string
		input  string
		output string
	}{
		{"WithURL", "Here's my spammy page: http://hehefouls.netHAHAHA see you.", "Here's my spammy page: http://******************* see you."},
		{"NoURL", "No URL here.", "No URL here."},
		{"SimpleURL", "Visit http://example.com for more info.", "Visit http://*********** for more info."},
		{"MultipleURLs", "Много ссылок в одной строке http://link1.com и http://link2.com", "Много ссылок в одной строке http://********* и http://*********"},
		{"EndSpace", "http://link1.com ", "http://********* "},
		{"EndWithURL", "http://link1.com", "http://*********"},
	}

	for _, test := range tests {
		result := mask(test.input)
		if result != test.output {
			t.Errorf("Test %s failed: For input %q, expected %q, but got %q", test.name, test.input, test.output, result)
		}
	}
}
