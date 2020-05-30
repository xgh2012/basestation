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

type Menu struct {
}

//登录模块
func (this *Menu) Router(router *gin.Engine) {
	router.GET("/menu/init", this.GetMenu)
}

type child struct {
	Title string `json:"title"`
	Href string `json:"href"`
	Icon string `json:"icon"`
	Target string `json:"target"`
	Child []child `json:"child"`
}
type menuInfo struct {
	Title string `json:"title"`
	Icon string `json:"icon"`
	Href string `json:"href"`
	Target string `json:"target"`
	Child []child `json:"child"`

}
type Init struct {
	HomeInfo struct{
		Title string `json:"title"`
		Href string `json:"href"`
	} `json:"homeInfo"`
	LogoInfo struct{
		Title string `json:"title"`
		Image string `json:"image"`
	} `json:"logoInfo"`
	MenuInfo []menuInfo `json:"menuInfo"`
}

func (this *Menu) GetMenu(ctx *gin.Context) {
	a := &Init{
		HomeInfo: struct {
			Title string `json:"title"`
			Href  string `json:"href"`
		}{"首页","page/welcome-1.html?t=1"},
		LogoInfo: struct {
			Title string `json:"title"`
			Image string `json:"image"`
		}{"LAYUI MINI","images/logo.png"},
	}

	chdChd := child{
		Title: "站点列表",
		Href: "/station/list",
		Icon: "fa fa-tachometer",
		Target: "_self",
	}

	chd := child{
		Title:  "主页模板",
		Href:   "",
		Icon:   "fa fa-home",
		Target: "_self",
	}
	chd.Child=append(chd.Child,chdChd)

	chd.Child=append(chd.Child,child{
		Title: "主页二",
		Href: "page/welcome-2.html",
		Icon: "fa fa-tachometer",
		Target: "_self",
	})

	mif:=menuInfo{
		Title:  "常规管理",
		Icon:   "fa fa-address-book",
		Href:   "",
		Target: "_self",
	}
	mif.Child = append(mif.Child,chd)

	a.MenuInfo = append(a.MenuInfo,mif)

	ctx.JSON(http.StatusOK,a)
}