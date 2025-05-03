package functions

import (
	"crm_system/internal/services"
	"crm_system/internal/utils"
	"database/sql"
	"errors"

	"github.com/gofiber/fiber/v2"
)

type EmployeeStatus struct {
	MessageOne         string       `json:"message" binding:"required"`
	SuspensionDuration sql.NullTime `json:"suspension_duration" binding:"required"`
	IsSuspended        bool
}

// verify if employee is suspended or banned
func VerifyEmployeeStatus(ctx *fiber.Ctx, employee_id int) (EmployeeStatus, error) {
	var ReturnData EmployeeStatus
	var service services.Services

	ipv4 := utils.GetIPAddress(ctx)
	data, err := service.GetEmployeeInformationNoClaims(employee_id, ipv4)
	if err != nil {
		ReturnData.MessageOne = err.Error()

		ctx.Status(400)
		return ReturnData, err
	}

	if data.IsBanned {
		ReturnData.MessageOne = "You have been banned from the system"

		ctx.Status(401)
		return ReturnData, errors.New("you have been banned from the system")
	}

	if data.IsSuspended {
		// cancel suspension if duration its completed
		exist, err := utils.IsSuspensionCompleted(data.SuspensionDuration)
		if err != nil {
			ReturnData.MessageOne = err.Error()

			ctx.Status(400)
			return ReturnData, err
		}
		if exist {
			err := service.CancelSuspension(data.EmployeeID)
			if err != nil {
				ReturnData.MessageOne = err.Error()

				ctx.Status(400)
				return ReturnData, err
			}
		}

		ReturnData.MessageOne = "You have been suspended"
		ReturnData.SuspensionDuration = data.SuspensionDuration
		ReturnData.IsSuspended = true

		ctx.Status(401)
		return ReturnData, errors.New("you have been suspended from the system")
	}

	return EmployeeStatus{}, nil
}
