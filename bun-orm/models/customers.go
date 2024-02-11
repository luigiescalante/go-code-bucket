package models

import (
	"github.com/uptrace/bun"
	"time"
)

type Customers struct {
	bun.BaseModel `bun:"table:customers,alias:cs"`
	ID            uint      `bun:"id,pk,autoincrement"`
	FirstName     string    `bun:"first_name"`
	LastName      string    `bun:"last_name"`
	Email         string    `bun:"email"`
	Age           uint      `bun:"age"`
	CreatedAt     time.Time `bun:"-"`
	UpdatedAt     time.Time `bun:"-"`
	DeletedAt     time.Time `bun:",soft_delete,nullzero"`
}
