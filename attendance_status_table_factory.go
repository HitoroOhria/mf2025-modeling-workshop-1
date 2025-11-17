package main

import "time"

func CreateSpecifiedAttendanceStatusTable() *AttendanceStatusTable {
	december3 := NewProposedDate(mustTimeInJST("2025-12-03T00:00:00+09:00"))
	december7 := NewProposedDate(mustTimeInJST("2025-12-07T00:00:00+09:00"))
	december8 := NewProposedDate(mustTimeInJST("2025-12-08T00:00:00+09:00"))

	sato := NewMember("佐藤")
	suzuki := NewMember("鈴木")
	takahashi := NewMember("高橋")

	attendances := Attendances{
		NewAttendance(december3, sato, AnswerAttendance),
		NewAttendance(december7, sato, AnswerAttendance),
		NewAttendance(december8, sato, AnswerAbsence),
		NewAttendance(december3, suzuki, AnswerAttendance),
		NewAttendance(december7, suzuki, AnswerAttendance),
		NewAttendance(december8, suzuki, AnswerAttendance),
		NewAttendance(december3, takahashi, AnswerAbsence),
		NewAttendance(december7, takahashi, AnswerAttendance),
		NewAttendance(december8, takahashi, AnswerAttendance),
	}

	return NewAttendanceStatusTable(attendances)
}

func mustTimeInJST(value string) time.Time {
	t, err := time.Parse(time.RFC3339, value)
	if err != nil {
		panic(err)
	}

	return t
}
