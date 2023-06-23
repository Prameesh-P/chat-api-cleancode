package main

import (
	."log"

	cfg "github.com/Prameesh-P/chat-api-cleancode/pkg/config"
	"github.com/Prameesh-P/chat-api-cleancode/pkg/database"
)

var (
	lf=Fatalf
)

func Init(){

	

}

func main(){


	cfg.LoadConfig()
	database.ConnectDatabase()
	database.SyncDatabase()

}