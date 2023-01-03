package config

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/arwansa/echo-ent/ent"
	_ "github.com/go-sql-driver/mysql"
)

var (
	client *ent.Client
)

func GetClient() *ent.Client {
	return client
}

func SetClient(newClient *ent.Client) {
	client = newClient
}

func NewEntClient() (*ent.Client, error) {
	cfg := Get()
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?parseTime=true",
		cfg.DB.User, cfg.DB.Pass, cfg.DB.Host, cfg.DB.Port, cfg.DB.Name,
	)

	client, err := ent.Open("mysql", dsn, ent.Log(func(i ...interface{}) {
		for _, v := range i {
			fmt.Println(time.Now().Format("2006-01-02 15:04:05"), v)
			fmt.Print("\n")
		}
	}))

	if err != nil {
		log.Fatalf("failed opening connection to mysql: %v", err)
	}

	// Run the auto migration tool.
	if err := client.Schema.Create(context.Background()); err != nil {
		log.Fatalf("failed creating schema resources: %v", err)
	}

	return client, err
}
