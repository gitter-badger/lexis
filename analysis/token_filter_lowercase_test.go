package analysis

import (
	"reflect"
	"testing"
)

func TestTokenFilterLowercase(t *testing.T) {
	filter := NewTokenFilterLowercase()

	msgs := []struct {
		in  [][]byte
		out [][]byte
	}{
		{
			[][]byte{
				[]byte("ThE"),
				[]byte("quiCk"),
				[]byte("brown"),
				[]byte("fOx"),
				[]byte("jumpS"),
				[]byte("OVER"),
				[]byte("the"),
				[]byte("LAzy"),
				[]byte("doG"),
			},
			[][]byte{
				[]byte("the"),
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
		if !reflect.DeepEqual(filter.Process(msg.in), msg.out) {
			t.Error("Invalid response")
		}
	}
}

func BenchmarkTokenFilterLowercase(b *testing.B) {
	filter := NewTokenFilterLowercase()

	msgs := [][]byte{
		[]byte("ThE"),
		[]byte("quiCk"),
		[]byte("brown"),
		[]byte("fOx"),
		[]byte("jumpS"),
		[]byte("OVER"),
		[]byte("the"),
		[]byte("LAzy"),
		[]byte("doG"),
	}

	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		filter.Process(msgs)
	}
}
