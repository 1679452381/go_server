/**
 * @author X
 * @date 2023/3/30
 * @description
 **/
package config

import "fmt"

type Mysql struct {
	Host     string `yaml:"host"`
	Port     int    `yaml:"port"`
	Config   string `yaml:"config"` // 高级配置 charset
	Db       string `yaml:"db"`
	User     string `yaml:"user"`
	Password string `yaml:"password"`
	LogLevel string `yaml:"log_level"`
}

func (m Mysql) Dsn() string {
	return fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?%s", m.User, m.Password, m.Host, m.Port, m.Db, m.Config)
}
