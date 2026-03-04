//GENERATED CODE

package model

import (
	"fmt"
	"time"
)

type ISO8601Time struct {
	time.Time
}

// Список форматов, которые мы готовы принять от фронта
var layouts = []string{
	"2006-01-02",               // Только дата
	"2006-01-02T15:04:05Z",     // ISO без миллисекунд
	"2006-01-02T15:04:05.000Z", // ISO с миллисекундами (JS стандарт)
	time.RFC3339,               // Стандарт Go
}

func (t *ISO8601Time) UnmarshalText(text []byte) error {
	s := string(text)
	if s == "" || s == "null" {
		return nil
	}

	var lastErr error
	for _, layout := range layouts {
		parsedTime, err := time.Parse(layout, s)
		if err == nil {
			t.Time = parsedTime
			return nil
		}
		lastErr = err
	}
	return fmt.Errorf("invalid date format in form: %s", lastErr)
}
