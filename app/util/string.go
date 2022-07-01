package util

import (
	"fmt"
	"math/rand"
	"regexp"
	"time"
)

const account_rule = "^\\w{2,12}$"

func CheckAccount(account string) bool {
	r, _ := regexp.Compile(account_rule)
	return r.MatchString(account)
}

// 随机头像
const avatar_format = "https://www.thiswaifudoesnotexist.net/example-%d.jpg"

func GetAvatar() string {
	rand.Seed(time.Now().UnixNano())
	return fmt.Sprintf(avatar_format, rand.Intn(10000))
}
