package entity

import "time"

type Roledata struct {
	RoleID                  int       `json:"role_id" binding:"required"`
	RoleTitle               string    `json:"role_title" binding:"required"`
	RolePermissions         []string  `json:"role_permissions" binding:"required"`
	RoleCreatedByEmployeeID int       `json:"role_created_by_employee_id" binding:"required"`
	CreatedAt               time.Time `json:"created_at" binding:"reauired"`
}
type RolePayload struct {
	RoleTitle               string   `json:"role_title" binding:"required"`
	RolePermissions         []string `json:"role_permissions" binding:"required"`
	RoleCreatedByEmployeeID int      `json:"role_created_by_employee_id" binding:"required"`
}
type RoleApply struct {
	EmployeeID int      `json:"employee_id" binding:"required"`
	RoleTitles []string `json:"role_title" binding:"required"`
}
