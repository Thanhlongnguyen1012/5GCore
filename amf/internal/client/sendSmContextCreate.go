package client

import (
	"amf/api"
	"amf/models"
	"fmt"
)

func SendSmContextCreate(data models.SMContextCreateData) {
	resp, err := api.PostSmCreate(data)
	if err == nil && resp.StatusCode == 201 {
		fmt.Println("send SmContextcreate ok !")
	} else {
		fmt.Println(err)
	}
}
