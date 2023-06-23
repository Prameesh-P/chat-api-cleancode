package database

import (
	"os"
	."log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)
var (
	DB *gorm.DB
	err error
	lf=Fatalf
)
func ConnectDatabase(){
	dsn:=os.Getenv("DSN")
	DB,err=gorm.Open(postgres.Open(dsn),&gorm.Config{})
	if err !=nil{
		lf("can't connect with database, %v",err)
		return
	}
	Println("Database connection successfully..")
}