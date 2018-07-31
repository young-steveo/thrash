package result

// Code for exit
type Code int

// Codes
const (
	ExOK        Code = 0
	ExError     Code = 1
	ExDataError Code = 65
	ExNoInput   Code = 66
	ExSoftware  Code = 70
)
