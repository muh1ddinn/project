package model

type Customers struct {
	Id         string
	First_name string
	Last_name  string
	Gmail      string
	Phone      string
	Is_blocked bool
	Created_at string
	Updated_at string
	Deleted_at int
}

type GetAllCustomersResponse struct {
	Customers []Customers `json:"customers"`
	Count     int16       `json:"count"`
}

type GetAllCustomerRequest struct {
	Search string `json:"search"`
	Page   uint64 `json:"page"`
	Limit  uint64 `json:"limit"`
}
