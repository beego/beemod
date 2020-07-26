package datasource

import (
	"bytes"
	. "github.com/smartystreets/goconvey/convey"
	"github.com/spf13/viper"
	"gopkg.in/ini.v1"
	"testing"
)

type Oss struct {
	CdnName string `ini:"cdnName"`
}

func TestDatasourceToml(t *testing.T) {
	viper.SetConfigType("toml")
	err := viper.ReadConfig(bytes.NewReader([]byte(`[beego.oss.myoss]
		mode = "file"
		debug = true
        isDeleteSrcPath = false
        cdnName = "http://127.0.0.1:8080/oss/"
        fileBucket = "oss"`)))

	Convey("unmarshal toml", t, func() {
		So(err, ShouldBeNil)
	})
	ds := &Toml{
		Viper: viper.GetViper(),
	}

	Convey("unmarshal toml to struct", t, func() {
		ds.Range("beego.oss", func(key string, name string) bool {
			var oss Oss
			So(name, ShouldEqual, "myoss")
			err = ds.Unmarshal(key, &oss)
			So(err, ShouldBeNil)
			So(oss.CdnName, ShouldEqual, "http://127.0.0.1:8080/oss/")
			return true
		})
	})
}

func TestDatasourceIni(t *testing.T) {
	f, err := ini.Load([]byte(`[beego.oss.myoss]
		debug = true
        isDeleteSrcPath = false
        cdnName = "http://127.0.0.1:8080/oss/"
        fileBucket = "oss"`))
	Convey("unmarshal ini", t, func() {
		So(err, ShouldBeNil)
	})
	ds := &Ini{
		Ini: f,
	}

	Convey("unmarshal ini to struct", t, func() {
		ds.Range("beego.oss", func(key string, name string) bool {
			var oss Oss
			So(name, ShouldEqual, "myoss")
			err = ds.Unmarshal(key, &oss)
			So(err, ShouldBeNil)
			So(oss.CdnName, ShouldEqual, "http://127.0.0.1:8080/oss/")
			return true
		})
	})
}
