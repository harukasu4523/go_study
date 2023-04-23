package database

import (
	"context"
	"fmt"

	supa "github.com/nedpals/supabase-go"
)

type Auth struct{}

func (a *Auth) SignInSupabase(client *supa.Client, email string, password string) (*supa.User, error) {
	ctx := context.Background()
	SignInResponse, err := client.Auth.SignIn(ctx, supa.UserCredentials{
		Email:    email,
		Password: password,
	})
	if err != nil {
		return nil, err
	}
	fmt.Println("ログインに成功しました。")
	fmt.Println(SignInResponse.User)
	return &SignInResponse.User, nil
}

func (a *Auth) SignUpSupabase(client *supa.Client) {

}
