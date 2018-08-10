package ast

import (
	"fmt"
)

// Listen to STDIN
type Listen struct {
	Var Expression // identifier to put data into
}

func (l *Listen) String() string {
	if l.Var != nil {
		return fmt.Sprintf("Listen: %s", l.Var)
	}
	return fmt.Sprintf("Listen")
}

// Source fulfils the Expression interface
func (l *Listen) Source() []byte {
	var variable []byte
	var to []byte
	if l.Var != nil {
		to = []byte(` to `)
		variable = l.Var.Source()
	} else {
		to = []byte{}
		variable = to
	}
	return concatBytes([]byte(`Listen`), to, variable)
}
