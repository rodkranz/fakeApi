// Package slack contains the errors available
package slack

import "fmt"

// ErrSlackPayload error if can't convert struct to json.
type ErrSlackPayload struct {
	Err error
}

// IsErrSlackPayload returns if the error type is ErrSlackPayload or not.
func IsErrSlackPayload(err error) bool {
	_, ok := err.(ErrSlackPayload)
	return ok
}

// Error interface for error.
func (e ErrSlackPayload) Error() string {
	return fmt.Sprintf("Error slack payload: %v", e.Err.Error())
}

// ErrSlackIsNotActivated struct error
type ErrSlackIsNotActivated struct{}

// IsErrSlackIsNotActivated returns if the error type is ErrSlackIsNotActivated or not.
func IsErrSlackIsNotActivated(err error) bool {
	_, ok := err.(ErrSlackIsNotActivated)
	return ok
}

// Error interface for error.
func (e ErrSlackIsNotActivated) Error() string {
	return "Slack is not activated for notify"
}

// ErrSlackPost struct error
type ErrSlackPost struct {
	Err error
}

// IsErrSlackPost returns if the error type is ErrSlackPost or not.
func IsErrSlackPost(err error) bool {
	_, ok := err.(ErrSlackPost)
	return ok
}

// Error interface for error.
func (e ErrSlackPost) Error() string {
	return fmt.Sprintf("Error to try to post payload: %v", e.Err.Error())
}

// ErrSlackReadBody struct error
type ErrSlackReadBody struct {
	Err error
}

// IsErrSlackReadBody returns if the error type is ErrSlackReadBody or not.
func IsErrSlackReadBody(err error) bool {
	_, ok := err.(ErrSlackReadBody)
	return ok
}

// Error interface for error.
func (e ErrSlackReadBody) Error() string {
	return fmt.Sprintf("Error to try to read payload of response: %v", e.Err.Error())
}
