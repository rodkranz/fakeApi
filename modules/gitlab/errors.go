// Package gitlab contains the errors available
package gitlab

import "fmt"

// ErrGitLabUnmarshal error if can't possible unmarshal gitHub input slack-bot
type ErrGitLabUnmarshal struct {
	Err error
}

// IsErrGitLabUnmarshal returns if the error type is ErrGitLabUnmarshal or not.
func IsErrGitLabUnmarshal(err error) bool {
	_, ok := err.(ErrGitLabUnmarshal)
	return ok
}

// Error interface for error.
func (e ErrGitLabUnmarshal) Error() string {
	return fmt.Sprintf("Error to parse gitlab payload: %v", e.Err.Error())
}
