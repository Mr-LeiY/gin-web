package cache_service

import (
	"gin-web/models"
	"gin-web/pkg/global"
	"gin-web/pkg/request"
	"strings"
)

func (rd RedisService) FindRoleIdBySort(currentRoleSort uint) []uint {
	if !global.Conf.Redis.Enable || !global.Conf.Redis.EnableService {
		return rd.mysql.FindRoleIdBySort(currentRoleSort)
	}
	roles := make([]models.SysRole, 0)
	roleIds := make([]uint, 0)
	rd.Q.
		Table("sys_role").
		Where("sort", ">=", currentRoleSort).
		Find(&roles)
	for _, role := range roles {
		roleIds = append(roleIds, role.Id)
	}
	return roleIds
}

func (rd RedisService) FindRole(req *request.RoleReq) []models.SysRole {
	if !global.Conf.Redis.Enable || !global.Conf.Redis.EnableService {
		return rd.mysql.FindRole(req)
	}
	list := make([]models.SysRole, 0)
	query := rd.Q.
		Table("sys_role").
		Order("created_at DESC").
		Where("sort", ">=", req.CurrentRoleSort)
	name := strings.TrimSpace(req.Name)
	if name != "" {
		query = query.Where("name", "contains", name)
	}
	keyword := strings.TrimSpace(req.Keyword)
	if keyword != "" {
		query = query.Where("keyword", "contains", keyword)
	}
	if req.Status != nil {
		if *req.Status > 0 {
			query = query.Where("status", "=", 1)
		} else {
			query = query.Where("status", "=", 0)
		}
	}
	rd.Q.FindWithPage(query, &req.Page, &list)
	return list
}
