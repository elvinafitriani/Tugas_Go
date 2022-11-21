package connection

import (
	"fmt"
	"relation/model"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var Db *gorm.DB

func Connect() *gorm.DB {
	var err error
	db := "root:@tcp(127.0.0.1:3306)/test"

	Db, err = gorm.Open(mysql.Open(db), &gorm.Config{})
	if err != nil {
		fmt.Println("Can't Connect to Database")
		return nil
	}
	Db.AutoMigrate(&model.Kategori{}, &model.Barang{}, &model.Jual{})

	return Db
}
