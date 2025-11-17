package main

import "time"

// ProposedDate は候補日
type ProposedDate = time.Time

func NewProposedDate(t time.Time) ProposedDate {
	return ProposedDate(t)
}
