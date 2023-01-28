package models

import "sync"

type Session struct {
	TournamentKey   string
	TournamentID    int64
	Questions       []int64
	Current         int
	QuestionMessage int
	Next            chan struct{}
	Stop            chan struct{}
	Mu              *sync.RWMutex
}
