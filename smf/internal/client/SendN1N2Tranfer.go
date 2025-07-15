package client

import (
	"fmt"
	"smf/api"
	"smf/models"
)

func SendN1N2tranfer(data models.N1N2MessageTransferReqData) {
	resp, err := api.PostN1N2Tranfer(data)
	if err == nil && resp.StatusCode == 200 {
		//fmt.Println("send N1N2TranferReq OK !")
	} else {
		fmt.Println(err)
	}
}
