/**
 * @author X
 * @date 2023/3/30
 * @description
 **/
package config

type Config struct {
	Mysql    Mysql    `yaml:"mysql"`
	Logger   Logger   `yaml:"logger"`
	System   System   `yaml:"system"`
	SiteInfo SiteInfo `yaml:"site_info"`
	QiNiu    QiNiu    `yaml:"qi_niu"`
	Local    Local    `yaml:"local"`
}
