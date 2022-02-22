// +----------------------------------------------------------------------
// | EasyGoAdmin敏捷开发框架 [ EasyGoAdmin ]
// +----------------------------------------------------------------------
// | 版权所有 2019~2022 EasyGoAdmin深圳研发中心
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
 * 操作日志-服务类
 * @author 半城风雨
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
