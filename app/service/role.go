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
 * 角色管理-服务类
 * @author 半城风雨
 * @since 2021/7/15
 * @File : role
 */
package service

import (
	"easygoadmin/app/dao"
	"easygoadmin/app/model"
	"easygoadmin/app/utils"
	"easygoadmin/app/utils/convert"
	"github.com/gogf/gf/errors/gerror"
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/os/gtime"
)

// 中间件管理服务
var Role = new(roleService)

type roleService struct{}

func (s *roleService) GetList(req *model.RolePageReq) ([]model.Role, int, error) {
	// 创建查询对象
	query := dao.Role.Where("mark=1")
	// 查询条件
	if req != nil {
		// 角色名称
		if req.Name != "" {
			query = query.Where("name like ?", "%"+req.Name+"%")
		}
	}
	// 获取记录总数
	count, err := query.Count()
	if err != nil {
		return nil, 0, err
	}
	// 排序
	query = query.Order("sort asc")
	// 分页
	query = query.Page(req.Page, req.Limit)
	// 对象转换
	var list []model.Role
	query.Structs(&list)
	return list, count, nil
}

func (s *roleService) Add(req *model.RoleAddReq, userId int) (int64, error) {
	if utils.AppDebug() {
		return 0, gerror.New("演示环境，暂无权限操作")
	}
	// 实例化模型
	var entity model.Role
	entity.Name = req.Name
	entity.Code = req.Code
	entity.Status = req.Status
	entity.Sort = req.Sort
	entity.Note = req.Note
	entity.CreateUser = userId
	entity.CreateTime = gtime.Now()
	entity.Mark = 1

	// 插入数据
	result, err := dao.Role.Insert(entity)
	if err != nil {
		return 0, err
	}

	// 获取插入ID
	id, err := result.RowsAffected()
	if err != nil || id <= 0 {
		return 0, err
	}
	return id, nil
}

func (s *roleService) Update(req *model.RoleUpdateReq, userId int) (int64, error) {
	if utils.AppDebug() {
		return 0, gerror.New("演示环境，暂无权限操作")
	}
	// 获取记录信息
	info, err := dao.Role.FindOne("id=?", req.Id)
	if err != nil {
		return 0, err
	}
	if info == nil {
		return 0, gerror.New("记录不存在")
	}

	// 对象赋值
	info.Name = req.Name
	info.Code = req.Code
	info.Status = req.Status
	info.Sort = req.Sort
	info.Note = req.Note
	info.UpdateUser = userId
	info.UpdateTime = gtime.Now()
	result, err := dao.Role.Save(info)
	if err != nil {
		return 0, err
	}

	// 获取受影响的行数
	rows, err := result.RowsAffected()
	if err != nil {
		return 0, err
	}
	return rows, nil
}

func (s *roleService) Delete(ids string) (int64, error) {
	if utils.AppDebug() {
		return 0, gerror.New("演示环境，暂无权限操作")
	}
	// 记录ID
	idsArr := convert.ToInt64Array(ids, ",")
	result, err := dao.Role.Delete("id in (?)", idsArr)
	if err != nil {
		return 0, err
	}

	// 获取受影响记录数
	count, err := result.RowsAffected()
	if err != nil {
		return 0, err
	}
	return count, nil
}

func (s *roleService) Status(req *model.RoleStatusReq, userId int) (int64, error) {
	if utils.AppDebug() {
		return 0, gerror.New("演示环境，暂无权限操作")
	}
	// 查询记录
	info, err := dao.Role.FindOne("id=?", req.Id)
	if err != nil {
		return 0, err
	}
	if info == nil {
		return 0, gerror.New("记录不存在")
	}

	// 设置状态
	result, err := dao.Role.Data(g.Map{
		"status":      req.Status,
		"update_user": userId,
		"update_time": gtime.Now(),
	}).Where(dao.Role.Columns.Id, info.Id).Update()
	if err != nil {
		return 0, err
	}

	// 获取受影响行数
	count, err := result.RowsAffected()
	if err != nil {
		return 0, err
	}
	return count, nil
}
