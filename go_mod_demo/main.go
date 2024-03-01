// main.go

package main

import (
	"context"
	"fmt"
	"github.com/go-redis/redis/v8"
	"go_mod_demo/pkg1"
	"go_mod_demo/pkg2"
)

func main() {
	pkg1.Func1()
	pkg2.Func2()

	// 创建实例
	client := redis.NewClient(&redis.Options{
		Addr:     "127.0.0.1:25001", // Redis 服务器地址
		Password: "",                // Redis 服务器密码
		DB:       0,                 // Redis 数据库索引
	})

	err := client.Set(context.Background(), "mykey", "myvalue", 0).Err()
	if err != nil {
		panic(err)
	}

	val, err := client.Get(context.Background(), "mykey").Result()
	if err != nil {
		panic(err)
	}
	fmt.Println("mykey:", val)
}
