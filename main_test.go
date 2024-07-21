package main

import (
    "testing"
)

func TestMask(t *testing.T) {
    tests := []struct {
        input  string
        output string
    }{
        {"Here's my spammy page: http://hehefouls.netHAHAHA see you.", "Here's my spammy page: http://******************* see you."},
        {"No URL here.", "No URL here."},
        {"Visit http://example.com for more info.", "Visit http://*********** for more info."},
        {"Много ссылок в одной строке http://link1.com и http://link2.com", "Много ссылок в одной строке http://********* и http://*********"},
    }

    for _, test := range tests {
        result := mask(test.input)
        if result != test.output {
            t.Errorf("For input %q, expected %q, but got %q", test.input, test.output, result)
        }
    }
}
