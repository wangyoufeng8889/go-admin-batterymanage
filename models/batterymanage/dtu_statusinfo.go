package batterymanage

import (
	orm "go-admin/global"
	"go-admin/models"
	"go-admin/tools"
	"time"
)
type Dtu_statusinfo struct {
	Dtu_statusInfoId     int    `json:"dtu_statusInfoId" gorm:"size:10;primary_key;AUTO_INCREMENT"`
	Dtu_uptime time.Time  `json:"dtu_uptime"`
	Pkg_id   string `json:"pkg_id" gorm:"size:20;"`
	Dtu_id      string `json:"dtu_id" gorm:"size:20;"`
	Dtu_latitudeType   string    `json:"dtu_latitudeType" gorm:"Type：size:2"`
	Dtu_longitudeType   string    `json:"dtu_longitudeType" gorm:"Type：size:2"`
	Dtu_latitude   string    `json:"dtu_latitude" gorm:"Type：size:20"`
	Dtu_longitude   string    `json:"dtu_longitude" gorm:"Type：size:20"`
	Dtu_csq   uint8    `json:"dtu_csq" gorm:"Type：uint8"`
	Dtu_locateMode   uint8    `json:"dtu_locateMode" gorm:"Type：uint8"`
	Dtu_gpsSateCnt   uint8    `json:"dtu_gpsSateCnt" gorm:"Type：uint8"`
	Dtu_speed   uint16    `json:"dtu_speed" gorm:"Type：uint16"`
	Dtu_altitude   uint16    `json:"dtu_altitude" gorm:"Type：uint16"`
	Dtu_pluginVoltage   uint8    `json:"dtu_pluginVoltage" gorm:"Type：uint8"`
	Dtu_selfInVoltage   uint8    `json:"dtu_selfInVoltage" gorm:"Type：uint8"`
	Dtu_errStatus   uint8    `json:"dtu_errStatus" gorm:"Type：uint8"`
	Dtu_errNbr   uint8    `json:"dtu_errNbr" gorm:"Type：uint8"`
	Dtu_errCode   uint16    `json:"dtu_errCode" gorm:"Type：uint16"`
	DataScope  string `json:"dataScope" gorm:"-"`
	models.BaseModel
}
func (Dtu_statusinfo) TableName() string {
	return "user_dtu_statusinfo"
}
func (e *Dtu_statusinfo) Getdtu_statusinfo(startdate time.Time, enddate time.Time) ([]Dtu_statusinfo,int, error) {
	var doc []Dtu_statusinfo

	table := orm.Eloquent.Select("*").Table(e.TableName())
	if e.Dtu_statusInfoId != 0 {
		table = table.Where("dtu_statusInfoId = ?", e.Dtu_statusInfoId)
	}
	if e.Pkg_id != "" {
		table = table.Where("pkg_id = ?", e.Pkg_id)
	}else {
		table = table.Not("pkg_id = ?", "0")
	}
	if e.Dtu_id != "" {
		table = table.Where("dtu_id = ?", e.Dtu_id)
	}
	table = table.Where("dtu_uptime BETWEEN ? AND ?",startdate,enddate)
	// 数据权限控制
	dataPermission := new(models.DataPermission)
	dataPermission.UserId, _ = tools.StringToInt(e.DataScope)
	table, err := dataPermission.GetDataScope(e.TableName(), table)
	if err != nil {
		return nil, 0, err
	}
	var count int

	if err := table.Order("dtu_uptime").Find(&doc).Error; err != nil {
		return nil, 0, err
	}
	table.Where("`deleted_at` IS NULL").Count(&count)
	return doc, count, nil
}