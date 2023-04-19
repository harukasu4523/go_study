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
	if err != nil {
		// エラーが int の範囲外であることを確認
		if numErr, ok := err.(*strconv.NumError); ok && numErr.Err == strconv.ErrRange {
			return 0, errors.New("数値が int の範囲を超えています")
		} else {
			return 0, errors.New("数値を入力してください")
		}
	}
	if validNum < 0 {
		return 0, errors.New("正の整数を入力してください")
	}

	return validNum, nil
}
