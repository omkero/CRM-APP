package utils

import (
	"crm_system/internal/constants"
	"crm_system/internal/entity"
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
)

func JwtCreateToken(email_address string, user_id int, secrete_key string) (string, error) {

	jwtToken := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"email_address": email_address,
		"user_id":       user_id,
		"exp":           time.Now().Add(constants.SIGNIN_DURATION * time.Hour).Unix(),
	})

	token, err := jwtToken.SignedString([]byte(secrete_key))
	if err != nil {
		return "", err
	}

	return token, nil
}

func JwtVerifySignature(token_string string, secrete_key string) (*jwt.Token, error) {
	token, err := jwt.Parse(token_string, func(token *jwt.Token) (interface{}, error) {
		return []byte(secrete_key), nil
	})

	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	if !token.Valid {
		return nil, fmt.Errorf(constants.INVALID_TOKEN_SIGNATURE)
	}

	return token, nil
}

func ExtractToken(ctx *fiber.Ctx) (entity.EmployeeTokenClaims, error) {

	authorization := ctx.Get("Authorization")

	bearerAuth := strings.Replace(authorization, "Bearer", "", -1)
	bearerToken := strings.Replace(bearerAuth, " ", "", -1)

	var secretKey string = os.Getenv("SIGN_IN_PRIVATE_KEY")
	token, err := jwt.Parse(bearerToken, func(token *jwt.Token) (interface{}, error) {
		return []byte(secretKey), nil
	})

	if err != nil {
		if errors.Is(err, jwt.ErrTokenInvalidClaims) {
			return entity.EmployeeTokenClaims{}, errors.New(constants.INVALID_TOKEN_SIGNATURE)
		}

		if errors.Is(err, jwt.ErrTokenMalformed) {
			return entity.EmployeeTokenClaims{}, errors.New(constants.INVALID_TOKEN_SIGNATURE)
		}
		fmt.Println(err)
		return entity.EmployeeTokenClaims{}, errors.New(constants.SOMETHING_WENT_WRONG)
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		fmt.Println("not ok")
		return entity.EmployeeTokenClaims{}, errors.New("token Claims is Not ok")
	}

	var CalimsData entity.EmployeeTokenClaims
	num, err := strconv.Atoi(fmt.Sprint(claims["user_id"]))
	if err != nil {
		fmt.Println(err)
		return entity.EmployeeTokenClaims{}, err
	}

	CalimsData.EmployeeEmailAddress = fmt.Sprint(claims["email_address"])
	CalimsData.UserID = num

	return CalimsData, nil
}
