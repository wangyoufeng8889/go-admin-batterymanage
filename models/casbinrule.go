package models

//casbin_rule
type CasbinRule struct {
	Id          int    `json:"Id" gorm:"primary_key;AUTO_INCREMENT"` // 编码
	PType string `json:"p_type" gorm:"size:100;"`
	V0    string `json:"v0" gorm:"size:100;"`
	V1    string `json:"v1" gorm:"size:100;"`
	V2    string `json:"v2" gorm:"size:100;"`
	V3    string `json:"v3" gorm:"size:100;"`
	V4    string `json:"v4" gorm:"size:100;"`
	V5    string `json:"v5" gorm:"size:100;"`
}

func (CasbinRule) TableName() string {
	return "casbin_rule"
}
