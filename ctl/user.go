/**
* @Author: XGH
* @Email: 55821284@qq.com
* @Date: 2020/5/19 16:03
 */
package ctl

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type User struct {
}

//用户模块
func (this *User) Router(router *gin.Engine) {
	r := router.Group("/user/")
	r.GET("setting", this.setting)
	r.GET("password", this.password)
}

func (this *User) setting(ctx *gin.Context) {
	ctx.HTML(http.StatusOK, "user_setting.html", gin.H{
		"title": "用户资料",
	})
}

func (this *User) password(ctx *gin.Context) {
	ctx.HTML(http.StatusOK, "user_password.html", gin.H{
		"title": "修改密码",
	})
}