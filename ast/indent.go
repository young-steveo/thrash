package ast

import (
	"strings"
)

var level int

func indent() string {
	return strings.Repeat(`    `, level)
}
