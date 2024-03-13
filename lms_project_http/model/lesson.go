package model

type Lesson struct {
	ID          string
	Schedule_id string
	Group_id    string
	From        string
	To          string
	Theme       string
	Created_at  string
	Updated_at  string
}

type GEtlesson struct {
	Lesson []Lesson
	Count  int64
}
