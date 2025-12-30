package main

import (
	"fmt"
	"lelForum/controller"
	"lelForum/database/postgres"
	"lelForum/database/redis"
	"lelForum/logger"
	"lelForum/pkg/snowflake"
	"lelForum/routers"
	"lelForum/settings"
)

func main() {
	// Load configuration
	if err := settings.Init(); err != nil {
		fmt.Printf("load config failed, err:%v\n", err)
		return
	}

	// Get current configuration
	cfg := settings.GetConf()

	if err := logger.Init(cfg.LogConfig, cfg.Mode); err != nil {
		fmt.Printf("init logger failed, err:%v\n", err)
		return
	}
	if err := postgres.Init(cfg.PostgreSQLConfig); err != nil {
		fmt.Printf("init postgres failed, err:%v\n", err)
		return
	}
	defer postgres.Close() // Close the database connection when the program exits

	if err := redis.Init(cfg.RedisConfig); err != nil {
		fmt.Printf("init redis failed, err:%v\n", err)
		return
	}
	defer redis.Close() // Close the Redis connection when the program exits

	if err := snowflake.Init(1); err != nil {
		fmt.Printf("init snowflake failed, err:%v\n", err)
		return
	}
	// Initialize the custom parameter validator to support internationalization
	if err := controller.InitTrans("en"); err != nil {
		fmt.Printf("init translator failed, err:%v\n", err)
		return
	}
	// Register routes
	r := routers.SetupRoutes(settings.GetConf().Mode)
	err := r.Run(fmt.Sprintf(":%d", settings.GetConf().Port))
	if err != nil {
		fmt.Printf("run server failed, err:%v\n", err)
		return
	}
}
