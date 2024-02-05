package middlewares

import (
	"com.github.goscaffold/config"
	"com.github.goscaffold/internal/dao"
	"com.github.goscaffold/internal/model/UserModel"
	"com.github.goscaffold/pkg/utils"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
	"net/http"
	"strings"
	"time"
)

func JwtAuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenString := c.GetHeader("Authorization")
		if tokenString == "" || !strings.HasPrefix(tokenString, "Bearer ") {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "token未获取"})
			c.Abort()
			return
		}

		token := tokenString[7:]
		// 解析JWT令牌
		getToken, err := utils.ParseToken(token, []byte(config.SysYamlconfig.Jwt.PublicKey))

		if getToken != nil && getToken.Valid {
			fmt.Println(getToken.Claims.(*utils.UserClaim).UserId)
			userId := getToken.Claims.(*utils.UserClaim).UserId

			// 将userId 转为int， 并调用 dao.DaoGetter.FindUserById
			user := UserModel.New()
			_, err := dao.UserGetter.FindUserById(userId, user)
			if err != nil {
				c.JSON(http.StatusUnauthorized, gin.H{"error": "用户不存在"})
				c.Abort()
			} else {
				//// 生成 token
				prikey := []byte(config.SysYamlconfig.Jwt.PrivateKey)
				curTime := time.Now().Add(time.Second * 60 * 60 * 24)
				token, _ := utils.GenerateToken(user.Id, prikey, curTime)

				c.Set("token", token)
				c.Next()
			}
		} else if ve, ok := err.(*jwt.ValidationError); ok {
			if ve.Errors&jwt.ValidationErrorMalformed != 0 {
				fmt.Println("错误的token")
				c.JSON(http.StatusUnauthorized, gin.H{"error": "错误的token"})
			} else if ve.Errors&(jwt.ValidationErrorExpired|jwt.ValidationErrorNotValidYet) != 0 {
				fmt.Println("token过期或未启用")
				c.JSON(http.StatusUnauthorized, gin.H{"error": "token过期或未启用"})
			} else {
				fmt.Println("Couldn't handle this token:", err)
				c.JSON(http.StatusUnauthorized, gin.H{"error": "无法解析此token"})
			}
		} else {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "无法解析此token"})
		}
		c.Abort()
	}
}
