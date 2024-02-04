package main

import (
	"com.github.goscaffold/config"
	"com.github.goscaffold/pkg/utils"
	"fmt"
	"github.com/golang-jwt/jwt/v4"
	"time"
)

func main() {

	//priKey, _ := jwt.ParseRSAPrivateKeyFromPEM([]byte(config.SysYamlconfig.Jwt.PrivateKey))
	//
	//user := UserClaim{UserId: "10086"}
	//user.ExpiresAt = time.Now().Add(time.Second * 5).Unix()
	//token_obj := jwt.NewWithClaims(jwt.SigningMethodRS256, user)
	//// 使用私钥签名生成token
	//token, _ := token_obj.SignedString(priKey)
	prikey := []byte(config.SysYamlconfig.Jwt.PrivateKey)
	curTime := time.Now().Add(time.Second * 60 * 60 * 24)
	token, _ := utils.GenerateToken(10, prikey, curTime)
	fmt.Println(token)

	i := 0
	for {
		//// 使用公钥解析token
		//getToken, err := jwt.ParseWithClaims(token, &uc, func(token *jwt.Token) (i interface{}, e error) {
		//	return pubKey, nil
		//})
		getToken, err := utils.ParseToken(token, []byte(config.SysYamlconfig.Jwt.PublicKey))

		if getToken != nil && getToken.Valid {
			fmt.Println(getToken.Claims.(*utils.UserClaim).UserId)
		} else if ve, ok := err.(*jwt.ValidationError); ok {
			if ve.Errors&jwt.ValidationErrorMalformed != 0 {
				fmt.Println("错误的token")
			} else if ve.Errors&(jwt.ValidationErrorExpired|jwt.ValidationErrorNotValidYet) != 0 {

				fmt.Println("token过期或未启用")
			} else {
				fmt.Println("Couldn't handle this token:", err)
			}
		} else {
			fmt.Println("无法解析此token", err)
		}

		i++
		fmt.Println(i)
		time.Sleep(time.Second * 1)
	}
}

/*
alexrh
1
alexrh
2
alexrh
3
alexrh
4
alexrh
5
token过期或未启用
6
token过期或未启用
7
token过期或未启用
8
...
...
...
*/
