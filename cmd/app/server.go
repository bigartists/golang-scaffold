package app

import (
	dbs "com.github.goscaffold/internal/dao"
	middlewares "com.github.goscaffold/pkg/middlewares"
	"com.github.goscaffold/pkg/validators"
	"com.github.goscaffold/web/controllers"
	"com.github.goscaffold/web/routes"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/spf13/cobra"
)

// 本函数用于初始化gin
// 有一个测试路由 GET  /test
// 函数接收一个参数，用于指定监听的端口

func Run(port int) error {
	// 执行命令行
	dbs.InitDB()
	r := gin.New()

	controllers.SetUpAuthController(r)

	r.Use(middlewares.JwtAuthMiddleware())
	r.Use(middlewares.ErrorHandler())

	// 登录注册接口 跳过 JwtAuthMiddleware 验证

	// 加载路由
	//handlers.Build(r)
	routes.Build(r)
	// 加载 validator
	validators.Build()

	err := r.Run(fmt.Sprintf(":%d", port))
	if err != nil {
		return err
	}
	return nil
}

// NewApiServerCommand 初始化命令行参数
func NewApiServerCommand() (cmd *cobra.Command) {
	// 集成 cobra命令
	cmd = &cobra.Command{
		Use: "appServer",
		RunE: func(cmd *cobra.Command, args []string) error {
			port, err := cmd.Flags().GetInt("port")
			if err != nil {
				return err
			}
			return Run(port)
		},
	}
	// 添加 flag, name=port, 默认值是 9090
	cmd.Flags().Int("port", 9090, "appserver port")
	return
}
