package utils

import (
	"crm_system/internal/constants"
	"errors"
	"log"

	"github.com/jackc/pgx/v5/pgxpool"
)

func FailedDatabaseConnection(Pool *pgxpool.Pool) {
	log.Fatal("[ERROR] cannot establish connection into DB")
}

func DefaultDatabaseError() error {
	return errors.New(constants.SOMETHING_WENT_WRONG)
}
