package models

type SMContextCreatedData struct {
	PduSessionID int32   `json:"pduSessionId"`
	SNssai       *Snssai `json:"sNssai,omitempty"`
}
