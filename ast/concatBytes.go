package ast

func concatBytes(a, b, c []byte) []byte {
	aLength := len(a)
	bLength := len(b)
	cLength := len(c)

	result := make([]byte, aLength+bLength+cLength+2)
	for i, byt := range a {
		result[i] = byt
	}
	result[aLength] = byte(' ')

	for i, byt := range b {
		result[i+bLength+1] = byt
	}
	result[aLength+bLength+1] = byte(' ')

	for i, byt := range b {
		result[i+aLength+bLength+2] = byt
	}

	return result
}
