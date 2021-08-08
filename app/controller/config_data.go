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
 * 配置项管理-控制器
 * @author 半城风雨
 * @since 2021/7/21
 * @File : config_data
 */
package controller

import (
	"easygoadmin/app/model"
	"easygoadmin/app/service"
	"easygoadmin/app/utils"
	"easygoadmin/app/utils/common"
	"github.com/gogf/gf/net/ghttp"
)

// 控制器管理对象
var ConfigData = new(configDataCtl)

type configDataCtl struct{}

func (c *configDataCtl) List(r *ghttp.Request) {
	// 参数验证
	var req *model.ConfigDataPageReq
	if err := r.Parse(&req); err != nil {
		r.Response.WriteJsonExit(common.JsonResult{
			Code: -1,
			Msg:  err.Error(),
		})
	}

	// 调用查询方法
	list, count, err := service.ConfigData.GetList(req)
	if err != nil {
		r.Response.WriteJsonExit(common.JsonResult{
			Code: -1,
			Msg:  err.Error(),
		})
	}

	// 返回结果
	r.Response.WriteJsonExit(common.JsonResult{
		Code:  0,
		Data:  list,
		Msg:   "操作成功",
		Count: count,
	})
}

func (c *configDataCtl) Add(r *ghttp.Request) {
	// 参数验证
	var req *model.ConfigDataAddReq
	if err := r.Parse(&req); err != nil {
		r.Response.WriteJsonExit(common.JsonResult{
			Code: -1,
			Msg:  err.Error(),
		})
	}

	// 调用添加方法
	id, err := service.ConfigData.Add(req, utils.Uid(r.Session))
	if err != nil || id == 0 {
		r.Response.WriteJsonExit(common.JsonResult{
			Code: -1,
			Msg:  err.Error(),
		})
	}

	// 返回结果
	r.Response.WriteJsonExit(common.JsonResult{
		Code: 0,
		Msg:  "添加成功",
	})
}

func (c *configDataCtl) Update(r *ghttp.Request) {
	// 参数验证
	var req *model.ConfigDataUpdateReq
	if err := r.Parse(&req); err != nil {
		r.Response.WriteJsonExit(common.JsonResult{
			Code: -1,
			Msg:  err.Error(),
		})
	}

	// 调用更新方法
	rows, err := service.ConfigData.Update(req, utils.Uid(r.Session))
	if err != nil || rows == 0 {
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

func (c *configDataCtl) Delete(r *ghttp.Request) {
	// 参数验证
	var req *model.ConfigDataDeleteReq
	if err := r.Parse(&req); err != nil {
		r.Response.WriteJsonExit(common.JsonResult{
			Code: -1,
			Msg:  err.Error(),
		})
	}

	// 调用删除方法
	rows, err := service.ConfigData.Delete(req.Ids)
	if err != nil || rows == 0 {
		r.Response.WriteJsonExit(common.JsonResult{
			Code: -1,
			Msg:  err.Error(),
		})
	}

	// 返回结果
	r.Response.WriteJsonExit(common.JsonResult{
		Code: 0,
		Msg:  "删除成功",
	})
}

func (c *configDataCtl) Status(r *ghttp.Request) {
	// 参数验证
	var req *model.ConfigDataStatusReq
	if err := r.Parse(&req); err != nil {
		r.Response.WriteJsonExit(common.JsonResult{
			Code: -1,
			Msg:  err.Error(),
		})
	}
	result, err := service.ConfigData.Status(req, utils.Uid(r.Session))
	if err != nil || result == 0 {
		r.Response.WriteJsonExit(common.JsonResult{
			Code: -1,
			Msg:  err.Error(),
		})
	}
	// 保存成功
	r.Response.WriteJsonExit(common.JsonResult{
		Code: 0,
		Msg:  "设置成功",
	})
}
