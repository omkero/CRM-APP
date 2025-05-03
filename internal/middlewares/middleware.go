package middlewares

import (
	"crm_system/internal/constants"
	"crm_system/internal/services"
	"crm_system/internal/utils"
	"os"
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
)

func VerifyEmployeeAuthToken(ctx *fiber.Ctx) error {
	authorization := ctx.Get("Authorization")

	bearerAuth := strings.Replace(authorization, "Bearer", "", -1)
	bearerToken := strings.Replace(bearerAuth, " ", "", -1)

	var secretKey string = os.Getenv("SIGN_IN_PRIVATE_KEY")
	token, err := jwt.Parse(bearerToken, func(token *jwt.Token) (interface{}, error) {
		return []byte(secretKey), nil
	})

	if err != nil {
		ctx.Status(401)
		if strings.Contains(err.Error(), "token is malformed") {
			return ctx.JSON(fiber.Map{
				"message": constants.INVALID_TOKEN_SIGNATURE,
			})
		}
		if strings.Contains(err.Error(), "token is expired") {
			return ctx.JSON(fiber.Map{
				"message": constants.TOKEN_EXPIRED,
			})
		}
		return ctx.JSON(fiber.Map{
			"message": constants.INVALID_TOKEN_SIGNATURE,
		})
	}

	if !token.Valid {
		ctx.Status(401)
		return ctx.JSON(fiber.Map{
			"message": constants.INVALID_TOKEN_SIGNATURE,
		})
	}
	return ctx.Next()
}

func IsEmployeeSuspended(ctx *fiber.Ctx) error {
	claims, err := utils.ExtractToken(ctx)
	if err != nil {
		ctx.Status(401)
		return ctx.JSON(fiber.Map{
			"message": err.Error(),
		})
	}
	service := services.Services{}

	data, err := service.Repo.SelectEmployeeSuspension(claims.UserID)
	if err != nil {
		ctx.Status(400)
		return ctx.JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	if data.IsSuspended {
		ctx.Status(400)
		return ctx.JSON(fiber.Map{
			"message":  "You have been suspended",
			"duration": data.SuspensionDuration,
		})
	}

	return ctx.Next()
}
