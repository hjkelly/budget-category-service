package views

import (
	"github.com/hjkelly/budget-category-service/categories"
	"github.com/hjkelly/budget-category-service/common"
	"github.com/satori/go.uuid"
	"strings"
	"time"
)

type UserCategoryInput struct {
	Name string `json:"name"`
	Type string `json:"type"`
}

func (c UserCategoryInput) ValidationErrors() error {
	errors := []common.APIError{}

	// validate: name
	if len(strings.TrimSpace(c.Name)) == 0 {
		errors = append(errors, common.APIError{
			Field:   "name",
			Message: "Must contain at least one visible character.",
		})
	}

	// validate: type
	typeIsValid := false
	for _, validType := range categories.Types {
		if validType == c.Type {
			typeIsValid = true
			break
		}
	}
	if !typeIsValid {
		errors = append(errors, common.APIError{
			Field:   "type",
			Message: "Must be one of the following: income, expense, goal",
		})
	}

	if len(errors) > 0 {
		return common.NewValidationError(errors...)
	} else {
		return nil
	}
}

func (c UserCategoryInput) AsCategory() categories.Category {
	return categories.Category{
		Name: c.Name,
		Type: c.Type,
	}
}

type UserCategoryOutput struct {
	UserCategoryInput
	ID           uuid.UUID `json:"id"`
	Created      time.Time `json:"created"`
	LastModified time.Time `json:"lastModified"`
}

func NewUserCategoryOutput(c categories.Category) UserCategoryOutput {
	return UserCategoryOutput{
	// TODO
	}
}
