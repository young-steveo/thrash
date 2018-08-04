package token

// Type of token
type Type int

// Token Types
const (
	Error Type = iota
	EOF
	Identifier

	Number       // DEC64
	PoeticNumber // words!
	Pronoun      // it, he, she, him, her, they, them, ze, hir, zie, zir, xe, xem, ve, ver,
	Mysterious   // undefined
	Null         // nothing, nowhere, nobody, empty, gone,
	True         // true, right, yes, ok,
	False        // false, wrong, no, lies
	String       // sequence of integers
	Plus         // +
	Minus        // -
	Times        // *
	Over         // /
	Greater      // higher, greater, bigger, stronger
	Less         // lower, less, smaller, weaker
	Than         // terminates Greater / Less
	GreaterEqual // High, Great, Big, Strong
	LessEqual    // low, little, small, weak
	As           // starts and terminates greater than or equal to
	Listen       // STDIN
	Say          // STDOUT, alias shout, whisper, scream
	If           // Condition
	Else         // Condition block
	Not          // Comparison invert `is not`
	Aint         // `is not` alias
	Is           // Poetic Literal assignment // aliases was, were
	Put          // Assignment
	Into         // Assignment
	Says         // Poetic String Literal assignment
	While        // Loop, alias Until
	Break        // alias: Break it down
	Continue     // Alias Take it to the top
	Newline      // \n
	Build        // Increment
	Up           // Terminate increment
	Knock        // Decrement
	Down         // Terminate Decrement
	Takes        // Function Definition infix
	And          // Function Definition param seperator
	Taking       // Function call infix
	Comma        // Function call param seperator
	Return       // Give back

	// Reserved words
	Maybe
	DefinitelyMaybe
)
