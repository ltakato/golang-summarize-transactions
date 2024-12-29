package core

import (
	"time"
)

func IsValidPartialISO8601(date string) bool {
	_, err := time.Parse("2006-01", date)
	return err == nil
}
