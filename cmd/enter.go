/**
 * @author X
 * @date 2023/4/2
 * @description
 **/
package cmd

import sys_flag "flag"

type Option struct {
	DB bool
}

// 解析命令参数
func Parse() Option {
	//flag.Bool()它返回布尔变量的地址，该变量存储定义的标志的值。
	db := sys_flag.Bool("db", false, "初始化数据库")
	//解析命令行参数写进注册的flag中
	sys_flag.Parse()
	return Option{
		*db,
	}

}

// 是否停掉web
func IsWebSrop(option Option) bool {
	if option.DB {
		return true
	}
	return false
}

func SwitchOption(option Option) {
	if option.DB {
		Makemigrations()
	}
}
