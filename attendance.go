package main

// Attendances は出欠の集合
type Attendances []*Attendance

// Attendance は出欠
type Attendance struct {
	Date   ProposedDate
	Member *Member
	Answer Answer
}

func NewAttendance(date ProposedDate, member *Member, answer Answer) *Attendance {
	return &Attendance{
		Date:   date,
		Member: member,
		Answer: answer,
	}
}
