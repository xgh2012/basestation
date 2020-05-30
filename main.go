/**
* @Author: XGH
* @Email: 55821284@qq.com
* @Date: 2020/5/19 15:51
 */
package main

import (
	"basestation/ctl"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

func main() {
	gin.SetMode(gin.DebugMode)
	router := gin.Default()

	router.LoadHTMLGlob("web/view/**/**/*")
	router.StaticFS("/public", http.Dir("web/static"))

	RegistRouter(router)
	err := router.Run(":9527")
	log.Println(err)
}

//注册
func RegistRouter(router *gin.Engine) {
	new(ctl.Login).Router(router)
	new(ctl.Index).Router(router)
	new(ctl.Menu).Router(router)
	new(ctl.Station).Router(router)
	new(ctl.User).Router(router)
	new(ctl.Front).Router(router)
}
