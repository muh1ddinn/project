package model

type Schedule struct {
	Group_id   string
	Group_type string
	Start_time string
	End_time   string
	Date       string
	Branch_id  string
	Teacher_id string
	Created_at string
	Updated_at string
}

type Getschedule struct {
	Schedule []Schedule
	Count    int64
}
