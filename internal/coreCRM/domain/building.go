package models

import "time"

type Building struct {
	ID        int
	Title     string
	Price     float64
	Owner     User
	Rieltor   User
	CreatedAt time.Time
}
