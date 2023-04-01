package model

import "time"

type Candidates struct {
	Id     int
	Option int
	Name   string
	//Image
	CreatedAt time.Time
	UpdatedAt time.Time

}
