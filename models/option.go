package models

import "time"

func init() {
	Register(new(Options))
}

type Options struct {
	ID       int64
	Question string
	Answer   int64
	Desc     string
	Options  []string
	Created  time.Time
	Updated  time.Time
}
