package main

import (
	"com.github.goscaffold/pkg/utils"
	"log"
)

func main() {
	err := utils.GenRSAPubAndPri(1024, "./resources/pem")
	if err != nil {
		log.Fatal(err)
	}
}
