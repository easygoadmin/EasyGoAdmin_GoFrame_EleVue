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
 * 栏目管理-控制器
 * @author 半城风雨
 * @since 2021/7/24
 * @File : item_cate
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
var ItemCate = new(itemCateCtl)

type itemCateCtl struct{}

func (c *itemCateCtl) List(r *ghttp.Request) {
	// 参数验证
	var req *model.ItemCateQueryReq
	if err := r.Parse(&req); err != nil {
		r.Response.WriteJsonExit(common.JsonResult{
			Code: -1,
			Msg:  err.Error(),
		})
	}

	// 调用查询列表方法
	list := service.ItemCate.GetList(req)

	// 返回结果
	r.Response.WriteJsonExit(common.JsonResult{
		Code: 0,
		Msg:  "查询成功",
		Data: list,
	})
}

func (c *itemCateCtl) Add(r *ghttp.Request) {
	// 参数验证
	var req *model.ItemCateAddReq
	if err := r.Parse(&req); err != nil {
		r.Response.WriteJsonExit(common.JsonResult{
			Code: -1,
			Msg:  err.Error(),
		})
	}

	// 调用添加方法
	id, err := service.ItemCate.Add(req, utils.Uid(r))
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

func (c *itemCateCtl) Update(r *ghttp.Request) {
	// 参数验证
	var req *model.ItemCateUpdateReq
	if err := r.Parse(&req); err != nil {
		r.Response.WriteJsonExit(common.JsonResult{
			Code: -1,
			Msg:  err.Error(),
		})
	}

	// 调用更新方法
	rows, err := service.ItemCate.Update(req, utils.Uid(r))
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

func (c *itemCateCtl) Delete(r *ghttp.Request) {
	// 参数验证
	var req *model.ItemCateDeleteReq
	if err := r.Parse(&req); err != nil {
		r.Response.WriteJsonExit(common.JsonResult{
			Code: -1,
			Msg:  err.Error(),
		})
	}

	// 调用删除方法
	rows, err := service.ItemCate.Delete(req.Ids)
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

//
//func (c *itemCateCtl) GetCateTreeList(r *ghttp.Request) {
//	itemId := r.GetQueryInt("item_id")
//	list, err := service.ItemCate.GetCateTreeList(itemId, 0)
//	if err != nil {
//		r.Response.WriteJsonExit(common.JsonResult{
//			Code: -1,
//			Msg:  err.Error(),
//		})
//	}
//	// 数据源转换
//	result := service.ItemCate.MakeList(list)
//	// 返回结果
//	r.Response.WriteJsonExit(common.JsonResult{
//		Code: 0,
//		Msg:  "操作成功",
//		Data: result,
//	})
//}

func (c *itemCateCtl) GetCateList(r *ghttp.Request) {
	// 查询栏目列表
	list, _ := dao.ItemCate.Where("status=1 and mark=1").Order("sort asc").All()
	// 返回结果
	r.Response.WriteJsonExit(common.JsonResult{
		Code: 0,
		Msg:  "查询成功",
		Data: list,
	})
}
