package db

import "errors"

var (
	ErrEntityVersionConflicted = errors.New("updating entity version conflicted with the current state in database")
	ErrNoReturningResult       = errors.New("no returning result set, expect the returning result")
)
