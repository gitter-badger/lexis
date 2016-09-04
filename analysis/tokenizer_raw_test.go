package analysis

import (
	"reflect"
	"testing"
)

func TestTokenizerRaw(t *testing.T) {
	tokenizer := NewTokenizerRaw()

	msgs := [][]byte{
		[]byte("The quick brown fox jumps over the lazy dog"),
	}

	for _, msg := range msgs {
		if !reflect.DeepEqual(tokenizer.Process(msg), [][]byte{msg}) {
			t.Error("Invalid response")
		}
	}
}
