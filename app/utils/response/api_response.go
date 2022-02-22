// +----------------------------------------------------------------------
// | EasyGoAdmin敏捷开发框架 [ EasyGoAdmin ]
// +----------------------------------------------------------------------
// | 版权所有 2019~2022 EasyGoAdmin深圳研发中心
// +----------------------------------------------------------------------
// | 官方网站: http://www.easygoadmin.vip
// +----------------------------------------------------------------------
// | Author: 半城风雨 <easygoadmin@163.com>
// +----------------------------------------------------------------------
// | 免责声明:
// | 本软件框架禁止任何单位和个人用于任何违法、侵害他人合法利益等恶意的行为，禁止用于任何违
// | 反我国法律法规的一切平台研发，任何单位和个人使用本软件框架用于产品研发而产生的任何意外
// | 、疏忽、合约毁坏、诽谤、版权或知识产权侵犯及其造成的损失 (包括但不限于直接、间接、附带
// | 或衍生的损失等)，本团队不承担任何法律责任。本软件框架只能用于公司和个人内部的法律所允
// | 许的合法合规的软件产品研发，详细声明内容请阅读《框架免责声明》附件；
// +----------------------------------------------------------------------

/**
 *
 * @author 半城风雨
 * @since 2021/5/20
 * @File : api_response
 */
package response

import (
	"easygoadmin/app/utils/common"
	"github.com/gogf/gf/net/ghttp"
)

// 通用API响应
type ApiResp struct {
	c *common.JsonResult
	r *ghttp.Request
}

// 返回一个成功的消息体
func SucessResp(r *ghttp.Request) *ApiResp {
	msg := common.JsonResult{
		Code:  0,
		Btype: common.BOther,
		Msg:   "操作成功",
	}
	var a = ApiResp{
		c: &msg,
		r: r,
	}
	return &a
}

// 返回一个错误的消息体
func ErrorResp(r *ghttp.Request) *ApiResp {
	msg := common.JsonResult{
		Code:  500,
		Btype: common.BOther,
		Msg:   "操作失败",
	}
	var a = ApiResp{
		c: &msg,
		r: r,
	}
	return &a
}

// 返回一个拒绝访问的消息体
func ForbiddenResp(r *ghttp.Request) *ApiResp {
	msg := common.JsonResult{
		Code:  403,
		Btype: common.BOther,
		Msg:   "无操作权限",
	}
	var a = ApiResp{
		c: &msg,
		r: r,
	}
	return &a
}

// 设置消息体的内容
func (resp *ApiResp) SetMsg(msg string) *ApiResp {
	resp.c.Msg = msg
	return resp
}

// 设置消息体的编码
func (resp *ApiResp) SetCode(code int) *ApiResp {
	resp.c.Code = code
	return resp
}

// 设置消息体的数据
func (resp *ApiResp) SetData(data interface{}) *ApiResp {
	resp.c.Data = data
	return resp
}

// 设置消息体的业务类型
func (resp *ApiResp) SetBtype(btype common.BunissType) *ApiResp {
	resp.c.Btype = btype
	return resp
}

// 输出json到客户端
func (resp *ApiResp) WriteJsonExit() {
	resp.r.Response.WriteJsonExit(resp.c)
}
