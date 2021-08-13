package boot

import (
	_ "easygoadmin/packed"
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/os/gview"
	"github.com/gogf/swagger"
)

// 用于应用初始化。
func init() {
	s := g.Server()
	s.Plugin(&swagger.Swagger{})

	gview.Instance().BindFuncMap(gview.FuncMap{

	})
}
