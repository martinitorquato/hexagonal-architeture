package model

import (
	"time"
)

type User struct {
	ID        int64     `db:"id" json:"user_id"`
	Username  string    `db:"username" json:"username"`
	Email     string    `db:"email" json:"email"`
	Password  string    `db:"hash_password" json:"password,omitempty"`
	CreatedAt time.Time `db:"created_at" json:"created_at"`
	LastLogin time.Time `db:"last_login" json:"last_login"`
	Status    int8      `db:"status" json:"status"`
}
