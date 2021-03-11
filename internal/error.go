package internal

import (
	"errors"
	"log"
	"time"
)

var (
	// ErrRequiredToken will throw if discord bot token is empty.
	ErrRequiredToken = errors.New("required discord bot token")
	// ErrInvalidInterval will throw if interval is invalid.
	ErrInvalidInterval = errors.New("interval must be 1 or higher")
	// ErrRequiredPrefix will throw if command prefix is empty..
	ErrRequiredPrefix = errors.New("required command prefix")
	// ErrRequiredDB will throw if db connection is empty.
	ErrRequiredDB = errors.New("required database connection")
	// ErrInvalidTimezone will throw if timezone is invalid.
	ErrInvalidTimezone = errors.New("invalid timezone")
)

// HandleError to send error to log.
func HandleError(l Logger, err error) {
	if l == nil || err == nil {
		return
	}
	if errLog := l.Send("nxd-error", LogError{
		Error:     err.Error(),
		CreatedAt: time.Now(),
	}); errLog != nil {
		log.Println(errLog)
	}
}
