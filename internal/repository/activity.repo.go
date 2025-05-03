package repository

import (
	"context"
	"crm_system/config"
	"crm_system/internal/entity"
	"crm_system/internal/utils"
	"fmt"
)

func (R *Repository) InsertActivity(data *entity.ActivityPayload) error {
	if config.Pool == nil {
		utils.FailedDatabaseConnection(config.Pool)
	}

	const SQL_QUERY string = `
    insert into Activity
	(
     activity_employee_id,
     activity_action,
     activity_type,
     activity_log,
     activity_ipv4
    )
	VALUES (
	$1, $2, $3, $4, $5
	)
	`
	_, err := config.Pool.Exec(context.Background(), SQL_QUERY, data.ActivityEmployeeID,
		data.ActivityAction, data.ActivityType, data.ActivityLog, data.ActivityIPV4)

	if err != nil {
		fmt.Println(err)
		return err
	}

	return nil
}
