package check

import (
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)
type DummyModel struct {
	Id uint
	UId  int `gorm:"primaryKey"`
	Name string
	Price int
}

func CheckImportFunction(x int) {
	fmt.Println("received variable x: ", x)
}
func InitDb(){
	fmt.Print("adasdsd")
	dsn := "user=abc password=password dbname=mydb sslmode=disable TimeZone=Asia/Kolkata"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	db.AutoMigrate(&DummyModel{})
	if err != nil {
		panic("Failed to connect to the database")
	}
}
