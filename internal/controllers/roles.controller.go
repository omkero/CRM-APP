package controllers

import (
	"crm_system/internal/constants"
	"crm_system/internal/entity"
	"crm_system/internal/functions"
	"crm_system/internal/permissions"
	"crm_system/internal/utils"

	"github.com/gofiber/fiber/v2"
)

func (C *Controllers) CreateRole(ctx *fiber.Ctx) error {
	var BodyData entity.RolePayload
	err := ctx.BodyParser(&BodyData)
	if err != nil {
		ctx.Status(400)
		return ctx.JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	claims, err := utils.ExtractToken(ctx)
	if err != nil {
		ctx.Status(401)
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

	err = C.Service.CreateNewRole(BodyData, claims)
	if err != nil {
		ctx.Status(400)
		return ctx.JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	return ctx.JSON(fiber.Map{
		"message": "role created",
	})
}

func (C *Controllers) UpdateRole(ctx *fiber.Ctx) error {
	var BodyData entity.RolePayload
	err := ctx.BodyParser(&BodyData)
	if err != nil {
		ctx.Status(400)
		return ctx.JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	claims, err := utils.ExtractToken(ctx)
	if err != nil {
		ctx.Status(401)
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

	wrongPermissions, err := permissions.IsPermissionsNotFound(BodyData.RolePermissions)
	if err != nil {
		ctx.Status(400)
		return ctx.JSON(fiber.Map{
			"message":           err.Error(),
			"wrong_permissions": wrongPermissions,
		})
	}

	err = C.Service.UpdateRole(BodyData, claims)
	if err != nil {
		ctx.Status(400)
		return ctx.JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	return ctx.JSON(fiber.Map{
		"message": "role updated",
	})
}

func (C *Controllers) GetSystemRoles(ctx *fiber.Ctx) error {
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

	rolse, err := C.Service.GetRolesList()
	if err != nil {
		ctx.Status(400)
		return ctx.JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	return ctx.JSON(fiber.Map{
		"data":    rolse,
		"message": "roles",
	})
}

func (C *Controllers) ApplyRole(ctx *fiber.Ctx) error {
	var BodyData entity.RoleApply
	err := ctx.BodyParser(&BodyData)
	if err != nil {
		ctx.Status(400)
		return ctx.JSON(fiber.Map{
			"message": constants.SOMETHING_WENT_WRONG,
		})
	}

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

	if BodyData.EmployeeID == 0 {
		ctx.Status(400)
		return ctx.JSON(fiber.Map{
			"message": "missing employee_id ",
		})
	}
	if BodyData.RoleTitles == nil {
		ctx.Status(400)
		return ctx.JSON(fiber.Map{
			"message": "missing role_title",
		})
	}

	systemRoles, err := C.Service.GetRolesList()
	if err != nil {
		ctx.Status(400)
		return ctx.JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	wrongRoles, err := permissions.IsRoleNotFound(BodyData.RoleTitles, systemRoles)
	if err != nil {
		ctx.Status(400)
		return ctx.JSON(fiber.Map{
			"message":     err.Error(),
			"wrong_roles": wrongRoles,
		})
	}

	err = C.Service.ApplyRole(BodyData, claims)
	if err != nil {
		ctx.Status(400)
		return ctx.JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	return ctx.JSON(fiber.Map{
		"message": "role applied to employee.",
	})
}

func (C *Controllers) GetSystemPermissions(ctx *fiber.Ctx) error {
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

	var permissionsList = permissions.Permissionlist

	return ctx.JSON(fiber.Map{
		"data":       permissionsList,
		"message":    "system permissions",
		"statusCode": 200,
	})
}

// now after creating roles and permission next step is to check if employee has the right to do this permission based on his roles
// has the right to do this permission based on his roles
