package connection

import(
	"pt9/models"
	"log"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func ConnectToDb() *gorm.DB{
	dsn := "root:@tcp(127.0.0.1:3306)/dagang"

	db, err :=gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil{
		log.Fatalln(err)
	}
	db.AutoMigrate(&models.Kategori{},&models.Barang{},&models.Jual{})

	return db
}