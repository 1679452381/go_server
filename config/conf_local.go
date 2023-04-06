/**
 * @author X
 * @date 2023/4/6
 * @description
 **/
package config

type Local struct {
	Path     string `yaml:"path" json:"path"`
	FileSize int    `yaml:"file_size" json:"file_size"`
}
