package database

import (
	"fmt"

	supa "github.com/nedpals/supabase-go"
)

type Getdata struct{}

func (g *Getdata) CheckAuth(client *supa.Client, user *supa.User) bool {
	var role []map[string]interface{}

	err := client.DB.From("user_roles").Select("Role").Eq("UserId", user.ID).Execute(&role)
	if err != nil {
		fmt.Println(err.Error())
		return false
	}
	fmt.Println(len(role))
	fmt.Println(role[0]["Role"].(string))
	// fmt.Println(role[1])
	return true
}
