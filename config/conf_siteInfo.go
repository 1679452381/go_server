/**
 * @author X
 * @date 2023/4/5
 * @description
 **/
package config

type SiteInfo struct {
	CreatedAt   string `yaml:"created_at" json:"created_at"`
	BeiAN       string `yaml:"bei_an" json:"bei_an"`
	QQImage     string `yaml:"qq_image" json:"qq_image"`
	WechatImage string `yaml:"wechat_image" json:"wechat_image"`
	Version     string `yaml:"version" json:"version"`
	Email       string `yaml:"email" json:"email"`
	Name        string `yaml:"name" json:"name"`
	Job         string `yaml:"job" json:"job"`
	Addr        string `yaml:"addr" json:"addr"`
	Slogan      string `yaml:"slogan" json:"slogan"`
	SloganEn    string `yaml:"sloganEn" json:"slogan_en"`
	Web         string `yaml:"web" json:"web"`
	GithubUrl   string `yaml:"github_url" json:"github_url"`
}
