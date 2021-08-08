package middleware

import "github.com/gogf/gf/net/ghttp"

//跨域处理中间件
func CORS(r *ghttp.Request) {
	r.Response.CORSDefault()
	r.Middleware.Next()
}
