// ==========================================================================
// This is auto-generated by gf cli tool. DO NOT EDIT THIS FILE MANUALLY.
// ==========================================================================

package internal

import (
    "github.com/gogf/gf/os/gtime"
)

// Menu is the golang structure for table sys_menu.
type Menu struct {
    Id         int         `orm:"id,primary"  json:"id"`         // 唯一性标识                     
    Name       string      `orm:"name"        json:"name"`       // 菜单名称                       
    Icon       string      `orm:"icon"        json:"icon"`       // 图标                           
    Url        string      `orm:"url"         json:"url"`        // URL地址                        
    Param      string      `orm:"param"       json:"param"`      // 参数                           
    Pid        int         `orm:"pid"         json:"pid"`        // 上级ID                         
    Type       int         `orm:"type"        json:"type"`       // 类型：1模块 2导航 3菜单 4节点  
    Permission string      `orm:"permission"  json:"permission"` // 权限标识                       
    Status     int         `orm:"status"      json:"status"`     // 是否显示：1显示 2不显示        
    Target     int         `orm:"target"      json:"target"`     // 打开方式：1内部打开 2外部打开  
    Note       string      `orm:"note"        json:"note"`       // 菜单备注                       
    Sort       int         `orm:"sort"        json:"sort"`       // 显示顺序                       
    CreateUser int         `orm:"create_user" json:"createUser"` // 添加人                         
    CreateTime *gtime.Time `orm:"create_time" json:"createTime"` // 添加时间                       
    UpdateUser int         `orm:"update_user" json:"updateUser"` // 更新人                         
    UpdateTime *gtime.Time `orm:"update_time" json:"updateTime"` // 更新时间                       
    Mark       int         `orm:"mark"        json:"mark"`       // 有效标识                       
}