// Package analysis is responsible to receive a `[]byte` content and process to
// normalize the data into a way it can be searched later.
package analysis

// CharFilter process incomming content before the tokenizer.
type CharFilter interface {
	Process([]byte) []byte
}

// Tokenizer splits the incomming content into tokens.
type Tokenizer interface {
	Process([]byte) [][]byte
}

// TokenFilter process the tokens generated from tokenizer.
type TokenFilter interface {
	Process([][]byte) [][]byte
}
