package mysql

import "github.com/astaxie/beego/orm"

var TableName string

func init() {
	orm.RegisterModel(new(AccessToken))
}

type AccessToken struct {
	Jti        int    `orm:"pk;auto" json:"jti" form:"jti"`
	Sub        int    `orm:"unique" json:"sub" form:"sub"`
	IaTime     int64  `orm:"" json:"ia_time" form:"ia_time"`
	ExpTime    int64  `orm:"" json:"exp_time" form:"exp_time"`
	Ip         string `orm:"" json:"ip" form:"ip"`
	CreateTime int64  `orm:"" json:"create_time" form:"create_time"`
	IsLogout   int    `orm:"" json:"is_logout" form:"is_logout"`
	IsInvalid  int    `orm:"" json:"is_invalid" form:"is_invalid"`
	LogoutTime int64  `orm:"" json:"logout_time" form:"logout_time"`
}

func (a *AccessToken) TableName() string {
	return "access_token"
}
