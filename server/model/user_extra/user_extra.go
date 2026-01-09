
// 自动生成模板UserExtra
package user_extra
import (
)

// user extra info 结构体  UserExtra
type UserExtra struct {
  Id  *int `json:"id" form:"id" gorm:"primarykey;column:id;size:10;"`  //id
  SysUserId  *int `json:"sysUserId" form:"sysUserId" gorm:"column:sys_user_id;size:10;"`  //系统用户 ID
  AgentCode  *string `json:"agentCode" form:"agentCode" gorm:"column:agent_code;size:100;"`  //agentCode字段
  AgentName  *string `json:"agentName" form:"agentName" gorm:"column:agent_name;size:100;"`  //agentName字段
  ServiceType  *string `json:"serviceType" form:"serviceType" gorm:"column:service_type;size:100;"`  //serviceType字段
  UccId  *string `json:"uccId" form:"uccId" gorm:"column:ucc_id;size:100;"`  //uccId字段
}


// TableName user extra info UserExtra自定义表名 user_extra
func (UserExtra) TableName() string {
    return "user_extra"
}






