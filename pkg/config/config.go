package config

import (
	"github.com/joho/godotenv"
	."log"
)

var (
	lf=Fatalf
)

func LoadConfig(){
	if err :=godotenv.Load(".env");err!=nil{
		lf("can't load env file.. %v",err)
	}
}
