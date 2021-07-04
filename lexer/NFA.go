package lexer

import "io"

type NFA interface {
	DFA() (DFA, error)
}

type NFAFactory func (r io.ReadCloser) NFA
