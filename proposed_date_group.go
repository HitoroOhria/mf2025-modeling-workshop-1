package main

// ProposedDateGroup は候補日グループ (命名適当)
// 候補日ごとにメンバーと回答をまとめたもの
type ProposedDateGroup struct {
	ProposedDate ProposedDate
	Items        []*ProposedDateGroupItem
}

func NewProposedDateGroup(date ProposedDate, items []*ProposedDateGroupItem) *ProposedDateGroup {
	return &ProposedDateGroup{
		ProposedDate: date,
		Items:        items,
	}
}

func (pdg *ProposedDateGroup) Count() int64 {
	var count int64

	for _, item := range pdg.Items {
		if item.IsAttendance() {
			count++
		}
	}

	return count
}

// ProposedDateGroupItem は候補日グループのアイテム (命名適当)
// メンバーと回答の組み合わせ
type ProposedDateGroupItem struct {
	Member *Member
	Answer Answer
}

func NewProposedDateGroupItem(member *Member, answer Answer) *ProposedDateGroupItem {
	return &ProposedDateGroupItem{
		Member: member,
		Answer: answer,
	}
}

func (pdg *ProposedDateGroupItem) IsAttendance() bool {
	return pdg.Answer == AnswerAttendance
}
