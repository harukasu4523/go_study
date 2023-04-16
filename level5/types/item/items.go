package item

type Items struct {
	ItemList []Item
	ItemMap  map[string]Item
}

func (items *Items) Initialize() {
	items.ItemList = []Item{}
	items.ItemMap = map[string]Item{}
}
