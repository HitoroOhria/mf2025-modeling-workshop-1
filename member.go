package main

// Member は参加者
type Member struct {
	Name string
}

func NewMember(name string) *Member {
	return &Member{
		Name: name,
	}
}
