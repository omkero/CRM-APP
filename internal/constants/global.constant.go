package constants

import (
	"fmt"
)

const (
	SIGNIN_DURATION    = 24 // 24hour 1day
	PASSWORD_LENGTH    = 14
	TASKS_PER_PAGE     = 7
	EMPLOYEES_PER_PAGE = 10
	CUSTOMERS_PER_PAGE = 10
	PRODUCTS_LIMIT     = 10
	SIZE_LIMIT         = 15 // per bytes 15MB
)

var (
	SHORT_PASSWORD          = "Password must be at least" + fmt.Sprintf(" %d ", PASSWORD_LENGTH) + "characters long."
	INVALID_JSON_FORMAT     = "invalid JSON Format"
	SOMETHING_WENT_WRONG    = "[ERROR] oops something went wrong."
	INVALID_TOKEN_SIGNATURE = "Invalid Token Signature"
	TOKEN_EXPIRED           = "Token Has Expired"
	NO_TOKEN_PROVIDED       = "Token Not Found"
	WRONG_PASSWORD          = "Password is Incorrect try again"
	WRONG_EMAIL_ADDRESS     = "Email Address is Wrong try again"
	INVALID_EMAIL_ADDRESS   = "Invalid Email Address try again"
	USER_NOT_FOUND          = "User not found."
	USER_EXIST              = "This user already exists please try another one"
	CUSTOMERS_NOT_FOUND     = "Customers Not Found"
	INVALID_INPUT_TYPE      = "Invalid input try another body type"
)

const (
	API_BASE_NAME     string = "/api/v1"
	PRODUCT_SAVE_PATH        = "public/static/product/"
)
