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
	"github.com/jackc/pgx/v5/pgconn"
)

func (R *Repository) InsertEmployee(EmployeeData *entity.EmployeeParserStruct) error {
	if config.Pool == nil {
		utils.FailedDatabaseConnection(config.Pool)
	}

	var SQL_QUERY string = `
	INSERT INTO Employee (
	employee_username,
    employee_uuid,
    employee_position,
    employee_full_name,
    employee_phone_number,
    employee_email_address,
    employee_password,
    employee_role,
    created_by_employee_id
	) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9)
	`
	_, err := config.Pool.Exec(context.Background(), SQL_QUERY,
		&EmployeeData.EmplyeeUserName, &EmployeeData.EmployeeUUID, &EmployeeData.EmployeePosition,
		&EmployeeData.EmployeeFullName, &EmployeeData.EmployeePhoneNumber, &EmployeeData.EmployeeEmailAddress,
		&EmployeeData.EmployeePassword, &EmployeeData.EmployeeRole, &EmployeeData.CreatedByEmployeeId)

	if err != nil {
		var pgErr *pgconn.PgError
		if errors.As(err, &pgErr) {
			if pgErr.Code == `23505` {
				return errors.New(constants.USER_EXIST)
			}
		}

		return utils.DefaultDatabaseError()
	}

	return nil
}

func (R *Repository) SelectEmployeeByEmail(email_address string) (entity.EmployeeData, error) {
	var SelectedEmployee entity.EmployeeData

	var SQL_QUERY string = `
	SELECT * FROM Employee WHERE employee_email_address = $1
	`
	err := config.Pool.QueryRow(context.Background(), SQL_QUERY, email_address).Scan(
		&SelectedEmployee.EmployeeID,
		&SelectedEmployee.EmplyeeUserName,
		&SelectedEmployee.EmployeeUUID,
		&SelectedEmployee.EmployeePosition,
		&SelectedEmployee.EmployeeFullName,
		&SelectedEmployee.EmployeePhoneNumber,
		&SelectedEmployee.EmployeeEmailAddress,
		&SelectedEmployee.EmployeePassword,
		&SelectedEmployee.CreatedAt,
		&SelectedEmployee.CreatedByEmployeeId,
		&SelectedEmployee.EmployeeRole,
		&SelectedEmployee.IsBanned,
		&SelectedEmployee.IsSuspended,
		&SelectedEmployee.SuspensionDuration,
	)

	if err != nil {
		if err == pgx.ErrNoRows {
			return entity.EmployeeData{}, errors.New(constants.USER_NOT_FOUND)
		}
		fmt.Println(err)
		return entity.EmployeeData{}, errors.New(constants.SOMETHING_WENT_WRONG)
	}

	return SelectedEmployee, nil
}

func (R *Repository) SelectEmployeeByID(employee_id int) (entity.EmployeeDataPublic, error) {
	var SelectedEmployee entity.EmployeeDataPublic

	var SQL_QUERY string = `
	SELECT 
	
	employee_id,
	employee_username,
	employee_uuid,
	employee_position,
	employee_full_name,
	employee_phone_number,
	employee_email_address,
	created_at,
	created_by_employee_id,
	employee_role,
	is_banned,
	is_suspended,
	suspension_duration
	
	FROM Employee WHERE employee_id = $1
	`
	err := config.Pool.QueryRow(context.Background(), SQL_QUERY, employee_id).Scan(
		&SelectedEmployee.EmployeeID,
		&SelectedEmployee.EmplyeeUserName,
		&SelectedEmployee.EmployeeUUID,
		&SelectedEmployee.EmployeePosition,
		&SelectedEmployee.EmployeeFullName,
		&SelectedEmployee.EmployeePhoneNumber,
		&SelectedEmployee.EmployeeEmailAddress,
		&SelectedEmployee.CreatedAt,
		&SelectedEmployee.CreatedByEmployeeId,
		&SelectedEmployee.EmployeeRole,
		&SelectedEmployee.IsBanned,
		&SelectedEmployee.IsSuspended,
		&SelectedEmployee.SuspensionDuration)

	if err != nil {
		fmt.Println(err)
		if err == pgx.ErrNoRows {
			return entity.EmployeeDataPublic{}, errors.New(constants.USER_NOT_FOUND)
		}
		return entity.EmployeeDataPublic{}, errors.New(constants.SOMETHING_WENT_WRONG)
	}

	return SelectedEmployee, nil
}

