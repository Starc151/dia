package user

import "github.com/Starc151/dia/pkg/ydb"

type User struct {
	Id_user  uint64
	NName    string
	FName    string
	LName    string
	Email    string
	Password string
}

func RegUser() {
	user := User{
		Id_user:  1,
		NName:    "Starc",
		FName:    "Aleksey",
		LName:    "Paskov",
		Email:    "dr.Starc@ya.ru",
		Password: "123bH987",
	}
	ydb.CreateTable(user.NName)
}

