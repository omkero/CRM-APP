package controllers

import (
	"crm_system/internal/constants"
	"crm_system/internal/entity"
	"crm_system/internal/utils"
	"errors"
	"fmt"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

func (c *Controllers) GetCustomerInfo(ctx *fiber.Ctx) error {
	claims, err := utils.ExtractToken(ctx)
	if err != nil {
		return ctx.JSON(fiber.Map{
			"message": err.Error(),
		})
	}
	ipv4 := utils.GetIPAddress(ctx)
	employeeData, err := c.Service.GetEmployeeInformation(int(claims.UserID), claims, ipv4)
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
			err := c.Service.CancelSuspension(employeeData.EmployeeID)
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

	param := ctx.Query("id")
	if param == "" {
		return ctx.JSON(fiber.Map{
			"message": "missing query id",
		})
	}
	nn, err := strconv.ParseInt(param, 10, 64)
	if err != nil {
		return ctx.JSON(fiber.Map{
			"message": "missing query id",
		})
	}

	resp, err := c.Service.GetCustomerBuID(nn, claims)
	if err != nil {
		ctx.SendStatus(404)
		return ctx.JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	return ctx.JSON(fiber.Map{
		"data":   resp,
		"status": "Sucessfully",
	})
}

func (C *Controllers) CreateNewCustomer(ctx *fiber.Ctx) error {
	var BodyData = new(entity.CustomerDataParserType)
	err := ctx.BodyParser(BodyData)
	if err != nil {
		return ctx.JSON(fiber.Map{
			"message": errors.New(constants.INVALID_JSON_FORMAT),
		})
	}

	claims, err := utils.ExtractToken(ctx)
	if err != nil {
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

	if BodyData.CustomerUsername == "" {
		return ctx.JSON(fiber.Map{
			"message": "Missing Username",
		})
	}

	if BodyData.CustomerEmailAddress == "" {
		return ctx.JSON(fiber.Map{
			"message": "Missing EmailAddress",
		})
	}

	if BodyData.CustomerPosition == "" {
		return ctx.JSON(fiber.Map{
			"message": "Missing Position",
		})
	}

	if BodyData.CustomerPhoneNumber == "" {
		return ctx.JSON(fiber.Map{
			"message": "Missing PhoneNumber",
		})
	}

	if BodyData.CustomerFullName == "" {
		return ctx.JSON(fiber.Map{
			"message": "Missing FullName",
		})
	}

	if err := utils.IsValidEmail(BodyData.CustomerEmailAddress); err != nil {
		return ctx.JSON(fiber.Map{
			"message": constants.INVALID_EMAIL_ADDRESS,
		})
	}

	err = C.Service.CreateCustomer(BodyData, claims)
	if err != nil {
		return ctx.JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	return ctx.JSON(fiber.Map{
		"message": "ok inserted",
	})
}

func (C *Controllers) DeleteCustomer(ctx *fiber.Ctx) error {
	claims, err := utils.ExtractToken(ctx)
	if err != nil {
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
	err = C.Service.DeleteCustomer(9, claims)
	if err != nil {
		return ctx.JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	return ctx.JSON(fiber.Map{
		"message": "okkk",
	})
}

func (C *Controllers) GetAllCustomers(ctx *fiber.Ctx) error {
	QueryParamString := ctx.Query("pageNumber")
	if QueryParamString == "" || QueryParamString == "0" {
		return ctx.JSON(fiber.Map{
			"message": "no pageNumber provided",
		})
	}

	QueryParam, err := strconv.Atoi(QueryParamString)
	if err != nil {
		fmt.Println(err)
	}

	claims, err := utils.ExtractToken(ctx)
	if err != nil {
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

	var page = QueryParam
	var limit = 5
	var offset = (page - 1) * limit

	data, count, err := C.Service.GetAllCustomers(limit, offset, claims)
	if err != nil {
		return ctx.JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	return ctx.JSON(fiber.Map{
		"data":            data,
		"current_page":    page,
		"total_customers": count,
	})
}
