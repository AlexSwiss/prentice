package formaterror

import (
	"errors"
	"strings"
)

// FormatError returns error
func FormatError(err string) error {

	if strings.Contains(err, "email") {
		return errors.New("Email Already Taken")
	}

	if strings.Contains(err, "phone") {
		return errors.New("Phone number already used")
	}

	if strings.Contains(err, "name") {
		return errors.New("Name Already Taken")
	}
	if strings.Contains(err, "hashedPassword") {
		return errors.New("Incorrect Password, Check password and retry")
	}
	return errors.New("Incorrect Details")
}
