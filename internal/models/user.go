package models

import (
	"time"

	"github.com/google/uuid"
)

type User struct {
	ID          uuid.UUID `db:"id"`
	Username    string    `db:"username"`
	Password    string    `db:"password"`
	FullName    string    `db:"full_name"`
	Email       string    `db:"email"`
	Nim         string    `db:"nim"`
	PhoneNumber string    `db:"phone_number"`
	CreatedAt   time.Time `db:"created_at"`
	UpdatedAt   time.Time `db:"updated_at"`
}
