package prompt

import (
	"bufio"
	"fmt"
	"level5/database"
	"level5/types/appmode"
	"level5/types/item"
	"os"

	supa "github.com/nedpals/supabase-go"
)

type Prompt struct{}

func (prompt *Prompt) Scan() (string, error) {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	return scanner.Text(), scanner.Err()

}

/*
fmt.Printf("\x1b[1m%s\x1b[0m\n", "1:BOLD")
//文字色
fmt.Printf("\x1b[30m%s\x1b[0m\n", "30:Black (ForeColer)")
fmt.Printf("\x1b[31m%s\x1b[0m\n", "31:Red (ForeColer)")
fmt.Printf("\x1b[32m%s\x1b[0m\n", "32:Green (ForeColer)")
fmt.Printf("\x1b[33m%s\x1b[0m\n", "33:Yellow (ForeColer)")
fmt.Printf("\x1b[34m%s\x1b[0m\n", "34:Blue (ForeColer)")
fmt.Printf("\x1b[35m%s\x1b[0m\n", "35:Magenta (ForeColer)")
fmt.Printf("\x1b[36m%s\x1b[0m\n", "36:Cyan (ForeColer)")
fmt.Printf("\x1b[37m%s\x1b[0m\n", "37:White (ForeColer)")
//ここから背景色。色によって文字色を白(37)/黒(30)に切り替え
fmt.Printf("\x1b[37m\x1b[40m%s\x1b[0m\n", "40:Black (BackColor)")
fmt.Printf("\x1b[37m\x1b[41m%s\x1b[0m\n", "41:Red (BackColor)")
fmt.Printf("\x1b[37m\x1b[42m%s\x1b[0m\n", "42:Green (BackColor)")
fmt.Printf("\x1b[37m\x1b[43m%s\x1b[0m\n", "43:Yellow (BackColor)")
fmt.Printf("\x1b[37m\x1b[44m%s\x1b[0m\n", "44:Blue (BackColor)")
fmt.Printf("\x1b[37m\x1b[45m%s\x1b[0m\n", "45:Magenta (BackColor)")
fmt.Printf("\x1b[30m\x1b[46m%s\x1b[0m\n", "46:Cyan (BackColor)")
fmt.Printf("\x1b[30m\x1b[47m%s\x1b[0m\n", "47:White (BackColor)")*/

func (prompt *Prompt) coloredPrintln(color string, value string) {
	fmt.Printf("\x1b["+color+"m%s\x1b[0m\n", value)
}

func (prompt *Prompt) PrintlnGreen(value string) {
	prompt.coloredPrintln("32", value)
}

func (prompt *Prompt) PrintlnBlue(value string) {
	prompt.coloredPrintln("34", value)
}

func (prompt *Prompt) PrintlnYellow(value string) {
	prompt.coloredPrintln("33", value)
}

func (prompt *Prompt) PrintlnRed(value string) {
	prompt.coloredPrintln("31", value)
}

func (prompt *Prompt) selectionNotValid(mode string) {
	fmt.Println()
	prompt.PrintlnRed("[" + mode + "]" + "は適切な入力ではありません")
	fmt.Println()
}

func (p *Prompt) SignInOrSignUp() string {
	for {
		p.PrintlnGreen("ログイン:1 登録:2")
		input, err := p.Scan()
		if err != nil {
			p.PrintlnRed("文字の読み取りに失敗しました。\nもう一度入力してください")
			continue
		}
		if input == "1" || input == "2" {
			return input
		} else {
			p.PrintlnRed("1 か 2 を入力してください")
			continue
		}
	}
}

func (p *Prompt) PromptModeSelect(client *supa.Client, option string) (string, error) {
	// sign Upならそのままログインした後に無条件で一般を付与
	db := database.Auth{}
	for {
		p.PrintlnGreen("メールアドレスを入力してください")
		email, mailErr := p.Scan()
		if mailErr != nil {
			p.PrintlnRed("文字の読み取りに失敗しました")
			continue
		}
		p.PrintlnGreen("パスワードを入力してください")
		password, passErr := p.Scan()
		if passErr != nil {
			p.PrintlnRed("文字の読み取りに失敗しました")
			continue
		}
		if option == "1" {
			user, err := db.SignInSupabase(client, email, password)
			if err != nil {
				return "err", err
			}
			p.PrintlnYellow(user.ID)
			break
		} else {
			break
		}
	}

	return "admin", nil
}

func (prompt *Prompt) PromptModeAdmin() (appmode.AppModeType, error) {
	for {
		prompt.PrintlnYellow("〇選択一覧")
		prompt.PrintlnGreen("購入：Enter / 登録(register)：\"r\" / 商品一覧(list)：\"l\" / 終了(quit)：\"q\"")
		mode, err := prompt.Scan()
		if err != nil {
			return appmode.Purchase, err
		}
		switch mode {
		case appmode.AppModeKeys.Register:
			return appmode.Register, err
		case appmode.AppModeKeys.List:
			return appmode.List, err
		case appmode.AppModeKeys.Purchase:
			return appmode.Purchase, err
		case appmode.AppModeKeys.Quit:
			return appmode.Quit, err
		default:
			prompt.selectionNotValid(mode)
			continue
		}
	}
}

func (p *Prompt) ToBeContinue() bool {
	for {
		p.PrintlnGreen("続けて入力する場合は:1 終了する場合は:0 を入力してください")
		input, err := p.Scan()
		if err != nil {
			p.PrintlnRed("入力エラーが発生しました。再入力してください")
			continue
		}
		if input == "1" {
			return true
		} else if input == "0" {
			return false
		} else {
			p.PrintlnRed("1か0で入力してください")
			continue
		}
	}
}

func (prompt *Prompt) PromptRegister(items *item.Items) error {
	prompt.PrintlnGreen("商品の入力をしてください")
	for {
		prompt.PrintlnGreen("商品名,価格,在庫数 のようにカンマ区切りで入力してください")
		input, err := prompt.Scan()
		if err != nil {
			prompt.PrintlnRed("商品読み取りに失敗しました")
			continue
		}
		// アイテムの登録
		item := item.Item{}
		err = item.RegistItem(input)
		if err != nil {
			prompt.PrintlnRed(err.Error())
			continue
		} else {
			if err := items.AddItem(item); err != nil {
				prompt.PrintlnRed(err.Error())
			} else {
				if prompt.ToBeContinue() {
					continue
				} else {
					return nil
				}
			}
			fmt.Println()
		}
		return nil
	}
}
