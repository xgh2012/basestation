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

type Index struct {
}

//登录模块
func (this *Index) Router(router *gin.Engine) {
	router.GET("/", this.Index)
	router.GET("index", this.Index)
}

func (this *Index) Index(ctx *gin.Context) {
	ctx.HTML(http.StatusOK, "layout_index.html", gin.H{
		"title": "管理后台",
	})
}