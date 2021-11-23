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
 * 登录日志-服务类
 * @author 半城风雨
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
