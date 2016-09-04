package analysis

import (
	"reflect"
	"testing"
)

func TestTokenizerWhitespace(t *testing.T) {
	tokenizer := NewTokenizerWhitespace()

	msgs := []struct {
		in  []byte
		out [][]byte
	}{
		{
			[]byte("The quick brown fox jumps over the lazy dog"),
			[][]byte{
				[]byte("The"),
				[]byte("quick"),
				[]byte("brown"),
				[]byte("fox"),
				[]byte("jumps"),
				[]byte("over"),
				[]byte("the"),
				[]byte("lazy"),
				[]byte("dog"),
			},
		},
	}

	for _, msg := range msgs {
		if !reflect.DeepEqual(tokenizer.Process(msg.in), msg.out) {
			t.Error("Invalid response")
		}
	}
}
