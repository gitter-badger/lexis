package analysis

import "bytes"

// TokenFilterLowercase lowercase all tokens.
type TokenFilterLowercase struct{}

// Process execute the filter.
func (filter *TokenFilterLowercase) Process(content [][]byte) [][]byte {
	for i := range content {
		content[i] = bytes.ToLower(content[i])
	}

	return content
}

// NewTokenFilterLowercase returns a instance of TokenFilterLowercase.
func NewTokenFilterLowercase() *TokenFilterLowercase {
	return &TokenFilterLowercase{}
}
