/**
* @Author: XGH
* @Email: 55821284@qq.com
* @Date: 2020/5/19 16:03
 */
package ctl

import (
	"github.com/gin-gonic/gin"
	"github.com/spf13/cast"
	"basestation/model"
	"log"
	"net/http"
	"time"
)

type Station struct {
}

//登录模块
func (this *Station) Router(router *gin.Engine) {
	r := router.Group("/station/")
	r.GET("list", this.list)
	r.GET("edit", this.edit)
	r.GET("map", this.maps)
	r.POST("savemap", this.saveMap)

	r.GET("getlist", this.getList)
}

func (this *Station) list(ctx *gin.Context) {
	ctx.HTML(http.StatusOK, "station_list.html", gin.H{
		"title": "基站列表",
	})
}
func (this *Station) edit(ctx *gin.Context) {
	ctx.HTML(http.StatusOK, "station_edit.html", gin.H{
		"title": "基站编辑",
	})
}

func (this *Station) maps(ctx *gin.Context) {
	//https://api.map.baidu.com/api?v=3.0&ak=GsQTZyGZqnGQnTwGGhGNgc40clKKwxxm
	//102.221499,31.908822
	info := &model.StStations{Id:1}
	info.GetOne()
	log.Printf("%+v\n",info)
	ctx.HTML(http.StatusOK, "station_map.html", gin.H{
		"title": "基站地图",
		"id": 1,
		"lat": info.Lat,
		"lng": info.Lng,
	})

	/*ctx.HTML(http.StatusOK, "station_map.html", gin.H{
		"title": "基站地图",
		"id": 1,
		"lat": 31.908822,
		"lng": 102.221499,
	})*/
}

func (this *Station) saveMap(ctx *gin.Context) {
	ctx.Request.ParseForm()

	var postParams = make(map[string]string)
	for k, v := range ctx.Request.PostForm {
		postParams[k] = v[0]
	}
	log.Printf("%+v\n",postParams)

	time.Sleep(1*time.Second)
	a := &model.StStations{Id:1}

	a.Lat = cast.ToFloat64(postParams["lat"])
	a.Lng = cast.ToFloat64(postParams["lng"])
	a.Update()
	ctx.JSON(http.StatusOK,gin.H{
		"result":"success",
	})
}
func (this *Station) getList(ctx *gin.Context)  {

}