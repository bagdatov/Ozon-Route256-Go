package usecase

import "errors"

var (
	ErrStarted         = errors.New("session already exists")
	ErrNotStarted      = errors.New("session not started")
	ErrAnswered        = errors.New("answer was already submitted")
	ErrNotFound        = errors.New("models not found")
	ErrChgkUnavailable = errors.New("chgk question base is unavailable")
	ErrChgkNotFound    = errors.New("chgk question base doesn't have this model")
	ErrChgkBadRequst   = errors.New("chgk tournament key must not contain dots")
)
