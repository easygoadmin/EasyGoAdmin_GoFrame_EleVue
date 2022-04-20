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
 * 用户管理-服务类
 * @author 半城风雨
 * @since 2021/7/27
 * @File : user
 */
package service

import (
	"easygoadmin/app/dao"
	"easygoadmin/app/model"
	"easygoadmin/app/utils"
	"easygoadmin/app/utils/convert"
	"github.com/gogf/gf/errors/gerror"
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
	"github.com/gogf/gf/os/gtime"
	"github.com/gogf/gf/util/gconv"
)

// 中间件管理服务
var User = new(userService)

type userService struct{}

func (s *userService) GetList(req *model.UserPageReq) ([]model.UserInfoVo, int, error) {
	// 创建查询实例
	query := dao.User.Where("mark=1")
	// 查询条件
	if req != nil {
		// 用户姓名
		if req.Realname != "" {
			query = query.Where("realname like ?", "%"+req.Realname+"%")
		}
		// 性别
		if req.Gender > 0 {
			query = query.Where("gender=?", req.Gender)
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
	var list []model.User
	query.Structs(&list)

	// 获取职级列表
	levelList, _ := dao.Level.Where("mark=1").Fields("id,name").All()
	var levelMap = make(map[int]string)
	for _, v := range levelList {
		levelMap[v.Id] = v.Name
	}
	// 获取岗位列表
	positionList, _ := dao.Position.Where("mark=1").Fields("id,name").All()
	var positionMap = make(map[int]string)
	for _, v := range positionList {
		positionMap[v.Id] = v.Name
	}
	// 部门
	deptList, _ := dao.Dept.Where("mark=1").Fields("id,name").All()
	var deptMap = make(map[int]string)
	for _, v := range deptList {
		deptMap[v.Id] = v.Name
	}

	// 数据处理
	var result []model.UserInfoVo
	for _, v := range list {
		item := model.UserInfoVo{}
		item.User = v
		// 头像
		if v.Avatar != "" {
			item.Avatar = utils.GetImageUrl(v.Avatar)
		}
		// 性别
		if v.Gender > 0 {
			item.GenderName = utils.GENDER_LIST[v.Gender]
		}
		// 职级
		if v.LevelId > 0 {
			item.LevelName = levelMap[v.LevelId]
		}
		// 岗位
		if v.PositionId > 0 {
			item.PositionName = positionMap[v.PositionId]
		}
		// 部门
		if v.DeptId > 0 {
			item.DeptName = deptMap[v.DeptId]
		}
		// 角色列表
		roleList := UserRole.GetUserRoleList(v.Id)
		if len(roleList) > 0 {
			item.RoleList = roleList
		} else {
			item.RoleList = make([]model.Role, 0)
		}
		// 省市区
		cityList := make([]string, 0)
		// 省份编号
		cityList = append(cityList, item.ProvinceCode)
		// 城市编号
		cityList = append(cityList, item.CityCode)
		// 县区编号
		cityList = append(cityList, item.DistrictCode)
		item.City = cityList
		// 加入数组
		result = append(result, item)
	}
	return result, count, nil
}

func (s *userService) Add(req *model.UserAddReq, userId int) (int64, error) {
	if utils.AppDebug() {
		return 0, gerror.New("演示环境，暂无权限操作")
	}
	// 实例化对象
	var entity model.User
	entity.Realname = req.Realname
	entity.Nickname = req.Nickname
	entity.Gender = req.Gender
	entity.Avatar = req.Avatar
	entity.Mobile = req.Mobile
	entity.Email = req.Email
	entity.Birthday = req.Birthday
	entity.DeptId = req.DeptId
	entity.LevelId = req.LevelId
	entity.PositionId = req.PositionId
	entity.Address = req.Address
	entity.Username = req.Username
	entity.Intro = req.Intro
	entity.Status = req.Status
	entity.Note = req.Note
	entity.Sort = req.Sort

	// 省市区处理
	if len(req.City) == 3 {
		entity.ProvinceCode = req.City[0]
		entity.CityCode = req.City[1]
		entity.DistrictCode = req.City[2]
	}

	// 密码
	if req.Password != "" {
		password, _ := utils.Md5(req.Password + req.Username)
		entity.Password = password
	}

	// 头像处理
	if req.Avatar != "" {
		avatar, err := utils.SaveImage(req.Avatar, "user")
		if err != nil {
			return 0, err
		}
		entity.Avatar = avatar
	}
	entity.CreateUser = userId
	entity.CreateTime = gtime.Now()
	entity.Mark = 1

	// 插入对象
	result, err := dao.User.Insert(entity)
	if err != nil {
		return 0, err
	}

	// 获取插入ID
	id, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}

	// 删除用户角色关系
	dao.UserRole.Delete("user_id=?", id)
	// 创建人员角色关系
	for _, v := range req.RoleIds {
		if v == 0 {
			continue
		}
		var userRole model.UserRole
		userRole.UserId = gconv.Int(id)
		userRole.RoleId = gconv.Int(v)
		dao.UserRole.Insert(userRole)
	}
	return id, nil
}

func (s *userService) Update(req *model.UserUpdateReq, userId int) (int64, error) {
	if utils.AppDebug() {
		return 0, gerror.New("演示环境，暂无权限操作")
	}
	// 查询记录
	info, err := dao.User.FindOne("id=?", req.Id)
	if err != nil {
		return 0, err
	}
	if info == nil {
		return 0, gerror.New("记录不存在")
	}
	// 设置对象
	info.Realname = req.Realname
	info.Nickname = req.Nickname
	info.Gender = req.Gender
	info.Avatar = req.Avatar
	info.Mobile = req.Mobile
	info.Email = req.Email
	info.Birthday = req.Birthday
	info.DeptId = req.DeptId
	info.LevelId = req.LevelId
	info.PositionId = req.PositionId
	info.Address = req.Address
	info.Username = req.Username
	info.Intro = req.Intro
	info.Status = req.Status
	info.Note = req.Note
	info.Sort = req.Sort

	// 省市区处理
	if len(req.City) == 3 {
		info.ProvinceCode = req.City[0]
		info.CityCode = req.City[1]
		info.DistrictCode = req.City[2]
	}

	// 密码
	if req.Password != "" {
		password, _ := utils.Md5(req.Password + req.Username)
		info.Password = password
	}

	// 头像处理
	if req.Avatar != "" {
		avatar, err := utils.SaveImage(req.Avatar, "user")
		if err != nil {
			return 0, err
		}
		info.Avatar = avatar
	}
	info.CreateUser = userId
	info.CreateTime = gtime.Now()

	// 更新记录
	result, err := dao.User.Save(info)
	if err != nil {
		return 0, err
	}

	// 删除用户角色关系
	dao.UserRole.Delete("user_id=?", info.Id)
	// 创建人员角色关系
	for _, v := range req.RoleIds {
		if v <= 0 {
			continue
		}
		var userRole model.UserRole
		userRole.UserId = info.Id
		userRole.RoleId = gconv.Int(v)
		dao.UserRole.Insert(userRole)
	}

	// 获取受影响行数
	rows, err := result.RowsAffected()
	if err != nil {
		return 0, err
	}
	return rows, nil
}

func (s *userService) Delete(ids string) (int64, error) {
	if utils.AppDebug() {
		return 0, gerror.New("演示环境，暂无权限操作")
	}
	// 记录ID
	idsArr := convert.ToInt64Array(ids, ",")
	// 删除记录
	result, err := dao.User.Delete("id in (?)", idsArr)
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

func (s *userService) Status(req *model.UserStatusReq, userId int) (int64, error) {
	if utils.AppDebug() {
		return 0, gerror.New("演示环境，暂无权限操作")
	}
	info, err := dao.User.FindOne("id=?", req.Id)
	if err != nil {
		return 0, err
	}
	if info == nil {
		return 0, gerror.New("记录不存在")
	}

	// 设置状态
	result, err := dao.User.Data(g.Map{
		"status":      req.Status,
		"update_user": userId,
		"update_time": gtime.Now(),
	}).Where(dao.User.Columns.Id, info.Id).Update()
	if err != nil {
		return 0, err
	}
	res, err := result.RowsAffected()
	if err != nil {
		return 0, err
	}
	return res, nil
}

func (s *userService) ResetPwd(id int, userId int) (int64, error) {
	if utils.AppDebug() {
		return 0, gerror.New("演示环境，暂无权限操作")
	}
	// 查询记录
	info, err := dao.User.FindOne("id=?", id)
	if err != nil {
		return 0, err
	}
	if info == nil {
		return 0, gerror.New("记录不存在")
	}
	// 设置初始密码
	password, err := utils.Md5("123456" + info.Username)
	if err != nil {
		return 0, err
	}

	// 初始化密码
	result, err := dao.User.Data(g.Map{
		"password":    password,
		"update_user": userId,
		"update_time": gtime.Now(),
	}).Where(dao.User.Columns.Id, info.Id).Update()
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

func (s *userService) UpdateUserInfo(req *model.UserInfoReq, r *ghttp.Request) (int64, error) {
	if utils.AppDebug() {
		return 0, gerror.New("演示环境，暂无权限操作")
	}

	// 头像处理
	avatar := ""
	if req.Avatar != "" {
		image, err := utils.SaveImage(req.Avatar, "user")
		if err != nil {
			return 0, err
		}
		avatar = image
	}

	// 更新用户信息
	result, err := dao.User.Data(g.Map{
		"avatar":   avatar,
		"realname": req.Realname,
		"nickname": req.Nickname,
		"gender":   req.Gender,
		"mobile":   req.Mobile,
		"email":    req.Email,
		"address":  req.Address,
		"intro":    req.Intro,
	}).Where("id", utils.Uid(r)).Update()
	if err != nil {
		return 0, err
	}

	// 获取受影响行数
	rows, err := result.RowsAffected()
	if err != nil {
		return 0, gerror.New("更新失败")
	}

	// 获取信息
	userInfo, _ := dao.User.FindOne("id=?", utils.Uid(r))
	// 设置SESSON
	r.Session.Set("userInfo", userInfo)
	return rows, nil
}

func (s *userService) UpdatePwd(req *model.UpdatePwd, userId int) (int64, error) {
	if utils.AppDebug() {
		return 0, gerror.New("演示环境，暂无权限操作")
	}
	// 查询信息
	info, err := dao.User.FindOne("id=?", userId)
	if err != nil {
		return 0, err
	}
	if info == nil {
		return 0, gerror.New("记录不存在")
	}
	// 比对旧密码
	oldPwd, err := utils.Md5(req.OldPassword + info.Username)
	if err != nil {
		return 0, err
	}
	if oldPwd != info.Password {
		return 0, gerror.New("旧密码不正确")
	}

	// 设置新密码
	if req.NewPassword != req.RePassword {
		return 0, gerror.New("两次输入的新密码不一致")
	}
	newPwd, err := utils.Md5(req.NewPassword + info.Username)
	if err != nil {
		return 0, err
	}

	result, err := dao.User.Data(g.Map{
		"password": newPwd,
	}).Where(dao.User.Columns.Id, userId).Update()
	if err != nil {
		return 0, err
	}

	// 获取受影响函数
	rows, err := result.RowsAffected()
	if err != nil {
		return 0, err
	}
	return rows, nil
}

func (s *userService) CheckUser(req *model.CheckUserReq) (*model.User, error) {
	user, err := dao.User.Where("username=? and mark=1", req.Username).FindOne()
	if err != nil {
		return nil, err
	}
	return user, nil
}
