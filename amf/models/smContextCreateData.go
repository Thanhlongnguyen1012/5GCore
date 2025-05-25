package models

type SMContextCreateData struct {
	Supi         string  `json:"supi,omitempty" gorm:"column:supi"`
	Gpsi         string  `json:"gpsi,omitempty" gorm:"column:gpsi"`
	PduSessionId int32   `json:"pduSessionId,omitempty" gorm:"column:pduSessionId"`
	Dnn          string  `json:"dnn,omitempty" gorm:"column:dnn"`
	SNssai       *Snssai `json:"sNssai,omitempty" gorm:"column:sNssai"`
	ServingNfId  string  `json:"servingNfId" gorm:"column: servingNfId"`
	AnType       string  `json:"anType" gorm:"column:anType"`
	//UeLocation     *UserLocation `json:"ueLocation,omitempty"`
}
