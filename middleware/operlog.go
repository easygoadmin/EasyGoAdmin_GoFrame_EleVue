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
 * 操作日志
 * @author 半城风雨
 * @since 2021/8/13
 * @File : operlog
 */
package middleware

import (
	"easygoadmin/app/dao"
	"easygoadmin/app/model"
	"easygoadmin/app/utils"
	"fmt"
	"github.com/gogf/gf/net/ghttp"
	"github.com/gogf/gf/os/gtime"
	"github.com/gogf/gf/text/gstr"
)

func OperLog(r *ghttp.Request) {
	// 后置中间件
	r.Middleware.Next()
	// 中间件处理逻辑
	fmt.Println("操作日志中间件")

	// 分析请求URL地址
	urlArr := gstr.Split(r.URL.String(), "?")
	urlItem := gstr.Split(urlArr[0], "/")
	if len(urlItem) < 3 {
		return
	}
	// 拼接节点
	permission := "sys:" + urlItem[1] + ":" + urlItem[2]
	// 查询节点信息
	info, err := dao.Menu.Where("permission=?", permission).FindOne()
	if err != nil || info == nil {
		return
	}

	// 创建日志对象
	var entity model.OperLog
	entity.Model = info.Title
	if urlItem[2] == "add" || urlItem[2] == "addz" {
		// 新增
		entity.OperType = 1
	} else if urlItem[2] == "update" {
		// 修改
		entity.OperType = 2
	} else if urlItem[2] == "delete" || urlItem[2] == "dall" {
		// 删除
		entity.OperType = 3
	} else if urlItem[2] == "list" {
		// 查询
		entity.OperType = 4
	} else if urlItem[2] == "status" {
		// 设置状态
		entity.OperType = 5
	} else if urlItem[2] == "import" {
		// 导入
		entity.OperType = 6
	} else if urlItem[2] == "export" {
		// 导出
		entity.OperType = 7
	} else if urlItem[2] == "permission" {
		// 设置权限
		entity.OperType = 8
	} else if urlItem[2] == "resetPwd" {
		// 设置密码
		entity.OperType = 9
	} else {
		// 其他
		entity.OperType = 0
	}
	entity.OperMethod = r.Method
	entity.OperName = utils.UInfo(r).Realname
	entity.Username = utils.UInfo(r).Username
	entity.OperUrl = r.URL.String()
	entity.OperIp = r.GetClientIp()
	entity.OperLocation = utils.GetIpCity(entity.OperIp)
	entity.RequestParam = string(r.GetBody())
	entity.Status = 0
	entity.UserAgent = r.UserAgent()
	entity.CreateUser = utils.Uid(r)
	entity.CreateTime = gtime.Now()
	entity.Mark = 1
	dao.OperLog.Insert(entity)
}
