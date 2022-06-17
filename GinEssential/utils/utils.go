package utils

import (
	"awesomeProject1/model"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
	"log"
	"math/rand"
	"time"
)

func RandomString(n int) string {
	var letters = []byte("qwertyuiopasdfghjklzxcvbnmQWERTYUIOPASDFGHJKLZXCVBNM")
	result := make([]byte, n)

	//让每次随机的值都和上次不一样
	rand.Seed(time.Now().Unix())
	for i := range result {
		result[i] = letters[rand.Intn(len(letters))]
	}
	return string(result)
}

// SecretPassword 对用户密码进行哈希加密
func SecretPassword(password string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		log.Println("password encryption failed,err:", err)
		return "", err
	}
	return string(hash), err
}

// VerifyPassword 用bcrypt进行密码比对
func VerifyPassword(secretPassword, password []byte) bool {
	err := bcrypt.CompareHashAndPassword(secretPassword, password)
	if err != nil {
		return false
	}
	return true
}

// IsTelephoneExist 判断手机号是否存在
func IsTelephoneExist(db *gorm.DB, telephone string) bool {
	var user model.User
	db.Where(
		"telephone = ?",
		telephone,
	).First(&user)
	if user.ID != 0 {
		return true
	}
	return false
}
