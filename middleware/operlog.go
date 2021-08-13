/**
 *
 * @author 摆渡人
 * @since 2021/8/13
 * @File : operlog
 */
package middleware

import (
	"fmt"
	"github.com/gogf/gf/net/ghttp"
)

func OperLog(r *ghttp.Request) {
	// 后置中间件
	r.Middleware.Next()
	// 中间件处理逻辑
	fmt.Println("日志处理")
}
