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
 * 站点管理-控制器
 * @author 半城风雨
 * @since 2021/7/24
 * @File : item
 */
package controller

import (
	"easygoadmin/app/dao"
	"easygoadmin/app/model"
	"easygoadmin/app/service"
	"easygoadmin/app/utils"
	"easygoadmin/app/utils/common"
	"github.com/gogf/gf/net/ghttp"
)

// 控制器管理对象
var Item = new(itemCtl)

type itemCtl struct{}

func (c *itemCtl) List(r *ghttp.Request) {
	// 参数验证
	var req *model.ItemPageReq
	if err := r.Parse(&req); err != nil {
		r.Response.WriteJsonExit(common.JsonResult{
			Code: -1,
			Msg:  err.Error(),
		})
	}

	// 调用分页查询方法
	list, count, err := service.Item.GetList(req)
	if err != nil {
		r.Response.WriteJsonExit(common.JsonResult{
			Code: -1,
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

func (c *itemCtl) Add(r *ghttp.Request) {
	// 参数验证
	var req *model.ItemAddReq
	if err := r.Parse(&req); err != nil {
		r.Response.WriteJsonExit(common.JsonResult{
			Code: -1,
			Msg:  err.Error(),
		})
	}

	// 调用添加方法
	id, err := service.Item.Add(req, utils.Uid(r))
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

func (c *itemCtl) Update(r *ghttp.Request) {
	// 参数验证
	var req *model.ItemUpdateReq
	if err := r.Parse(&req); err != nil {
		r.Response.WriteJsonExit(common.JsonResult{
			Code: -1,
			Msg:  err.Error(),
		})
	}

	// 调用更新方法
	rows, err := service.Item.Update(req, utils.Uid(r))
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

func (c *itemCtl) Delete(r *ghttp.Request) {
	// 参数验证
	var req *model.ItemDeleteReq
	if err := r.Parse(&req); err != nil {
		r.Response.WriteJsonExit(common.JsonResult{
			Code: -1,
			Msg:  err.Error(),
		})
	}

	// 调用删除方法
	rows, err := service.Item.Delete(req.Ids)
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

func (c *itemCtl) Status(r *ghttp.Request) {
	var req *model.ItemStatusReq
	if err := r.Parse(&req); err != nil {
		r.Response.WriteJsonExit(common.JsonResult{
			Code: -1,
			Msg:  err.Error(),
		})
	}
	result, err := service.Item.Status(req, utils.Uid(r))
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

func (c *itemCtl) GetItemList(r *ghttp.Request) {
	// 查询站点列表
	list, _ := dao.Item.Where("status=1 and mark=1").Order("sort asc").All()
	// 返回结果
	r.Response.WriteJsonExit(common.JsonResult{
		Code: 0,
		Msg:  "查询成功",
		Data: list,
	})
}