func (R *Repository) SelectAllEmployees(limit int, offset int) ([]entity.EmployeeDataPublic, int, error) {
	var SelectedEmployeesData []entity.EmployeeDataPublic
	var totalEmployees int = 0

	var SQL_QUERY string = `
	SELECT 
	
	employee_id,
	employee_username,
	employee_uuid,
	employee_position,
	employee_full_name,
	employee_phone_number,
	employee_email_address,
	employee_role,
	created_at,
	created_by_employee_id,
	is_banned,
	is_suspended,
	suspension_duration
	
	FROM Employee LIMIT $1 OFFSET $2
	`
	var SQL_COUNT string = `
	SELECT COUNT(*) FROM employee
	`

	Rows, err := config.Pool.Query(context.Background(), SQL_QUERY, limit, offset)

	if err != nil {
		fmt.Println(err)
		if err == pgx.ErrNoRows {
			return []entity.EmployeeDataPublic{}, 0, errors.New(constants.USER_NOT_FOUND)
		}
		return []entity.EmployeeDataPublic{}, 0, errors.New(constants.SOMETHING_WENT_WRONG)
	}

	for Rows.Next() {
		var temp entity.EmployeeDataPublic
		err = Rows.Scan(
			&temp.EmployeeID,
			&temp.EmplyeeUserName,
			&temp.EmployeeUUID,
			&temp.EmployeePosition,
			&temp.EmployeeFullName,
			&temp.EmployeePhoneNumber,
			&temp.EmployeeEmailAddress,
			&temp.EmployeeRole,
			&temp.CreatedAt,
			&temp.CreatedByEmployeeId,
			&temp.IsBanned,
			&temp.IsSuspended,
			&temp.SuspensionDuration,
		)
		if err != nil {
			return []entity.EmployeeDataPublic{}, 0, errors.New(constants.SOMETHING_WENT_WRONG)
		}

		SelectedEmployeesData = append(SelectedEmployeesData, temp)
	}
	err = config.Pool.QueryRow(context.Background(), SQL_COUNT).Scan(&totalEmployees)
	if err != nil {
		if err == pgx.ErrNoRows {
			return []entity.EmployeeDataPublic{}, 0, fmt.Errorf("employee not found")
		}
		return []entity.EmployeeDataPublic{}, 0, fmt.Errorf(constants.SOMETHING_WENT_WRONG)
	}

	return SelectedEmployeesData, totalEmployees, nil
}

func (R *Repository) SelectEmployeePosition(employee_id int) (entity.EmployeePosition, error) {
	if config.Pool == nil {
		utils.FailedDatabaseConnection(config.Pool)
	}
	var EmployeePosition entity.EmployeePosition
	var SQL_QUERY string = `
	SELECT employee_position FROM Employee WHERE employee_id = $1
	`
	err := config.Pool.QueryRow(context.Background(), SQL_QUERY, employee_id).Scan(&EmployeePosition.EmployeePosition)
	if err != nil {
		if err == pgx.ErrNoRows {
			return entity.EmployeePosition{}, errors.New(constants.USER_NOT_FOUND)
		}
		return entity.EmployeePosition{}, errors.New(constants.SOMETHING_WENT_WRONG)
	}

	return EmployeePosition, nil
}

func (R *Repository) SelectEmployeeUUIDByUUID(employee_uuid string) (entity.EmployeeUUID, error) {
	if config.Pool == nil {
		utils.FailedDatabaseConnection(config.Pool)
	}
	var EmployeeUUID entity.EmployeeUUID
	var SQL_QUERY string = `
	SELECT employee_uuid FROM Employee WHERE employee_uuid = $1
	`
	err := config.Pool.QueryRow(context.Background(), SQL_QUERY, employee_uuid).Scan(&EmployeeUUID.EmployeeUUID)
	if err != nil {
		if err == pgx.ErrNoRows {
			return entity.EmployeeUUID{}, errors.New("task_to_employee_uuid not exist")
		}
		return entity.EmployeeUUID{}, errors.New(constants.SOMETHING_WENT_WRONG)
	}
	return EmployeeUUID, nil
}

func (R *Repository) SelectEmployeeIDByUUID(employee_uuid string) (entity.EmployeeID, error) {
	if config.Pool == nil {
		utils.FailedDatabaseConnection(config.Pool)
	}
	var EmployeeUUID entity.EmployeeID
	var SQL_QUERY string = `
	SELECT employee_id FROM Employee WHERE employee_uuid = $1
	`
	err := config.Pool.QueryRow(context.Background(), SQL_QUERY, employee_uuid).Scan(&EmployeeUUID.EmployeeID)
	if err != nil {
		if err == pgx.ErrNoRows {
			return entity.EmployeeID{}, errors.New("employee_id not exist")
		}
		return entity.EmployeeID{}, errors.New(constants.SOMETHING_WENT_WRONG)
	}
	return EmployeeUUID, nil
}

