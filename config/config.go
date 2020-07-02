package config

import (
	"fmt"
	"gopkg.in/yaml.v2"
	"io/ioutil"
)

type DBC struct {
	Dbms string `yaml:"dbms"`
	DbHost string `yaml:"dbHost"`
	DbPort int `yaml:"dbPort"`
	DbUser string `yaml:"dbUser"`
	DbPwd string `yaml:"dbPwd"`
	DbName string `yaml:"dbName"`
	DbCharSet string `yaml:"dbCharSet"`
}


type Config struct {
	Version string `yaml:"version"`
	DbConf DBC `yaml:"dbConf"`
}

func readConfFile()[]byte{
	body,err := ioutil.ReadFile("./config/config.yml")
	if err != nil{
		panic("读取配置文件错误")
	}
	return body
}

func LoadConfig()Config{
	var config Config
	err:= yaml.Unmarshal(readConfFile(),&config)
	if err != nil{
		panic(fmt.Sprintf("读取配置文件错误:%s",err))
	}
	return config
}