/**
 * @author X
 * @date 2023/4/1
 * @description
 **/
package ctype

import "encoding/json"

type Role int

const (
	PermissionAdmin       Role = 1
	PermissionUser        Role = 2
	PermissionVisitor     Role = 3
	PermissionDisableUser Role = 4
)

var AuthMap = map[Role]string{
	PermissionAdmin:       "管理员",
	PermissionUser:        "用户",
	PermissionVisitor:     "游客",
	PermissionDisableUser: "被禁言",
}

func (role Role) MarshalJSON() ([]byte, error) {
	return json.Marshal(role.String())
}

func (role Role) String() string {
	var str string
	switch role {
	case PermissionAdmin:
		str = AuthMap[PermissionAdmin]
	case PermissionUser:
		str = AuthMap[PermissionUser]
	case PermissionVisitor:
		str = AuthMap[PermissionVisitor]
	case PermissionDisableUser:
		str = AuthMap[PermissionDisableUser]
	default:
		str = "其他"
	}
	return str
}
