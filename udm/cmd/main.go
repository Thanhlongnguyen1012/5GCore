package main

import (
	"database/sql/driver"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"gorm.io/gorm/logger"

	"github.com/gin-gonic/gin"
	"golang.org/x/net/http2"
	"golang.org/x/net/http2/h2c"
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

var (
	dsn = os.Getenv("MYSQL_DSN")
	db  *gorm.DB
)

func main() {
	// connect MySQL
	//dsn := "root:my-secret-pw@tcp(udm-mysql-core:3306)/udm?charset=utf8mb4&parseTime=True&loc=Local"
	//dsn := "root:my-secret-pw@tcp(127.0.0.1:3307)/udm?charset=utf8mb4&parseTime=True&loc=Local"
	// dsn := os.Getenv("MYSQL_DSN")
	// if dsn == "" {
	// 	fmt.Println("MYSQL_DSN environment variable not set")
	// 	return
	// }
	var err error
	db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: logger.New(
			log.New(os.Stdout, "\r\n", log.LstdFlags),
			logger.Config{
				SlowThreshold: time.Millisecond * 200,
				LogLevel:      logger.Info,
				Colorful:      true,
			},
		),
	})
	if err != nil {
		fmt.Println("Failed to connect to database: ", err)
		return
	}
	sqlDB, _ := db.DB()
	sqlDB.SetMaxOpenConns(32)
	sqlDB.SetMaxIdleConns(32)
	sqlDB.SetConnMaxLifetime(10 * time.Minute)
	// Start server
	//r := gin.Default()
	//r := gin.New()
	r := gin.Default()
	// r.GET("/nudm-sdm/v2/imsi-452040989692072/sm-data", func(c *gin.Context) {
	// 	var data SmContextCreateData
	// 	db.First(&data, "supi = ? ", "imsi-452040989692072")
	// 	c.JSON(http.StatusOK, &data)
	// })
	/*r.Use(func(c *gin.Context) {
		fmt.Println("Protocol:", c.Request.Proto) // <-- In ra HTTP/1.1 hoặc HTTP/2.0
		c.Next()
	})*/
	r.GET("/nudm-sdm/v2/imsi-452040989692072/sm-data", handlerSQL(db))
	//udmUrl := os.Getenv("UDM_URL")
	//r.Run("0.0.0.0:8082")
	h2s := &http2.Server{}
	server := &http.Server{
		Addr:    ":8082",
		Handler: h2c.NewHandler(r, h2s),
	}
	error := server.ListenAndServe()
	//error := r.RunTLS(":8082", "cert.pem", "key.pem")
	if error != nil {
		fmt.Println("Failed to run TLS Server ")
	}
}
func handlerSQL(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var data SmContextCreateData
		db.Raw("SELECT * FROM udm WHERE supi = ?", "imsi-452040989692072").Scan(&data)
		//db.First(&data, "supi = ? ", "imsi-452040989692072")
		c.JSON(http.StatusOK, &data)
	}
}
