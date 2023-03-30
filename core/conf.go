/**
 * @author X
 * @date 2023/3/30
 * @description
 **/
package core

import (
	"fmt"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"log"
)

type Server struct {
	Mysql struct {
		Host string `yaml:"host"`
		Port string `yaml:"port"`
	} `yaml:"mysql"`
}

// 读取yaml文件配置
func InitConf() {
	const ConfigFile = "settings.yaml"
	c := &Server{}
	yamlConf, err := ioutil.ReadFile(ConfigFile)
	if err != nil {
		panic(fmt.Errorf("get yaml config err:%s", err))
	}
	err = yaml.Unmarshal(yamlConf, c)
	if err != nil {
		log.Fatalf("config Init Unmarshal:%v ", err)
	}
	log.Println("config yaml file log init success")
	fmt.Println(c)
}
