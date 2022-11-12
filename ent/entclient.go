package ent

// var EntClient *ent.Client

import (
	"context"
	"log"
	"sync"

	"github.com/4meepo/tiktok-tools/ent/migrate"
	_ "github.com/go-sql-driver/mysql"
)

var instance *Client
var once sync.Once

func GetInstance() *Client {
	once.Do(func() {
		client, err := Open("mysql", "root:pass@tcp(ecs:3306)/tiktok?parseTime=True")
		if err != nil {
			log.Fatalf("failed opening connection to mysql: %v", err)
		}
		// defer client.Close()
		// Run the auto migration tool.
		if err := client.Debug().Schema.Create(context.Background(),
			migrate.WithForeignKeys(false), // 不使用数据库外键
			migrate.WithDropIndex(true),    // 启用删除索引
			migrate.WithDropColumn(true),   // 启用删除列
		); err != nil {
			log.Fatalf("failed creating schema resources: %v", err)
		}
		log.Printf("数据库初始化成功...")
		instance = client
	})
	return instance
}
