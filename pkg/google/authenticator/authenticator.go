/**
* @Author: myxy99 <myxy99@foxmail.com>
* @Date: 2020/8/6 18:42
 */
package authenticator

import (
	"crypto/hmac"
	"crypto/sha1"
	"encoding/base32"
	"fmt"
	"github.com/beego/beemod/pkg/datasource"
	"github.com/beego/beemod/pkg/module"
	"hash"
	"math/rand"
	"net/url"
	"strings"
	"sync"
	"time"
)

var defaultInvoker = &descriptor{
	Name: module.Authenticator,
	Key:  module.ConfigPrefix + module.Authenticator,
}

type descriptor struct {
	Name  string
	Key   string
	store sync.Map
	cfg   map[string]InvokerCfg
}

type GoogleAuth struct {
	user, issuer, secret string
}

type Client struct {
	Auth *GoogleAuth
	cfg  InvokerCfg
}

// default invoker build
func DefaultBuild() module.Invoker {
	return defaultInvoker
}

// invoker
func Invoker(name string) *Client {
	obj, ok := defaultInvoker.store.Load(name)
	if !ok {
		return nil
	}
	return obj.(*Client)
}

func (c *descriptor) InitCfg(ds datasource.Datasource) error {
	c.cfg = make(map[string]InvokerCfg, 0)
	ds.Range(c.Key, func(key string, name string) bool {
		config := DefaultInvokerCfg
		if err := ds.Unmarshal(key, &config); err != nil {
			return false
		}
		c.cfg[name] = config
		return true
	})
	return nil
}

func (c *descriptor) Run() error {
	for name, cfg := range c.cfg {
		db := provider(cfg)
		c := &Client{
			db,
			cfg,
		}
		defaultInvoker.store.Store(name, c)
	}
	return nil
}

func provider(cfg InvokerCfg) *GoogleAuth {
	return newGoogleAuth(cfg.User, cfg.Iss)
}

var (
	codes   = "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz"
	codeLen = len(codes)
)

func getCode(key []byte, value []byte) uint32 {
	var (
		hmacSha1         hash.Hash
		bytes, hashParts []byte
		offset           uint8
		number, pwd      uint32
	)
	hmacSha1 = hmac.New(sha1.New, key)
	hmacSha1.Write(value)
	bytes = hmacSha1.Sum(nil)
	offset = bytes[len(bytes)-1] & 0x0F
	hashParts = bytes[offset : offset+4]
	hashParts[0] = hashParts[0] & 0x7F
	number = toUint32(hashParts)
	pwd = number % 1000000
	return pwd
}
func toBytes(value int64) []byte {
	var (
		result []byte
		mask   int64
		shifts [8]uint16
	)
	mask = int64(0xFF)
	shifts = [8]uint16{56, 48, 40, 32, 24, 16, 8, 0}
	for _, shift := range shifts {
		result = append(result, byte((value>>shift)&mask))
	}
	return result
}
func toUint32(bytes []byte) uint32 {
	return (uint32(bytes[0]) << 24) + (uint32(bytes[1]) << 16) +
		(uint32(bytes[2]) << 8) + uint32(bytes[3])
}
func (g *GoogleAuth) getCode() (code string, err error) {
	var (
		key      []byte
		secret   string
		codeUI32 uint32
	)
	secret = strings.ToUpper(strings.Replace(g.secret, " ", "", -1))
	key, err = base32.StdEncoding.DecodeString(secret)
	if err != nil {
		return "", nil
	}
	codeUI32 = getCode(key, toBytes(time.Now().Unix()/30))
	code = fmt.Sprintf("%0*d", 6, codeUI32)
	return
}

//生成url
func (g *GoogleAuth) ProvisionURL() (codeUrl string) {
	auth := "totp/"
	q := make(url.Values)
	q.Add("secret", g.secret)
	if g.issuer != "" {
		q.Add("issuer", g.issuer)
		auth += g.issuer + ":"
	}
	return "otpauth://" + auth + g.user + "?" + q.Encode()
}

//生成验证code正确
func (g *GoogleAuth) Authenticate(code string) (ok bool, err error) {
	var (
		nowCode string
	)
	nowCode, err = g.getCode()
	return nowCode == code, err
}

//生成密钥
func (g *GoogleAuth) GenerateKey() string {
	data := make([]byte, 32)
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < 32; i++ {
		idx := rand.Intn(codeLen)
		data[i] = byte(codes[idx])
	}
	g.secret = string(data)
	return string(data)
}

//设置密钥
func (g *GoogleAuth) SetKey(key string) {
	g.secret = key
	return
}
func newGoogleAuth(user string, issuer ...string) *GoogleAuth {
	var iss string
	if len(issuer) == 1 {
		iss = issuer[0]
	}
	return &GoogleAuth{
		user:   user,
		issuer: iss,
	}
}
