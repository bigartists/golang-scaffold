package main

import (
	"com.github.goscaffold/internal/model/UserModel"
	"fmt"
	"golang.org/x/crypto/bcrypt"
	"log"
)

func main() {
	user := UserModel.New()
	user.Password = "123456"
	err := user.GeneratePassword()
	if err != nil {
		return

	}

	fmt.Println(user.Password)

	// 输出加密后的密码

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte("1234456"))
	if err != nil {
		log.Println("密码错误")
	} else {
		fmt.Println("密码正确")
	}
}
