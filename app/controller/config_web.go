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
 * 网站配置-控制器
 * @author 半城风雨
 * @since 2021/7/28
 * @File : config_set
 */
package controller

import (
	"easygoadmin/app/dao"
	"easygoadmin/app/utils"
	"easygoadmin/app/utils/common"
	"encoding/json"
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
	"github.com/gogf/gf/text/gstr"
	"github.com/gogf/gf/util/gconv"
	"reflect"
	"regexp"
)

// 控制器管理对象
var ConfigWeb = new(configWeb)

type configWeb struct{}

func (c *configWeb) Index(r *ghttp.Request) {
	// 获取配置列表
	configData, _ := dao.Config.Where("mark=1").All()
	configList := make([]map[string]interface{}, 0)
	for _, v := range configData {
		item := make(map[string]interface{})
		item["config_id"] = v.Id
		item["config_name"] = v.Name

		// 查询配置项列表
		itemData, _ := dao.ConfigData.Where("config_id=? and status=1 and mark=1", v.Id).Order("sort asc").All()
		itemList := make([]map[string]interface{}, 0)
		for _, v := range itemData {
			item := make(map[string]interface{})
			item["id"] = v.Id
			item["title"] = v.Title
			item["code"] = v.Code
			item["value"] = v.Value
			item["type"] = v.Type

			if v.Type == "checkbox" {
				// 复选框
				re := regexp.MustCompile(`\r?\n`)
				list := gstr.Split(re.ReplaceAllString(v.Options, "|"), "|")
				optionsList := make(map[string]string)
				for _, val := range list {
					re2 := regexp.MustCompile(`:|：|\s+`)
					item := gstr.Split(re2.ReplaceAllString(val, "|"), "|")
					optionsList[item[0]] = item[1]
				}
				// 选择项
				item["optionsList"] = optionsList
				// 选择值
				item["value"] = gstr.Split(v.Value, ",")

			} else if v.Type == "radio" {
				// 单选框
				re := regexp.MustCompile(`\r?\n`)
				list := gstr.Split(re.ReplaceAllString(v.Options, "|"), "|")
				optionsList := make(map[string]string)
				for _, v := range list {
					re2 := regexp.MustCompile(`:|：|\s+`)
					item := gstr.Split(re2.ReplaceAllString(v, "|"), "|")
					optionsList[item[0]] = item[1]
				}
				item["optionsList"] = optionsList

			} else if v.Type == "select" {
				// 下拉选择框
				re := regexp.MustCompile(`\r?\n`)
				list := gstr.Split(re.ReplaceAllString(v.Options, "|"), "|")
				optionsList := make(map[string]string)
				for _, v := range list {
					re2 := regexp.MustCompile(`:|：|\s+`)
					item := gstr.Split(re2.ReplaceAllString(v, "|"), "|")
					optionsList[item[0]] = item[1]
				}
				item["optionsList"] = optionsList
			} else if v.Type == "image" {
				// 单图片
				item["value"] = utils.GetImageUrl(v.Value)
			} else if v.Type == "images" {
				// 多图片
				list := gstr.Split(v.Value, ",")
				itemList := make([]string, 0)
				for _, v := range list {
					// 图片地址
					item := utils.GetImageUrl(v)
					itemList = append(itemList, item)
				}
				item["value"] = itemList
			}
			itemList = append(itemList, item)
		}
		item["itemList"] = itemList
		// 加入数组
		configList = append(configList, item)
	}

	// 返回结果
	r.Response.WriteJsonExit(common.JsonResult{
		Code: 0,
		Msg:  "操作成功",
		Data: configList,
	})
}

func (c *configWeb) Save(r *ghttp.Request) {
	if utils.AppDebug() {
		r.Response.WriteJsonExit(common.JsonResult{
			Code: -1,
			Msg:  "演示环境，暂无权限操作",
		})
	}
	// key：string类型，value：interface{}  类型能存任何数据类型
	var jsonObj map[string]interface{}
	json.Unmarshal(r.GetBody(), &jsonObj)
	for key, val := range jsonObj {
		// 数组处理
		if reflect.ValueOf(val).Kind() == reflect.Slice {
			if reflect.TypeOf(val).String() == "[]interface {}" {
				// 初始化URL数组
				item := make([]string, 0)
				for _, v := range val.([]interface{}) {
					value := gconv.String(v)
					// 判断是否http(s)开头
					if gstr.SubStr(value, 0, 4) == "http" ||
						gstr.SubStr(value, 0, 5) == "https" {
						// 图片地址
						if gstr.Contains(value, utils.ImgUrl()) {
							url, _ := utils.SaveImage(value, "config")
							item = append(item, url)
						}
					} else {
						// 复选框处理
						item = append(item, value)
					}
				}
				// 逗号拼接
				val = gstr.Join(item, ",")
			}
		} else {
			// 图片处理
			if gstr.Contains(gconv.String(val), "http://") ||
				gstr.Contains(gconv.String(val), "https://") {
				// 图片地址
				if gstr.Contains(gconv.String(val), utils.ImgUrl()) {
					val, _ = utils.SaveImage(gconv.String(val), "config")
				}
			}
		}

		// 查询记录
		info, err := dao.ConfigData.FindOne("code=?", key)
		if err != nil || info == nil {
			continue
		}

		// 更新记录
		result, err := dao.ConfigData.Data(g.Map{
			"value": val,
		}).Where("code=?", key).Update()
		if err != nil {
			continue
		}

		// 获取受影响行数
		rows, _ := result.RowsAffected()
		if rows == 0 {
			continue
		}
	}

	// 返回结果
	r.Response.WriteJsonExit(common.JsonResult{
		Code: 0,
		Msg:  "保存成功",
	})
}
