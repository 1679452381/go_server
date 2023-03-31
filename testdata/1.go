/**
 * @author X
 * @date 2023/4/1
 * @description
 **/
package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	res, _ := json.Marshal("zxc")
	fmt.Println(string(res))
}
