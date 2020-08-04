package mysql

import (
	"github.com/astaxie/beego/orm"
	"github.com/beego/beemod/pkg/logger"
	"github.com/beego/beemod/pkg/token/standard"
	"time"
)

// 如果你希望使用这个实现来作为token的实现，那么需要在配置文件里面设置：
// [muses.logger.system]
//    ...logger的配置
// [muses.mysql.default]
//    ...mysql的配置
// [muses.token.jwt.mysql]
//    logger = "system"
//    client = "default"
// 而后将Register()方法注册进去muses.Container(...)中
type mysqlTokenAccessor struct {
	standard.JwtTokenAccessor
	logger *logger.Client
	db     orm.Ormer
}

func InitTokenAccessor(logger *logger.Client, db orm.Ormer) standard.TokenAccessor {
	return &mysqlTokenAccessor{
		JwtTokenAccessor: standard.JwtTokenAccessor{},
		logger:           logger,
		db:               db,
	}
}

func (accessor *mysqlTokenAccessor) CreateAccessToken(uid int, startTime int64) (resp standard.AccessTokenTicket, err error) {
	AccessTokenData := &AccessToken{
		Sub:        uid,
		IaTime:     startTime,
		ExpTime:    startTime + standard.AccessTokenExpireInterval,
		Ip:         "",
		CreateTime: time.Now().Unix(),
		IsLogout:   0,
		IsInvalid:  0,
		LogoutTime: 0,
	}
	qs := accessor.db.QueryTable("access_token")
	accessToken := AccessToken{}
	if err = qs.Filter("sub__exact", uid).One(&accessToken); err != nil {
		if _, err = accessor.db.Insert(AccessTokenData); err != nil {
			accessor.logger.Error("create accessToken create error", err.Error())
			return
		}
	} else {
		AccessTokenData.Jti = accessToken.Jti
		if _, err = accessor.db.Update(AccessTokenData); err != nil {
			accessor.logger.Error("Update accessToken create error", err.Error())
			return
		}
	}

	tokenString, err := accessor.EncodeAccessToken(AccessTokenData.Jti, uid, startTime)
	if err != nil {
		return
	}
	resp.AccessToken = tokenString
	resp.ExpiresIn = standard.AccessTokenExpireInterval
	return
}

func (accessor *mysqlTokenAccessor) CheckAccessToken(tokenStr string) bool {
	sc, err := accessor.DecodeAccessToken(tokenStr)
	if err != nil {
		accessor.logger.Error("access_token CheckAccessToken error1", err.Error())
		return false
	}
	var resp AccessToken
	qs := accessor.db.QueryTable("access_token")
	if err = qs.
		Filter("jti__exact", sc["jti"]).
		Filter("sub__exact", sc["sub"]).
		Filter("exp_time__gte", sc["exp"]).
		Filter("is_invalid__exact", 0).
		Filter("is_logout__exact", 0).
		One(&resp);
		err != nil {
		accessor.logger.Error("access_token CheckAccessToken error2", err.Error())
		return false
	}
	return true
}

func (accessor *mysqlTokenAccessor) RefreshAccessToken(tokenStr string, startTime int64) (resp standard.AccessTokenTicket, err error) {
	sc, err := accessor.DecodeAccessToken(tokenStr)
	if err != nil {
		accessor.logger.Error("access_token CheckAccessToken error1", err.Error())
		return
	}

	jti := sc["jti"].(int)
	refreshToken, err := accessor.EncodeAccessToken(jti, sc["uid"].(int), startTime)

	if err != nil {
		return
	}
	qs := accessor.db.QueryTable("access_token")
	_, err = qs.Filter("jti__exact", jti).Update(map[string]interface{}{
		"exp_time": startTime + standard.AccessTokenExpireInterval,
	})

	if err != nil {
		return
	}

	resp.AccessToken = refreshToken
	resp.ExpiresIn = standard.AccessTokenExpireInterval
	return
}
