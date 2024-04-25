package reverse

// String function reverses a string
func String(s string) string {
	// convert to the slice of runes to reverse string in utf-8 encoding
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	// convert back to string
	return string(runes)
}
