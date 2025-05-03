package utils

import (
	"database/sql"
	"errors"
	"time"
)

func IsSuspensionCompleted(targetTime sql.NullTime) (bool, error) {
	if !targetTime.Valid {
		// No suspension time means it's not suspended or not applicable
		return true, errors.New("suspension_duration is not valid try again later")
	}

	timeNow := time.Now().UnixNano()
	parsedTargetTime := targetTime.Time.UnixNano()

	// Suspension is completed if the target time is in the past
	return parsedTargetTime <= timeNow, nil
}
