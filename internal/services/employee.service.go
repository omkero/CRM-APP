package services

import (
	"crm_system/internal/constants"
	activity_data "crm_system/internal/data"
	"crm_system/internal/entity"
	"crm_system/internal/permissions"
	"crm_system/internal/utils"
	"errors"
	"fmt"
	"os"

	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

func (S *Services) CreateNewEmployee(EmployeeData *entity.EmployeeParserStruct, claims entity.EmployeeTokenClaims) error {

	err := permissions.CanDo(claims.UserID, permissions.CreateEmployee)
	if err != nil {
		return err
	}
	// Generate a new UUID.
	uuID, err := uuid.NewRandom()
	if err != nil {
		fmt.Println("Error generating UUID:", err)
	}

	hash, err := bcrypt.GenerateFromPassword([]byte(EmployeeData.EmployeePassword), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	// before we send it so the DB we modify some inputs
	EmployeeData.EmployeeUUID = uuID
	EmployeeData.CreatedByEmployeeId = claims.UserID
	EmployeeData.EmployeePassword = string(hash)

	err = S.Repo.InsertEmployee(EmployeeData)
	if err != nil {
		return err
	}

	return nil
}

func (S *Services) SignInEmployee(email_address string, password string, ipv4 string) (string, entity.EmployeeData, error) {
	data, err := S.Repo.SelectEmployeeByEmail(email_address)

	// create activity payload and change the ActivityLog in each error senario
	payload := entity.ActivityPayload{
		ActivityEmployeeID: data.EmployeeID,
		ActivityAction:     activity_data.ACTION_AUTH_SIGNIN,
		ActivityType:       activity_data.TYPE_SIGNIN,
		ActivityLog:        "",
		ActivityIPV4:       ipv4,
	}

	if err != nil {
		payload.ActivityLog = fmt.Sprintf("attempt log-in from ipv4:  %s from email_address: %s, reason: %s", ipv4, email_address, err.Error())
		err = S.Repo.InsertActivity(&payload)
		if err != nil {
			return "", entity.EmployeeData{}, err
		}
		return "", entity.EmployeeData{}, err
	}

	if data.EmployeeEmailAddress != email_address {
		payload.ActivityLog = fmt.Sprintf("attempt log-in from ipv4:  %s from email_address: %s", ipv4, email_address)
		err = S.Repo.InsertActivity(&payload)
		if err != nil {
			return "", entity.EmployeeData{}, err
		}
		return "", entity.EmployeeData{}, errors.New(constants.WRONG_EMAIL_ADDRESS)
	}

	err = bcrypt.CompareHashAndPassword([]byte(data.EmployeePassword), []byte(password))
	if err != nil {
		if errors.Is(err, bcrypt.ErrMismatchedHashAndPassword) {
			payload.ActivityLog = fmt.Sprintf("attempt log-in from ipv4:  %s, from email_address: %s, reason: %s", ipv4, email_address, err.Error())
			err = S.Repo.InsertActivity(&payload)
			if err != nil {
				return "", entity.EmployeeData{}, err
			}
			return "", entity.EmployeeData{}, errors.New(constants.WRONG_PASSWORD)
		}
		return "", entity.EmployeeData{}, err
	}

	var secret string = os.Getenv("SIGN_IN_PRIVATE_KEY")
	// here encrypt email and user_id before creating the token
	token, err := utils.JwtCreateToken(data.EmployeeEmailAddress, data.EmployeeID, secret)
	if err != nil {
		payload.ActivityLog = fmt.Sprintf("attempt log-in from ipv4:  %s, from email_address: %s, reason: %s", ipv4, email_address, err.Error())
		err = S.Repo.InsertActivity(&payload)
		if err != nil {
			return "", entity.EmployeeData{}, err
		}

		return "", entity.EmployeeData{}, errors.New(constants.SOMETHING_WENT_WRONG)
	}

	payload.ActivityLog = fmt.Sprintf("attempt with success log-in from ipv4:  %s, from email_address: %s", ipv4, email_address)
	err = S.Repo.InsertActivity(&payload)
	if err != nil {
		return "", entity.EmployeeData{}, err
	}

	return token, data, nil
}

func (S *Services) GetEmployeeInformation(employee_id int, claims entity.EmployeeTokenClaims, ipv4 string) (entity.EmployeeDataPublic, error) {
	payload := entity.ActivityPayload{
		ActivityEmployeeID: claims.UserID,
		ActivityAction:     activity_data.ACTION_GET_INFO,
		ActivityType:       activity_data.TYPE_GET_EMPLOYEE_INFO,
		ActivityLog:        "",
		ActivityIPV4:       ipv4,
	}

	err := permissions.CanDo(claims.UserID, permissions.GetEmployeeInformation)
	if err != nil {
		return entity.EmployeeDataPublic{}, err
	}

	data, err := S.Repo.SelectEmployeeByID(employee_id)
	if err != nil {
		return entity.EmployeeDataPublic{}, err
	}

	payload.ActivityLog = fmt.Sprintf("attempt with success GetEmployeeInformation from ipv4:  %s, from employee_id: %d", ipv4, claims.UserID)
	err = S.Repo.InsertActivity(&payload)
	if err != nil {
		return entity.EmployeeDataPublic{}, err
	}

	return data, nil
}

func (S *Services) GetEmployeeInformationNoClaims(employee_id int, ipv4 string) (entity.EmployeeDataPublic, error) {
	payload := entity.ActivityPayload{
		ActivityEmployeeID: employee_id,
		ActivityAction:     activity_data.ACTION_GET_INFO,
		ActivityType:       activity_data.TYPE_GET_EMPLOYEE_INFO,
		ActivityLog:        "",
		ActivityIPV4:       ipv4,
	}

	err := permissions.CanDo(employee_id, permissions.GetEmployeeInformation)
	if err != nil {
		fmt.Println(err)
		return entity.EmployeeDataPublic{}, err
	}

	data, err := S.Repo.SelectEmployeeByID(employee_id)
	if err != nil {
		return entity.EmployeeDataPublic{}, err
	}

	payload.ActivityLog = fmt.Sprintf("attempt with success GetEmployeeInformation from ipv4:  %s, from employee_id: %d", ipv4, employee_id)
	err = S.Repo.InsertActivity(&payload)
	if err != nil {
		return entity.EmployeeDataPublic{}, err
	}

	return data, nil
}

func (S *Services) GetAllEmployeesInformation(limit int, offset int, claims entity.EmployeeTokenClaims, ipv4 string) ([]entity.EmployeeDataPublic, int, error) {

	err := permissions.CanDo(claims.UserID, permissions.GetAllEmployeesInformation)
	if err != nil {
		return []entity.EmployeeDataPublic{}, 0, err
	}

	data, count, err := S.Repo.SelectAllEmployees(limit, offset)
	if err != nil {
		return []entity.EmployeeDataPublic{}, 0, err
	}

	payload := entity.ActivityPayload{
		ActivityEmployeeID: claims.UserID,
		ActivityAction:     "GET_INFO",
		ActivityType:       "GET_ALL_EMPLOYEE_INFO",
		ActivityLog:        fmt.Sprintf("attempt with success GetAllEmployeesInformation from ipv4:  %s, from employee_id: %d", ipv4, claims.UserID),
		ActivityIPV4:       ipv4,
	}
	err = S.Repo.InsertActivity(&payload)
	if err != nil {
		return []entity.EmployeeDataPublic{}, 0, err
	}

	return data, count, nil
}

func (S *Services) ChangeEmployeePosition(target_employee_id int, new_position string, claims entity.EmployeeTokenClaims, ipv4 string) error {

	payload := entity.ActivityPayload{
		ActivityEmployeeID: claims.UserID,
		ActivityAction:     activity_data.ACTION_GET_INFO,
		ActivityType:       activity_data.TYPE_GET_EMPLOYEE_INFO,
		ActivityLog:        "",
		ActivityIPV4:       ipv4,
	}

	err := permissions.CanDo(claims.UserID, permissions.ChangeEmployeePosition)
	if err != nil {
		return err
	}

	err = S.Repo.UpdateEmployeePosition(target_employee_id, new_position)
	if err != nil {
		return err
	}

	payload.ActivityLog = fmt.Sprintf("attempt with success ChangeEmployeePosition from ipv4:  %s, from employee_id: %d", ipv4, claims.UserID)
	err = S.Repo.InsertActivity(&payload)
	if err != nil {
		return err
	}

	return nil
}

func (S *Services) GetEmployeeRoles(claims entity.EmployeeTokenClaims) ([]string, error) {
	err := permissions.CanDo(claims.UserID, permissions.GetEmployeeRoles)
	if err != nil {
		return []string{}, err
	}

	data, err := S.Repo.SelectEmployeeRoles(claims.UserID)
	if err != nil {
		return []string{}, err
	}

	return data, nil
}

func (S *Services) SuspendEmployee(claims entity.EmployeeTokenClaims, BodyData entity.EmployeeSuspend) error {
	err := permissions.CanDo(claims.UserID, permissions.SuspendEmployee)
	if err != nil {
		return err
	}

	err = S.Repo.SetSuspendEmployeeByID(BodyData)
	if err != nil {
		return err
	}

	sus, err := S.Repo.SelectEmployeeSuspension(BodyData.EmployeeID)
	if err != nil {
		return err
	}

	fmt.Println(sus)
	return nil
}

func (S *Services) GetEmployeeSuspension(employee_id int, BodyData entity.EmployeeSuspend) (entity.SuspensionData, error) {
	err := permissions.CanDo(employee_id, permissions.SuspendEmployee)
	if err != nil {
		return entity.SuspensionData{}, err
	}

	data, err := S.Repo.SelectEmployeeSuspension(employee_id)
	if err != nil {
		return entity.SuspensionData{}, err
	}

	return data, nil
}

func (S *Services) CancelSuspension(employee_id int) error {
	err := permissions.CanDo(employee_id, permissions.CancelSuspension)
	if err != nil {
		return err
	}
	err = S.Repo.SetCancelEmployeeSuspension(employee_id)
	if err != nil {
		return err
	}

	return nil
}

func (S *Services) BanEmployee(employee_id int) error {
	err := permissions.CanDo(employee_id, permissions.BanEmployee)
	if err != nil {
		return err
	}
	err = S.Repo.SetEmployeeBanned(employee_id)
	if err != nil {
		return err
	}

	return nil
}

func (S *Services) UnBanEmployee(employee_id int) error {
	err := permissions.CanDo(employee_id, permissions.BanEmployee)
	if err != nil {
		return err
	}
	err = S.Repo.SetEmployeeUnBanned(employee_id)
	if err != nil {
		return err
	}

	return nil
}
