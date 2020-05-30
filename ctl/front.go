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

type Front struct {
}

//登录模块
func (this *Front) Router(router *gin.Engine) {
	router.GET("/f/index", this.index)
	router.GET("/f/view", this.view)
	router.GET("/f/list/:id", this.list)
}

func (this *Front) index(ctx *gin.Context) {
	ctx.HTML(http.StatusOK, "index_index.html", gin.H{
		"title": "前端",
	})
}


func (this *Front) list(ctx *gin.Context) {
	cid := ctx.Param("id")
	ctx.JSON(http.StatusOK, gin.H{
		"title": "前端",
		"cid": cid,
	})
}

func (this *Front) view(ctx *gin.Context) {
	ctx.HTML(http.StatusOK, "index_view.html", gin.H{
		"title": "前端",
	})
}