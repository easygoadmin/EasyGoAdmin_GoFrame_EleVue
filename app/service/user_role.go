// +----------------------------------------------------------------------
// | EasyGoAdmin敏捷开发框架 [ EasyGoAdmin ]
// +----------------------------------------------------------------------
// | 版权所有 2021 EasyGoAdmin深圳研发中心
// +----------------------------------------------------------------------
// | 官方网站: http://www.easygoadmin.vip
// +----------------------------------------------------------------------
// | Author: 半城风雨 <easygoadmin@163.com>
// +----------------------------------------------------------------------

/**
 * 用户角色-服务类
 * @author 半城风雨
 * @since 2021/7/27
 * @File : user
 */
package service

import (
	"easygoadmin/app/dao"
	"easygoadmin/app/model"
)

var UserRole = new(userRoleService)

type userRoleService struct{}

// 获取用户角色列表
func (s *userRoleService) GetUserRoleList(userId int) []model.Role {
	query := dao.Role.As("r").Clone()
	// 内联查询
	query = query.InnerJoin("sys_user_role as ur", "r.id=ur.role_id")
	query = query.Where("ur.user_id=? AND r.mark=1", userId)
	// 获取字段
	query.Fields("r.*")
	// 排序
	query = query.Order("r.sort asc")
	// 数据转换
	var list []model.Role
	query.Structs(&list)
	return list
}
