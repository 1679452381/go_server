/**
 * @author X
 * @date 2023/4/6
 * @description
 **/
package config

type QiNiu struct {
	Enable    string `yaml:"enable" json:"enable"` //启用七牛存储
	AccessKey string `yaml:"accessKey" json:"access_key"`
	SecretKey string `yaml:"secretKey" json:"secret_key"`
	Bucket    string `yaml:"bucket" json:"bucket"` //存储桶的名字
	CDN       string `yaml:"cdn" json:"cdn"`       // 访问图片地址前缀
	Zone      string `yaml:"zone" json:"zone"`     //存储的地区
	Size      string `yaml:"size" json:"size"`     //存储的大小 单位MB
}
