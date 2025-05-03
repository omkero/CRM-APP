package utils

import "net/mail"

func IsValidEmail(email_address string) error {
	_, err := mail.ParseAddress(email_address)

	if err != nil {
		return err
	}

	return nil
}
