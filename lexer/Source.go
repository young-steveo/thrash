package lexer

// Source represents a lexable slice of bytes
type Source struct {
	Data    []byte
	Length  int
	Start   int
	Current int
	Line    int
}

// FromBytes returns a Source
func FromBytes(data []byte) *Source {
	return &Source{Data: data, Length: len(data), Start: 0, Current: 0, Line: 1}
}

// IsAtEnd checks if the source is consumed.
func (s *Source) IsAtEnd() bool {
	return s.Current == s.Length
}

// Advance Current and return the previous byte
func (s *Source) Advance() byte {
	s.Current++
	return s.Data[s.Current-1]
}

// Retreat Current and return the previous byte
func (s *Source) Retreat() byte {
	s.Current--
	return s.Data[s.Current-1]
}

// Match a byte and advance if matched.
func (s *Source) Match(c byte) (byte, bool) {
	if s.IsAtEnd() || s.Data[s.Current] != c {
		return c, false
	}
	return s.Advance(), true
}

// Peek at the current byte
func (s *Source) Peek() byte {
	return s.PeekAt(0)
}

// PeekAt a byte at position ahead of current
func (s *Source) PeekAt(position int) byte {
	offset := s.Current + position
	if offset >= s.Length {
		return byte(0)
	}
	return s.Data[offset]
}

// PeekSlice returns a slice of bytes from current to count
func (s *Source) PeekSlice(count int) []byte {
	offset := s.Current + count
	if offset >= s.Length {
		return []byte{byte(0)}
	}
	return s.Data[s.Current:offset]
}
