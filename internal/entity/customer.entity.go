package entity

import (
	"time"

	"github.com/google/uuid"
)

type CustomerData struct {
	CustomerID                  int       `json:"customer_id"`
	CustomerUsername            string    `json:"customer_username"`
	CustomerUUID                string    `json:"customer_uuid"`
	CustomerPosition            string    `json:"customer_position"`
	CustomerFullName            string    `json:"customer_full_name"`
	CustomerPhoneNumber         string    `json:"customer_phone_number"`
	CustomerEmailAddress        string    `json:"customer_email_address"`
	CustomerCreatedByEmployeeID int       `json:"customer_created_by_employee_id"`
	CreatedAt                   time.Time `json:"created_at"`
}

type CustomerDataParserType struct {
	CustomerUsername            string    `json:"customer_username"`
	CustomerUUID                uuid.UUID `json:"customer_uuid"`
	CustomerPosition            string    `json:"customer_position"`
	CustomerFullName            string    `json:"customer_full_name"`
	CustomerPhoneNumber         string    `json:"customer_phone_number"`
	CustomerEmailAddress        string    `json:"customer_email_address"`
	CustomerCreatedByEmployeeID int       `json:"customer_created_by_employee_id"`
}
