package standard

var AccessTokenExpireInterval int64

func InitAccessTokenExpireInterval(accessTokenExpireInterval int64) {
	AccessTokenExpireInterval = accessTokenExpireInterval
}

type AccessTokenTicket struct {
	AccessToken string `json:"accessToken"`
	ExpiresIn   int64  `json:"expiresIn"`
}

type TokenAccessor interface {
	CreateAccessToken(uid int, startTime int64) (resp AccessTokenTicket, err error)
	CheckAccessToken(tokenStr string) bool
	RefreshAccessToken(tokenStr string, startTime int64) (resp AccessTokenTicket, err error)
	DecodeAccessToken(tokenStr string) (resp map[string]interface{}, err error)
}
