package model

type Payment struct {
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

type GetPayment struct {
	Payment []Payment
	Count   int64
}
