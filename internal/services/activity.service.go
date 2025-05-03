package services

import (
	"crm_system/internal/entity"
)

func (Service *Services) CreateActivity(data entity.ActivityPayload, employee_id int) error {
	// activity dont need permissions the system needs to collect logs
	err := Service.Repo.InsertActivity(&data)
	if err != nil {
		return err
	}
	return nil
}