func (R *Repository) UpdateEmployeePosition(employee_id int, new_employee_position string) error {
	if config.Pool == nil {
		utils.FailedDatabaseConnection(config.Pool)
	}

	var SQL_QUERY string = `
	UPDATE Employee SET employee_position = $1 WHERE employee_id = $2
	`

	_, err := config.Pool.Exec(context.Background(), SQL_QUERY, new_employee_position, employee_id)
	if err != nil {
		return err
	}

	return nil
}

func (R *Repository) SelectEmployeeRoles(employee_id int) ([]string, error) {
	if config.Pool == nil {
		utils.FailedDatabaseConnection(config.Pool)
	}
	var EmployeeRoles []string
	var SQL_QUERY string = `
	SELECT employee_role FROM Employee WHERE employee_id = $1
	`
	err := config.Pool.QueryRow(context.Background(), SQL_QUERY, employee_id).Scan(&EmployeeRoles)
	if err != nil {
		fmt.Println(err)
		if err == pgx.ErrNoRows {
			return []string{}, errors.New("no employee_role found")
		}
		return []string{}, errors.New(constants.SOMETHING_WENT_WRONG)
	}

	return EmployeeRoles, nil
}

func (R *Repository) SetSuspendEmployeeByID(BodyData entity.EmployeeSuspend) error {
	if config.Pool == nil {
		utils.FailedDatabaseConnection(config.Pool)
	}

	var SQL_QUERY string = `
	UPDATE Employee SET is_suspended = true, suspension_duration = $1 WHERE employee_id = $2
	`

	Exec, err := config.Pool.Exec(context.Background(), SQL_QUERY, BodyData.SuspensionDuration, BodyData.EmployeeID)
	if err != nil {
		fmt.Println(err)
		return err
	}
	if Exec.RowsAffected() < 1 {
		return errors.New("employee_id not found or something went wrong")
	}

	return nil
}

func (R *Repository) SelectEmployeeSuspension(employee_id int) (entity.SuspensionData, error) {
	if config.Pool == nil {
		utils.FailedDatabaseConnection(config.Pool)
	}
	var SuspensionData entity.SuspensionData
	var SQL_QUERY string = `
	SELECT is_suspended, suspension_duration FROM Employee WHERE employee_id = $1
	`
	err := config.Pool.QueryRow(context.Background(), SQL_QUERY, employee_id).Scan(&SuspensionData.IsSuspended, &SuspensionData.SuspensionDuration)
	if err != nil {
		fmt.Println(err)
		if err == pgx.ErrNoRows {
			return entity.SuspensionData{}, errors.New("no employee_role found")
		}
		return entity.SuspensionData{}, errors.New(constants.SOMETHING_WENT_WRONG)
	}

	return SuspensionData, nil
}

func (R *Repository) SetCancelEmployeeSuspension(employee_id int) error {
	if config.Pool == nil {
		utils.FailedDatabaseConnection(config.Pool)
	}

	var SQL_QUERY string = `
	UPDATE Employee SET is_suspended = false, suspension_duration = null WHERE employee_id = $1
	`

	_, err := config.Pool.Exec(context.Background(), SQL_QUERY, employee_id)
	if err != nil {
		return err
	}

	return nil
}

func (R *Repository) SetEmployeeBanned(employee_id int) error {
	if config.Pool == nil {
		utils.FailedDatabaseConnection(config.Pool)
	}

	var SQL_QUERY string = `
	UPDATE Employee SET is_banned = true WHERE employee_id = $1
	`

	_, err := config.Pool.Exec(context.Background(), SQL_QUERY, employee_id)
	if err != nil {
		return err
	}

	return nil
}

func (R *Repository) SetEmployeeUnBanned(employee_id int) error {
	if config.Pool == nil {
		utils.FailedDatabaseConnection(config.Pool)
	}

	var SQL_QUERY string = `
	UPDATE Employee SET is_banned = false WHERE employee_id = $1
	`

	_, err := config.Pool.Exec(context.Background(), SQL_QUERY, employee_id)
	if err != nil {
		return err
	}

	return nil
}
