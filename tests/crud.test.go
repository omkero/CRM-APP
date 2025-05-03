package test

/*
import (
	"context"
	"crm_system/config"
	"crm_system/internal/types"
	"fmt"

	"github.com/gofiber/fiber/v2"
)

func HandleUserPost(ctx *fiber.Ctx) error {

	if config.Pool == nil {
		return fmt.Errorf("error while connecting to DB")
	}

	var userType types.UserModel
	err := ctx.BodyParser(&userType)
	if err != nil {
		ctx.JSON(fiber.Map{
			"message": err,
		})
	}

	_, err = config.Pool.Exec(context.Background(), "INSERT INTO Test (customer_name, customer_role, prefered_product) VALUES ($1, $2, $3)", userType.CustomerName, userType.CustomerRole, userType.PreferedProduct)
	if err != nil {
		return ctx.JSON(fiber.Map{
			"message": err,
		})

	}
	return ctx.JSON(fiber.Map{
		"message": "customer inserted",
	})
}

func GetAllCustomers(ctx *fiber.Ctx) error {
	if config.Pool == nil {
		return fmt.Errorf("error while connecting to DB")
	}

	var CustomersList []types.CustomerModel

	Rows, err := config.Pool.Query(context.Background(), "SELECT * FROM Test")
	if err != nil {
		return ctx.JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	// it will iteratin in each item in Rows
	for Rows.Next() {
		// creating temp instance
		tempUser := new(types.CustomerModel)

		// scanning the data and put it in this tempUser
		err = Rows.Scan(
			&tempUser.CustomerID,
			&tempUser.CustomerName,
			&tempUser.CustomerRole,
			&tempUser.PreferedProduct,
			&tempUser.CreatedAt,
		)

		// if the scanning goes wrong it will return an error
		if err != nil {
			return ctx.JSON(fiber.Map{
				"message": err.Error(),
			})
		}
		CustomersList = append(CustomersList, *tempUser) // finally append this item to the list
	}

	return ctx.JSON(fiber.Map{
		"data": CustomersList,
	})
}

*/
