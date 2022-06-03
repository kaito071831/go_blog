package utility

import (
	"fmt"
	"os"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var Db *gorm.DB

func init() {
	user := os.Getenv("MYSQL_USER")
	pw := os.Getenv("MYSQL_PASSWORD")
	db_name := os.Getenv("MYSQL_DATABASE")
	dsn := fmt.Sprintf("%s:%s@tcp(db:3306)/%s?charset=utf8&parseTime=true", user, pw, db_name)
	dialector := mysql.Open(dsn)
	var err error
	if Db, err = gorm.Open(dialector); err != nil {
		dbConnect(dialector, 100)
	}
	fmt.Println("DB Connect!")
}

func dbConnect(dialector gorm.Dialector, count uint) {
	var err error
	if Db, err = gorm.Open(dialector); err != nil {
		if count > 1 {
			time.Sleep(time.Second * 2)
			count--
			fmt.Printf("Retry... count:%v\n", count)
			dbConnect(dialector, count)
			return
		}
	}
}