server:
    port: 8090
    name: admin-api

jwt:
    public_key: |
        -----BEGIN PUBLIC KEY-----
        MIGfMA0GCSqGSIb3DQEBAQUAA4GNADCBiQKBgQC+5sWGL58X8PQx5d3X3Vc53fwL
        6JuKaIBYATmjD0OA0ay0HQsreI2xa98yJ0lH7mZUw0MctJMx+VyI67n8dkoSApXu
        EAPxsMHIdE98Z+LkUQezAT7vEbwlL8gtpysdMl5XyUE+y7njXtSBL5GiUC/Uz21F
        JHJ24nh36X8ZAPWwLwIDAQAB
        -----END PUBLIC KEY-----
    private_key: |
        -----BEGIN RSA PRIVATE KEY-----
        MIICXgIBAAKBgQC+5sWGL58X8PQx5d3X3Vc53fwL6JuKaIBYATmjD0OA0ay0HQsr
        eI2xa98yJ0lH7mZUw0MctJMx+VyI67n8dkoSApXuEAPxsMHIdE98Z+LkUQezAT7v
        EbwlL8gtpysdMl5XyUE+y7njXtSBL5GiUC/Uz21FJHJ24nh36X8ZAPWwLwIDAQAB
        AoGBAKdouM9z85CLPZqEeodTE6srgFzxH3XBLsv+Rw031XqiZVGOJr14esmcT58r
        -----END RSA PRIVATE KEY-----


上面才是将pem放在 yaml 的正确的格式；


验证程序如下

```go
package main

func main() {

	priKey, _ := jwt.ParseRSAPrivateKeyFromPEM([]byte(config.SysYamlconfig.Jwt.PrivateKey))
	pubKey, _ := jwt.ParseRSAPublicKeyFromPEM([]byte(config.SysYamlconfig.Jwt.PublicKey))

	user := UserClaim{Username: "alexrh"}
	user.ExpiresAt = time.Now().Add(time.Second * 5).Unix()
	token_obj := jwt.NewWithClaims(jwt.SigningMethodRS256, user)
	// 使用私钥签名生成token
	token, _ := token_obj.SignedString(priKey)
	fmt.Println(token)

	i := 0
	for {
		uc := UserClaim{}
		// 使用公钥解析token
		getToken, err := jwt.ParseWithClaims(token, &uc, func(token *jwt.Token) (i interface{}, e error) {
			return pubKey, nil
		})
		if getToken != nil && getToken.Valid {
			fmt.Println(getToken.Claims.(*UserClaim).Username)
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
```