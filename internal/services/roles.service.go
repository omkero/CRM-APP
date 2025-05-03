package services

import (
	"crm_system/internal/entity"
	"crm_system/internal/permissions"
)

func (Service *Services) CreateNewRole(body entity.RolePayload, claims entity.EmployeeTokenClaims) error {
	err := permissions.CanDo(claims.UserID, permissions.CreateRole)
	if err != nil {
		return err
	}
	// var rolesList = permissions.Permissionlist
	var data = entity.RolePayload{
		RoleTitle:               body.RoleTitle,
		RolePermissions:         body.RolePermissions,
		RoleCreatedByEmployeeID: claims.UserID,
	}

	err = Service.Repo.InsertNewRole(claims.UserID, data)
	if err != nil {
		return err
	}

	return nil
}

func (Service *Services) UpdateRole(body entity.RolePayload, claims entity.EmployeeTokenClaims) error {
	err := permissions.CanDo(claims.UserID, permissions.UpdateRole)
	if err != nil {
		return err
	}
	// var rolesList = permissions.Permissionlist
	var data = entity.RolePayload{
		RoleTitle:               body.RoleTitle,
		RolePermissions:         body.RolePermissions,
		RoleCreatedByEmployeeID: claims.UserID,
	}

	err = Service.Repo.UpdateRole(data)
	if err != nil {
		return err
	}

	return nil
}

func (S *Services) GetRolesList() ([]entity.Roledata, error) {
	data, err := S.Repo.SelectAllSystemRoles()
	if err != nil {
		return nil, err
	}

	return data, err
}

func (S *Services) ApplyRole(BodyData entity.RoleApply, claims entity.EmployeeTokenClaims) error {
	err := permissions.CanDo(claims.UserID, permissions.ApplyRole)
	if err != nil {
		return err
	}
	_, err = S.Repo.SelectEmployeeByID(BodyData.EmployeeID)
	if err != nil {
		return err
	}

	err = S.Repo.UpdateEmployeeRoles(BodyData)
	if err != nil {
		return err
	}
	return nil
}

func (S *Services) GetPermissionsWhereRoles(empoyee_roles []string) ([]string, error) {
	data, err := S.Repo.SelectPermissionsWhereEmployeeRoles(empoyee_roles)
	if err != nil {
		return nil, err
	}

	return data, err
}
