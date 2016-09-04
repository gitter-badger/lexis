package analysis

import "bytes"

// TokenizerWhitespace split the content into tokens every whitespace.
type TokenizerWhitespace struct{}

// Process execute the tokenizer.
func (w *TokenizerWhitespace) Process(content []byte) [][]byte {
	return bytes.Fields(content)
}

// NewTokenizerWhitespace returns a instance of TokenizerWhitespace.
func NewTokenizerWhitespace() *TokenizerWhitespace {
	return &TokenizerWhitespace{}
}
