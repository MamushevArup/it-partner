package count

import "testing"

func TestSymbol(t *testing.T) {
	tests := map[string]struct {
		input    string
		expected int
	}{
		"empty string": {
			input:    "",
			expected: 0,
		},
		"ASCII characters": {
			input:    "hello",
			expected: 5,
		},
		"non-ASCII characters": {
			input:    "привет",
			expected: 6,
		},
		"mix of ASCII and non-ASCII characters": {
			input:    "hello, привет!",
			expected: 14,
		},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			output := Symbol(tc.input)
			if output != tc.expected {
				t.Errorf("Symbol(%q) = %d; want %d", tc.input, output, tc.expected)
			}
		})
	}
}
