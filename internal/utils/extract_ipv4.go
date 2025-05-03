package utils

import (
	"github.com/gofiber/fiber/v2"
)

// IP returns the remote IP address of the request.
func GetIPAddress(ctx *fiber.Ctx) string {
	IPv4 := ctx.IP()
	return IPv4
}
