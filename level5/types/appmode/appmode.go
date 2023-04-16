package appmode

type AppModeType int

const (
	Register AppModeType = iota
	List
	Purchase
	Quit
)

func (mode AppModeType) String() string {
	switch mode {
	case Register:
		return "〇商品登録モード"
	case List:
		return "〇商品一覧/商品編集モード"
	case Purchase:
		return "〇商品購入モード"
	case Quit:
		return "終了"
	default:
		return ""
	}
}

var AppModeKeys = struct {
	Register string
	List     string
	Purchase string
	Quit     string
}{
	Register: "r",
	List:     "l",
	Purchase: "",
	Quit:     "q",
}
