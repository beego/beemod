package datasource

import (
	"github.com/spf13/viper"
	"gopkg.in/ini.v1"
	"strings"
)

var sep = "."

type Datasource interface {
	Range(key string, f func(key string, name string) bool)
	Unmarshal(key string, value interface{}) error
}

type Toml struct {
	Viper *viper.Viper
}

func (t *Toml) Range(prefix string, f func(key string, name string) bool) {
	for _, k := range t.Viper.AllKeys() {
		//k beego.oss.myoss.debug key+sep = beego.oss.
		if !strings.HasPrefix(k, prefix+sep) {
			continue
		}

		// nv myoss.debug
		nv := strings.TrimPrefix(k, prefix+sep)
		arr := strings.Split(nv, sep)
		if len(arr) != 2 {
			continue
		}

		//  arr[0] myoss
		if !f(prefix+sep+arr[0], arr[0]) {
			continue
		}
	}
}
func (t *Toml) Unmarshal(key string, value interface{}) error {
	return t.Viper.UnmarshalKey(key, &value)
}

type Ini struct {
	Ini *ini.File
}

func (t *Ini) Range(prefix string, f func(key string, name string) bool) {
	for _, value := range t.Ini.Sections() {
		// value.name beego.oss.myoss
		if !strings.HasPrefix(value.Name(), prefix+sep) {
			continue
		}

		n := strings.TrimPrefix(value.Name(), prefix+sep)
		if strings.Contains(n, sep) {
			continue
		}
		if !f(value.Name(), n) {
			continue
		}
	}
}

func (t *Ini) Unmarshal(key string, value interface{}) error {
	return t.Ini.Section(key).MapTo(value)
}
