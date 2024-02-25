package check

import (
	"Gol/model"
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func CheckImportFunction(x int) {
	fmt.Println("received variable x: ", x)
}

func InitDb(){
	dsn := "user=abc password=password dbname=mydb sslmode=disable TimeZone=Asia/Kolkata"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	db.AutoMigrate(&model.DummyModel{})
	if err != nil {
		panic("Failed to connect to the database")
	}
	InsertData()
}