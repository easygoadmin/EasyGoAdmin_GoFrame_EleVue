// ==========================================================================
// This is auto-generated by gf cli tool. Fill this file as you wish.
// ==========================================================================

package model

import (
	"easygoadmin/app/model/internal"
)

// Config is the golang structure for table sys_config.
type Config internal.Config

// Fill with you ideas below.

// 列表查询条件
type ConfigQueryReq struct {
	Name string `p:"name"` // 配置名称
}

// 添加配置
type ConfigAddReq struct {
	Name string `p:"name"  v:"required#配置名称不能为空"` // 配置名称
	Sort int    `p:"sort"  v:"required#配置排序不能为空"` // 显示顺序
}

// 修改配置
type ConfigUpdateReq struct {
	Id   int    `p:id v:"required#主键ID不能为空"`      // 主键ID
	Name string `p:"name"  v:"required#配置名称不能为空"` // 配置名称
	Sort int    `p:"sort"  v:"required#配置排序不能为空"` // 显示顺序
}

// 删除配置
type ConfigDeleteReq struct {
	Ids string `p:ids v:"required#请选择需要删除的数据记录"`
}
