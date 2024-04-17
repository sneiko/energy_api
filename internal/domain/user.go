package domain

import "time"

type User struct {
	Token        string
	LastActivity time.Time
}

func (u *User) UpdateLastActivity() {
	u.LastActivity = time.Now()
}
