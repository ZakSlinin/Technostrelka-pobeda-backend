//GENERATED CODE

package model

import (
	"fmt"
	"strings"
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

func (t *ISO8601Time) UnmarshalJSON(b []byte) error {
	s := strings.Trim(string(b), "\"") // Убираем кавычки JSON
	if s == "null" || s == "" {
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
	return fmt.Errorf("неверный формат даты: %s. Ожидается YYYY-MM-DD или ISO8601. Ошибка: %v", s, lastErr)
}
