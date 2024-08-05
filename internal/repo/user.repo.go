package repo

import "fmt"

type UserRepo struct{}

func NewUserRepo() *UserRepo {
	return &UserRepo{}
}

func (ur *UserRepo) GetUserByIdRepo(id int) string {
	fmt.Println(id)
	return "HiuMX"
}
