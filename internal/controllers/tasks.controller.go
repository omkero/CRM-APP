package controllers

import (
	"crm_system/internal/constants"
	"crm_system/internal/entity"
	"crm_system/internal/functions"
	"crm_system/internal/permissions"
	"crm_system/internal/utils"
	"fmt"
	"math"
	"strconv"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

func (C *Controllers) GetEmployeeTasks(ctx *fiber.Ctx) error {

	claims, err := utils.ExtractToken(ctx)
	if err != nil {
		ctx.Status(403)
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

	var BodyData entity.EmployeeUUID
	err = ctx.BodyParser(&BodyData)
	if err != nil {
		return ctx.JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	if BodyData.EmployeeUUID == "" {
		ctx.Status(403)
		return ctx.JSON(fiber.Map{
			"message": "Missing Task To employee_uuid",
		})
	}

	tasks, err := C.Service.GetEmployeeTasks(claims.UserID, BodyData.EmployeeUUID)
	if err != nil {
		ctx.Status(403)
		return ctx.JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	return ctx.JSON(fiber.Map{
		"data": tasks,
	})
}

func (C *Controllers) CreateTaskToEmployee(ctx *fiber.Ctx) error {
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
	var BodyData entity.TaskParser
	err = ctx.BodyParser(&BodyData)
	if err != nil {
		fmt.Println(err)
		if BodyData.TaskStartFrom.IsZero() {
			ctx.Status(403)
			return ctx.JSON(fiber.Map{
				"message": "TaskStartFrom Is Not valid Date",
			})
		}
		if BodyData.TaskEndIn.IsZero() {
			ctx.Status(403)
			return ctx.JSON(fiber.Map{
				"message": "TaskEndIn Is Not valid Date",
			})
		}
		return ctx.JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	if BodyData.TaskTitle == "" {
		ctx.Status(403)
		return ctx.JSON(fiber.Map{
			"message": "Missing Task Title",
		})
	}

	if BodyData.TaskToEmployeeUUID == "" {
		ctx.Status(403)
		return ctx.JSON(fiber.Map{
			"message": "Missing Task To Employee UUID",
		})
	}

	if err := uuid.Validate(string(BodyData.TaskToEmployeeUUID)); err != nil {
		fmt.Println(err)
		if uuid.IsInvalidLengthError(err) {
			ctx.Status(403)
			return ctx.JSON(fiber.Map{
				"message": "Invalid UUID Format task_to_employee_uuid",
			})
		}
	}

	if BodyData.TaskDescription == "" {
		ctx.Status(403)
		return ctx.JSON(fiber.Map{
			"message": "Missing Description",
		})
	}
	if BodyData.TaskStartFrom.IsZero() {
		ctx.Status(403)
		return ctx.JSON(fiber.Map{
			"message": "Missing Task Start From",
		})
	}
	if BodyData.TaskEndIn.IsZero() {
		ctx.Status(403)
		return ctx.JSON(fiber.Map{
			"message": "Missing Task End In",
		})
	}
	if BodyData.TaskPriority == "" {
		ctx.Status(403)
		return ctx.JSON(fiber.Map{
			"message": "Missing task_priority",
		})
	}

	// verfiy if task_end_in is less than task_start_from by converting the data into Unix Timestamp
	locationZone, _ := time.LoadLocation("UTC")

	// make sure date to be parsed into utc
	startTime := BodyData.TaskStartFrom.Format(time.RFC3339)
	parsedStartTime, err := time.ParseInLocation(time.RFC3339Nano, startTime, locationZone)
	if err != nil {
		ctx.Status(400)
		return ctx.JSON(fiber.Map{
			"message": "task_start_from is not UTC Format",
		})
	}

	endTime := BodyData.TaskEndIn.Format(time.RFC3339)
	parsedEndTime, err := time.ParseInLocation(time.RFC3339Nano, endTime, locationZone)
	if err != nil {
		ctx.Status(400)
		return ctx.JSON(fiber.Map{
			"message": "task_end_in is not UTC Format",
		})
	}

	if parsedEndTime.UnixNano() <= parsedStartTime.UnixNano() {
		ctx.Status(403)
		return ctx.JSON(fiber.Map{
			"message":  "task_end_in cannot be less or equal than task_start_from",
			"required": "[task_end_in must be greater than task_start_from]",
		})
	}

	// verify if Difference is less than default duration of task
	var timeDifference = int64(math.Abs(float64(parsedEndTime.UnixNano()) - float64(parsedStartTime.UnixNano())))
	var minDifferenceTarget = int64(time.Hour) // system miniumum task duration

	if int64(timeDifference) < int64(minDifferenceTarget) {
		ctx.Status(403)
		return ctx.JSON(fiber.Map{
			"message": "the task must be at least one hour long.",
		})
	}

	err = C.Service.CreateTaskToEmployee(&BodyData, claims)
	if err != nil {
		ctx.Status(403)
		return ctx.JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	return ctx.JSON(fiber.Map{
		"message": "Task Created",
		"data":    BodyData,
	})
}

func (C *Controllers) DeleteTask(ctx *fiber.Ctx) error {
	claims, err := utils.ExtractToken(ctx)
	if err != nil {
		ctx.Status(403)

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

	var BodyData entity.TaskID
	err = ctx.BodyParser(&BodyData)
	if err != nil {
		ctx.Status(400)
		return ctx.JSON(fiber.Map{
			"message": constants.SOMETHING_WENT_WRONG,
		})
	}

	if BodyData.TaskID == 0 {
		ctx.Status(400)
		return ctx.JSON(fiber.Map{
			"message": "task_id is missing",
		})
	}

	err = C.Service.DeleteTask(claims.UserID, BodyData.TaskID)
	if err != nil {
		ctx.Status(400)
		return ctx.JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	return ctx.JSON(fiber.Map{
		"message": fmt.Sprintf("task_id %d deleted", BodyData.TaskID),
	})
}

func (C *Controllers) ApplyTask(ctx *fiber.Ctx) error {
	claims, err := utils.ExtractToken(ctx)
	if err != nil {
		ctx.Status(403)
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

	var BodyData entity.TaskApply
	err = ctx.BodyParser(&BodyData)
	if err != nil {
		ctx.Status(403)
		return ctx.JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	if BodyData.TaskID == 0 {
		ctx.Status(403)
		return ctx.JSON(fiber.Map{
			"message": "missing task_id",
		})
	}

	result := permissions.CheckOperation(BodyData.TaskOperation)
	if !result {
		ctx.Status(403)
		return ctx.JSON(fiber.Map{
			"message": fmt.Sprintf("task_operation: %s not found", BodyData.TaskOperation),
		})
	}

	err = C.Service.ApplyTask(claims.UserID, BodyData.TaskID, BodyData.TaskOperation)
	if err != nil {
		ctx.Status(403)
		return ctx.JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	return ctx.JSON(fiber.Map{
		"message": fmt.Sprintf("You asked for %s", BodyData.TaskOperation),
	})
}

func (C *Controllers) GetAllTasks(ctx *fiber.Ctx) error {
	param := ctx.Query("page")
	if param == "" {
		return ctx.JSON(fiber.Map{
			"message": "missing query page",
		})
	}
	if param == "0" {
		return ctx.JSON(fiber.Map{
			"message": "missing query page",
		})
	}
	pageNum, err := strconv.ParseInt(param, 10, 64)

	if err != nil {
		ctx.Status(403)
		return ctx.JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	claims, err := utils.ExtractToken(ctx)
	if err != nil {
		ctx.Status(403)
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

	var limit = constants.TASKS_PER_PAGE
	var offset = (int(pageNum) - 1) * limit

	allTasks, err := C.Service.GetAllTasks(claims.UserID, limit, offset)
	if err != nil {
		ctx.Status(403)
		return ctx.JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	return ctx.JSON(fiber.Map{
		"data": allTasks,
	})
}
