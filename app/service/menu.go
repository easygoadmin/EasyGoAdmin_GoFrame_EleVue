// +----------------------------------------------------------------------
// | EasyGoAdmin敏捷开发框架 [ EasyGoAdmin ]
// +----------------------------------------------------------------------
// | 版权所有 2021 EasyGoAdmin深圳研发中心
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
 * 菜单管理-服务类
 * @author 半城风雨
 * @since 2021/5/19
 * @File : menu
 */
package service

import (
	"easygoadmin/app/dao"
	"easygoadmin/app/model"
	"easygoadmin/app/utils"
	"easygoadmin/app/utils/convert"
	"errors"
	"github.com/gogf/gf/errors/gerror"
	"github.com/gogf/gf/os/gtime"
	"github.com/gogf/gf/text/gstr"
	"github.com/gogf/gf/util/gconv"
	"reflect"
	"strings"
)

// 中间件管理服务
var Menu = new(menuService)

type menuService struct{}

// 获取菜单权限列表
func (s *menuService) GetPermissionMenuList(userId int) interface{} {
	if userId == 1 {
		// 管理员(拥有全部权限)
		menuList, _ := Menu.GetTreeList()
		return menuList
	} else {
		// 非管理员
		// 创建查询实例
		query := dao.Menu.As("m").Clone()
		// 内联查询
		query = query.InnerJoin("sys_role_menu as r", "m.id = r.menu_id")
		query = query.InnerJoin("sys_user_role ur", "ur.role_id=r.role_id")
		query = query.Where("ur.user_id=? AND m.type=0 AND m.`status`=1 AND m.mark=1", userId)
		// 获取字段
		query.Fields("m.*")
		// 排序
		query = query.Order("m.id asc")
		// 数据转换
		var list []*model.Menu
		query.Structs(&list)
		// 数据处理
		var menuNode model.TreeNode
		makeTree(list, &menuNode)
		return menuNode.Children
	}
}

// 获取子级菜单
func (s *menuService) GetTreeList() ([]*model.TreeNode, error) {
	var menuNode model.TreeNode
	data, err := dao.Menu.Where("type=0 and mark=1").Order("sort asc").FindAll()
	if err != nil {
		return nil, errors.New("系统错误")
	}
	makeTree(data, &menuNode)
	return menuNode.Children, nil
}

//递归生成分类列表
func makeTree(menu []*model.Menu, tn *model.TreeNode) {
	for _, c := range menu {
		if c.ParentId == tn.Id {
			child := &model.TreeNode{}
			child.Menu = *c
			tn.Children = append(tn.Children, child)
			makeTree(menu, child)
		}
	}
}

func (s *menuService) GetList(req *model.MenuQueryReq) []model.Menu {
	// 创建查询条件
	query := dao.Menu.Where("mark=1")
	// 查询条件
	if req != nil {
		// 菜单标题
		if req.Title != "" {
			query = query.Where("title like ?", "%"+req.Title+"%")
		}
	}
	// 排序
	query = query.Order("sort asc")
	// 对象转换
	var list []model.Menu
	query.Structs(&list)
	return list
}

func (s *menuService) Add(req *model.MenuAddReq, userId int) (int64, error) {
	if utils.AppDebug() {
		return 0, gerror.New("演示环境，暂无权限操作")
	}
	// 实例化对象
	var entity model.Menu
	entity.ParentId = req.ParentId
	entity.Title = req.Title
	entity.Icon = req.Icon
	entity.Path = req.Path
	entity.Component = req.Component
	entity.Target = req.Target
	entity.Permission = req.Permission
	entity.Type = req.Type
	entity.Status = req.Status
	entity.Hide = req.Hide
	entity.Note = req.Note
	entity.Sort = req.Sort
	entity.CreateUser = userId
	entity.CreateTime = gtime.Now()
	entity.Mark = 1

	// 插入记录
	result, err := dao.Menu.Insert(entity)
	if err != nil {
		return 0, err
	}

	// 获取插入ID
	id, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}

	// 添加节点
	setPermission(req.Type, req.CheckedList, req.Title, req.Path, gconv.Int(id), userId)

	return id, nil
}

