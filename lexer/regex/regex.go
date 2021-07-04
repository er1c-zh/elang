package regex

import "errors"

type nfa struct {
}

func Init() {

}


var (
	ErrInvalidRunePattern = errors.New("invalid rune pattern")
)

type runeMatcher func(r rune) bool

func runeMatcherFactory(pattern string) (runeMatcher, error) {
	panic("implement me")
}
