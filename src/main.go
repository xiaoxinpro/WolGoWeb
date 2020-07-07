package main

import (
	"flag"
	"fmt"
	"github.com/gin-gonic/gin"
)

var (
	WebPort = flag.Int("port", 9090, "wol web port.")
)

func main()  {
	flag.Parse()

	gin.SetMode(gin.ReleaseMode)

	r:=gin.Default()

	r.GET("/", GetIndex)
	r.GET("/index", GetIndex)
	r.GET("/wol", GetWol)

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
