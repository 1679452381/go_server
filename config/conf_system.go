/**
 * @author X
 * @date 2023/3/30
 * @description
 **/
package config

type System struct {
	Host string `yaml:"host"`
	Port int    `yaml:"port"`
	Env  string `yaml:"env"`
}
