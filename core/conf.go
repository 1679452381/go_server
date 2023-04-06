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
	"io/fs"
	"log"
	"os"
)

const ConfigFile = "settings.yaml"

// 读取yaml文件配置
func InitConf() {
	c := &config.Config{}
	yamlConf, err := os.ReadFile(ConfigFile)
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

// 修改yaml文件
func SetConf() error {
	//转化为yaml格式数据
	byteData, err := yaml.Marshal(global.Config)
	if err != nil {
		return err
	}
	err = os.WriteFile(ConfigFile, byteData, fs.ModePerm)
	if err != nil {
		return err
	}
	global.Log.Info("配置yaml文件成功")
	return nil
}
