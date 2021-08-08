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
 * 角色菜单-控制器
 * @author 半城风雨
 * @since 2021/7/15
 * @File : rolemenu
 */
package controller

import (
	"easygoadmin/app/model"
	"easygoadmin/app/service"
	"easygoadmin/app/utils/common"
	"github.com/gogf/gf/net/ghttp"
)

// 控制器管理对象
var RoleMenu = new(roleMenuCtl)

type roleMenuCtl struct{}

func (c *roleMenuCtl) Index(r *ghttp.Request) {
	// 角色ID
	roleId := r.GetQueryInt64("roleId")
	if roleId <= 0 {
		r.Response.WriteJsonExit(common.JsonResult{
			Code: -1,
			Msg:  "角色ID不能为空",
		})
	}

	// 获取角色菜单权限列表
	list, err := service.RoleMenu.GetRoleMenuList(roleId)
	if err != nil {
		r.Response.WriteJsonExit(common.JsonResult{
			Code: -1,
			Msg:  err.Error(),
		})
	}
	r.Response.WriteJsonExit(common.JsonResult{
		Code: 0,
		Data: list,
		Msg:  "操作成功",
	})
}

func (c *roleMenuCtl) Save(r *ghttp.Request) {
	var req *model.RoleMenuSaveReq
	if err := r.Parse(&req); err != nil {
		r.Response.WriteJsonExit(common.JsonResult{
			Code: -1,
			Msg:  err.Error(),
		})
	}

	// 调用保存方法
	err := service.RoleMenu.Save(req)
	if err != nil {
		r.Response.WriteJsonExit(common.JsonResult{
			Code: -1,
			Msg:  err.Error(),
		})
	}
	r.Response.WriteJsonExit(common.JsonResult{
		Code: 0,
		Msg:  "保存成功",
	})
}
