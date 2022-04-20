// +----------------------------------------------------------------------
// | 版权和免责声明:
// | 本团队对该软件框架产品拥有知识产权（包括但不限于商标权、专利权、著作权、商业秘密等）
// | 均受到相关法律法规的保护，任何个人、组织和单位不得在未经本团队书面授权的情况下对所授权
// | 软件框架产品本身申请相关的知识产权，禁止用于任何违法、侵害他人合法权益等恶意的行为，禁
// | 止用于任何违反我国法律法规的一切项目研发，任何个人、组织和单位用于项目研发而产生的任何
// | 意外、疏忽、合约毁坏、诽谤、版权或知识产权侵犯及其造成的损失 (包括但不限于直接、间接、
// | 附带或衍生的损失等)，本团队不承担任何法律责任，本软件框架禁止任何单位和个人、组织用于
// | 任何违法、侵害他人合法利益等恶意的行为，如有发现违规、违法的犯罪行为，本团队将无条件配
// | 合公安机关调查取证同时保留一切以法律手段起诉的权利，本软件框架只能用于公司和个人内部的
// | 法律所允许的合法合规的软件产品研发，详细声明内容请阅读《框架免责声明》附件；
// +----------------------------------------------------------------------

/**
 * 角色菜单-服务类
 * @author 半城风雨
 * @since 2021/7/15
 * @File : rolemenu
 */
package service

import (
	"easygoadmin/app/dao"
	"easygoadmin/app/model"
	"easygoadmin/app/utils"
	"easygoadmin/app/utils/function"
	"github.com/gogf/gf/errors/gerror"
	"github.com/gogf/gf/util/gconv"
	"github.com/gogf/gf/util/gutil"
)

var RoleMenu = new(roleMenuService)

type roleMenuService struct{}

func (s *roleMenuService) GetRoleMenuList(roleId int64) ([]model.RoleMenuInfo, error) {
	// 获取全部菜单列表
	var menuList []model.Menu
	dao.Menu.Where("status=1").Where("mark=1").Order("sort asc").Structs(&menuList)
	if len(menuList) == 0 {
		return nil, gerror.New("菜单列表不存在")
	}
	// 获取角色菜单权限列表
	var roleMenuList []model.RoleMenu
	dao.RoleMenu.Where("role_id=?", roleId).Structs(&roleMenuList)
	idList := gutil.ListItemValuesUnique(&roleMenuList, "MenuId")

	// 对象处理
	var list []model.RoleMenuInfo
	if len(menuList) > 0 {
		for _, m := range menuList {
			var info model.RoleMenuInfo
			info.Id = m.Id
			info.Title = m.Title
			info.Open = true
			info.ParentId = m.ParentId
			// 节点选中值
			if function.InArray(gconv.String(m.Id), idList) {
				info.Checked = true
			}
			list = append(list, info)
		}
	}
	return list, nil
}

func (s *roleMenuService) Save(req *model.RoleMenuSaveReq) error {
	if utils.AppDebug() {
		return gerror.New("演示环境，暂无权限操作")
	}
	itemArr := req.MenuIds
	if len(itemArr) == 0 {
		return gerror.New("请选择权限节点")
	}
	// 删除现有的角色权限数据
	dao.RoleMenu.Delete("role_id=?", req.RoleId)
	// 遍历创建新角色权限数据
	for _, v := range itemArr {
		var entity model.RoleMenu
		entity.RoleId = req.RoleId
		entity.MenuId = v
		dao.RoleMenu.Insert(entity)
	}
	// 批量插入
	//dao.RoleMenu.Data(list).Batch(2).Insert()
	return nil
}
