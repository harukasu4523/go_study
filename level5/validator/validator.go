package validator

import (
	"errors"
	"strconv"
)

type Validator struct{}

func (validator *Validator) CheckString(name string, value string) (string, error) {
	if len(value) == 0 {
		return value, errors.New(name + "を入力してください")
	}
	return value, nil
}

func (validator *Validator) CheckNumeric(name string, value string) (int, error) {
	if len(value) == 0 {
		return 0, errors.New(name + "を入力してください")
	}
	validNum, err := strconv.Atoi(value)
	if err != nil  {
		return 0 , errors.New("数値を入力してください")
	}
	if validNum < 0 {
		return 0, errors.New("正の整数を入力してください")
	} 

	return validNum, nil
}
