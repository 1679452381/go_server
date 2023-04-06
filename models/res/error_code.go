/**
 * @author X
 * @date 2023/3/31
 * @description
 **/
package res

type ErrorCode int

const (
	SettingErr ErrorCode = 1001 //系统错误
	ParamsErr  ErrorCode = 1002
)

var (
	ErrorMap = map[ErrorCode]string{
		SettingErr: "系统错误",
		ParamsErr:  "参数错误",
	}
)
