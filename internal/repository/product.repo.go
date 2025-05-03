package repository

import (
	"context"
	"crm_system/config"
	"crm_system/internal/constants"
	"crm_system/internal/entity"
	"crm_system/internal/utils"
	"fmt"

	"github.com/jackc/pgx/v5"
)

func (Repo *Repository) InsertProduct(BodyData entity.ProductPayload) error {
	if config.Pool == nil {
		utils.FailedDatabaseConnection(config.Pool)
	}

	var SQL_QUERY string = `
	INSERT INTO product 
	(
	product_title,
	product_description,
	product_price,
	created_by_employee_id,
	product_cover
	) 
	VALUES ($1, $2, $3, $4, $5)
	`

	_, err := config.Pool.Exec(context.Background(), SQL_QUERY, BodyData.ProductTitle,
		BodyData.ProductDescription, BodyData.ProductPrice, BodyData.ProdcutCreatedByEmployeeID, BodyData.ProductCoverPath)

	if err != nil {
		fmt.Println(err)
		return fmt.Errorf(constants.SOMETHING_WENT_WRONG)
	}

	return nil
}

func (Repo *Repository) SelectProduct(product_id int) (entity.ProductData, error) {
	if config.Pool == nil {
		utils.FailedDatabaseConnection(config.Pool)
	}
	var response = entity.ProductData{}

	var SQL_QUERY string = `
	SELECT * FROM product WHERE product_id = $1
	`

	err := config.Pool.QueryRow(context.Background(), SQL_QUERY, product_id).Scan(&response.ProductID,
		&response.ProductTitle, &response.ProductDescription, &response.ProductPrice, &response.ProdcutCreatedByEmployeeID, &response.ProductCoverPath,
		&response.CreatedAt,
	)
	if err != nil {
		if err == pgx.ErrNoRows {
			return entity.ProductData{}, fmt.Errorf("product not found")
		}
		return entity.ProductData{}, fmt.Errorf(constants.SOMETHING_WENT_WRONG)
	}

	return response, nil
}
func (Repo *Repository) SelectAllProducts(limit int, offset int) ([]entity.ProductData, int, error) {
	if config.Pool == nil {
		utils.FailedDatabaseConnection(config.Pool)
	}
	var response = []entity.ProductData{}
	var totalProducts int = 0

	var SQL_QUERY string = `
	SELECT * FROM product ORDER BY created_at DESC LIMIT $1 OFFSET $2
	`
	var SQL_COUNT string = `
	SELECT COUNT(*) FROM product
	`

	Row, err := config.Pool.Query(context.Background(), SQL_QUERY, limit, offset)
	if err != nil {
		if err == pgx.ErrNoRows {
			return []entity.ProductData{}, 0, fmt.Errorf("product not found")
		}
		return []entity.ProductData{}, 0, fmt.Errorf(constants.SOMETHING_WENT_WRONG)
	}

	for Row.Next() {
		var tempProduct entity.ProductData
		err := Row.Scan(
			&tempProduct.ProductID,
			&tempProduct.ProductTitle,
			&tempProduct.ProductDescription,
			&tempProduct.ProductPrice,
			&tempProduct.ProdcutCreatedByEmployeeID,
			&tempProduct.ProductCoverPath,

			&tempProduct.CreatedAt,
		)
		if err != nil {
			fmt.Println(err)
			return []entity.ProductData{}, 0, fmt.Errorf(constants.SOMETHING_WENT_WRONG)
		}
		response = append(response, tempProduct)
	}

	err = config.Pool.QueryRow(context.Background(), SQL_COUNT).Scan(&totalProducts)
	if err != nil {
		if err == pgx.ErrNoRows {
			return []entity.ProductData{}, 0, fmt.Errorf("product not found")
		}
		return []entity.ProductData{}, 0, fmt.Errorf(constants.SOMETHING_WENT_WRONG)
	}

	return response, totalProducts, nil
}

func (Repo *Repository) RemoveProduct(product_id int, employee_id int) error {
	if config.Pool == nil {
		utils.FailedDatabaseConnection(config.Pool)
	}

	var SQL_QUERY string = `
	DELETE FROM product WHERE product_id = $1 AND created_by_employee_id = $2
	`

	_, err := config.Pool.Exec(context.Background(), SQL_QUERY, product_id, employee_id)

	if err != nil {
		fmt.Println(err)
		return fmt.Errorf(constants.SOMETHING_WENT_WRONG)
	}

	return nil
}

func (Repo *Repository) UpdateSetProduct(BodyData entity.ProductPayloadUpdate, employee_id int) error {
	if config.Pool == nil {
		utils.FailedDatabaseConnection(config.Pool)
	}

	var SQL_QUERY string = `
	UPDATE product SET product_title = $1, product_description = $2, product_price = $3, product_cover = $4 
	WHERE product_id = $5 AND created_by_employee_id = $6
	`

	_, err := config.Pool.Exec(context.Background(), SQL_QUERY, BodyData.ProductTitle, BodyData.ProductDescription,
		BodyData.ProductPrice, BodyData.ProductCoverPath, BodyData.ProductID, employee_id,
	)

	if err != nil {
		fmt.Println(err)
		return fmt.Errorf(constants.SOMETHING_WENT_WRONG)
	}

	return nil
}
