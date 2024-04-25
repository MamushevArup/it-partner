package reverse

import "testing"

func TestString(t *testing.T) {
	tests := map[string]struct {
		input    string
		expected string
	}{
		"empty string": {
			input:    "",
			expected: "",
		},
		"single character": {
			input:    "a",
			expected: "a",
		},
		"word": {
			input:    "hello",
			expected: "olleh",
		},
		"sentence": {
			input:    "hello world",
			expected: "dlrow olleh",
		},
		"non-ascii": {
			input:    "привет мир!",
			expected: "!рим тевирп",
		},
		"odd-len": {
			input:    "gofmt",
			expected: "tmfog",
		},
		"even-len": {
			input:    "gofunc",
			expected: "cnufog",
		},
		"same": {
			input:    "hhhhhhhh",
			expected: "hhhhhhhh",
		},
		"zero-width characters": {
			input:    "a\u200Bb\u200Bc",
			expected: "c\u200Bb\u200Ba",
		},
		"string with special characters": {
			input:    "!@#$%^&*()",
			expected: ")(*&^%$#@!",
		},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			output := String(tc.input)
			if output != tc.expected {
				t.Errorf("String(%q) = %q; want %q", tc.input, output, tc.expected)
			}
		})
	}
}
