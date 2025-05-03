package repository

import (
	"context"
	"crm_system/config"
	"crm_system/internal/constants"
	"crm_system/internal/entity"
	"crm_system/internal/utils"
	"errors"
	"fmt"

	"github.com/jackc/pgx/v5"
)

func (Repo *Repository) SelectAllTasksByEmployeeUUID(employee_uuid string) ([]entity.TaskType, error) {
	if config.Pool == nil {
		utils.FailedDatabaseConnection(config.Pool)
	}

	var TasksArray []entity.TaskType

	var SQL_QUERY string = `
	SELECT * FROM Tasks WHERE task_to_employee_uuid = $1
	`
	Row, err := config.Pool.Query(context.Background(), SQL_QUERY, employee_uuid)
	if err != nil {
		fmt.Println(err)
		return []entity.TaskType{}, err
	}

	for Row.Next() {
		var TaskItem entity.TaskType
		err = Row.Scan(
			&TaskItem.TskID,
			&TaskItem.TaskTitle,
			&TaskItem.TaskDescription,
			&TaskItem.TaskStartFrom,
			&TaskItem.TaskEndIn,
			&TaskItem.TaskCreatedByEmployeeID,
			&TaskItem.TaskToEmployeeUUID,
			&TaskItem.IsTaskFinsihed,
			&TaskItem.IsTaskCenceled,
			&TaskItem.CanceledReason,
			&TaskItem.FinishedAt,
			&TaskItem.CanceledAt,
			&TaskItem.CreatedAt,
			&TaskItem.TaskPriority,
		)
		if err != nil {
			fmt.Println(err)
			return []entity.TaskType{}, errors.New(constants.SOMETHING_WENT_WRONG)
		}
		TasksArray = append(TasksArray, TaskItem)
	}
	return TasksArray, nil
}

func (Repo *Repository) SelectTasksFilter() {

}

func (Repo *Repository) SelectAllTasks(limit int, offset int) ([]entity.TaskType, error) {
	if config.Pool == nil {
		utils.FailedDatabaseConnection(config.Pool)
	}

	var TasksArray []entity.TaskType

	var SQL_QUERY string = `
	SELECT * FROM Tasks LIMIT $1 OFFSET $2
	`
	Row, err := config.Pool.Query(context.Background(), SQL_QUERY, limit, offset)
	if err != nil {
		fmt.Println(err)
		return []entity.TaskType{}, err
	}

	for Row.Next() {
		var TaskItem entity.TaskType
		err = Row.Scan(
			&TaskItem.TskID,
			&TaskItem.TaskTitle,
			&TaskItem.TaskDescription,
			&TaskItem.TaskStartFrom,
			&TaskItem.TaskEndIn,
			&TaskItem.TaskCreatedByEmployeeID,
			&TaskItem.TaskToEmployeeUUID,
			&TaskItem.IsTaskFinsihed,
			&TaskItem.IsTaskCenceled,
			&TaskItem.CanceledReason,
			&TaskItem.FinishedAt,
			&TaskItem.CanceledAt,
			&TaskItem.CreatedAt,
			&TaskItem.TaskPriority,
		)
		if err != nil {
			fmt.Println(err)
			return []entity.TaskType{}, errors.New(constants.SOMETHING_WENT_WRONG)
		}
		TasksArray = append(TasksArray, TaskItem)
	}
	return TasksArray, nil
}

func (Repo *Repository) InsertNewTask(BodyData *entity.TaskParser) error {
	if config.Pool == nil {
		utils.FailedDatabaseConnection(config.Pool)
	}

	var SQL_QUERY string = `
	INSERT INTO Tasks (
	task_title,
	task_description,
	task_start_from,
	task_end_in,
	task_created_by_employee_id,
	task_to_employee_uuid,
	priority
	) values (
	 $1, $2, $3, $4, $5, $6, $7
	)
	`
	_, err := config.Pool.Exec(context.Background(), SQL_QUERY,
		&BodyData.TaskTitle, &BodyData.TaskDescription, &BodyData.TaskStartFrom, &BodyData.TaskEndIn,
		&BodyData.TaskCreatedByEmployeeID, &BodyData.TaskToEmployeeUUID, &BodyData.TaskPriority)

	if err != nil {
		fmt.Println(err)
		return errors.New(constants.SOMETHING_WENT_WRONG)
	}

	return nil
}

func (Repo *Repository) CancelTask(task_id int) error {
	if config.Pool == nil {
		utils.FailedDatabaseConnection(config.Pool)
	}

	const SQL_QUERY string = `
	UPDATE Tasks SET is_task_canceled = true, is_task_finished = false, finished_at = null,  canceled_at = now() WHERE task_id = $1;
	`
	_, err := config.Pool.Exec(context.Background(), SQL_QUERY, task_id)
	if err != nil {
		fmt.Println(err)
		return err
	}

	return nil
}

func (Repo *Repository) TaskFinsih(task_id int) error {
	if config.Pool == nil {
		utils.FailedDatabaseConnection(config.Pool)
	}

	const SQL_QUERY string = `
	UPDATE Tasks SET is_task_finished = true,is_task_canceled = false, canceled_at = null, finished_at = now() WHERE task_id = $1;
	`
	_, err := config.Pool.Exec(context.Background(), SQL_QUERY, task_id)
	if err != nil {
		fmt.Println(err)
		return err
	}

	return nil
}

func (R *Repository) SelectEmployeeUUIDByTaskID(task_id int) (entity.EmployeeUUID, error) {
	if config.Pool == nil {
		utils.FailedDatabaseConnection(config.Pool)
	}
	var EmployeeUUID entity.EmployeeUUID
	var SQL_QUERY string = `
	SELECT task_to_employee_uuid FROM Tasks WHERE task_id = $1
	`
	err := config.Pool.QueryRow(context.Background(), SQL_QUERY, task_id).Scan(&EmployeeUUID.EmployeeUUID)
	if err != nil {
		if err == pgx.ErrNoRows {
			return entity.EmployeeUUID{}, errors.New("task not exist")
		}
		return entity.EmployeeUUID{}, errors.New(constants.SOMETHING_WENT_WRONG)
	}

	return EmployeeUUID, nil
}

func (R *Repository) DeleteTaskByTAskID(task_id int) error {
	if config.Pool == nil {
		utils.FailedDatabaseConnection(config.Pool)
	}
	var SQL_QUERY string = `
	DELETE FROM Tasks WHERE task_id = $1
	`
	exe, err := config.Pool.Exec(context.Background(), SQL_QUERY, task_id)

	if exe.RowsAffected() < 1 {
		// task_id not found
		return fmt.Errorf("task_id %d not found or failed to delete", task_id)
	}
	if err != nil {
		fmt.Println(err)
		return errors.New(constants.SOMETHING_WENT_WRONG)
	}

	return nil
}
