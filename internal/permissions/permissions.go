package permissions

import (
	"crm_system/internal/entity"
	"crm_system/internal/repository"
	"errors"
	"fmt"
)

const (
	// customers permissions
	GetCustomer = "GetCustomer"

	GetAllCustomers            = "GetAllCustomers"
	DeleteCustomer             = "DeleteCustomer"
	CreateCustomer             = "CreateCustomer"
	CreateEmployee             = "CreateEmployee"
	SuspendEmployee            = "SuspendEmployee"
	CancelSuspension           = "CancelSuspension"
	BanEmployee                = "BanEmployee"
	UnBanEmployee              = "UnBanEmployee"
	GetEmployeeInformation     = "GetEmployeeInformation"
	GetAllEmployeesInformation = "GetAllEmployeesInformation"
	ChangeEmployeePosition     = "ChangeEmployeePosition"

	// emlpoyee permissions
	CreateTaskToEmployee = "CreateTaskToEmployee"
	ApplyTask            = "ApplyTask"
	DeleteTask           = "DeleteTask"
	GetTaskByID          = "GetTaskByID"
	GetEmployeeTasks     = "GetEmployeeTasks"
	GetAllTasks          = "GetAllTasks"

	// roles permissions
	CreateRole       = "CreateRole"
	GetEmployeeRoles = "GetEmployeeRoles"
	UpdateRole       = "UpdateRole"
	ApplyRole        = "ApplyRole"

	// product
	CreateProduct  = "CreateProduct"
	UpdateProduct  = "UpdateProduct"
	DeleteProduct  = "DeleteProduct"
	GetProduct     = "GetProduct"
	GetAllProducts = "GetAllProducts"
)

var Permissionlist = []string{
	GetCustomer,
	GetAllCustomers,
	DeleteCustomer,
	CreateCustomer,
	CreateEmployee,
	SuspendEmployee,
	BanEmployee,
	UnBanEmployee,
	GetEmployeeInformation,
	GetAllEmployeesInformation,
	ChangeEmployeePosition,
	CreateTaskToEmployee,
	ApplyTask,
	DeleteTask,
	GetTaskByID,
	GetEmployeeTasks,
	GetAllTasks,
	CreateProduct,
	GetProduct,
	GetAllProducts,
	DeleteProduct,
}

// operation types
const (
	CompletedTask = "completed"
	CanceledTask  = "canceled"
)

func IsRoleNotFound(data []string, systemRoles []entity.Roledata) ([]string, error) {
	var notFoundItems []string

	// Convert systemRoles to a map for fast lookup
	permissionMap := make(map[string]bool)
	for _, role := range systemRoles {
		permissionMap[role.RoleTitle] = true
	}

	for _, item := range data {
		if !permissionMap[item] {
			notFoundItems = append(notFoundItems, item)
		}
	}

	if len(notFoundItems) > 0 {
		return notFoundItems, errors.New("this role(s) not found")
	}

	return []string{}, nil
}

func IsPermissionsNotFound(data []string) ([]string, error) {
	var notFoundItems []string

	// Convert Permissionlist to a map for fast lookup
	permissionMap := make(map[string]bool)
	for _, role := range Permissionlist {
		permissionMap[role] = true
	}

	for _, item := range data {
		if !permissionMap[item] {
			notFoundItems = append(notFoundItems, item)
		}
	}

	if len(notFoundItems) > 0 {
		fmt.Println(notFoundItems)
		return notFoundItems, fmt.Errorf("this permission(s) not found in the system %s", notFoundItems)
	}

	return []string{}, nil
}

func CheckOperation(task_operation string) bool {
	found := 0
	opeartionsSlice := []string{
		CompletedTask,
		CanceledTask,
	}

	for _, item := range opeartionsSlice {
		if task_operation == item {
			found += 1
		}
	}

	return found >= 1
}

func isAllowed(employeePermissions []string, actionType string) bool {
	var foundCount = 0
	for _, permission := range employeePermissions {
		if permission == actionType {
			foundCount += 1
		}
	}

	return foundCount >= 1
}
func CanDo(employee_id int, actionType string) error {
	var repo = repository.Repository{}
	employeeRoles, err := repo.SelectEmployeeRoles(employee_id)
	if err != nil {
		return err
	}

	employeePermissions, err := repo.SelectPermissionsWhereEmployeeRoles(employeeRoles)
	if err != nil {
		return err
	}

	_, err = IsPermissionsNotFound(employeePermissions)
	if err != nil {
		return err
	}

	resault := isAllowed(employeePermissions, actionType)
	if !resault {
		return errors.New("You dont have permission to do this")
	}

	return nil
}
