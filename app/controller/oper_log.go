/**
 *
 * @author 摆渡人
 * @since 2021/8/14
 * @File : oper_log
 */
package controller

import (
	"easygoadmin/app/model"
	"easygoadmin/app/service"
	"easygoadmin/app/utils/common"
	"github.com/gogf/gf/net/ghttp"
)

// 控制器管理对象
var OperLog = new(operLogCtl)

type operLogCtl struct{}

func (c *operLogCtl) List(r *ghttp.Request) {
	// 参数验证
	var req *model.OperLogPageReq
	if err := r.Parse(&req); err != nil {
		r.Response.WriteJsonExit(common.JsonResult{
			Code: -1,
			Msg:  err.Error(),
		})
	}

	// 调用获取列表方法
	list, count, err := service.OperLog.GetList(req)
	if err != nil {
		r.Response.WriteJsonExit(common.JsonResult{
			Code: 0 - 1,
			Msg:  err.Error(),
		})
	}

	// 返回结果
	r.Response.WriteJsonExit(common.JsonResult{
		Code:  0,
		Msg:   "查询成功",
		Data:  list,
		Count: count,
	})
}
