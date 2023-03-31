/**
 * @author X
 * @date 2023/3/31
 * @description
 **/
package res

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type Resonse struct {
	Code int
	Data map[string]interface{}
	Msg  string
}

const (
	Success = 0
	Error   = 1
)

func Result(code int, data map[string]interface{}, msg string, c *gin.Context) {
	c.JSON(http.StatusOK, Resonse{
		code,
		data,
		msg,
	})
}

func Ok(data map[string]interface{}, msg string, c *gin.Context) {
	Result(Success, data, msg, c)
}
func OkWithMsg(msg string, c *gin.Context) {
	Result(Success, map[string]interface{}{}, msg, c)
}
func OkWithData(data map[string]interface{}, c *gin.Context) {
	Result(Success, data, "成功", c)
}

func Fail(data map[string]interface{}, msg string, c *gin.Context) {
	Result(Error, data, msg, c)
}
func FailWithMsg(msg string, c *gin.Context) {
	Result(Error, map[string]interface{}{}, msg, c)
}
func FailWithData(data map[string]interface{}, c *gin.Context) {
	Result(Error, data, "成功", c)
}
func FailWithCode(code ErrorCode, c *gin.Context) {
	msg, ok := ErrorMap[code]
	if ok {
		Result(int(code), map[string]interface{}{}, msg, c)
		return
	}
	Result(Error, map[string]interface{}{}, "未知错误", c)

}
