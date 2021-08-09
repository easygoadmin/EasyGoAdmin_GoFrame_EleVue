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
 * 系统主页
 * @author 半城风雨
 * @since 2021/5/19
 * @File : index
 */
package controller

import (
	"easygoadmin/app/dao"
	"easygoadmin/app/model"
	"easygoadmin/app/service"
	"easygoadmin/app/utils"
	"easygoadmin/app/utils/common"
	"easygoadmin/app/utils/function"
	"easygoadmin/app/utils/response"
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
)

// 用户API管理对象
var Index = new(indexCtl)

type indexCtl struct{}

// 获取系统菜单
func (c *indexCtl) Menu(r *ghttp.Request) {
	// Jwt获取用户ID
	userId := 1
	// 获取菜单列表
	menuList := service.Menu.GetPermissionList(userId)
	// 返回结果
	r.Response.WriteJsonExit(common.JsonResult{
		Code: 0,
		Msg:  "操作成功",
		Data: menuList,
	})
}

func (c *indexCtl) User(r *ghttp.Request) {
	// 获取用户信息
	userInfo, _ := dao.User.FindOne(utils.Uid(r.Session))
	// 头像
	userInfo.Avatar = utils.GetImageUrl(userInfo.Avatar)
	// 返回结果
	r.Response.WriteJsonExit(common.JsonResult{
		Code: 0,
		Msg:  "操作成功",
		Data: userInfo,
	})
}

func (c *indexCtl) Index(r *ghttp.Request) {
	if function.IsLogin(r.Session) {
		// 已登录
		// 获取用户信息
		userInfo := service.Login.GetProfile(r.Session)
		// 获取菜单列表
		menuList := service.Menu.GetPermissionList(userInfo.Id)
		// 渲染模板并绑定数据
		response.BuildTpl(r, "index.html").WriteTpl(g.Map{
			"userInfo": userInfo,
			"menuList": menuList,
		})
	} else {
		// 未登录
		r.Response.RedirectTo("/login")
	}
}

func (c *indexCtl) Main(r *ghttp.Request) {
	response.BuildTpl(r, "main.html").WriteTpl()
}

// 个人中心
func (c *indexCtl) UserInfo(r *ghttp.Request) {
	if r.Method == "PUT" {
		// 参数验证
		var req *model.UserInfoReq
		if err := r.Parse(&req); err != nil {
			r.Response.WriteJsonExit(common.JsonResult{
				Code: -1,
				Msg:  err.Error(),
			})
		}
		// 更新信息
		_, err := service.User.UpdateUserInfo(req, r.Session)
		if err != nil {
			r.Response.WriteJsonExit(common.JsonResult{
				Code: -1,
				Msg:  err.Error(),
			})
		}

		// 返回结果
		r.Response.WriteJsonExit(common.JsonResult{
			Code: 0,
			Msg:  "更新成功",
		})
	}
	// 获取用户信息
	userInfo := service.Login.GetProfile(r.Session)
	// 渲染模板
	response.BuildTpl(r, "public/layout.html").WriteTpl(g.Map{
		"mainTpl":  "user_info.html",
		"userInfo": userInfo,
	})
}

// 更新密码
func (c *indexCtl) UpdatePwd(r *ghttp.Request) {
	// 参数验证
	var req *model.UpdatePwd
	if err := r.Parse(&req); err != nil {
		r.Response.WriteJsonExit(common.JsonResult{
			Code: -1,
			Msg:  err.Error(),
		})
	}

	// 调用更新密码方法
	rows, err := service.User.UpdatePwd(req, utils.Uid(r.Session))
	if err != nil || rows == 0 {
		r.Response.WriteJsonExit(common.JsonResult{
			Code: -1,
			Msg:  err.Error(),
		})
	}

	// 返回结果
	r.Response.WriteJsonExit(common.JsonResult{
		Code: 0,
		Msg:  "更新密码成功",
	})
}

// 退出登录
func (c *indexCtl) Logout(r *ghttp.Request) {
	// 设置SESSION为空
	r.Session.Set("userId", "")
	r.Session.Set("userInfo", "")
	// 重定向至登录页面
	r.Response.RedirectTo("/login")
}
