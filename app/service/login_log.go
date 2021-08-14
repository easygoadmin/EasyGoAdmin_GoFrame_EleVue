/**
 *
 * @author 摆渡人
 * @since 2021/8/14
 * @File : login_log
 */
package service

import (
	"easygoadmin/app/dao"
	"easygoadmin/app/model"
	"easygoadmin/app/utils"
	"easygoadmin/app/utils/convert"
	"github.com/gogf/gf/errors/gerror"
)

// 中间件管理服务
var LoginLog = new(loginLogService)

type loginLogService struct{}

func (s *loginLogService) GetList(req *model.LoginLogPageReq) ([]model.LoginLog, int, error) {
	// 实例化查询条件
	query := dao.LoginLog.Clone()
	query = query.Where("mark=1")
	// 查询条件
	if req != nil {
		// 操作用户
		if req.Username != "" {
			query = query.Where("username=?", req.Username)
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
	var list []model.LoginLog
	query.Structs(&list)
	return list, count, nil
}

// 删除
func (s *loginLogService) Delete(Ids string) (int64, error) {
	if utils.AppDebug() {
		return 0, gerror.New("演示环境，暂无权限操作")
	}
	idsArr := convert.ToInt64Array(Ids, ",")
	result, err := dao.LoginLog.Delete("id in (?)", idsArr)
	if err != nil {
		return 0, err
	}
	// 获取受影响行数
	rows, err := result.RowsAffected()
	return rows, nil
}
