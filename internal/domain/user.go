package domain

import (
	"time"

	"github.com/uptrace/bun"
)

type User struct {
	bun.BaseModel `bun:"table:users"`

	ID           int       `bun:"id,pk"`
	Token        string    `bun:"token"`
	LastActivity time.Time `bun:"last_activity"`
}

func (u *User) UpdateLastActivity() {
	u.LastActivity = time.Now()
}
