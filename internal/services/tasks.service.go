package services

import (
	"crm_system/internal/entity"
	"crm_system/internal/permissions"
	"errors"
)

func (Service *Services) GetEmployeeTasks(employee_id int, employee_uuid string) ([]entity.TaskType, error) {
	err := permissions.CanDo(employee_id, permissions.GetEmployeeTasks)
	if err != nil {
		return []entity.TaskType{}, err
	}

	response, err := Service.Repo.SelectAllTasksByEmployeeUUID(employee_uuid)
	if err != nil {
		return []entity.TaskType{}, err
	}

	return response, nil
}

func (Service *Services) CreateTaskToEmployee(BodyData *entity.TaskParser, claims entity.EmployeeTokenClaims) error {

	err := permissions.CanDo(claims.UserID, permissions.CreateTaskToEmployee)
	if err != nil {
		return err
	}

	// verify if this uuid exist
	_, err = Service.Repo.SelectEmployeeUUIDByUUID(BodyData.TaskToEmployeeUUID)
	if err != nil {
		return err
	}

	BodyData.TaskCreatedByEmployeeID = claims.UserID
	BodyData.IsTaskCenceled = false
	BodyData.IsTaskFinsihed = false

	err = Service.Repo.InsertNewTask(BodyData)
	if err != nil {
		return err
	}

	return nil
}

func (Service *Services) DeleteTask(employee_id int, task_id int) error {

	err := permissions.CanDo(employee_id, permissions.DeleteTask)
	if err != nil {
		return err
	}

	response, err := Service.Repo.SelectEmployeeUUIDByTaskID(task_id)
	if err != nil {
		return err
	}

	empid, err := Service.Repo.SelectEmployeeIDByUUID(response.EmployeeUUID)
	if err != nil {
		return err
	}

	if empid.EmployeeID != employee_id {
		return errors.New("you dont have permission to delete this task")
	}

	err = Service.Repo.DeleteTaskByTAskID(task_id)
	if err != nil {
		return err
	}
	return nil
}

func (Service *Services) ApplyTask(employee_id int, task_id int, TaskOperation string) error {
	err := permissions.CanDo(employee_id, permissions.ApplyTask)
	if err != nil {
		return err
	}

	response, err := Service.Repo.SelectEmployeeUUIDByTaskID(task_id)
	if err != nil {
		return err
	}

	empid, err := Service.Repo.SelectEmployeeIDByUUID(response.EmployeeUUID)
	if err != nil {
		return err
	}

	if empid.EmployeeID != employee_id {
		return errors.New("you dont have permission to apply this task")
	}

	if TaskOperation == "completed" {
		err = Service.Repo.TaskFinsih(task_id)
		if err != nil {
			return err
		}
		return nil
	}

	if TaskOperation == "canceled" {
		err = Service.Repo.CancelTask(task_id)
		if err != nil {
			return err
		}
		return nil
	}

	return errors.New("task not submitted")
}

func (Service *Services) GetAllTasks(employee_id int, limit int, offset int) ([]entity.TaskType, error) {

	err := permissions.CanDo(employee_id, permissions.GetAllTasks)
	if err != nil {
		return []entity.TaskType{}, err
	}

	response, err := Service.Repo.SelectAllTasks(limit, offset)
	if err != nil {
		return []entity.TaskType{}, err
	}

	return response, nil
}
