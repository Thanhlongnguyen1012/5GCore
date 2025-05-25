package models

type N1N2MessageTransferReqData struct {
	//N1MessageContainer *N1messageContainer `json:"n1MessageContainer"`
	//N2InfoContainer    *N2infoContainer    `json:"n2InfoContainer"`
	PduSessionId int32 `json:"pduSessionId"`
}
