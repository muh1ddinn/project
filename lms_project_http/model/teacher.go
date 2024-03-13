package model

type Teacher struct {
	Id         string
	Full_name  string
	Email      string
	Age        int
	Paid_sum   float64
	Status     string
	Login      string
	Password   string
	Created_at string
	Updated_at string
}

type Getteacher struct {
	Teacher []Teacher
	Count   int64
}
