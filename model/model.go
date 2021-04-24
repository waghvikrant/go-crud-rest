package model

import (
	"time"
)

type Employee struct {
	ID          uint64    `gorm:"primarykey;auto_increment;not_null" json:"id"`
	CreatedAt   time.Time `json:"createdAt"`
	UpdatedAt   time.Time `json:"updatedAt"`
	FirstName   string    `json:"firstName"`
	LastName    string    `json:"lastName"`
	Designation string    `json:"designation"`
	Salary      float32   `json:"salary"`
}
