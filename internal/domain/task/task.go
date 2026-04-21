package task

import "time"

type Task struct {
	ID              int64       `json:"id"`
	Title           string      `json:"title"`
	Description     string      `json:"description"`
	Status          Status      `json:"status"`
	CreatedAt       time.Time   `json:"created_at"`
	UpdatedAt       time.Time   `json:"updated_at"`
	// Новые поля
	RecurrenceType  string      `json:"recurrence_type,omitempty"`  // daily, monthly, parity, dates
	RecurrenceValue int         `json:"recurrence_value,omitempty"` // число для интервала
	Parity          string      `json:"parity,omitempty"`           // even/odd
	SpecificDates   []time.Time `json:"specific_dates,omitempty"`   // список дат
}
