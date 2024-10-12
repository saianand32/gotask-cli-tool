package models

import "time"

type Item struct {
	Id          string
	Group       string
	Task        string
	Done        bool
	CreatedAt   time.Time
	CompletedAt time.Time
}
