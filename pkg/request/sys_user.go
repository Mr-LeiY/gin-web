package request

import (
	"gin-web/models"
	"github.com/piupuer/go-helper/pkg/req"
	"github.com/piupuer/go-helper/pkg/resp"
)

// User login structure
type RegisterAndLoginReq struct {
	Username string `json:"username" validate:"required"`
	Password string `json:"password" validate:"required"`
}

func (s RegisterAndLoginReq) FieldTrans() map[string]string {
	m := make(map[string]string, 0)
	m["Username"] = "用户名"
	m["Password"] = "密码"
	return m
}

type ChangePwdReq struct {
	OldPassword string `json:"oldPassword" form:"oldPassword"`
	NewPassword string `json:"newPassword" form:"newPassword"`
}

type UserReq struct {
	CurrentRole  models.SysRole `json:"currentRole"`
	Username     string         `json:"username" form:"username"`
	Mobile       string         `json:"mobile" form:"mobile"`
	Avatar       string         `json:"avatar" form:"avatar"`
	Nickname     string         `json:"nickname" form:"nickname"`
	UsernameOr   string         `json:"usernameOr" form:"usernameOr"`
	MobileOr     string         `json:"mobileOr" form:"mobileOr"`
	NicknameOr   string         `json:"nicknameOr" form:"nicknameOr"`
	Introduction string         `json:"introduction" form:"introduction"`
	Status       *uint          `json:"status" form:"status"`
	RoleId       uint           `json:"roleId" form:"roleId"`
	resp.Page
}

type CreateUserReq struct {
	Username     string        `json:"username" validate:"required"`
	Password     string        `json:"password"`
	InitPassword string        `json:"initPassword" validate:"required"` // 不使用SysUser的Password字段, 避免请求劫持绕过系统校验
	NewPassword  string        `json:"newPassword"`
	Mobile       string        `json:"mobile" validate:"required"`
	Avatar       string        `json:"avatar"`
	Nickname     string        `json:"nickname"`
	Introduction string        `json:"introduction"`
	Status       *req.NullUint `json:"status"`
	RoleId       uint          `json:"roleId" validate:"required"`
}

func (s CreateUserReq) FieldTrans() map[string]string {
	m := make(map[string]string, 0)
	m["Username"] = "用户名"
	m["InitPassword"] = "初始密码"
	m["Mobile"] = "手机号"
	m["RoleId"] = "角色"
	return m
}

type UpdateUserReq struct {
	Username     *string       `json:"username"`
	Password     *string       `json:"password"`
	InitPassword *string       `json:"initPassword"`
	NewPassword  *string       `json:"newPassword"`
	Mobile       *string       `json:"mobile"`
	Avatar       *string       `json:"avatar"`
	Nickname     *string       `json:"nickname"`
	Introduction *string       `json:"introduction"`
	Status       *req.NullUint `json:"status"`
	RoleId       *uint         `json:"roleId"`
}
