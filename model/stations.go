/**
* @Author: XGH
* @Email: 55821284@qq.com
* @Date: 2020/5/20 14:27
 */

package model

import (
	_ "github.com/go-sql-driver/mysql"
)

type StStations struct {
	Id  int64
	Lng float64
	Lat float64
}

func (this *StStations) GetOne() {
	_,_ =Engine.Get(this)
}

func (this *StStations) Update() {
	_, err := Engine.ID(this.Id).Cols("lng,lat").Update(this)
	if err != nil {
	}
}