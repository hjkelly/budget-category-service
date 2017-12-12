package categories

import (
	"github.com/satori/go.uuid"
	"time"
)

var Types = []string{
	"income",
	"expense",
	"goal",
}

type Category struct {
	Name         string
	Type         string
	ID           string
	Created      time.Time
	LastModified time.Time
	UserID       string
}

func NewCategory(name, catType, userID string) Category {
	now := time.Now()
	return Category{
		ID:           uuid.NewV4().String(),
		Name:         name,
		Type:         catType,
		UserID:       userID,
		Created:      now,
		LastModified: now,
	}
}

func (c Category) IsValid() (bool, error) {
	return true, nil // TODO
}
