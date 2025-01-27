package main

import (
	"crypto/md5"
	"flag"
	"fmt"
	"github.com/gin-gonic/gin"
	"os"
	"strconv"
	"time"
)

var (
	VERSION = "1.8.72"
)

var (
	ConfigSource string
	WebMode      string
	WebPort      int
	WebEnable    bool
	WebUsername  string
	WebPassword  string
	ApiKey       string
)

var (
	vkBakDict = make(map[string]int64)
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
	flag.StringVar(&WebMode, "mode", gin.ReleaseMode, "wol web mode: debug, release, test.")
	flag.IntVar(&WebPort, "port", 9090, "wol web port: 0-65535")
	flag.BoolVar(&WebEnable, "web", true, "wol web page switch: true or false.")
	flag.StringVar(&WebUsername, "username", "", "wol web page login username.")
	flag.StringVar(&WebPassword, "password", "", "wol web page login password.")
	flag.StringVar(&ApiKey, "key", "false", "wol web api key, key length greater than 6.")
}

func main() {
	flag.Parse()

	fmt.Printf("Start Run WolGoWeb...\n\n")
	fmt.Printf("Version: %s\n\n", VERSION)

	names, err := NetworkInterfaceNames()
	if err == nil {
		fmt.Printf("Network Interface Names: %+q\n\n", names)
	}

	if ConfigSource == "env" {
		WebMode = getEnvString("MODE", WebMode)
		WebPort = getEnvInt("PORT", WebPort)
		WebEnable = getEnvString("WEB", strconv.FormatBool(WebEnable)) == "true"
		WebUsername = getEnvString("USERNAME", WebUsername)
		WebPassword = getEnvString("PASSWORD", WebPassword)
		ApiKey = getEnvString("KEY", ApiKey)
	}

	gin.SetMode(WebMode)

	r := gin.Default()

	// 添加静态文件服务
	r.StaticFile("/", "./src/index.html")
	r.StaticFile("/index", "./src/index.html")

	if WebEnable {
		if WebUsername != "" && WebPassword != "" {
			ginUserAccount := gin.Accounts{WebUsername: WebPassword}
			r.GET("/", gin.BasicAuth(ginUserAccount), func(c *gin.Context) {
				c.File("./src/index.html")
			})
			r.GET("/index", gin.BasicAuth(ginUserAccount), func(c *gin.Context) {
				c.File("./src/index.html")
			})
			fmt.Printf("BasicAuth\n  username:%s\n  password:%s\n\n", WebUsername, WebPassword)
		}
	}
	r.GET("/wol", GetWol)

	fmt.Printf("WolGoWeb Runing [port:%d, key:%s, web:%s]\n", WebPort, ApiKey, strconv.FormatBool(WebEnable))

	err = r.Run(fmt.Sprintf(":%d", WebPort))
	if err != nil {
		fmt.Println(err.Error())
	}
}

func VerifyAuth(key string, mac string, vk int64, token string) (int, string) {
	err := 0
	message := "OK"
	if len(key) >= 6 {
		timeUnix := time.Now().Unix()
		fmt.Printf("now=%d, vk=%d\n", timeUnix, vk)
		if len(token) != 32 {
			err = 101
			message = "No authority."
		} else if timeUnix-vk > 30 || vk-timeUnix > 1 {
			err = 102
			message = "The value of Time is no longer in the valid range."
		} else if bakVK, ok := vkBakDict[mac]; ok && bakVK == vk {
			err = 103
			message = "Time value repetition."
		} else if MD5(ApiKey+mac+fmt.Sprintf("%d", vk)) != token {
			err = 104
			message = "No authority token."
		} else {
			vkBakDict[mac] = vk
		}
	}
	return err, message
}

func GetWol(c *gin.Context) {
	mac := c.Query("mac")
	ip := c.DefaultQuery("ip", "255.255.255.255")
	port := c.DefaultQuery("port", "9")
	network := c.DefaultQuery("network", "")
	token := c.DefaultQuery("token", "")
	vk, _ := strconv.ParseInt(c.DefaultQuery("time", "0"), 10, 64)
	if errAuth, messageAuth := VerifyAuth(ApiKey, mac, vk, token); errAuth == 0 {
		err := Wake(mac, ip, port, network)
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
