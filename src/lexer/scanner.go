package lexer

type Scanner struct {
	code     []byte
	lenght   int
	position int
	line     int
	err      func(int, int, int)
}

func (sc *Scanner) Init(code []byte, err func(int, int, int)) {
	sc.code = code
	sc.lenght = len(code)
	sc.position = 0
	sc.line = 1
	sc.err = err
}

func (sc *Scanner) NextLine() {
	sc.line++
}

func (sc *Scanner) Line() int {
	return sc.line
}

func (sc *Scanner) HasNext() bool {
	return (sc.position < sc.lenght)
}

func (sc *Scanner) Next() {
	sc.position++
}

func (sc *Scanner) Peek() byte {
	if sc.HasNext() {
		return sc.code[sc.position]
	} else {
		sc.err(sc.lenght, sc.position, sc.line)
		return 0
	}
}

func (sc *Scanner) Until(char byte) bool {
	return sc.Peek() != char
}

func (sc *Scanner) PeekEquals(char byte) bool {
	return sc.Peek() == char
}

func (sc *Scanner) PeekBetween(first byte, second byte) bool {
	return (sc.Peek() >= first && sc.Peek() <= second)
}
