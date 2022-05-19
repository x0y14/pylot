package pylot

func ToJson(code string) (string, error) {
	pyAst, err := DumpAstWithCode(code)
	if err != nil {
		return "", err
	}
	tokenizer := NewTokenizer()
	tokens, err := tokenizer.Tokenize(pyAst, true)
	if err != nil {
		return "", err
	}
	return V2(tokens), nil
}
