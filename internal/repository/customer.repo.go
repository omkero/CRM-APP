package repository

import (
	"context"
	"crm_system/config"
	"crm_system/internal/constants"
	"crm_system/internal/entity"
	"crm_system/internal/utils"
	"errors"
	"fmt"
	"strings"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"
)

func (Repo *Repository) SelectCustomerByID(customer_id int64) (*entity.CustomerData, error) {
	if config.Pool == nil {
		utils.FailedDatabaseConnection(config.Pool)
	}
	var largeCustomerID int64 = int64(customer_id)
	var CustomerData entity.CustomerData

	err := config.Pool.QueryRow(context.Background(), "SELECT * FROM Customer WHERE customer_id = $1", largeCustomerID).Scan(
		&CustomerData.CustomerID,
		&CustomerData.CustomerUsername,
		&CustomerData.CustomerUUID,
		&CustomerData.CustomerPosition,
		&CustomerData.CustomerFullName,
		&CustomerData.CustomerPhoneNumber,
		&CustomerData.CustomerEmailAddress,
		&CustomerData.CustomerCreatedByEmployeeID,
		&CustomerData.CreatedAt,
	)
	if err != nil {
		if err == pgx.ErrNoRows {
			return nil, errors.New("this customer not been found")
		}
		if strings.Contains(err.Error(), "failed to encode") {
			return nil, errors.New("this number is not valid or its too big")
		}
		return nil, err
	}
	return &CustomerData, nil
}

func (Repo *Repository) InsertCustomer(BodyData *entity.CustomerDataParserType) error {
	if config.Pool == nil {
		utils.FailedDatabaseConnection(config.Pool)
	}
	var SQL_QUERY string = `
	INSERT INTO Customer (
	customer_username,
	customer_uuid,
	customer_position,
	customer_full_name,
	customer_phone_number,
	customer_email_address,
	customer_created_by_employee_id
	) VALUES (
	 $1, $2, $3 , $4 , $5 , $6 , $7
	)
	`
	_, err := config.Pool.Exec(context.Background(), SQL_QUERY, &BodyData.CustomerUsername,
		&BodyData.CustomerUUID, &BodyData.CustomerPosition, &BodyData.CustomerFullName, &BodyData.CustomerPhoneNumber,
		&BodyData.CustomerEmailAddress, &BodyData.CustomerCreatedByEmployeeID)

	if err != nil {
		var pgErr *pgconn.PgError
		if errors.As(err, &pgErr) {
			if pgErr.Code == `23505` {
				return errors.New(constants.USER_EXIST)
			}
		}
		return err
	}
	return nil
}

func (Repo *Repository) RemoveCustomer(customer_id int) error {
	if config.Pool == nil {
		utils.FailedDatabaseConnection(config.Pool)
	}

	var SQL_QUERY string = `
	DELETE FROM Customer WHERE customer_id = $1
	`
	_, err := config.Pool.Exec(context.Background(), SQL_QUERY, customer_id)
	if err != nil {
		return err
	}

	return nil
}

func (Repo *Repository) SelectAllCustomers(limit int, offset int) ([]entity.CustomerData, int, error) {
	if config.Pool == nil {
		utils.FailedDatabaseConnection(config.Pool)
	}

	var Customers []entity.CustomerData
	var totalCustomers int

	var SQL_QUERY string = `
	SELECT * FROM Customer ORDER BY created_at DESC LIMIT $1 OFFSET $2
	`
	var SQL_COUNT string = `
	SELECT COUNT(*) FROM customer
	`

	Row, err := config.Pool.Query(context.Background(), SQL_QUERY, limit, offset)
	if err != nil {
		if err == pgx.ErrNoRows {
			return []entity.CustomerData{}, 0, errors.New(constants.CUSTOMERS_NOT_FOUND)
		}
		return []entity.CustomerData{}, 0, errors.New(constants.SOMETHING_WENT_WRONG)
	}

	for Row.Next() {
		var CustomerData entity.CustomerData
		err = Row.Scan(
			&CustomerData.CustomerID,
			&CustomerData.CustomerUsername,
			&CustomerData.CustomerUUID,
			&CustomerData.CustomerPosition,
			&CustomerData.CustomerFullName,
			&CustomerData.CustomerPhoneNumber,
			&CustomerData.CustomerEmailAddress,
			&CustomerData.CustomerCreatedByEmployeeID,
			&CustomerData.CreatedAt,
		)
		if err != nil {
			return []entity.CustomerData{}, 0, errors.New(constants.SOMETHING_WENT_WRONG)
		}

		Customers = append(Customers, CustomerData)
	}

	err = config.Pool.QueryRow(context.Background(), SQL_COUNT).Scan(&totalCustomers)
	if err != nil {
		if err == pgx.ErrNoRows {
			return []entity.CustomerData{}, 0, fmt.Errorf("product not found")
		}
		return []entity.CustomerData{}, 0, fmt.Errorf(constants.SOMETHING_WENT_WRONG)
	}

	return Customers, totalCustomers, nil
}
