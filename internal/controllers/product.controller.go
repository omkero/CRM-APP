package controllers

import (
	"crm_system/internal/constants"
	"crm_system/internal/entity"
	"crm_system/internal/functions"
	"crm_system/internal/utils"
	"fmt"
	"os"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

func (C *Controllers) CreateProduct(ctx *fiber.Ctx) error {
	claims, err := utils.ExtractToken(ctx)
	if err != nil {

		ctx.Status(403)
		return ctx.JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	var BodyData = entity.ProductPayload{}

	err = ctx.BodyParser(&BodyData)
	if err != nil {
		ctx.Status(400)
		return ctx.JSON(fiber.Map{
			"message": constants.INVALID_INPUT_TYPE,
		})
	}

	if BodyData.ProductTitle == "" {
		ctx.Status(400)
		return ctx.JSON(fiber.Map{
			"message": "missing product_title",
		})
	}
	if BodyData.ProductDescription == "" {
		ctx.Status(400)
		return ctx.JSON(fiber.Map{
			"message": "missing product_description",
		})
	}
	if BodyData.ProductPrice == 0 {
		ctx.Status(400)
		return ctx.JSON(fiber.Map{
			"message": "missing product_price",
		})
	}

	file, err := ctx.FormFile("product_cover")
	if err != nil {
		ctx.Status(400)
		return ctx.JSON(fiber.Map{
			"message": "missing product_cover",
		})
	}
	var pathName = fmt.Sprintf("%s%s", constants.PRODUCT_SAVE_PATH, file.Filename)
	BodyData.ProductCover = file
	BodyData.ProductCoverPath = pathName
	err = ctx.SaveFile(file, pathName)
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

	err = C.Service.CreateProduct(BodyData, claims.UserID)
	if err != nil {
		ctx.Status(400)
		return ctx.JSON(fiber.Map{
			"message": err.Error(),
		})
	}
	return ctx.JSON(fiber.Map{
		"message": "product has been created",
		"data":    BodyData.ProductCover.Filename,
	})
}

func (C *Controllers) UpdateProduct(ctx *fiber.Ctx) error {
	claims, err := utils.ExtractToken(ctx)
	if err != nil {

		ctx.Status(403)
		return ctx.JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	var BodyData = entity.ProductPayloadUpdate{}

	err = ctx.BodyParser(&BodyData)
	if err != nil {

		ctx.Status(400)
		return ctx.JSON(fiber.Map{
			"message": constants.INVALID_INPUT_TYPE,
		})
	}
	if BodyData.ProductID == 0 {
		ctx.Status(400)
		return ctx.JSON(fiber.Map{
			"message": "missing product_id",
		})
	}
	if BodyData.ProductTitle == "" {
		ctx.Status(400)
		return ctx.JSON(fiber.Map{
			"message": "missing product_title",
		})
	}
	if BodyData.ProductDescription == "" {
		ctx.Status(400)
		return ctx.JSON(fiber.Map{
			"message": "missing product_description",
		})
	}
	if BodyData.ProductPrice == 0 {
		ctx.Status(400)
		return ctx.JSON(fiber.Map{
			"message": "missing product_price",
		})
	}

	file, err := ctx.FormFile("product_cover")
	if err != nil {
		ctx.Status(400)
		return ctx.JSON(fiber.Map{
			"message": "missing product_cover",
		})
	}
	var pathName = fmt.Sprintf("%s%s", constants.PRODUCT_SAVE_PATH, file.Filename)
	BodyData.ProductCover = file
	BodyData.ProductCoverPath = pathName

	err = ctx.SaveFile(file, pathName)
	if err != nil {
		ctx.Status(400)
		return ctx.JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	response, err := functions.VerifyEmployeeStatus(ctx, claims.UserID)
	if err != nil {
		err = os.Remove(pathName)
		if err != nil {
			fmt.Println(err)
		}
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

	err = C.Service.UpdateProduct(BodyData, claims.UserID)
	if err != nil {
		// remove the file
		osErr := os.Remove(pathName)
		if osErr != nil {
			fmt.Println(osErr)
		}

		ctx.Status(400)
		return ctx.JSON(fiber.Map{
			"message": err.Error(),
		})
	}
	return ctx.JSON(fiber.Map{
		"message": "product has been updated",
	})
}

func (C *Controllers) DeleteProduct(ctx *fiber.Ctx) error {
	claims, err := utils.ExtractToken(ctx)
	if err != nil {
		ctx.Status(403)
		return ctx.JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	var BodyData = entity.InputProductType{}

	err = ctx.BodyParser(&BodyData)
	if err != nil {
		ctx.Status(400)
		return ctx.JSON(fiber.Map{
			"message": constants.INVALID_INPUT_TYPE,
		})
	}

	if BodyData.ProductID == 0 {
		ctx.Status(400)
		return ctx.JSON(fiber.Map{
			"message": "missing missing employee_id",
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

	err = C.Service.DeleteProduct(BodyData, claims.UserID)
	if err != nil {
		ctx.Status(400)
		return ctx.JSON(fiber.Map{
			"message": err.Error(),
		})
	}
	return ctx.JSON(fiber.Map{
		"message": "product has been deleted",
	})
}

func (C *Controllers) GetAllProducts(ctx *fiber.Ctx) error {
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

	param := ctx.Query("page")
	if param == "" {
		return ctx.JSON(fiber.Map{
			"message": "missing page query",
		})
	}
	page, err := strconv.ParseInt(param, 10, 64)
	if err != nil {
		return ctx.JSON(fiber.Map{
			"message": "missing page query",
		})
	}

	var limit = constants.PRODUCTS_LIMIT
	var offset = (int(page) - 1) * limit

	data, total, err := C.Service.GetAllProducts(claims.UserID, limit, offset)
	if err != nil {
		ctx.Status(400)
		return ctx.JSON(fiber.Map{
			"message": err.Error(),
		})
	}
	return ctx.JSON(fiber.Map{
		"data":           data,
		"current_page":   page,
		"total_products": total,
	})
}

func (C *Controllers) GetProduct(ctx *fiber.Ctx) error {
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
			"message": "query id is not a number",
		})
	}

	data, err := C.Service.GetProduct(int(id), claims.UserID)
	if err != nil {
		ctx.Status(400)
		return ctx.JSON(fiber.Map{
			"message": err.Error(),
		})
	}
	return ctx.JSON(data)
}
