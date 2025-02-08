package dao

import (
	"database/sql/driver"
	"encoding/json"
	"errors"

	guuid "github.com/google/uuid"
)

type User struct {
	ID          guuid.UUID      `json:"id"`
	Name        string          `json:"name"`
	Email       string          `json:"email"`
	PhoneNumber string          `json:"phone_number"`
	Password    string          `json:"password"`
	PanNumbers  PanNumbersArray `json:"pan_numbers"`
}

type PanNumbersArray []string

// Implement the Scanner interface for reading JSONB from DB
func (p *PanNumbersArray) Scan(value interface{}) error {
	if value == nil {
		*p = []string{} // Set an empty slice if NULL
		return nil
	}
	bytes, ok := value.([]byte)
	if !ok {
		return errors.New("invalid type for PanNumbers")
	}
	return json.Unmarshal(bytes, p)
}

// Implement the Valuer interface for writing JSONB to DB
func (p PanNumbersArray) Value() (driver.Value, error) {
	return json.Marshal(p)
}
