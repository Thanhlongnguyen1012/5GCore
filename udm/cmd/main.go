package main

import (
	"database/sql/driver"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type SmContextCreateData struct {
	Supi         string  `json:"supi,omitempty" gorm:"column:supi"`
	Gpsi         string  `json:"gpsi,omitempty" gorm:"column:gpsi"`
	PduSessionId int32   `json:"pduSessionId,omitempty" gorm:"column:pduSessionId"`
	Dnn          string  `json:"dnn,omitempty" gorm:"column:dnn"`
	SNssai       *Snssai `json:"sNssai,omitempty" gorm:"column:sNssai"`
	ServingNfId  string  `json:"servingNfId" gorm:"column: servingNfId"`
	AnType       string  `json:"anType" gorm:"column:anType"`
	//UeLocation     *UserLocation `json:"ueLocation,omitempty"`
}

type Snssai struct {
	Sst int32  `json:"sst"`
	Sd  string `json:"sd,omitempty"`
}

func (s Snssai) Value() (driver.Value, error) {
	return json.Marshal(s)
}
func (s *Snssai) Scan(value interface{}) error {
	bytes, ok := value.([]byte)
	if !ok {
		return errors.New("type assertion to []byte failed")
	}
	return json.Unmarshal(bytes, s)
}
func (SmContextCreateData) TableName() string {
	return "udm"
}
func main() {
	// connect MySQL
	//dsn := "root:my-secret-pw@tcp(udm-mysql-core:3306)/udm?charset=utf8mb4&parseTime=True&loc=Local"
	dsn := "root:my-secret-pw@tcp(127.0.0.1:3307)/udm?charset=utf8mb4&parseTime=True&loc=Local"
	// dsn := os.Getenv("MYSQL_DSN")
	// if dsn == "" {
	// 	fmt.Println("MYSQL_DSN environment variable not set")
	// 	return
	// }

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Println("Failed to connect to database: ", err)
		return
	}
	// start server
	r := gin.Default()

	r.GET("/nudm-sdm/v2/imsi-452040989692072/sm-data", func(c *gin.Context) {
		var data SmContextCreateData
		db.First(&data, "supi = ? ", "imsi-452040989692072")
		c.JSON(http.StatusOK, &data)
	})

	r.Run(":8082")
}
