/**
* @Author: myxy99 <myxy99@foxmail.com>
* @Date: 2020/7/30 10:34
 */
package mongo

type CallerCfg struct {
	Debug    bool `ini:"debug"`
	URL      string `ini:"url"`
	Source   string `ini:"source"`
	User     string `ini:"user"`
	Password string `ini:"password"`
}
