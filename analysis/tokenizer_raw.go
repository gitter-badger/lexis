package analysis

// TokenizerRaw is a tokenizer that always returns the content without
// changing it.
type TokenizerRaw struct{}

// Process execut the tokenizer.
func (r *TokenizerRaw) Process(content []byte) [][]byte {
	return [][]byte{content}
}

// NewTokenizerRaw returns a instance of TokenizerRaw.
func NewTokenizerRaw() *TokenizerRaw {
	return &TokenizerRaw{}
}
