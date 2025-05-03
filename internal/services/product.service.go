package services

import (
	"crm_system/internal/entity"
	"crm_system/internal/permissions"
	"fmt"
	"os"
)

func (S *Services) CreateProduct(BodyData entity.ProductPayload, employee_id int) error {
	err := permissions.CanDo(employee_id, permissions.CreateProduct)
	if err != nil {
		return err
	}

	BodyData.ProdcutCreatedByEmployeeID = employee_id

	err = S.Repo.InsertProduct(BodyData)
	if err != nil {
		return err
	}

	return nil
}

func (S *Services) UpdateProduct(BodyData entity.ProductPayloadUpdate, employee_id int) error {
	err := permissions.CanDo(employee_id, permissions.CreateProduct)
	if err != nil {
		return err
	}

	response, err := S.Repo.SelectProduct(BodyData.ProductID)
	if err != nil {
		return err
	}

	if response.ProdcutCreatedByEmployeeID != employee_id {
		return fmt.Errorf("this product cannot be updated by you")
	}

	// remove the old file to update
	osErr := os.Remove(response.ProductCoverPath)
	if osErr != nil {
		fmt.Println(osErr)
	}

	err = S.Repo.UpdateSetProduct(BodyData, employee_id)
	if err != nil {
		return err
	}

	return nil
}

func (S *Services) GetProduct(product_id int, employee_id int) (entity.ProductData, error) {
	err := permissions.CanDo(employee_id, permissions.GetProduct)
	if err != nil {
		return entity.ProductData{}, err
	}
	fmt.Println(product_id)
	data, err := S.Repo.SelectProduct(product_id)
	if err != nil {
		return entity.ProductData{}, err
	}

	return data, nil
}

func (S *Services) GetAllProducts(employee_id int, limit int, offset int) ([]entity.ProductData, int, error) {
	err := permissions.CanDo(employee_id, permissions.GetAllProducts)
	if err != nil {
		return []entity.ProductData{}, 0, err
	}

	data, total, err := S.Repo.SelectAllProducts(limit, offset)
	if err != nil {
		return []entity.ProductData{}, 0, err
	}

	return data, total, nil
}

func (S *Services) DeleteProduct(BodyData entity.InputProductType, employee_id int) error {
	err := permissions.CanDo(employee_id, permissions.DeleteProduct)
	if err != nil {
		return err
	}

	response, err := S.Repo.SelectProduct(BodyData.ProductID)
	if err != nil {
		return err
	}

	if response.ProdcutCreatedByEmployeeID != employee_id {
		return fmt.Errorf("this product cannot be deleted by you")
	}

	err = S.Repo.RemoveProduct(BodyData.ProductID, employee_id)
	if err != nil {
		return err
	}

	// remove the file
	err = os.Remove(response.ProductCoverPath)
	if err != nil {
		fmt.Println(err)
	}

	return nil
}
