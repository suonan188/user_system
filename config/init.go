package config

import (
	"fmt"
	"io/ioutil"
	"os"

	"gopkg.in/yaml.v2"
)

//InitConfig :初始化全局配置
func InitConfig() {
	ENV := os.Getenv("ENV")

	if ENV == "" {
		ENV = "dev"
	}
	//读取配置文件
	file, err := ioutil.ReadFile("config/" + ENV + ".yaml")
	if err != nil {
		panic("文件读取失败")
	}
	//解析配置文件
	err = yaml.Unmarshal(file, &Config)
	if err != nil {
		panic("解析配置文件失败")
	}
	//打印配置
	p, _ := yaml.Marshal(&Config)
	fmt.Printf("--Config:\n%s\n", string(p))

}
