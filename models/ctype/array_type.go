/**
 * @author X
 * @date 2023/4/2
 * @description
 **/
package ctype

import (
	"database/sql/driver"
	"strings"
)

type Array []string

func (arr *Array) Scan(value interface{}) error {
	v, _ := value.([]byte)
	if string(v) == "" {
		*arr = []string{}
		return nil
	}
	*arr = strings.Split(string(v), "\n")
	return nil
}

func (arr Array) Value() (driver.Value, error) {
	//	将数值转化为字符串
	return strings.Join(arr, "\n"), nil
}
