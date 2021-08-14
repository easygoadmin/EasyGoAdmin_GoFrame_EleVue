/**
 *
 * @author 摆渡人
 * @since 2021/8/14
 * @File : oper_log
 */
package service

import (
	"easygoadmin/app/dao"
	"easygoadmin/app/model"
)

// 中间件管理服务
var OperLog = new(operLogService)

type operLogService struct{}

func (s *operLogService) GetList(req *model.OperLogPageReq) ([]model.OperLog, int, error) {
	// 实例化查询条件
	query := dao.OperLog.Clone()
	query = query.Where("mark=1")
	// 查询条件
	if req != nil {
		// 操作用户
		if req.Username != "" {
			query = query.Where("username=?", req.Username)
		}
		// 操作模块
		if req.Model != "" {
			query = query.Where("model like ?", "%"+req.Model+"%")
		}
	}
	// 查询记录总数
	count, err := query.Count()
	if err != nil {
		return nil, 0, err
	}
	// 排除
	query = query.Order("id desc")
	// 分页
	query = query.Page(req.Page, req.Limit)
	// 对象转换
	var list []model.OperLog
	query.Structs(&list)
	return list, count, nil
}
