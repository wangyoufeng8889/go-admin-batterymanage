package batterymanage
import (
	"go-admin/models"
	"time"
)
type Bms_temperatureinfo struct {
	Bms_temperatureinfoId     int    `json:"bms_temperatureinfoId" gorm:"size:10;primary_key;AUTO_INCREMENT"`
	Dtu_uptime time.Time  `json:"dtu_uptime"`
	Dtu_id      string `json:"dtu_id" gorm:"size:20;"`
	Pkg_id   string `json:"pkg_id" gorm:"size:20;"`
	Bms_temperature1   uint8    `json:"bms_temperature1" gorm:"Type：uint8"`
	Bms_temperature2   uint8    `json:"bms_temperature2" gorm:"Type：uint8"`
	Bms_temperature3   uint8    `json:"bms_temperature3" gorm:"Type：uint8"`
	Bms_temperature4   uint8    `json:"bms_temperature4" gorm:"Type：uint8"`
	Bms_temperature5   uint8    `json:"bms_temperature5" gorm:"Type：uint8"`
	Bms_temperature6   uint8    `json:"bms_temperature6" gorm:"Type：uint8"`
	models.BaseModel
}
func (Bms_temperatureinfo) TableName() string {
	return "user_bms_temperatureinfo"
}