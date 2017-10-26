package models

import "fmt"

// APIError is an error that contains information
// for API Repsonses
type APIError struct {
	Message, DeveloperMessage string
	Code                      int
}

func (e APIError) Error() string {
	return e.Message
}

// FullError returns a string of all the information for the APIError
func (e APIError) FullError() string {
	return fmt.Sprintf("Message: %s, DeveloperMessage: %s, Code: %d", e.Message, e.DeveloperMessage, e.Code)
}
