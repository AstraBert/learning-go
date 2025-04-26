package phonenumber

import (
	"strings"
	"errors"
	"fmt"
)

func Number(phoneNumber string) (string, error) {
	newPhone := ""
	for _,x := range phoneNumber {
		if strings.ContainsAny(string(x), "0123456789") {
			newPhone+=string(x)
		}
	}
	if len(newPhone) < 10 || len(newPhone) > 11 {
		return "", errors.New("This is not a valid NANP telephone number")
	} else if len(newPhone) == 11 {
		if string([]byte{newPhone[0]}) != "1" {
			return "", errors.New("Phone number cannot have a prefix different from 1")
		}
		if string([]byte{newPhone[1]}) == "0" || string([]byte{newPhone[1]}) == "1" || string([]byte{newPhone[4]}) == "0" || string([]byte{newPhone[4]}) == "1" {
			return "", errors.New("This is not a valid NANP telephone number")
		}
		return newPhone[1:], nil
	} else {
		if string([]byte{newPhone[0]}) == "0" || string([]byte{newPhone[0]}) == "1" || string([]byte{newPhone[3]}) == "0" || string([]byte{newPhone[3]}) == "1" {
			return "", errors.New("This is not a valid NANP telephone number")
		}
		return newPhone, nil
	}
}

func AreaCode(phoneNumber string) (string, error) {
	newPhone, err := Number(phoneNumber)
	if err == nil {
		return newPhone[:3], nil
	} else {
		return "", err
	}
}

func Format(phoneNumber string) (string, error) {
	newPhone, err := Number(phoneNumber)
	if err == nil {
		return fmt.Sprintf("(%s) %s-%s", newPhone[:3], newPhone[3:6], newPhone[6:]), nil
	} else {
		return "", err
	}
}
