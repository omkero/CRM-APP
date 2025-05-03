package repository

import (
	"context"
	"crm_system/config"
	"crm_system/internal/constants"
	"crm_system/internal/entity"
	"crm_system/internal/utils"
	"encoding/json"
	"errors"
	"fmt"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"
)

func (R *Repository) InsertNewRole(employee_id int, payload entity.RolePayload) error {
	if config.Pool == nil {
		utils.FailedDatabaseConnection(config.Pool)
	}

	var SQL_QUERY string = `
	INSERT INTO system_roles (role_title, role_permissions, role_created_by_employee_id)
	VALUES ($1, $2, $3)
	`
	_, err := config.Pool.Exec(context.Background(), SQL_QUERY, payload.RoleTitle, payload.RolePermissions, employee_id)
	if err != nil {
		var pgErr = &pgconn.PgError{}
		if errors.As(err, &pgErr) {
			if pgErr.Code == "23505" {
				return errors.New("this role already exist try something else or remove it")
			}
		}

		return errors.New(constants.SOMETHING_WENT_WRONG)
	}

	return nil
}

func (R *Repository) UpdateRole(payload entity.RolePayload) error {
	if config.Pool == nil {
		utils.FailedDatabaseConnection(config.Pool)
	}

	type Role struct {
		RoleTitle string
	}
	var data Role

	var SQL_QUERY_VERIFY = `SELECT role_title FROM system_roles WHERE role_title = $1`
	errRow := config.Pool.QueryRow(context.Background(), SQL_QUERY_VERIFY, payload.RoleTitle).Scan(&data.RoleTitle)
	if errRow != nil {
		fmt.Println(errRow)
		if errRow == pgx.ErrNoRows {
			return errors.New("role_title not found")
		}
		return errors.New(constants.SOMETHING_WENT_WRONG)
	}

	var SQL_QUERY string = `
	UPDATE system_roles SET role_title = $1, role_permissions = $2, role_created_by_employee_id = $3 WHERE role_title = $4;
	`
	_, err := config.Pool.Exec(context.Background(), SQL_QUERY, payload.RoleTitle,
		payload.RolePermissions, payload.RoleCreatedByEmployeeID, payload.RoleTitle)
	if err != nil {
		fmt.Println(err)
		return errors.New(constants.SOMETHING_WENT_WRONG)
	}

	return nil
}

func (R *Repository) SelectAllSystemRoles() ([]entity.Roledata, error) {
	if config.Pool == nil {
		utils.FailedDatabaseConnection(config.Pool)
	}
	var data []entity.Roledata

	var SQL_QUERY string = `
	SELECT * FROM system_roles;
	`
	Row, err := config.Pool.Query(context.Background(), SQL_QUERY)
	if err != nil {
		if err == pgx.ErrNoRows {
			return []entity.Roledata{}, errors.New("items not found")
		}

		return []entity.Roledata{}, errors.New(constants.SOMETHING_WENT_WRONG)
	}

	for Row.Next() {
		var tempRole entity.Roledata
		err := Row.Scan(
			&tempRole.RoleID,
			&tempRole.RoleTitle,
			&tempRole.RolePermissions,
			&tempRole.RoleCreatedByEmployeeID,
			&tempRole.CreatedAt,
		)
		if err != nil {
			return []entity.Roledata{}, errors.New(constants.SOMETHING_WENT_WRONG)

		}

		data = append(data, tempRole)
	}

	return data, nil
}

func (R *Repository) UpdateEmployeeRoles(BodyData entity.RoleApply) error {
	if config.Pool == nil {
		utils.FailedDatabaseConnection(config.Pool)
	}
	var SQL_QUERY string = `
	UPDATE employee SET employee_role = $1 WHERE employee_id = $2
	`
	response, err := config.Pool.Exec(context.Background(), SQL_QUERY, BodyData.RoleTitles, BodyData.EmployeeID)
	fmt.Println(response.RowsAffected())
	if err != nil {
		fmt.Println(err)
		return errors.New(constants.SOMETHING_WENT_WRONG)
	}

	return nil
}

func (r *Repository) SelectPermissionsWhereEmployeeRoles(EmployeeRoles []string) ([]string, error) {
	if config.Pool == nil {
		utils.FailedDatabaseConnection(config.Pool)
	}
	var SelectedList []byte // json_agg returns []byte
	var SQL_QUERY string = `
SELECT COALESCE(
    json_agg(DISTINCT permission), '[]'::json
) AS merged_permissions
FROM (
    SELECT unnest(role_permissions) AS permission
    FROM system_roles
    WHERE role_title = ANY($1)
) AS perms;

	`

	err := config.Pool.QueryRow(context.Background(), SQL_QUERY, EmployeeRoles).Scan(&SelectedList)
	if err != nil {
		return nil, err
	}

	var permissions []string
	if err := json.Unmarshal(SelectedList, &permissions); err != nil {

		fmt.Println(err)
		return nil, err
	}

	return permissions, nil
}