func (s *menuService) Update(req *model.MenuUpdateReq, userId int) (int64, error) {
	if utils.AppDebug() {
		return 0, gerror.New("演示环境，暂无权限操作")
	}
	// 查询记录
	info, err := dao.Menu.FindOne("id=?", req.Id)
	if err != nil {
		return 0, err
	}
	if info == nil {
		return 0, gerror.New("记录不存在")
	}
	// 设置参数值
	info.ParentId = req.ParentId
	info.Title = req.Title
	info.Icon = req.Icon
	info.Path = req.Path
	info.Component = req.Component
	info.Target = req.Target
	info.Permission = req.Permission
	info.Type = req.Type
	info.Status = req.Status
	info.Hide = req.Hide
	info.Note = req.Note
	info.Sort = req.Sort
	info.UpdateUser = userId
	info.UpdateTime = gtime.Now()

	// 更新数据
	result, err := dao.Menu.Save(info)
	if err != nil {
		return 0, err
	}

	// 添加节点
	setPermission(req.Type, req.CheckedList, req.Title, req.Path, req.Id, userId)

	// 获取数影响的行数
	rows, err := result.RowsAffected()
	if err != nil {
		return 0, err
	}
	return rows, nil
}

func (s *menuService) Delete(ids string) (int64, error) {
	if utils.AppDebug() {
		return 0, gerror.New("演示环境，暂无权限操作")
	}
	// 记录ID
	idsArr := convert.ToInt64Array(ids, ",")

	// 判断是否有子级
	child, err := dao.Menu.Where("parent_id in (?)", idsArr).Count()
	if err != nil {
		return 0, err
	}
	if child > 0 {
		return 0, gerror.New("有子级无法删除")
	}

	// 删除记录
	result, err := dao.Menu.Delete("id in (?)", idsArr)
	if err != nil {
		return 0, err
	}

	// 获取受影响行数
	rows, err := result.RowsAffected()
	if err != nil {
		return 0, err
	}
	return rows, nil
}

