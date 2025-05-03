package entity

import (
	"database/sql"
	"time"

	"github.com/google/uuid"
)

type EmployeeData struct {
	EmployeeID           int          `json:"employee_id"`
	EmplyeeUserName      string       `json:"employee_username" binding:"required"`
	EmployeeUUID         uuid.UUID    `json:"employee_uuid"`
	EmployeePosition     string       `json:"employee_position" binding:"required"`
	EmployeeFullName     string       `json:"employee_full_name" binding:"required"`
	EmployeePhoneNumber  string       `json:"employee_phone_number" binding:"required"`
	EmployeeEmailAddress string       `json:"employee_email_address" binding:"required"`
	EmployeeRole         []string     `json:"employee_role"`
	EmployeePassword     string       `json:"employee_password" binding:"required"`
	CreatedAt            time.Time    `json:"created_at"`
	CreatedByEmployeeId  int          `json:"created_by_employee_id"`
	IsBanned             bool         `json:"is_banned"  binding:"required"`
	IsSuspended          bool         `json:"is_suspended"  binding:"required"`
	SuspensionDuration   sql.NullTime `json:"suspension_duration"`
}

type EmployeeDataPublic struct {
	EmployeeID           int          `json:"employee_id"`
	EmplyeeUserName      string       `json:"employee_username" binding:"required"`
	EmployeeUUID         uuid.UUID    `json:"employee_uuid"`
	EmployeePosition     string       `json:"employee_position" binding:"required"`
	EmployeeFullName     string       `json:"employee_full_name" binding:"required"`
	EmployeePhoneNumber  string       `json:"employee_phone_number" binding:"required"`
	EmployeeEmailAddress string       `json:"employee_email_address" binding:"required"`
	CreatedAt            time.Time    `json:"created_at"`
	CreatedByEmployeeId  int          `json:"created_by_employee_id"`
	EmployeeRole         []string     `json:"employee_role"  binding:"required"`
	IsBanned             bool         `json:"is_banned"  binding:"required"`
	IsSuspended          bool         `json:"is_suspended"  binding:"required"`
	SuspensionDuration   sql.NullTime `json:"suspension_duration"` // sql.NullTime because it may be null in db not valid timestamptz
}

type EmployeeParserStruct struct {
	EmplyeeUserName      string    `json:"employee_username" binding:"required"`
	EmployeeUUID         uuid.UUID `json:"employee_uuid"`
	EmployeeRole         []string  `json:"employee_role"`
	EmployeePosition     string    `json:"employee_position" binding:"required"`
	EmployeeFullName     string    `json:"employee_full_name" binding:"required"`
	EmployeePhoneNumber  string    `json:"employee_phone_number" binding:"required"`
	EmployeeEmailAddress string    `json:"employee_email_address" binding:"required"`
	EmployeePassword     string    `json:"employee_password" binding:"required"`
	CreatedByEmployeeId  int       `json:"created_by_employee_id"`
}

type EmployeeSignIn struct {
	EmployeeEmailAddress string `json:"employee_email_address" binding:"required"`
	EmployeePassword     string `json:"employee_password" binding:"required"`
}

type EmployeeTokenClaims struct {
	UserID               int    `json:"user_id" binding:"required"`
	EmployeeEmailAddress string `json:"employee_email_address" binding:"required"`
}

type EmployeeSuspend struct {
	EmployeeID         int       `json:"employee_id" binding:"required"`
	SuspensionDuration time.Time `json:"suspension_duration" binding:"required"`
	SuspensionReason   string    `json:"suspension_reason" binding:"required"`
}

type EmployeePosition struct {
	EmployeePosition string `json:"employee_position" binding:"required"`
}
type EmployeeUUID struct {
	EmployeeUUID string `json:"employee_uuid" binding:"required"`
}
type EmployeeID struct {
	EmployeeID int `json:"employee_id" binding:"required"`
}
type SuspensionData struct {
	IsSuspended        bool      `json:"is_suspended"  binding:"required"`
	SuspensionDuration time.Time `json:"suspension_duration" binding:"required"`
}
