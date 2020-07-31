package logger

type CallerCfg struct {
	//	Type   = "console"
	//	Type      = "file"
	//	Type = "multifile"
	//	Type      = "smtp"
	//	Type      = "conn"
	//	Type        = "es"
	//	Type  = "jianliao"
	//	Type     = "slack"
	//	Type     = "alils"
	Type  string `ini:"type"`
	Level int    `ini:"level"`
	//file
	Path     string `ini:"path"`
	Maxlines int    `ini:"maxlines"`
	Maxsize  int    `ini:"maxsize"`
	Daily    bool   `ini:"daily"`
	Maxdays  int    `ini:"maxdays"`
	Rotate   bool   `ini:"rotate"`
	Perm     string `ini:"perm"`
	//multifile
	Separate string `ini:"separate"`
	//conn
	ReconnectOnMsg bool   `ini:"reconnectOnMsg"`
	Reconnect      bool   `ini:"reconnect"`
	Net            string `ini:"net"`
	Addr           string `ini:"addr"`
	//smtp
	Username string `ini:"username"`
	Password string `ini:"password"`
	Host     string `ini:"host"`
	SendTos  string `ini:"sendTos"`
	Subject  string `ini:"subject"`
	//ElasticSearch
	Dsn string `ini:"dsn"`
	//jianliao
	Authorname  string `ini:"authorname"`
	Title       string `ini:"title"`
	Webhookurl  string `ini:"webhookurl"`
	Redirecturl string `ini:"redirecturl"`
	Imageurl    string `ini:"imageurl"`
	//slack
	SlackWebhookurl string `ini:"slackWebhookurl"`
}

var DefaultInvokerCfg = CallerCfg{
	Type:            "console",
	Level:           7,
	Path:            "",
	Maxlines:        1000000,
	Maxsize:         256,
	Daily:           true,
	Maxdays:         7,
	Rotate:          true,
	Perm:            "0660",
	Separate:        "",
	ReconnectOnMsg:  false,
	Reconnect:       false,
	Net:             "",
	Addr:            "",
	Username:        "",
	Password:        "",
	Host:            "",
	SendTos:         "",
	Subject:         "Diagnostic message from server",
	Dsn:             "",
	Authorname:      "",
	Title:           "",
	Webhookurl:      "",
	Redirecturl:     "",
	Imageurl:        "",
	SlackWebhookurl: "",
}
