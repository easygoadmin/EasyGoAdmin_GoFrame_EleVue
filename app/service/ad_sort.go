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
 * 广告位管理-服务类
 * @author 半城风雨
 * @since 2021/7/24
 * @File : ad_sort
 */
package service

import (
	"easygoadmin/app/dao"
	"easygoadmin/app/model"
	"easygoadmin/app/utils"
	"easygoadmin/app/utils/common"
	"easygoadmin/app/utils/convert"
	"github.com/gogf/gf/errors/gerror"
	"github.com/gogf/gf/os/gtime"
	"github.com/gogf/gf/util/gconv"
)

// 中间件管理服务
var AdSort = new(adSortService)

type adSortService struct{}

func (s *adSortService) GetList(req *model.AdSortPageReq) ([]model.AdSortInfoVo, int, error) {
	// 创建查询实例
	query := dao.AdSort.Where("mark=1")
	// 查询条件
	if req != nil {
		// 广告位名称
		if req.Description != "" {
			query = query.Where("description like ?", "%"+req.Description+"%")
		}
	}
	// 查询记录总数
	count, err := query.Count()
	if err != nil {
		return nil, 0, err
	}
	// 排序
	query = query.Order("sort asc")
	// 分页
	query = query.Page(req.Page, req.Limit)
	// 对象转换
	var list []model.AdSort
	query.Structs(&list)

	// 数据处理
	var result = make([]model.AdSortInfoVo, 0)
	for _, v := range list {
		platformName, ok := common.ADSORT_PLATFORM_LIST[v.Platform]
		item := model.AdSortInfoVo{}
		item.AdSort = v
		if ok {
			item.PlatformName = platformName
		}
		// 站点名称
		if v.ItemId > 0 {
			info, err := dao.Item.FindOne("id=?", v.ItemId)
			if err == nil && info != nil {
				item.ItemName = info.Name
			}
		}

		// 栏目名称
		if v.CateId > 0 {
			cateName := ItemCate.GetCateName(v.CateId, ">>")
			item.CateName = cateName
		}

		// 加入数组
		result = append(result, item)
	}

	return result, count, nil
}

func (s *adSortService) Add(req *model.AdSortAddReq, userId int) (int64, error) {
	if utils.AppDebug() {
		return 0, gerror.New("演示环境，暂无权限操作")
	}
	// 实例化对象
	var entity model.AdSort
	entity.Description = req.Description
	entity.ItemId = req.ItemId
	entity.CateId = req.CateId
	entity.LocId = req.LocId
	entity.Platform = req.Platform
	entity.Sort = req.Sort
	entity.CreateUser = userId
	entity.CreateTime = gtime.Now()
	entity.Mark = 1

	// 插入数据
	result, err := dao.AdSort.Insert(entity)
	if err != nil {
		return 0, err
	}

	// 获取插入ID
	id, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}
	return id, nil
}

func (s *adSortService) Update(req *model.AdSortUpdateReq, userId int) (int64, error) {
	if utils.AppDebug() {
		return 0, gerror.New("演示环境，暂无权限操作")
	}
	// 查询记录
	info, err := dao.AdSort.FindOne("id=?", req.Id)
	if err != nil {
		return 0, err
	}
	if info == nil {
		return 0, gerror.New("记录不存在")
	}

	// 设置对象
	info.Description = req.Description
	info.ItemId = req.ItemId
	info.CateId = req.CateId
	info.LocId = req.LocId
	info.Platform = req.Platform
	info.Sort = req.Sort
	info.UpdateUser = userId
	info.UpdateTime = gtime.Now()

	// 更新记录
	result, err := dao.AdSort.Save(info)
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

func (s *adSortService) Delete(ids string) (int64, error) {
	if utils.AppDebug() {
		return 0, gerror.New("演示环境，暂无权限操作")
	}
	// 记录ID
	idsArr := convert.ToInt64Array(ids, ",")
	// 删除记录
	result, err := dao.AdSort.Delete("id in (?)", idsArr)
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

func (s *adSortService) GetAdSortList() []model.AdSortInfoVo {
	// 查询广告位列表
	var list []model.AdSort
	dao.AdSort.Where("mark=1").Order("sort asc").Structs(&list)
	// 数据处理
	result := make([]model.AdSortInfoVo, 0)
	for _, v := range list {
		item := model.AdSortInfoVo{}
		item.AdSort = v
		item.Description = v.Description + " >> " + gconv.String(v.LocId)
		result = append(result, item)
	}
	return result
}
