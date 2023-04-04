package main

import (
	"crypto/md5"
	"flag"
	"fmt"
	"math"
	"os"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

var (
	VERSION = "1.5.6"
)

var (
	ConfigSource string
	WebMode      string
	WebPort      int
	ApiKey       string
)

func MD5(str string) string {
	data := []byte(str)
	has := md5.Sum(data)
	md5str := fmt.Sprintf("%x", has)
	return md5str
}

func getEnvString(name string, value string) string {
	ret := os.Getenv(name)
	if ret == "" {
		return value
	} else {
		return ret
	}
}

func getEnvInt(name string, value int) int {
	env := os.Getenv(name)
	if ret, err := strconv.Atoi(env); env == "" || err != nil {
		return value
	} else {
		return ret
	}
}

func init() {
	flag.StringVar(&ConfigSource, "c", "default", "config source default or env.")
	flag.StringVar(&WebMode, "mode", gin.ReleaseMode, "wol web port.")
	flag.IntVar(&WebPort, "port", 9090, "wol web port.")
	flag.StringVar(&ApiKey, "key", "false", "wol web api key.")
}

func main() {
	flag.Parse()

	fmt.Printf("Start Run WolGoWeb...\n\n")
	fmt.Printf("Version: %s\n\n", VERSION)

	if ConfigSource == "env" {
		WebMode = getEnvString("MODE", WebMode)
		WebPort = getEnvInt("PORT", WebPort)
		ApiKey = getEnvString("KEY", ApiKey)
	}

	gin.SetMode(WebMode)

	r := gin.Default()

	r.GET("/", GetIndex)
	r.GET("/index", GetIndex)
	r.GET("/wol", GetWol)

	fmt.Printf("WolGoWeb Runing [port:%d, key:%s]\n", WebPort, ApiKey)

	r.Run(fmt.Sprintf(":%d", WebPort))
}

func GetIndex(c *gin.Context) {
	c.String(200, `
WOL唤醒工具

API: %s/wol

Params:
  mac  : 需要唤醒的MAC地址（必须）,
  ip   : 指定IP地址（默认：255.255.255.255）,
  port : 唤醒端口（默认：9）,
  time : 请求时间戳（配合授权验证使用）,
  token: 授权Token = MD5(key + mac + time)（必须存在key的情况下才有效，否则忽略。）,

Version: %s
`, c.Request.Host, VERSION)
}

func VerifyAuth(key string, mac string, vk int64, token string) (int, string) {
	errCode := 0
	message := "OK"
	if len(key) >= 6 {
		vkStr := strconv.FormatInt(vk, 10)
		vkLen := float64(19 - int64(len(vkStr)))
		vknano := vk * int64(math.Pow(10, vkLen))
		nanoTime := time.Now().UnixNano()
		fmt.Printf("now=%d, vk=%d\n", nanoTime, vk)
		if len(token) != 32 {
			errCode = 101
			message = "Token value invalid."
		} else if vknano > nanoTime+5*1e9 {
			errCode = 102
			message = "Time value invalid."
		} else if nanoTime-vknano > 30*1e9 {
			errCode = 103
			message = "Time value invalid."
		} else if MD5(ApiKey+mac+vkStr) != token {
			errCode = 104
			message = "No authority token."
		}
	}
	return errCode, message
}

func GetWol(c *gin.Context) {
	mac := c.Query("mac")
	ip := c.DefaultQuery("ip", "255.255.255.255")
	port := c.DefaultQuery("port", "9")
	token := c.DefaultQuery("token", "")
	vk, _ := strconv.ParseInt(c.DefaultQuery("time", "0"), 10, 64)
	if errAuth, messageAuth := VerifyAuth(ApiKey, mac, vk, token); errAuth == 0 {
		err := Wake(mac, ip, port)
		if err != nil {
			c.JSON(200, gin.H{
				"error":   100,
				"message": fmt.Sprintf("%s", err),
			})
		} else {
			c.JSON(200, gin.H{
				"error":   0,
				"message": fmt.Sprintf("Wake Success Mac:%s", mac),
			})
		}
	} else {
		c.JSON(200, gin.H{
			"error":   errAuth,
			"message": messageAuth,
		})
	}

}
