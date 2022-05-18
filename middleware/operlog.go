// +----------------------------------------------------------------------
// | EasyGoAdmin敏捷开发框架 [ 赋能开发者，助力企业发展 ]
// +----------------------------------------------------------------------
// | 版权所有 2019~2022 深圳EasyGoAdmin研发中心
// +----------------------------------------------------------------------
// | Licensed LGPL-3.0 EasyGoAdmin并不是自由软件，未经许可禁止去掉相关版权
// +----------------------------------------------------------------------
// | 官方网站: http://www.easygoadmin.vip
// +----------------------------------------------------------------------
// | Author: @半城风雨 团队荣誉出品
// +----------------------------------------------------------------------
// | 版权和免责声明:
// | 本团队对该软件框架产品拥有知识产权（包括但不限于商标权、专利权、著作权、商业秘密等）
// | 均受到相关法律法规的保护，任何个人、组织和单位不得在未经本团队书面授权的情况下对所授权
// | 软件框架产品本身申请相关的知识产权，禁止用于任何违法、侵害他人合法权益等恶意的行为，禁
// | 止用于任何违反我国法律法规的一切项目研发，任何个人、组织和单位用于项目研发而产生的任何
// | 意外、疏忽、合约毁坏、诽谤、版权或知识产权侵犯及其造成的损失 (包括但不限于直接、间接、
// | 附带或衍生的损失等)，本团队不承担任何法律责任，本软件框架禁止任何单位和个人、组织用于
// | 任何违法、侵害他人合法利益等恶意的行为，如有发现违规、违法的犯罪行为，本团队将无条件配
// | 合公安机关调查取证同时保留一切以法律手段起诉的权利，本软件框架只能用于公司和个人内部的
// | 法律所允许的合法合规的软件产品研发，详细声明内容请阅读《框架免责声明》附件；
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
