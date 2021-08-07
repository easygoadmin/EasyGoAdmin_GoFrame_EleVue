package boot

import (
	"easygoadmin/app/utils/function"
	"easygoadmin/app/widget"
	_ "easygoadmin/packed"
	"github.com/gogf/gf/os/gview"

	"github.com/gogf/gf/frame/g"
	"github.com/gogf/swagger"
)

// 用于应用初始化。
func init() {
	s := g.Server()
	s.Plugin(&swagger.Swagger{})

	gview.Instance().BindFuncMap(gview.FuncMap{
		"ossUrl":       function.GetOssUrl,
		"widget":       widget.Widget,
		"query":        widget.Query,
		"add":          widget.Add,
		"edit":         widget.Edit,
		"delete":       widget.Delete,
		"dall":         widget.Dall,
		"expand":       widget.Expand,
		"collapse":     widget.Collapse,
		"addz":         widget.Addz,
		"switch":       widget.Switch,
		"select":       widget.Select,
		"submit":       widget.Submit,
		"icon":         widget.Icon,
		"transfer":     widget.Transfer,
		"upload_image": widget.UploadImage,
		"album":        widget.Album,
		"item":         widget.Item,
		"kindeditor":   widget.Kindeditor,
		"date":         widget.Date,
		"checkbox":     widget.Checkbox,
		"radio":        widget.Radio,
		"city":         widget.City,
	})
}
