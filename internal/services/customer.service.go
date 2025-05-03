package services

import (
	"crm_system/internal/entity"
	"crm_system/internal/permissions"
	"fmt"

	"github.com/google/uuid"
)

func (Services *Services) GetCustomerBuID(customer_id int64, Claims entity.EmployeeTokenClaims) (*entity.CustomerData, error) {
	err := permissions.CanDo(Claims.UserID, permissions.GetCustomer)
	if err != nil {
		return nil, err
	}

	data, err := Services.Repo.SelectCustomerByID(customer_id)
	if err != nil {
		return nil, err
	}

	return data, nil
}

func (Services *Services) CreateCustomer(BodyData *entity.CustomerDataParserType, Claims entity.EmployeeTokenClaims) error {

	err := permissions.CanDo(Claims.UserID, permissions.CreateCustomer)
	if err != nil {
		return err
	}

	uuidV4, err := uuid.NewUUID()
	if err != nil {
		return err
	}

	BodyData.CustomerUUID = uuidV4
	BodyData.CustomerCreatedByEmployeeID = Claims.UserID

	err = Services.Repo.InsertCustomer(BodyData)
	if err != nil {
		fmt.Println(err)
		return err
	}

	return nil
}

func (Services *Services) DeleteCustomer(customer_id int, Claims entity.EmployeeTokenClaims) error {

	err := permissions.CanDo(Claims.UserID, permissions.DeleteCustomer)
	if err != nil {
		return err
	}

	err = Services.Repo.RemoveCustomer(customer_id)
	if err != nil {
		return err
	}

	return nil
}

func (Services *Services) GetAllCustomers(limit int, offset int, Claims entity.EmployeeTokenClaims) ([]entity.CustomerData, int, error) {

	err := permissions.CanDo(Claims.UserID, permissions.GetAllCustomers)
	if err != nil {
		return []entity.CustomerData{}, 0, err
	}

	ff, count, err := Services.Repo.SelectAllCustomers(limit, offset)
	if err != nil {
		return []entity.CustomerData{}, 0, err
	}

	return ff, count, nil
}
