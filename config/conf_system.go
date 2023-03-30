/**
 * @author X
 * @date 2023/3/30
 * @description
 **/
package config

import "fmt"

type System struct {
	Host string `yaml:"host"`
	Port int    `yaml:"port"`
	Env  string `yaml:"env"`
}

func (s *System) Addr() string {
	fmt.Println(s.Port)
	return fmt.Sprintf("%s:%d", s.Host, s.Port)
}
