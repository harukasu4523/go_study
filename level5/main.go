package main

import (
	"fmt"
	"level5/prompt"
	"level5/types/appmode"
	"level5/types/item"
)

func main() {
	items := item.Items{}
	items.Initialize()
	prompt := prompt.Prompt{}
	mode, err := prompt.PromptModeSelect()
	if err != nil {
		return
	}
	switch mode {
	case appmode.Register:
		fmt.Println()
		prompt.PrintlnGreen(appmode.Register.String())
		err := prompt.PromptRegister()
		if err != nil {
			return
		} else {
			// 佐治さんのやつだと入力したやつのリストが来ているけどいる？
			return
		}
	case appmode.List:
		fmt.Println()
		prompt.PrintlnGreen(appmode.List.String())
		return
	case appmode.Purchase:
		fmt.Println()
		prompt.PrintlnGreen(appmode.Purchase.String())
		return
	case appmode.Quit:
		prompt.PrintlnGreen(appmode.Quit.String())
		return
	}

	fmt.Println(mode, err)
}

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
