package models

import (
	"time"
)

type Scheduler struct {
	ID       uint64    `json:"id,omitempty"`
	UserId   uint64    `json:"user_id,omitempty"`
	ClientId uint64    `json:"client_id,omitempty"`
	Date     time.Time `json:"date,omitempty"`
}
