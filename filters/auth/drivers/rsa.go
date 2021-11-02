/*
 * @Author: your name
 * @Date: 2020-12-19 14:03:21
 * @LastEditTime: 2020-12-22 18:59:22
 * @LastEditors: Please set LastEditors
 * @Description: In User Settings Edit
 * @FilePath: /src/aglgin/filters/auth/drivers/jwt.go
 */
package drivers

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/base64"
	"encoding/json"
	"encoding/pem"
	"errors"
	"fmt"
	"gin/config"
	"gin/service"
	"io/ioutil"
	"net/http"

	"github.com/garyburd/redigo/redis"
	"github.com/gin-gonic/gin"
)

//openssl genrsa -out rsa_private_key.pem 1024
var adminPrivateKey = []byte(`-----BEGIN RSA PRIVATE KEY-----
MIICWwIBAAKBgQCuxk2OafXt888I0msU/k9xmyhdIHZuouFfDmtegGGdqoaM+sIt
A3C30gUvGL56yjysQYlf1kCdacXr2bF4f6z0u6rv84E59w8WeXfS0I3YeUjwGQ5t
mI8R+Rf8idre1DCkTpp5ffsA5z0SFzz1gsVyd5DoWl4oJY227H0HY3r5xQIDAQAB
AoGAAtHDaBiDEOZCUz124FQJDQwKpJQ7jgEzpjz68/+HLwdv1zkVM9BDUpc8quLC
FijcmUXqzgvn1RSHhXpECiSEjaXGLxo072FOFyslTQ/TbmKSN1EONOkLQKwy3r/f
ExxAvvs+saDWM4pbvOXDJmEbTFJyBLXzlWcedmyVRpBoovECQQD73OwogIFP8olM
Y1LrD3T8eWLZeGaOqT/pRrB3TZll9k3eLIHjH1dXPQDkT7U2J+i7weuWMwo7Umrn
N2Zcs1ANAkEAsaU50WSwbk0FiNfps2N1mJTFp4wkc1js/IE2GYwmwSDfdhhHPwW0
UtNdKdojpYsv0j7j3f1M0yyD/0GmXGYqmQJAFoJC9MevRtbVIGeMBIfoG5w5klfp
SnyjwpRXtwHPYMZnZSCzJvopExnXl4/sEP/2E7mb9VtwYabW+P0Bf+1ijQJAYdkC
ScXOMFMYU1GqFfcYlNyNKkZU5Xv7vPFm3ReHWSVEMIYa6Cm6M0zcqerPa6WIx6OA
W4vjvwVsBzMf8RENMQJAIYP30rtjf2boAFrkqrjc5SnfW1YrHFVUdX4iFq5FjSyG
nvLJgPlOW7F++ahqEm8VCn6Tv0vv/BpiX88u5XbBvA==
-----END RSA PRIVATE KEY-----
`)

//openssl
//openssl rsa -in rsa_private_key.pem -pubout -out rsa_public_key.pem
var adminPublicKey = []byte(`-----BEGIN PUBLIC KEY-----
MIGfMA0GCSqGSIb3DQEBAQUAA4GNADCBiQKBgQCuxk2OafXt888I0msU/k9xmyhd
IHZuouFfDmtegGGdqoaM+sItA3C30gUvGL56yjysQYlf1kCdacXr2bF4f6z0u6rv
84E59w8WeXfS0I3YeUjwGQ5tmI8R+Rf8idre1DCkTpp5ffsA5z0SFzz1gsVyd5Do
Wl4oJY227H0HY3r5xQIDAQAB
-----END PUBLIC KEY-----`)

type rsaAuthManager struct {
	sign string
}

/**
 * @description: Rsa加密
 * @param {*baseController} yxm --- 2020-12-21
 * @return {*} []byte, err
 */
func AuthRsaEncrypt(origData []byte) ([]byte, error) {
	//解密pem格式的公钥
	block, _ := pem.Decode(adminPublicKey)
	if block == nil {
		return nil, errors.New("public key error")
	}
	// 解析公钥
	pubInterface, err := x509.ParsePKIXPublicKey(block.Bytes)
	if err != nil {
		return nil, err
	}
	// 类型断言
	pub := pubInterface.(*rsa.PublicKey)
	//加密
	return rsa.EncryptPKCS1v15(rand.Reader, pub, origData)
	//return rsa.EncryptPKCS1v15(rand.Reader, pub, origData)
}

/**
 * @description: Rsa解密
 * @param {*baseController} yxm --- 2020-12-21
 * @return {*} []byte, err
 */
func AuthRsaDecrypt(ciphertext []byte) ([]byte, error) {
	//解密
	block, _ := pem.Decode(adminPrivateKey)
	if block == nil {
		return nil, errors.New("private key error!")
	}
	//p.Ctx.Output.JSON(err, false, true)
	//解析PKCS1格式的私钥
	priv, err := x509.ParsePKCS1PrivateKey(block.Bytes)
	if err != nil {
		return nil, err
	}
	// 解密
	return rsa.DecryptPKCS1v15(rand.Reader, priv, ciphertext)
}

func NewRsaAuthDriver() *rsaAuthManager {
	return &rsaAuthManager{
		sign: config.GetRsaConfig().SIGN,
	}
}

// 检查请求头的令牌是否有效。
func (rsaAuth *rsaAuthManager) Check(c *gin.Context) bool {
	//r.Ctx.Input.RequestBody   获取request请求数据类型为json的数据
	var authStatus bool = false
	jsonDatabytes := c.Request.Header.Get(rsaAuth.sign)

	signBytes, err := base64.StdEncoding.DecodeString(jsonDatabytes)

	if err != nil {
		return authStatus
	}
	origData, _ := AuthRsaDecrypt(signBytes)
	sign_map := make(map[string]string)
	err = json.Unmarshal(origData, &sign_map)
	if err != nil {
		return authStatus
	}
	//redis验证
	redis_pool, err := service.GetRedisConnection(2)
	if err != nil {
		return authStatus
	}
	defer redis_pool.Close() //函数运行结束 ，把连接放回连接池
	redis_key := "login:admin:uuid:" + sign_map["admin_uuid"]

	r, error := redis.String(redis_pool.Do("Get", redis_key))
	if error != nil {
		return authStatus
	}

	userinfo := make(map[string]interface{})

	err = json.Unmarshal([]byte(r), &userinfo)
	if err != nil {
		return authStatus
	}

	if userinfo["admin_token"].(string) != sign_map["admin_token"] {
		return authStatus
	}

	authStatus = true
	c.Set("rsa_auth_sign", sign_map)

	return authStatus
}

func (rsaAuth *rsaAuthManager) User(c *gin.Context) interface{} {
	var rsaAuthSign interface{}
	var exist bool
	if rsaAuthSign, exist = c.Get("rsa_auth_sign"); !exist {
		return nil
	}
	fmt.Print(rsaAuthSign)

	return rsaAuthSign
}

func (rsaAuth *rsaAuthManager) Login(http *http.Request, w http.ResponseWriter, user map[string]interface{}) interface{} {
	var sign string
	jsonDatabytes, _ := ioutil.ReadAll(http.Body)
	if jsonDatabytes != nil {
		data, _ := AuthRsaEncrypt(jsonDatabytes)
		sign = base64.StdEncoding.EncodeToString(data)

	} else {
		return nil
	}

	return sign
}

func (rsaAuth *rsaAuthManager) Logout(http *http.Request, w http.ResponseWriter) bool {

	return true
}
