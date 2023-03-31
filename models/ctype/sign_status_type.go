/**
 * @author X
 * @date 2023/4/1
 * @description
 **/
package ctype

import "encoding/json"

type SignStatus int

const (
	SignQQ    SignStatus = 1
	SignGitee SignStatus = 2
	SignEmail SignStatus = 3
)

var SignStatusMap = map[SignStatus]string{
	SignQQ:    "QQ登录",
	SignGitee: "Gitee登录",
	SignEmail: "Email登录",
}

func (s SignStatus) MarshalJSON() ([]byte, error) {
	return json.Marshal(s.String())
}
func (s SignStatus) String() string {
	var str string
	switch s {
	case SignQQ:
		str = SignStatusMap[SignQQ]
	case SignGitee:
		str = SignStatusMap[SignGitee]
	case SignEmail:
		str = SignStatusMap[SignEmail]
	default:
		str = "其他"
	}
	return str
}
