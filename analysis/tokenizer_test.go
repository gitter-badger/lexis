package analysis

import "testing"

func BenchmarkTokenizer(b *testing.B) {
	b.Run("whitespace", benchmarkTokenizer(NewTokenizerWhitespace()))
	b.Run("raw", benchmarkTokenizer(NewTokenizerRaw()))
}

func benchmarkTokenizer(tokenizer Tokenizer) func(*testing.B) {
	return func(b *testing.B) {
		content := []byte("The quick brown fox jumps over the lazy dog")

		b.ReportAllocs()
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			tokenizer.Process(content)
			b.SetBytes((int64)(len(content)))
		}
	}
}
