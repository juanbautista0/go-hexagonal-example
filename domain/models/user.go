package models

import "github.com/google/uuid"

type User struct {
	Id             uuid.UUID `json:"id,omitempty" gorm:"type:uuid;primary_key;"`
	DocumentNumber int       `json:"document_number"`
	Name           string    `json:"name"`
	Email          string    `json:"email"`
}

func (User) GetTableName() string {
	return "users"
}

func (User) GetPreloads() []string {
	return []string{}
}

func (User) GetPrimaryKey() string {
	return "id"
}
