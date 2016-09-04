package analysis

// Analyzer holds all the logic to process a given content.
type Analyzer struct {
	CharFilters  []CharFilter
	Tokenizer    Tokenizer
	TokenFilters []TokenFilter
}

// Process execute the whole chain of processing to generate the tokens.
func (analyzer *Analyzer) Process(content []byte) [][]byte {
	content = analyzer.execCharFilters(content)
	tokens := analyzer.Tokenizer.Process(content)
	return analyzer.execTokenFilters(tokens)
}

func (analyzer *Analyzer) execCharFilters(content []byte) []byte {
	if analyzer.CharFilters == nil {
		return content
	}

	for _, filter := range analyzer.CharFilters {
		content = filter.Process(content)
	}

	return content
}

func (analyzer *Analyzer) execTokenFilters(tokens [][]byte) [][]byte {
	if analyzer.TokenFilters == nil {
		return tokens
	}

	for _, filter := range analyzer.TokenFilters {
		tokens = filter.Process(tokens)
	}

	return tokens
}
