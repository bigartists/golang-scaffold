package main

import (
	"com.github.goscaffold/cmd/app"
	"com.github.goscaffold/config"
	_ "com.github.goscaffold/config"
	"fmt"
)

func main() {
	cmd := app.NewApiServerCommand()
	fmt.Println("config.SysYamlconfig.Server.Name = ", config.SysYamlconfig.Server.Name)
	cmd.Execute()
}

// go run cmd/apiserver.go apiserver --port=8888
