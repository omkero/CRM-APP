package controllers

import (
	"crm_system/internal/constants"
	"crm_system/internal/entity"
	"crm_system/internal/functions"
	"crm_system/internal/permissions"
	"crm_system/internal/utils"
	"errors"
	"fmt"
	"math"
	"strconv"
	"time"

	"github.com/gofiber/fiber/v2"
)

func (C *Controllers) SignUpEmployee(ctx *fiber.Ctx) error {
	claims, err := utils.ExtractToken(ctx)
	if err != nil {
		return ctx.JSON(fiber.Map{
			"message": err.Error(),
		})
	}
	response, err := functions.VerifyEmployeeStatus(ctx, claims.UserID)
	if err != nil {
		if response.IsSuspended {
			return ctx.JSON(fiber.Map{
				"message":             response.MessageOne,
				"suspension_duration": response.SuspensionDuration,
			})
		}

		return ctx.JSON(fiber.Map{
			"message": response.MessageOne,
		})
	}

	var EmployeeBody = new(entity.EmployeeParserStruct)
	err = ctx.BodyParser(EmployeeBody)

	if err != nil {
		return ctx.JSON(fiber.Map{
			"message": constants.INVALID_JSON_FORMAT,
		})
	}

	if EmployeeBody.EmployeeEmailAddress == "" {
		return ctx.JSON(fiber.Map{
			"message": "missing Email Address",
		})
	}
	if EmployeeBody.EmployeeFullName == "" {
		return ctx.JSON(fiber.Map{
			"message": "missing Full Name",
		})
	}
	if EmployeeBody.EmployeePassword == "" {
		return ctx.JSON(fiber.Map{
			"message": "missing Password",
		})
	}
	if EmployeeBody.EmployeePhoneNumber == "" {
		return ctx.JSON(fiber.Map{
			"message": "missing Phone Number",
		})
	}
	if EmployeeBody.EmplyeeUserName == "" {
		return ctx.JSON(fiber.Map{
			"message": "missing UserName",
		})
	}
	if EmployeeBody.EmployeePosition == "" {
		return ctx.JSON(fiber.Map{
			"message": "missing Position",
		})
	}
	if EmployeeBody.EmployeeRole == nil {
		return ctx.JSON(fiber.Map{
			"message": "missing roles",
		})
	}

	if len(EmployeeBody.EmployeePassword) < constants.PASSWORD_LENGTH {
		return ctx.JSON(fiber.Map{
			"message": constants.SHORT_PASSWORD,
		})
	}

	systemRoles, err := C.Service.GetRolesList()
	if err != nil {
		ctx.Status(400)
		return ctx.JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	notfoundRoles, err := permissions.IsRoleNotFound(EmployeeBody.EmployeeRole, systemRoles)
	if err != nil {
		ctx.Status(404)
		return ctx.JSON(fiber.Map{
			"message":       err.Error(),
			"invalid_roles": notfoundRoles,
		})
	}

	err = C.Service.CreateNewEmployee(EmployeeBody, claims)
	if err != nil {
		return ctx.JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	ctx.Status(201)
	return ctx.JSON(fiber.Map{
		"message": "Employee Created",
		"data":    EmployeeBody,
	})
}

func (C *Controllers) SignInEmployee(ctx *fiber.Ctx) error {
	var BodyData = new(entity.EmployeeSignIn)

	err := ctx.BodyParser(BodyData)
	if err != nil {
		ctx.Status(400)

		return ctx.JSON(fiber.Map{
			"message": constants.INVALID_JSON_FORMAT,
		})
	}

	if BodyData.EmployeeEmailAddress == "" {
		ctx.Status(400)
		return ctx.JSON(fiber.Map{
			"message": "missing Email Address",
		})
	}
	if BodyData.EmployeePassword == "" {
		ctx.Status(400)
		return ctx.JSON(fiber.Map{
			"message": "missing Password",
		})
	}

	err = utils.IsValidEmail(BodyData.EmployeeEmailAddress)
	if err != nil {
		ctx.Status(400)

		return ctx.JSON(fiber.Map{
			"message": "Email Address is not valid",
		})
	}

	if len(BodyData.EmployeePassword) < constants.PASSWORD_LENGTH {
		ctx.Status(400)
		return ctx.JSON(fiber.Map{
			"message": constants.SHORT_PASSWORD,
		})
	}
	ipv4 := utils.GetIPAddress(ctx)

	token, employeeData, err := C.Service.SignInEmployee(BodyData.EmployeeEmailAddress, BodyData.EmployeePassword, ipv4)
	if err != nil {
		fmt.Println(err)
		ctx.Status(400)
		return ctx.JSON(fiber.Map{
			"message": err.Error(),
		})
	}
	response, err := functions.VerifyEmployeeStatus(ctx, employeeData.EmployeeID)
	if err != nil {
		if response.IsSuspended {
			return ctx.JSON(fiber.Map{
				"message":             response.MessageOne,
				"suspension_duration": response.SuspensionDuration,
			})
		}

		return ctx.JSON(fiber.Map{
			"message": response.MessageOne,
		})
	}
	var expireDuration = time.Now().Add(constants.SIGNIN_DURATION * time.Hour) // until the cookie expire
	ctx.Cookie(&fiber.Cookie{
		HTTPOnly: true,
		Name:     "_session_token",
		Path:     "/",
		Expires:  expireDuration,
		Value:    token,
	})
	return ctx.JSON(fiber.Map{
		"message":      "logged in successfully",
		"access_token": token,
	})
}

func (C *Controllers) GetEmployeeInformation(ctx *fiber.Ctx) error {
	claims, err := utils.ExtractToken(ctx)
	if err != nil {
		ctx.Status(400)
		return ctx.JSON(fiber.Map{
			"message": err.Error(),
		})
	}
	ipv4 := utils.GetIPAddress(ctx)
	response, err := functions.VerifyEmployeeStatus(ctx, claims.UserID)
	if err != nil {
		if response.IsSuspended {
			return ctx.JSON(fiber.Map{
				"message":             response.MessageOne,
				"suspension_duration": response.SuspensionDuration,
			})
		}

		return ctx.JSON(fiber.Map{
			"message": response.MessageOne,
		})
	}

	param := ctx.Query("id")
	if param == "" {
		return ctx.JSON(fiber.Map{
			"message": "missing query id",
		})
	}
	if param == "0" {
		return ctx.JSON(fiber.Map{
			"message": "missing query id",
		})
	}
	id, err := strconv.ParseInt(param, 10, 64)

	if err != nil {
		ctx.Status(400)
		return ctx.JSON(fiber.Map{
			"message": "param id is not a number",
		})
	}

	data, err := C.Service.GetEmployeeInformation(int(id), claims, ipv4)
	if err != nil {
		ctx.Status(400)
		return ctx.JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	return ctx.JSON(data)
}

func (C *Controllers) GetAllEmployeesInformation(ctx *fiber.Ctx) error {
	param := ctx.Params("pageNum")
	if param == "" {
		return ctx.JSON(fiber.Map{
			"message": "missing page_number param.",
		})
	}
	if param == "0" {
		return ctx.JSON(fiber.Map{
			"message": "missing page_number param.",
		})
	}
	page, err := strconv.ParseInt(param, 10, 64)

	if err != nil {
		ctx.Status(400)
		return ctx.JSON(fiber.Map{
			"message": "param id is not a number",
		})
	}

	claims, err := utils.ExtractToken(ctx)
	if err != nil {
		ctx.Status(400)
		return ctx.JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	ipv4 := utils.GetIPAddress(ctx)
	employeeData, err := C.Service.GetEmployeeInformation(int(claims.UserID), claims, ipv4)
	if err != nil {
		ctx.Status(400)
		return ctx.JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	if employeeData.IsBanned {
		ctx.Status(401)
		return ctx.JSON(fiber.Map{
			"message": "You have been banned from the system",
		})
	}

	if employeeData.IsSuspended {
		// cancel suspension if duration its completed
		exist, err := utils.IsSuspensionCompleted(employeeData.SuspensionDuration)
		if err != nil {
			ctx.Status(400)
			return ctx.JSON(fiber.Map{
				"message": err.Error(),
			})
		}
		if exist {
			err := C.Service.CancelSuspension(employeeData.EmployeeID)
			if err != nil {
				ctx.Status(400)
				return ctx.JSON(fiber.Map{
					"message": err.Error(),
				})
			}
		}

		ctx.Status(401)
		return ctx.JSON(fiber.Map{
			"message":             "You have been suspended",
			"suspension_duration": employeeData.SuspensionDuration,
		})
	}

	var limit int = constants.EMPLOYEES_PER_PAGE
	var offset int = (int(page) - 1) * limit

	data, count, err := C.Service.GetAllEmployeesInformation(limit, offset, claims, ipv4)
	if err != nil {
		ctx.Status(400)
		return ctx.JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	return ctx.JSON(fiber.Map{
		"data":            data,
		"current_page":    page,
		"total_employees": count,
	})
}

func (C *Controllers) ChangeEmployeePosition(ctx *fiber.Ctx) error {
	claims, err := utils.ExtractToken(ctx)
	if err != nil {
		ctx.Status(400)
		return ctx.JSON(fiber.Map{
			"message": err.Error(),
		})
	}
	response, err := functions.VerifyEmployeeStatus(ctx, claims.UserID)
	if err != nil {
		if response.IsSuspended {
			return ctx.JSON(fiber.Map{
				"message":             response.MessageOne,
				"suspension_duration": response.SuspensionDuration,
			})
		}

		return ctx.JSON(fiber.Map{
			"message": response.MessageOne,
		})
	}

	type BodyStruct struct {
		EmployeeID          int    `json:"employee_id" binding:"required"`
		NewEmployeePosition string `json:"new_employee_position" binding:"required"`
	}
	var BodyData BodyStruct
	err = ctx.BodyParser(&BodyData)

	if err != nil {
		ctx.Status(400)
		return ctx.JSON(fiber.Map{
			"message": constants.SOMETHING_WENT_WRONG,
		})
	}
	if BodyData.EmployeeID == 0 {
		ctx.Status(400)
		return ctx.JSON(fiber.Map{
			"message": "missing employee_id",
		})
	}

	if BodyData.NewEmployeePosition == "" {
		ctx.Status(400)
		return ctx.JSON(fiber.Map{
			"message": "missing new_employee_position",
		})
	}
	ipv4 := utils.GetIPAddress(ctx)

	err = C.Service.ChangeEmployeePosition(BodyData.EmployeeID, BodyData.NewEmployeePosition, claims, ipv4)
	if err != nil {
		ctx.Status(400)
		return ctx.JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	return ctx.JSON(fiber.Map{
		"message": "ok",
	})
}

func (C *Controllers) SuspendEmployee(ctx *fiber.Ctx) error {
	claims, err := utils.ExtractToken(ctx)
	if err != nil {
		ctx.Status(400)
		return ctx.JSON(fiber.Map{
			"message": err.Error(),
		})
	}
	response, err := functions.VerifyEmployeeStatus(ctx, claims.UserID)
	if err != nil {
		if response.IsSuspended {
			return ctx.JSON(fiber.Map{
				"message":             response.MessageOne,
				"suspension_duration": response.SuspensionDuration,
			})
		}

		return ctx.JSON(fiber.Map{
			"message": response.MessageOne,
		})
	}

	_, err = utils.ExtractToken(ctx)
	if err != nil {
		return err
	}

	var BodyData entity.EmployeeSuspend
	err = ctx.BodyParser(&BodyData)
	if err != nil {
		if BodyData.SuspensionDuration.IsZero() {
			ctx.Status(403)
			return ctx.JSON(fiber.Map{
				"message": "missing suspension_duration or its Not valid Date",
			})
		}
		return errors.New(constants.SOMETHING_WENT_WRONG)
	}

	if BodyData.EmployeeID == 0 {
		ctx.Status(400)
		return ctx.JSON(fiber.Map{
			"message": "missing employee_id",
		})
	}
	if BodyData.EmployeeID == claims.UserID {
		ctx.Status(400)
		return ctx.JSON(fiber.Map{
			"message": "you cannot suspend yourself",
		})
	}
	if BodyData.SuspensionReason == "" {
		ctx.Status(400)
		return ctx.JSON(fiber.Map{
			"message": "missing suspension_reason",
		})
	}
	if BodyData.SuspensionDuration.IsZero() {
		ctx.Status(403)
		return ctx.JSON(fiber.Map{
			"message": "missing suspension_duration or its Not valid Date",
		})
	}

	// verfiy if task_end_in is less than task_start_from by converting the data into Unix Timestamp
	locationZone, _ := time.LoadLocation("UTC")

	// make sure date to be parsed into utc
	var fromNow int64 = time.Now().UnixNano()
	endTime := BodyData.SuspensionDuration.Format(time.RFC3339)
	parsedTimeEndTime, err := time.ParseInLocation(time.RFC3339Nano, endTime, locationZone)
	if err != nil {
		ctx.Status(400)
		return ctx.JSON(fiber.Map{
			"message": "suspension_duration is not UTC Format",
		})
	}

	if parsedTimeEndTime.UnixNano() <= fromNow {
		ctx.Status(403)
		return ctx.JSON(fiber.Map{
			"message": "suspension_duration cannot be less or equal from Now",
		})
	}

	// verify if Difference is less than default duration of task
	var timeDifference = int64(math.Abs(float64(parsedTimeEndTime.UnixNano()) - float64(fromNow)))
	var minDifferenceTarget = int64(time.Hour) // system miniumum task duration

	if int64(timeDifference) < int64(minDifferenceTarget) {
		ctx.Status(403)
		return ctx.JSON(fiber.Map{
			"message": "suspension_duration must be at least one hour long.",
		})
	}

	err = C.Service.SuspendEmployee(claims, BodyData)
	if err != nil {
		ctx.Status(400)
		return ctx.JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	return ctx.JSON(fiber.Map{
		"message": "employee has been supended",
	})
}

func (C *Controllers) CancelSuspension(ctx *fiber.Ctx) error {
	claims, err := utils.ExtractToken(ctx)
	if err != nil {
		ctx.Status(400)
		return ctx.JSON(fiber.Map{
			"message": err.Error(),
		})
	}
	response, err := functions.VerifyEmployeeStatus(ctx, claims.UserID)
	if err != nil {
		if response.IsSuspended {
			return ctx.JSON(fiber.Map{
				"message":             response.MessageOne,
				"suspension_duration": response.SuspensionDuration,
			})
		}

		return ctx.JSON(fiber.Map{
			"message": response.MessageOne,
		})
	}

	type BodyStruct struct {
		EmployeeID int `json:"employee_id" binding:"required"`
	}
	var BodyData BodyStruct
	err = ctx.BodyParser(&BodyData)

	if err != nil {
		ctx.Status(400)
		return ctx.JSON(fiber.Map{
			"message": constants.SOMETHING_WENT_WRONG,
		})
	}
	if BodyData.EmployeeID == 0 {
		ctx.Status(400)
		return ctx.JSON(fiber.Map{
			"message": "missing employee_id",
		})
	}
	if BodyData.EmployeeID == claims.UserID {
		ctx.Status(400)
		return ctx.JSON(fiber.Map{
			"message": "you cannot Cancel Suspension yourself",
		})
	}

	err = C.Service.CancelSuspension(BodyData.EmployeeID)
	if err != nil {
		ctx.Status(400)
		return ctx.JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	return ctx.JSON(fiber.Map{
		"message": "employee suspension has been canceled",
	})
}

func (C *Controllers) BanEmployee(ctx *fiber.Ctx) error {
	claims, err := utils.ExtractToken(ctx)
	if err != nil {
		ctx.Status(400)
		return ctx.JSON(fiber.Map{
			"message": err.Error(),
		})
	}
	response, err := functions.VerifyEmployeeStatus(ctx, claims.UserID)
	if err != nil {
		if response.IsSuspended {
			return ctx.JSON(fiber.Map{
				"message":             response.MessageOne,
				"suspension_duration": response.SuspensionDuration,
			})
		}

		return ctx.JSON(fiber.Map{
			"message": response.MessageOne,
		})
	}

	type BodyStruct struct {
		EmployeeID int `json:"employee_id" binding:"required"`
	}
	var BodyData BodyStruct
	err = ctx.BodyParser(&BodyData)

	if err != nil {
		ctx.Status(400)
		return ctx.JSON(fiber.Map{
			"message": constants.SOMETHING_WENT_WRONG,
		})
	}
	if BodyData.EmployeeID == 0 {
		ctx.Status(400)
		return ctx.JSON(fiber.Map{
			"message": "missing employee_id",
		})
	}
	if BodyData.EmployeeID == claims.UserID {
		ctx.Status(400)
		return ctx.JSON(fiber.Map{
			"message": "you cannot ban yourself",
		})
	}

	err = C.Service.BanEmployee(BodyData.EmployeeID)
	if err != nil {
		ctx.Status(400)
		return ctx.JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	return ctx.JSON(fiber.Map{
		"message": "employee has been banned",
	})
}

func (C *Controllers) UnBanEmployee(ctx *fiber.Ctx) error {
	claims, err := utils.ExtractToken(ctx)
	if err != nil {
		ctx.Status(400)
		return ctx.JSON(fiber.Map{
			"message": err.Error(),
		})
	}
	response, err := functions.VerifyEmployeeStatus(ctx, claims.UserID)
	if err != nil {
		if response.IsSuspended {
			return ctx.JSON(fiber.Map{
				"message":             response.MessageOne,
				"suspension_duration": response.SuspensionDuration,
			})
		}

		return ctx.JSON(fiber.Map{
			"message": response.MessageOne,
		})
	}

	type BodyStruct struct {
		EmployeeID int `json:"employee_id" binding:"required"`
	}
	var BodyData BodyStruct
	err = ctx.BodyParser(&BodyData)

	if err != nil {
		ctx.Status(400)
		return ctx.JSON(fiber.Map{
			"message": constants.SOMETHING_WENT_WRONG,
		})
	}
	if BodyData.EmployeeID == 0 {
		ctx.Status(400)
		return ctx.JSON(fiber.Map{
			"message": "missing employee_id",
		})
	}
	if BodyData.EmployeeID == claims.UserID {
		ctx.Status(400)
		return ctx.JSON(fiber.Map{
			"message": "you cannot un-ban yourself",
		})
	}

	err = C.Service.UnBanEmployee(BodyData.EmployeeID)
	if err != nil {
		ctx.Status(400)
		return ctx.JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	return ctx.JSON(fiber.Map{
		"message": "employee has been un-banned",
	})
}
