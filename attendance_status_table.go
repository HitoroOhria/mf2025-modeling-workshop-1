package main

// AttendanceStatusTable は出欠表
type AttendanceStatusTable struct {
	Attendances Attendances
}

func NewAttendanceStatusTable(attendances Attendances) *AttendanceStatusTable {
	// TODO: 出欠の日付・参加者が被っていないことのバリデーション

	return &AttendanceStatusTable{
		Attendances: attendances,
	}
}

func (ast *AttendanceStatusTable) Decide() ProposedDate {
	proposedDateGroups := ast.aggregateByProposedDate()

	proposedDateContMap := make(map[ProposedDate]int64, len(proposedDateGroups))
	for _, pdg := range proposedDateGroups {
		proposedDateContMap[pdg.ProposedDate] = pdg.Count()
	}
	var mostSuitableDate ProposedDate // TODO 初期値の対応する
	for date, count := range proposedDateContMap {
		currentCount := proposedDateContMap[mostSuitableDate]

		if currentCount < count {
			mostSuitableDate = date
		}
	}

	return mostSuitableDate
}

func (ast *AttendanceStatusTable) aggregateByProposedDate() []*ProposedDateGroup {
	proposedDateMap := make(map[ProposedDate][]*ProposedDateGroupItem, len(ast.Attendances))
	for _, attendance := range ast.Attendances {
		pdgi := NewProposedDateGroupItem(attendance.Member, attendance.Answer)

		if pdgis, ok := proposedDateMap[attendance.Date]; ok {
			proposedDateMap[attendance.Date] = append(pdgis, pdgi)
		} else {
			proposedDateMap[attendance.Date] = []*ProposedDateGroupItem{pdgi}
		}
	}

	proposedDateGroups := make([]*ProposedDateGroup, 0, len(proposedDateMap))
	for date, item := range proposedDateMap {
		proposedDateGroups = append(proposedDateGroups, NewProposedDateGroup(date, item))
	}

	return proposedDateGroups
}
