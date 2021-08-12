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
 * 系统路由
 * @author 半城风雨
 * @since 2021/7/26
 * @File : submit
 */
package router

import (
	"easygoadmin/app/controller"
	"easygoadmin/middleware"
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
)

func init() {
	s := g.Server()
	// 跨域处理
	s.Use(middleware.CORS)

	/* 文件上传 */
	s.Group("/upload", func(group *ghttp.RouterGroup) {
		// 上传图片
		group.POST("/uploadImage", controller.Upload.UploadImage)
	})

	/* 登录注册 */
	s.Group("/", func(group *ghttp.RouterGroup) {
		group.GET("/", controller.Login.Login)
		group.ALL("/login", controller.Login.Login)
		group.GET("/captcha", controller.Login.Captcha)
		group.GET("/index", controller.Index.Index)
		group.ALL("/updateUserInfo", controller.Index.UpdateUserInfo)
		group.ALL("/updatePwd", controller.Index.UpdatePwd)
		group.GET("/logout", controller.Index.Logout)
	})

	s.Group("index", func(group *ghttp.RouterGroup) {
		group.GET("/menu", controller.Index.Menu)
		group.GET("/user", controller.Index.User)
	})

	/* 用户管理 */
	s.Group("user", func(group *ghttp.RouterGroup) {
		group.GET("/list", controller.User.List)
		group.POST("/add", controller.User.Add)
		group.PUT("/update", controller.User.Update)
		group.DELETE("/delete/:ids", controller.User.Delete)
		group.PUT("/status", controller.User.Status)
		group.PUT("/resetPwd", controller.User.ResetPwd)
		group.GET("/checkUser", controller.User.CheckUser)
	})

	/* 职级管理 */
	s.Group("level", func(group *ghttp.RouterGroup) {
		group.GET("/list", controller.Level.List)
		group.POST("/add", controller.Level.Add)
		group.PUT("/update", controller.Level.Update)
		group.DELETE("/delete/:ids", controller.Level.Delete)
		group.PUT("/status", controller.Level.Status)
		group.GET("/getLevelList", controller.Level.GetLevelList)
	})

	/* 岗位路由 */
	s.Group("position", func(group *ghttp.RouterGroup) {
		group.GET("/list", controller.Position.List)
		group.POST("/add", controller.Position.Add)
		group.PUT("/update", controller.Position.Update)
		group.DELETE("/delete/:ids", controller.Position.Delete)
		group.PUT("/status", controller.Position.Status)
		group.GET("/getPositionList", controller.Position.GetPositionList)
	})

	/* 角色路由 */
	s.Group("role", func(group *ghttp.RouterGroup) {
		group.GET("/list", controller.Role.List)
		group.POST("/add", controller.Role.Add)
		group.PUT("/update", controller.Role.Update)
		group.DELETE("/delete/:ids", controller.Role.Delete)
		group.PUT("/status", controller.Role.Status)
		group.GET("/getRoleList", controller.Role.GetRoleList)
	})

	/* 角色菜单权限 */
	s.Group("rolemenu", func(group *ghttp.RouterGroup) {
		group.GET("/index", controller.RoleMenu.Index)
		group.POST("/save", controller.RoleMenu.Save)
	})

	/* 部门管理 */
	s.Group("dept", func(group *ghttp.RouterGroup) {
		group.GET("/list", controller.Dept.List)
		group.POST("/add", controller.Dept.Add)
		group.PUT("/update", controller.Dept.Update)
		group.DELETE("/delete/:ids", controller.Dept.Delete)
		group.GET("/getDeptList", controller.Dept.GetDeptList)
	})

	/* 菜单管理 */
	s.Group("menu", func(group *ghttp.RouterGroup) {
		group.GET("/list", controller.Menu.List)
		group.POST("/add", controller.Menu.Add)
		group.PUT("/update", controller.Menu.Update)
		group.DELETE("/delete/:ids", controller.Menu.Delete)
	})

	/* 城市管理 */
	s.Group("city", func(group *ghttp.RouterGroup) {
		group.GET("/list", controller.City.List)
		group.POST("/add", controller.City.Add)
		group.PUT("/update", controller.City.Update)
		group.DELETE("/delete/:ids", controller.City.Delete)
		group.POST("/getChilds", controller.City.GetChilds)
	})

	/* 字典管理 */
	s.Group("dict", func(group *ghttp.RouterGroup) {
		group.GET("/list", controller.Dict.List)
		group.POST("/add", controller.Dict.Add)
		group.PUT("/update", controller.Dict.Update)
		group.DELETE("/delete/:ids", controller.Dict.Delete)
	})

	/* 字典项管理 */
	s.Group("dictdata", func(group *ghttp.RouterGroup) {
		group.GET("/list", controller.DictData.List)
		group.POST("/add", controller.DictData.Add)
		group.PUT("/update", controller.DictData.Update)
		group.DELETE("/delete/:ids", controller.DictData.Delete)
	})

	/* 配置管理 */
	s.Group("config", func(group *ghttp.RouterGroup) {
		group.GET("/list", controller.Config.List)
		group.POST("/add", controller.Config.Add)
		group.PUT("/update", controller.Config.Update)
		group.DELETE("/delete/:ids", controller.Config.Delete)
	})

	/* 字典项管理 */
	s.Group("configdata", func(group *ghttp.RouterGroup) {
		group.GET("/list", controller.ConfigData.List)
		group.POST("/add", controller.ConfigData.Add)
		group.PUT("/update", controller.ConfigData.Update)
		group.DELETE("/delete/:ids", controller.ConfigData.Delete)
		group.PUT("/status", controller.ConfigData.Status)
	})

	/* 友链管理 */
	s.Group("link", func(group *ghttp.RouterGroup) {
		group.GET("/list", controller.Link.List)
		group.POST("/add", controller.Link.Add)
		group.PUT("/update", controller.Link.Update)
		group.DELETE("/delete/:ids", controller.Link.Delete)
		group.PUT("/status", controller.Link.Status)
	})

	/* 站点管理 */
	s.Group("item", func(group *ghttp.RouterGroup) {
		group.GET("/list", controller.Item.List)
		group.POST("/add", controller.Item.Add)
		group.PUT("/update", controller.Item.Update)
		group.DELETE("/delete/:ids", controller.Item.Delete)
		group.PUT("/status", controller.Item.Status)
		group.GET("/getItemList", controller.Item.GetItemList)
	})

	/* 栏目管理 */
	s.Group("itemcate", func(group *ghttp.RouterGroup) {
		group.GET("/list", controller.ItemCate.List)
		group.POST("/add", controller.ItemCate.Add)
		group.PUT("/update", controller.ItemCate.Update)
		group.DELETE("/delete/:ids", controller.ItemCate.Delete)
		//group.GET("/getCateTreeList", controller.ItemCate.GetCateTreeList)
		group.GET("/getCateList", controller.ItemCate.GetCateList)
	})

	/* 广告位管理 */
	s.Group("adsort", func(group *ghttp.RouterGroup) {
		group.GET("/list", controller.AdSort.List)
		group.POST("/add", controller.AdSort.Add)
		group.PUT("/update", controller.AdSort.Update)
		group.DELETE("/delete/:ids", controller.AdSort.Delete)
		group.GET("/getAdSortList", controller.AdSort.GetAdSortList)
	})

	/* 广告管理 */
	s.Group("ad", func(group *ghttp.RouterGroup) {
		group.GET("/list", controller.Ad.List)
		group.POST("/add", controller.Ad.Add)
		group.PUT("/update", controller.Ad.Update)
		group.DELETE("/delete/:ids", controller.Ad.Delete)
		group.PUT("/status", controller.Ad.Status)
	})

	/* 通知管理 */
	s.Group("notice", func(group *ghttp.RouterGroup) {
		group.GET("/list", controller.Notice.List)
		group.POST("/add", controller.Notice.Add)
		group.PUT("/update", controller.Notice.Update)
		group.DELETE("/delete/:ids", controller.Notice.Delete)
		group.PUT("/status", controller.Notice.Status)
	})

	/* 网站设置 */
	s.Group("configweb", func(group *ghttp.RouterGroup) {
		group.GET("/index", controller.ConfigWeb.Index)
		group.PUT("/save", controller.ConfigWeb.Save)
	})

	/* 会员等级 */
	s.Group("memberlevel", func(group *ghttp.RouterGroup) {
		group.GET("/list", controller.MemberLevel.List)
		group.POST("/add", controller.MemberLevel.Add)
		group.PUT("/update", controller.MemberLevel.Update)
		group.DELETE("/delete/:ids", controller.MemberLevel.Delete)
		group.GET("/getMemberLevelList", controller.MemberLevel.GetMemberLevelList)
	})

	/* 会员管理 */
	s.Group("member", func(group *ghttp.RouterGroup) {
		group.GET("/list", controller.Member.List)
		group.POST("/add", controller.Member.Add)
		group.PUT("/update", controller.Member.Update)
		group.DELETE("/delete/:ids", controller.Member.Delete)
		group.PUT("/status", controller.Member.Status)
	})

	/* 统计分析 */
	s.Group("analysis", func(group *ghttp.RouterGroup) {
		group.GET("/index", controller.Analysis.Index)
	})

	/* 代码生成器 */
	s.Group("generate", func(group *ghttp.RouterGroup) {
		group.GET("/list", controller.Generate.List)
		group.POST("/generate", controller.Generate.Generate)
	})

}
