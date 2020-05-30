/**
* @Author: XGH
* @Email: 55821284@qq.com
* @Date: 2020/5/19 16:03
 */
package ctl

import (
	"basestation/utils"
	"github.com/gin-gonic/gin"
	"net/http"
)

type Login struct {
}

//登录模块
func (this *Login) Router(router *gin.Engine) {
	r := router.Group("/login/")
	r.GET("index", this.Index)
	r.POST("do", this.DoLogin)
	r.GET("captcha", this.captcha)
	r.GET("getcaptcha/:source", this.getcaptcha)
}

func (this *Login) Index(ctx *gin.Context) {
	ctx.HTML(http.StatusOK, "login_index.html", gin.H{
		"title": "APP 加解密工具",
	})
}

func (this *Login) DoLogin(ctx *gin.Context) {

}

//获取验证码
func (this *Login) captcha(ctx *gin.Context) {
	utils.GetCaptcha(ctx)
}

//获取验证码
func (this *Login) getcaptcha(ctx *gin.Context) {
	utils.GetCaptchaPng(ctx)
}
