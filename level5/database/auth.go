package database

import (
	"context"
	"fmt"

	supa "github.com/nedpals/supabase-go"
)

// TODO:権限登録部分を作成する

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
	// fmt.Println(SignInResponse.User)
	return &SignInResponse.User, nil
}

func (a *Auth) GetRole(client *supa.Client, user *supa.User) (string, error) {
	//TODO:user_rolesテーブルからroleを取得する
	var results []map[string]interface{}
	err := client.DB.From("user_roles").Select("Role").Eq("UserId", user.ID).Execute(&results)
	if err != nil {
		panic(err)
	}
	role, _ := results[0]["Role"].(string)
	return role, nil
}

func (a *Auth) SignUpSupabase(client *supa.Client) {

}
