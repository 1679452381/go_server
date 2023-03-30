/**
 * @author X
 * @date 2023/3/30
 * @description
 **/
package core

import (
	"fmt"
	"go_server/config"
	"go_server/global"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"log"
)

// 读取yaml文件配置
func InitConf() {
	const ConfigFile = "settings.yaml"
	c := &config.Config{}
	yamlConf, err := ioutil.ReadFile(ConfigFile)
	if err != nil {
		panic(fmt.Errorf("get yaml config err:%s", err))
	}
	err = yaml.Unmarshal(yamlConf, c)
	if err != nil {
		log.Fatalf("config Init Unmarshal:%v ", err)
	}
	log.Println("config yaml file log init success")
	//定义全局变量存放配置数据
	global.Config = c
}
