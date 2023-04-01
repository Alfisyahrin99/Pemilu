package model

import "time"

type Votes struct {
	Id int
	ParticipantId int
	CandidateId int
	CreateAt time.Time
	UpdateAt time.Time
}