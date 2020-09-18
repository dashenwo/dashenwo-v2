package crypto

import (
	"crypto/md5"
	"fmt"
	"github.com/micro/go-micro/v2/util/log"
	"golang.org/x/crypto/bcrypt"
	"math/rand"
	"time"
)

// 随机生成指定长度字符串
func GetRandomString(l int) string {
	str := "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHJKLMMNOPQRSTUVWXYZ"
	bytes := []byte(str)
	result := []byte{}
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	for i := 0; i < l; i++ {
		result = append(result, bytes[r.Intn(len(bytes))])
	}
	return string(result)
}

// 密码加密
func HashAndSalt(pwd string, salt string) string {
	var newPwd = pwd + salt
	SaltPwd := []byte(newPwd)
	hash, err := bcrypt.GenerateFromPassword(SaltPwd, bcrypt.MinCost)
	if err != nil {
		log.Log(err)
	}
	return string(hash)
}

// 验证密码
func ComparePasswords(hashedPwd string, password string) bool {
	byteHash := []byte(hashedPwd)
	plainPwd := []byte(password)
	err := bcrypt.CompareHashAndPassword(byteHash, plainPwd)
	if err != nil {
		return false
	}
	return true
}

// md5
func Md5(str string) string {
	data := []byte(str)
	has := md5.Sum(data)
	md5str := fmt.Sprintf("%x", has)
	return md5str
}