// 添加节点
func setPermission(menuType int, checkedList []int, name string, url string, parentId int, userId int) {
	if menuType != 0 || len(checkedList) == 0 || url == "" {
		return
	}
	// 删除现有节点
	dao.Menu.Delete("parent_id=?", parentId)
	// 模块名称
	moduleTitle := gstr.Replace(name, "管理", "")
	// 创建权限节点
	urlArr := strings.Split(url, "/")

	if len(urlArr) >= 3 {
		// 模块名
		moduleName := urlArr[len(urlArr)-1]
		// 节点处理
		for _, v := range checkedList {
			// 实例化对象
			var entity model.Menu
			// 节点索引
			value := gconv.Int(v)
			if value == 1 {
				entity.Title = "查询" + moduleTitle
				entity.Path = "/" + moduleName + "/list"
				entity.Permission = "sys:" + moduleName + ":list"
				entity.Method = "GET"
			} else if value == 5 {
				entity.Title = "添加" + moduleTitle
				entity.Path = "/" + moduleName + "/add"
				entity.Permission = "sys:" + moduleName + ":add"
				entity.Method = "POST"
			} else if value == 10 {
				entity.Title = "修改" + moduleTitle
				entity.Path = "/" + moduleName + "/update"
				entity.Permission = "sys:" + moduleName + ":update"
				entity.Method = "PUT"
			} else if value == 15 {
				entity.Title = "删除" + moduleTitle
				entity.Path = "/" + moduleName + "/delete"
				entity.Permission = "sys:" + moduleName + ":delete"
				entity.Method = "DELETE"
			} else if value == 20 {
				entity.Title = moduleTitle + "详情"
				entity.Path = "/" + moduleName + "/detail"
				entity.Permission = "sys:" + moduleName + ":detail"
				entity.Method = "GET"
			} else if value == 25 {
				entity.Title = "设置状态"
				entity.Path = "/" + moduleName + "/status"
				entity.Permission = "sys:" + moduleName + ":status"
				entity.Method = "PUT"
			} else if value == 30 {
				entity.Title = "批量删除"
				entity.Path = "/" + moduleName + "/dall"
				entity.Permission = "sys:" + moduleName + ":dall"
				entity.Method = "DELETE"
			} else if value == 35 {
				entity.Title = "添加子级"
				entity.Path = "/" + moduleName + "/addz"
				entity.Permission = "sys:" + moduleName + ":addz"
				entity.Method = "POST"
			} else if value == 40 {
				entity.Title = "全部展开"
				entity.Path = "/" + moduleName + "/expand"
				entity.Permission = "sys:" + moduleName + ":expand"
				entity.Method = "GET"
			} else if value == 45 {
				entity.Title = "全部折叠"
				entity.Path = "/" + moduleName + "/collapse"
				entity.Permission = "sys:" + moduleName + ":collapse"
				entity.Method = "GET"
			} else if value == 50 {
				entity.Title = "导出" + moduleTitle
				entity.Path = "/" + moduleName + "/export"
				entity.Permission = "sys:" + moduleName + ":export"
				entity.Method = "GET"
			} else if value == 55 {
				entity.Title = "导入" + moduleTitle
				entity.Path = "/" + moduleName + "/import"
				entity.Permission = "sys:" + moduleName + ":import"
				entity.Method = "GET"
			} else if value == 60 {
				entity.Title = "分配权限"
				entity.Path = "/" + moduleName + "/permission"
				entity.Permission = "sys:" + moduleName + ":permission"
				entity.Method = "POST"
			} else if value == 65 {
				entity.Title = "重置密码"
				entity.Path = "/" + moduleName + "/resetPwd"
				entity.Permission = "sys:" + moduleName + ":resetPwd"
				entity.Method = "PUT"
			}
			entity.ParentId = parentId
			entity.Type = 1
			entity.Status = 1
			entity.Target = "_self"
			entity.Sort = value
			entity.CreateUser = userId
			entity.CreateTime = gtime.Now()
			entity.UpdateUser = userId
			entity.UpdateTime = gtime.Now()
			entity.Mark = 1

			// 插入节点
			dao.Menu.Insert(entity)
		}
	}
}

// 数据源转换
func (s *menuService) MakeList(data []*model.TreeNode) map[int]string {
	menuList := make(map[int]string, 0)
	if reflect.ValueOf(data).Kind() == reflect.Slice {
		// 一级栏目
		for _, val := range data {
			menuList[val.Id] = val.Title

			// 二级栏目
			for _, v := range val.Children {
				menuList[v.Id] = "|--" + v.Title

				// 三级栏目
				for _, vt := range v.Children {
					menuList[vt.Id] = "|--|--" + vt.Title
				}
			}
		}
	}
	return menuList
}

// 获取权限节点列表
func (s *menuService) GetPermissionsList(userId int) []string {
	if userId == 1 {
		// 管理员,管理员拥有全部权限
		list, _ := dao.Menu.Fields("permission").Where("type=1").Where("mark=1").Array()
		permissionList := gconv.Strings(list)
		return permissionList
	} else {
		// 非管理员
		// 创建查询实例
		query := dao.Menu.As("m").Clone()
		// 内联查询
		query = query.InnerJoin("sys_role_menu as r", "m.id = r.menu_id")
		query = query.InnerJoin("sys_user_role ur", "ur.role_id=r.role_id")
		query = query.Where("ur.user_id=? AND m.type=1 AND m.`status`=1 AND m.mark=1", userId)
		// 获取字段
		query = query.Fields("m.permission")
		list, _ := query.Array()
		permissionList := gconv.Strings(list)
		return permissionList
	}
}
