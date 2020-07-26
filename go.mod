module github.com/beego/beemod

go 1.14

require (
	github.com/BurntSushi/toml v0.3.1
	github.com/aliyun/alibaba-cloud-sdk-go v1.61.353
	github.com/aliyun/aliyun-oss-go-sdk v2.1.4+incompatible
	github.com/astaxie/beego v1.12.2
	github.com/baiyubin/aliyun-sts-go-sdk v0.0.0-20180326062324-cfa1a18b161f // indirect
	github.com/dgrijalva/jwt-go v3.2.0+incompatible
	github.com/go-sql-driver/mysql v1.5.0
	github.com/gomodule/redigo/redis v0.0.0-20200429221454-e14091dffc1b
	github.com/higker/qrcode-go v0.0.2
	github.com/satori/go.uuid v1.2.0 // indirect
	github.com/satori/uuid v1.2.0
	github.com/smartystreets/goconvey v1.6.4
	github.com/spf13/viper v1.7.0
	github.com/stretchr/testify v1.6.1
	github.com/tencentcloud/tencentcloud-sdk-go v3.0.219+incompatible
	gopkg.in/alexcesaro/quotedprintable.v3 v3.0.0-20150716171945-2caba252f4dc // indirect
	gopkg.in/gomail.v2 v2.0.0-20160411212932-81ebce5c23df
	gopkg.in/ini.v1 v1.57.0
)

replace github.com/gomodule/redigo v2.0.0+incompatible => github.com/gomodule/redigo/redis v0.0.0-20200429221454-e14091dffc1b
