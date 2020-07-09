package main

import (
	"crypto/md5"
	"flag"
	"fmt"
	"github.com/gin-gonic/gin"
)

var (
	WebMode = flag.String("mode", gin.ReleaseMode, "wol web port.")
	WebPort = flag.Int("port", 9090, "wol web port.")
	ApiKey = flag.String("key", "false", "wol web api key.")
)

func MD5(str string) string {
	data := []byte(str)
	has := md5.Sum(data)
	md5str := fmt.Sprintf("%x", has)
	return md5str
}

func main()  {
	flag.Parse()

	gin.SetMode(*WebMode)

	r:=gin.Default()

	r.GET("/", GetIndex)
	r.GET("/index", GetIndex)
	r.GET("/wol", GetWol)

	fmt.Printf("WolGoWeb Runing [port:%d]\n", *WebPort)

	r.Run(fmt.Sprintf(":%d", *WebPort))
}

func GetIndex(c *gin.Context) {
	c.String(200, `
WOL唤醒工具
API: %s/wol
Params:
  mac  : 需要唤醒的MAC地址（必须）,
  ip   : 指定IP地址（默认：255.255.255.255）,
  port : 唤醒端口（默认：9）,
`, c.Request.Host)
}

func GetWol(c *gin.Context)  {
	mac:=c.Query("mac")
	ip:=c.DefaultQuery("ip", "255.255.255.255")
	port:=c.DefaultQuery("port", "9")
	if *ApiKey != "false" {
		token:=c.DefaultQuery("token", "")
		vk:=c.DefaultQuery("vk", "")
		if len(token) != 32 {
			c.JSON(200, gin.H{
				"error": 101,
				"message": "No authority.",
			})
			return
		}
		if len(vk) < 6 {
			c.JSON(200, gin.H{
				"error": 102,
				"message": "VK requires a minimum of 6 chars.",
			})
			return
		}
		if MD5(*ApiKey + mac + vk) != token {
			c.JSON(200, gin.H{
				"error": 103,
				"message": "No authority token.",
			})
			return
		}
	}
	err:=Wake(mac,ip,port)
	if err != nil {
		c.JSON(200, gin.H{
			"error": 100,
			"message": fmt.Sprintf("%s", err),
		})
	} else {
		c.JSON(200, gin.H{
			"error": 0,
			"message": fmt.Sprintf("Wake Success Mac:%s", mac),
		})
	}
}

