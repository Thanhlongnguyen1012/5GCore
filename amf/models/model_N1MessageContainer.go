package models

type N1messageContainer struct {
	N1MessageClass   string            `json:"n1MessageClass"`
	N1MessageContent *N1messageContent `json:"N1MessageContent"`
	NfId             int32             `json:"nfId"`
}
