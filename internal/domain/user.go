package domain

import (
	"time"
)

type User struct {
	ID           int       `db:"id"`
	Token        string    `db:"token"`
	LastActivity time.Time `db:"last_activity"`
	CreatedAt    time.Time `db:"created_at"`
}

func (u *User) UpdateLastActivity() {
	u.LastActivity = time.Now()
}
