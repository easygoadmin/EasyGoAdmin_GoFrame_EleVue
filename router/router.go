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
		group.GET("/main", controller.Index.Main)
		group.ALL("/userInfo", controller.Index.UserInfo)
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
		group.DELETE("/delete", controller.User.Delete)
		group.PUT("/status", controller.User.Status)
		group.PUT("/resetPwd", controller.User.ResetPwd)
	})

	/* 职级管理 */
	s.Group("level", func(group *ghttp.RouterGroup) {
		group.GET("/list", controller.Level.List)
		group.POST("/add", controller.Level.Add)
		group.PUT("/update", controller.Level.Update)
		group.DELETE("/delete", controller.Level.Delete)
		group.PUT("/status", controller.Level.Status)
	})

	/* 岗位路由 */
	s.Group("position", func(group *ghttp.RouterGroup) {
		group.GET("/list", controller.Position.List)
		group.POST("/add", controller.Position.Add)
		group.PUT("/update", controller.Position.Update)
		group.DELETE("/delete", controller.Position.Delete)
	})

	/* 角色路由 */
	s.Group("role", func(group *ghttp.RouterGroup) {
		group.GET("/list", controller.Role.List)
		group.POST("/add", controller.Role.Add)
		group.PUT("/update", controller.Role.Update)
		group.DELETE("/delete", controller.Role.Delete)
		group.PUT("/status", controller.Role.Status)
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
		group.DELETE("/delete", controller.Dept.Delete)
	})

	/* 菜单管理 */
	s.Group("menu", func(group *ghttp.RouterGroup) {
		group.GET("/list", controller.Menu.List)
		group.POST("/add", controller.Menu.Add)
		group.PUT("/update", controller.Menu.Update)
		group.DELETE("/delete", controller.Menu.Delete)
	})

	/* 城市管理 */
	s.Group("city", func(group *ghttp.RouterGroup) {
		group.GET("/list", controller.City.List)
		group.POST("/add", controller.City.Add)
		group.PUT("/update", controller.City.Update)
		group.DELETE("/delete", controller.City.Delete)
		group.POST("/getChilds", controller.City.GetChilds)
	})

	/* 字典管理 */
	s.Group("dict", func(group *ghttp.RouterGroup) {
		group.GET("/list", controller.Dict.List)
		group.POST("/add", controller.Dict.Add)
		group.PUT("/update", controller.Dict.Update)
		group.DELETE("/delete", controller.Dict.Delete)
	})

	/* 字典项管理 */
	s.Group("dictdata", func(group *ghttp.RouterGroup) {
		group.GET("/list", controller.DictData.List)
		group.POST("/add", controller.DictData.Add)
		group.PUT("/update", controller.DictData.Update)
		group.DELETE("/delete", controller.DictData.Delete)
	})

	/* 配置管理 */
	s.Group("config", func(group *ghttp.RouterGroup) {
		group.GET("/list", controller.Config.List)
		group.POST("/add", controller.Config.Add)
		group.PUT("/update", controller.Config.Update)
		group.DELETE("/delete", controller.Config.Delete)
	})

	/* 字典项管理 */
	s.Group("configdata", func(group *ghttp.RouterGroup) {
		group.GET("/list", controller.ConfigData.List)
		group.POST("/add", controller.ConfigData.Add)
		group.PUT("/update", controller.ConfigData.Update)
		group.DELETE("/delete", controller.ConfigData.Delete)
		group.PUT("/status", controller.ConfigData.Status)
	})

	/* 友链管理 */
	s.Group("link", func(group *ghttp.RouterGroup) {
		group.GET("/list", controller.Link.List)
		group.POST("/add", controller.Link.Add)
		group.PUT("/update", controller.Link.Update)
		group.DELETE("/delete", controller.Link.Delete)
		group.PUT("/status", controller.Link.Status)
	})

	/* 站点管理 */
	s.Group("item", func(group *ghttp.RouterGroup) {
		group.GET("/list", controller.Item.List)
		group.POST("/add", controller.Item.Add)
		group.PUT("/update", controller.Item.Update)
		group.DELETE("/delete", controller.Item.Delete)
	})

	/* 栏目管理 */
	s.Group("itemcate", func(group *ghttp.RouterGroup) {
		group.GET("/list", controller.ItemCate.List)
		group.POST("/add", controller.ItemCate.Add)
		group.PUT("/update", controller.ItemCate.Update)
		group.DELETE("/delete", controller.ItemCate.Delete)
		group.GET("/getCateTreeList", controller.ItemCate.GetCateTreeList)
	})

	/* 广告位管理 */
	s.Group("adsort", func(group *ghttp.RouterGroup) {
		group.GET("/list", controller.AdSort.List)
		group.POST("/add", controller.AdSort.Add)
		group.PUT("/update", controller.AdSort.Update)
		group.DELETE("/delete", controller.AdSort.Delete)
	})

	/* 广告管理 */
	s.Group("ad", func(group *ghttp.RouterGroup) {
		group.GET("/list", controller.Ad.List)
		group.POST("/add", controller.Ad.Add)
		group.PUT("/update", controller.Ad.Update)
		group.DELETE("/delete", controller.Ad.Delete)
		group.PUT("/status", controller.Ad.Status)
	})

	/* 通知管理 */
	s.Group("notice", func(group *ghttp.RouterGroup) {
		group.GET("/list", controller.Notice.List)
		group.POST("/add", controller.Notice.Add)
		group.PUT("/update", controller.Notice.Update)
		group.DELETE("/delete", controller.Notice.Delete)
	})

	/* 网站设置 */
	s.Group("configweb", func(group *ghttp.RouterGroup) {
		group.ALL("/index", controller.ConfigWeb.Index)
	})

	/* 会员等级 */
	s.Group("memberlevel", func(group *ghttp.RouterGroup) {
		group.GET("/list", controller.MemberLevel.List)
		group.POST("/add", controller.MemberLevel.Add)
		group.PUT("/update", controller.MemberLevel.Update)
		group.DELETE("/delete", controller.MemberLevel.Delete)
	})

	/* 会员管理 */
	s.Group("member", func(group *ghttp.RouterGroup) {
		group.GET("/list", controller.Member.List)
		group.POST("/add", controller.Member.Add)
		group.PUT("/update", controller.Member.Update)
		group.DELETE("/delete", controller.Member.Delete)
		group.PUT("/status", controller.Member.Status)
	})

	/* 统计分析 */
	s.Group("analysis", func(group *ghttp.RouterGroup) {
		group.GET("/index", controller.Analysis.Index)
	})

	/* 代码生成器 */
	s.Group("generate", func(group *ghttp.RouterGroup) {
		group.GET("/index", controller.Generate.Index)
		group.GET("/list", controller.Generate.List)
		group.POST("/generate", controller.Generate.Generate)
	})

}
