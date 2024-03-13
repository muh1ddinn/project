package model

type Student struct {
	Id         string
	Full_name  string
	Email      string
	Age        int
	Paid_sum   float64
	Status     string
	Login      string
	Password   string
	Group_id   string
	Created_at string
	Updated_at string
}

type GetAllstudent struct {
	Student []Student
	Count   int64
}

type GetAllstudentRequest struct {
	Search string `json:"search"`
	Page   uint64 `json:"page"`
	Limit  uint64 `json:"limit"`
}
