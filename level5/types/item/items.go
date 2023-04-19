package item

import "errors"

type Items struct {
	ItemList []Item
	ItemMap  map[string]Item
}

func (items *Items) Initialize() {
	items.ItemList = []Item{}
	items.ItemMap = map[string]Item{}
}

func (items *Items) AddItem(item Item) error {
	_, exists := items.ItemMap[item.Name]
	if exists {
		return errors.New(item.Name + "はすでに登録されています。")
	} else {
		items.ItemList = append(items.ItemList, item)
		items.ItemMap[item.Name] = item
	}
	return nil
}
