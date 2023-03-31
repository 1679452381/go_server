/**
 * @author X
 * @date 2023/3/31
 * @description
 **/
package res

type ErrorCode int

const (
	SettingErr ErrorCode = 1001 //系统错误
)

var (
	ErrorMap = map[ErrorCode]string{
		SettingErr: "系统错误",
	}
)
