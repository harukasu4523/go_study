package main

import (
	"level5/prompt"
	"level5/types/item"

	supa "github.com/nedpals/supabase-go"
)

func main() {
	supabaseUrl := "<SUPABASEURL>"
	supabaseKey := "<SUPABASEAPIKEY>"
	client := supa.CreateClient(supabaseUrl, supabaseKey)

	items := item.Items{}
	items.Initialize()
	prompt := prompt.Prompt{}
	userChoice := prompt.SignInOrSignUp()
	role, err := prompt.PromptModeSelect(client, userChoice)
	if err != nil {
		prompt.PrintlnRed(err.Error())
		return
	}
	// products内のName,Price,Inventoryをmap
	// 管理者と一般で選択できるモードを変える。
	if role == "admin" {
		// 管理者権限
		return
	} else {
		// 一般モード
		return
	}
	/*
		for {
		switch mode {
		case appmode.Register:
			fmt.Println()
			prompt.PrintlnYellow(appmode.Register.String())
			err := prompt.PromptRegister(&items)
			if err != nil {
				return
			} else {
				mode, err = prompt.PromptModeSelect()
				if err != nil {
					return
				}
				continue
			}
		case appmode.List:
			fmt.Println()
			prompt.PrintlnYellow(appmode.List.String())
			return
		case appmode.Purchase:
			fmt.Println()
			prompt.PrintlnYellow(appmode.Purchase.String())
			return
		case appmode.Quit:
			prompt.PrintlnYellow(appmode.Quit.String())
			return
		}
	*/
	// fmt.Println(mode, err)
}

/*
DB設計
user_roles{
	id 		string
	roles string
}
products{
	id				int
	Name			string
	Price			int
	Inventory	int
}
cart {
	id		int
	product_id int 商品の外部キー
	quantity	int
}


ログイン
DB参照して管理者or一般を振り分け
管理者プロンプト
 -登録、購入、削除、一覧モード選択
  -登録
	-購入
	-削除
	-一覧
一般モード
 -購入、一覧モード選択
  -管理者と共同
*/

//	モード選択プロンプト（空文字で購入、rで登録、lで商品一覧プロンプト）
//	登録モード（r）なら
//		商品名,価格,在庫数入力プロンプト（カンマ区切り）
//		次商品有無入力プロンプト（Y/n）
//			次商品有り（空文字/Y）なら商品名,価格,在庫数入力プロンプト（カンマ区切り）
//			次商品無し（n）なら商品一覧プロンプト
//	商品一覧モード（l）なら
//		商品一覧プロンプト
//	購入モード（空文字）なら
//		商品選択プロンプト
//			商品選択（商品番号入力）なら
//				カート追加プロンプト（Y/n）
//					追加（空文字/Y）なら追加表示して商品選択プロンプト
//					キャンセル（n）なら商品選択プロンプト
//			商品選択終了（空文字）なら
//				決済確認プロンプト（領収金額入力）
//					領収金額>=決済金額ならお釣り表示して終了
//					領収金額<決済金額なら決済確認プロンプト

//	商品一覧プロンプト
//		修正有り（商品番号入力）なら
//			商品内容を表示して再入力（空文字でキャンセル、delで削除？）
//			商品一覧プロンプトへ戻る
//		修正無し（空文字）なら
//			モード選択プロンプトへ戻る
