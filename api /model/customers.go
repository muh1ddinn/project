package model

type Customers struct {
	Id         string     `json:"id"`
	First_name string     `json:"first_name"`
	Last_name  string     `json:"Last_name"`
	Gmail      string     `json:"gmail"`
	Phone      string     `json:"phone"`
	Is_blocked bool       `json:"is_blocked"`
	Created_at string     `json:"created_at"`
	Updated_at string     `json:"updated_at"`
	Deleted_at int        `json:"deleted_at"`
	Orders     []GetOrder `json:"orrders"`
}

type GetAllCustomersResponse struct {
	Customers []Customers `json:"customers"`
	Count     int64       `json:"count"`
}

type GetAllCustomerRequest struct {
	Search string `json:"search"`
	Page   uint64 `json:"page"`
	Limit  uint64 `json:"limit"`
}
