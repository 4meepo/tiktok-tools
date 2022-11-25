package ent

// var EntClient *ent.Client

import (
	"context"
	"fmt"
	"log"
	"os"
	"sync"

	"github.com/4meepo/tiktok-tools/ent/migrate"
	_ "github.com/go-sql-driver/mysql"
	"github.com/sirupsen/logrus"
)

var instance *Client
var once sync.Once

func GetInstance(user, password, host string) *Client {
	logrus.Printf("host: %s, user: %s, password: %s", host, user, password)
	once.Do(func() {
		client, err := Open("mysql", fmt.Sprintf("%s:%s@tcp(%s:3306)/tiktok?parseTime=True", user, password, host))
		if err != nil {
			log.Fatalf("failed opening connection to mysql: %v", err)
		}
		// defer client.Close()
		// Run the auto migration tool.
		client.Schema.WriteTo(context.Background(), os.Stdout,
			migrate.WithForeignKeys(false), // 不使用数据库外键
			migrate.WithDropIndex(true),    // 启用删除索引
			migrate.WithDropColumn(true),   // 启用删除列
		)
		if err := client.Schema.Create(context.Background(),
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
