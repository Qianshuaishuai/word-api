package helper

import (
	"crypto/md5"
	"crypto/rand"
	"crypto/sha1"
	"encoding/base64"
	"fmt"
	"io"
	mathrand "math/rand"
	"time"
)

//sha1
func Sha1(str string) string {
	return fmt.Sprintf("%x", sha1.Sum([]byte(str)))
}

//生成Guid字串(32位) 3a0f6874b37e0f12f8e9ea113985ad89
func GetGuid() string {
	b := make([]byte, 48)
	if _, err := io.ReadFull(rand.Reader, b); err != nil {
		return ""
	}
	return Md5(base64.URLEncoding.EncodeToString(b))
}

//md5
func Md5(str string) string {
	md5Str := fmt.Sprintf("%x", md5.Sum([]byte(str)))
	return md5Str
}

//md5(16位的)
func Md516(str string) string {
	data := md5.Sum([]byte(str))
	md5Str := string(data[0:16])
	return md5Str
}

func GetSaveToken() string {
	//token = 随机8位数 + 时间戳
	ntime := time.Now().Unix()
	rnd := mathrand.New(mathrand.NewSource(ntime))
	vcode := fmt.Sprintf("%08v", rnd.Int31n(10000000))

	return vcode + Int64ToString(ntime)
}
