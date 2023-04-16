package item

import (
	"errors"
	"level5/validator"
	"strings"
)

type Item struct {
	Name      string
	Price     int
	Inventory int
}

func (item *Item) RegistItem(input string) error {
	splitted := strings.Split(input, ",")
	if len(splitted) != 3 {
		return errors.New("商品名,価格,在庫数 のようにカンマ区切りで入力してください")
	}

	//値,数値,数値チェック
	name, price, inventory := splitted[0], splitted[1], splitted[2]
	validator := validator.Validator{}
	if validName, err := validator.CheckString("商品名", name); err != nil {
		return err
	} else {
		item.Name = validName
	}
	if validPrice, err := validator.CheckNumeric("価格", price); err != nil {
		return err
	} else {
		item.Price = validPrice
	}
	if validInventory, err := validator.CheckNumeric("在庫数", inventory); err != nil {
		return err
	} else {
		item.Inventory = validInventory
	}

	return nil
}
