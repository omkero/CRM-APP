package entity

import (
	"database/sql"
	"time"

	"github.com/google/uuid"
)

type TaskType struct {
	TskID                   int            `json:"task_id" binding:"required"`
	TaskTitle               string         `json:"task_title" binding:"required"`
	TaskDescription         string         `json:"task_description" binding:"required"`
	TaskStartFrom           time.Time      `json:"task_start_from" binding:"required"`
	TaskEndIn               time.Time      `json:"task_end_in" binding:"required"`
	TaskCreatedByEmployeeID int            `json:"task_created_by_employee_id" binding:"required"`
	TaskToEmployeeUUID      uuid.UUID      `json:"task_to_employee_uuid" binding:"required"`
	IsTaskFinsihed          bool           `json:"is_task_finished" binding:"required"`
	IsTaskCenceled          bool           `json:"is_task_canceled" binding:"required"`
	CanceledReason          sql.NullString `json:"canceled_reason" binding:"required"`
	FinishedAt              sql.NullTime   `json:"finished_at" binding:"required"`
	CanceledAt              sql.NullTime   `json:"canceled_at" binding:"required"`
	CreatedAt               time.Time      `json:"created_at" binding:"required"`
	TaskPriority            string         `json:"task_priority" binding:"required"`
}

type TaskParser struct {
	TaskTitle               string    `json:"task_title" binding:"required"`
	TaskDescription         string    `json:"task_description" binding:"required"`
	TaskStartFrom           time.Time `json:"task_start_from" binding:"required"`
	TaskEndIn               time.Time `json:"task_end_in" binding:"required"`
	TaskCreatedByEmployeeID int       `json:"task_created_by_employee_id"`
	TaskToEmployeeUUID      string    `json:"task_to_employee_uuid" binding:"required"`
	IsTaskFinsihed          bool      `json:"is_task_finished" binding:"required"`
	IsTaskCenceled          bool      `json:"is_task_canceled" binding:"required"`
	TaskPriority            string    `json:"task_priority" binding:"required"`
}

type TaskApply struct {
	TaskOperation string `json:"task_operation" binding:"required"`
	TaskID        int    `json:"task_id" binding:"required"`
}
type TaskID struct {
	TaskID int `json:"task_id" binding:"required"`
}
