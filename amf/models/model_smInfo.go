package models

type Sminfo struct {
	PduSessionId  int32          `json:"pduSessionId"`
	N2InfoContent *N2infoContent `json:"n2InfoContent"`
	SNssai        *Snssai        `json:"sNssai"`
}
