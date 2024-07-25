package db

import (
	"context"
	"fmt"

	_ "github.com/lib/pq"
	"pentag.kr/distimer/configs"
	"pentag.kr/distimer/ent"
	"pentag.kr/distimer/utils/logger"
)

var (
	dbClient *ent.Client
)

func GetDBClient() *ent.Client {
	return dbClient
}

func ConnectDBClient() {

	dsn := fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=disable",
		configs.Env.DBHost, configs.Env.DBPort, configs.Env.DBUser, configs.Env.DBName, configs.Env.DBPass)

	client, err := ent.Open("postgres", dsn)
	if err != nil {
		logger.Fatal(err)
	}
	if err := client.Schema.Create(context.Background()); err != nil {
		logger.Fatal(err)
	}
	dbClient = client
}
