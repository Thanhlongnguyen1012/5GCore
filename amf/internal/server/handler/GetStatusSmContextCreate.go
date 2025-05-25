package handler

import (
	"amf/api"
	"amf/models"

	"github.com/gin-gonic/gin"
)

var request = models.SMContextCreateData{
	Supi:         "imsi-452040916843227",
	Gpsi:         "msisdn-84867220452",
	PduSessionId: 5,
	Dnn:          "v-internet",
	SNssai: &models.Snssai{
		Sst: 1,
		Sd:  "000001",
	},
	ServingNfId: "",
	AnType:      "3GPP_ACCESS",
}

func GetStatusSmContextCreate(c *gin.Context) {
	status, response := api.GetSmContextCreate(request)

	c.JSON(status, response)
}
