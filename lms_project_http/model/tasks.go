package model

type Tasks struct {
	Id         string
	Lesson_id  string
	Group_id   string
	Task       string
	Score      string
	Created_at string
	Updated_at string
}

type Gettask struct {
	Task  []Tasks
	Count int64
}
