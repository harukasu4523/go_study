package database

import (
	"context"
	"fmt"
	"level5/prompt"
	"reflect"

	supa "github.com/nedpals/supabase-go"
)

type Auth struct{}

func (a *Auth) SignInSupabase(client *supa.Client, email string, password string) {
	prompt.PrintlnGreen("メールアドレスを入力してください")
	email, mailErr := prompt.Scan()
	prompt.PrintlnGreen("パスワードを入力してください")
	password, passErr := prompt.Scan()

	ctx := context.Background()
	signUpResponse, err := client.Auth.SignIn(ctx, supa.UserCredentials{
		Email:    email,
		Password: password,
	})
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Println("Created admin user:", signUpResponse.User.ID)
	fmt.Println("type is :", reflect.TypeOf(signUpResponse.User.ID))

}

func (a *Auth) SignUpSupabase(client *supa.Client) {

}
