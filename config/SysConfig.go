package config

import (
	"fmt"
	"gopkg.in/yaml.v3"
	"log"
	"os"
	"path/filepath"
)

var SysYamlconfig *SysConfig

type ServerConfig struct {
	Port int32  `yaml:"port"`
	Name string `yaml:"name"`
}

type DatabaseConfig struct {
	Dsn      string `yaml:"dsn"`
	Host     string `yaml:"host"`
	Port     int32  `yaml:"port"`
	Username string `yaml:"username"`
	Password string `yaml:"password"`
	Database string `yaml:"database"`
}

type JwtConfig struct {
	PublicKey  string `yaml:"public_key"`
	PrivateKey string `yaml:"private_key"`
}

type SysConfig struct {
	Server   *ServerConfig   `yaml:"server"`
	Database *DatabaseConfig `yaml:"database"`
	Jwt      *JwtConfig      `yaml:"jwt"`
}

func DefaultSysConfig() *SysConfig {
	return &SysConfig{Server: &ServerConfig{

		Port: 8080,
		Name: "admin-Web-Api",
	}, Database: &DatabaseConfig{
		Dsn:      "",
		Host:     "127.0.0.1",
		Port:     3306,
		Username: "root",
		Password: "root",
		Database: "root",
	}}
}

func InitConfig(yamlPath string) *SysConfig {
	// read config file
	yamlFile, err := os.ReadFile(yamlPath)
	if err != nil {
		log.Fatal("failed to read YAML file", err)
	}

	// default
	config := DefaultSysConfig()
	err = yaml.Unmarshal(yamlFile, &config)

	if err != nil {
		log.Fatal("failed to unmarshal YAML file", err)
	}

	return config
}

func init() {
	currentDir, _ := os.Getwd()
	resultDir := filepath.Join(currentDir, "application.yaml")
	fmt.Println("resultDir = ", resultDir)
	SysYamlconfig = InitConfig(resultDir)
}
