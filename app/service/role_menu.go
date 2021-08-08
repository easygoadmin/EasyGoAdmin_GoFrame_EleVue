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
 * 角色菜单-服务类
 * @author 半城风雨
 * @since 2021/7/15
 * @File : rolemenu
 */
package service

import (
	"easygoadmin/app/dao"
	"easygoadmin/app/model"
	"easygoadmin/app/utils"
	"easygoadmin/app/utils/function"
	"github.com/gogf/gf/errors/gerror"
	"github.com/gogf/gf/util/gconv"
	"github.com/gogf/gf/util/gutil"
)

var RoleMenu = new(roleMenuService)

type roleMenuService struct{}

func (s *roleMenuService) GetRoleMenuList(roleId int64) ([]model.RoleMenuInfo, error) {
	// 获取全部菜单列表
	var menuList []model.Menu
	dao.Menu.Where("status=1").Where("mark=1").Order("sort asc").Structs(&menuList)
	if len(menuList) == 0 {
		return nil, gerror.New("菜单列表不存在")
	}
	// 获取角色菜单权限列表
	var roleMenuList []model.RoleMenu
	dao.RoleMenu.Where("role_id=?", roleId).Structs(&roleMenuList)
	idList := gutil.ListItemValuesUnique(&roleMenuList, "MenuId")

	// 对象处理
	var list []model.RoleMenuInfo
	if len(menuList) > 0 {
		for _, m := range menuList {
			var info model.RoleMenuInfo
			info.Id = m.Id
			info.Title = m.Title
			info.Open = true
			info.ParentId = m.ParentId
			// 节点选中值
			if function.InArray(gconv.String(m.Id), idList) {
				info.Checked = true
			}
			list = append(list, info)
		}
	}
	return list, nil
}

func (s *roleMenuService) Save(req *model.RoleMenuSaveReq) error {
	if utils.AppDebug() {
		return gerror.New("演示环境，暂无权限操作")
	}
	itemArr := req.MenuIds
	if len(itemArr) == 0 {
		return gerror.New("请选择权限节点")
	}
	// 删除现有的角色权限数据
	dao.RoleMenu.Delete("role_id=?", req.RoleId)
	// 遍历创建新角色权限数据
	for _, v := range itemArr {
		var entity model.RoleMenu
		entity.RoleId = req.RoleId
		entity.MenuId = v
		dao.RoleMenu.Insert(entity)
	}
	// 批量插入
	//dao.RoleMenu.Data(list).Batch(2).Insert()
	return nil
}
