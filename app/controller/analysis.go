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
 * 统计分析-控制器
 * @author 半城风雨
 * @since 2021/7/24
 * @File : ad_sort
 */
package controller

import (
	"easygoadmin/app/utils/response"
	"github.com/gogf/gf/net/ghttp"
)

var Analysis = new(analysisCtl)

type analysisCtl struct{}

func (c *analysisCtl) Index(r *ghttp.Request) {
	// 渲染模板
	response.BuildTpl(r, "analysis/index.html").WriteTpl()
}
