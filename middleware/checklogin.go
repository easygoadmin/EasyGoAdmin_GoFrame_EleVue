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
 * 登录验证中间件
 * @author 半城风雨
 * @since 2021/8/13
 * @File : checklogin
 */
package middleware

import (
	"easygoadmin/app/utils"
	"easygoadmin/app/utils/common"
	"fmt"
	"github.com/gogf/gf/net/ghttp"
	"github.com/gogf/gf/text/gstr"
	"time"
)

func CheckLogin(r *ghttp.Request) {
	fmt.Println("登录验证中间件")
	// 放行设置
	urlItem := []string{"/captcha", "/login"}
	if !utils.InStringArray(r.RequestURI, urlItem) {
		// 从请求头中获取Token
		token := r.GetHeader("Authorization")
		// 字符串替换
		token = gstr.Replace(token, "Bearer ", "")
		claim, err := utils.ParseToken(token)
		if err != nil {
			fmt.Println("解析token出现错误：", err)
			r.Response.WriteJsonExit(common.JsonResult{
				Code: 401,
				Msg:  "Token已过期",
			})
		} else if time.Now().Unix() > claim.ExpiresAt {
			fmt.Println("时间超时")
			r.Response.WriteJsonExit(common.JsonResult{
				Code: 401,
				Msg:  "时间超时",
			})
		}
	}
	// 前置中间件
	r.Middleware.Next()
}
