package main

// Answer は回答
// "◯" と "✗" がある
type Answer int64

const (
	AnswerAttendance Answer = iota // ◯
	AnswerAbsence                  // ✗
)
