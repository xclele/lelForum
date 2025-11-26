package main

import (
	"fmt"
	"lelForum/controller"
	"lelForum/database/postgres"

	//"lelForum/database/postgres"
	//"lelForum/database/redis"
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
	if err := logger.Init(settings.Conf.LogConfig, settings.Conf.Mode); err != nil {
		fmt.Printf("init logger failed, err:%v\n", err)
		return
	}
	if err := postgres.Init(settings.Conf.PostgreSQLConfig); err != nil {
		fmt.Printf("init postgres failed, err:%v\n", err)
		return
	}
	defer postgres.Close() // Close the database connection when the program exits

	/* Redis
	if err := redis.Init(settings.Conf.RedisConfig); err != nil {
		fmt.Printf("init redis failed, err:%v\n", err)
		return
	}
	defer redis.Close() // Close the Redis connection when the program exits	*/
	if err := snowflake.Init(1); err != nil {
		fmt.Printf("init snowflake failed, err:%v\n", err)
		return
	}
	// Initialize the custom parameter validator to support internationalization
	if err := controller.InitTrans("zh"); err != nil {
		fmt.Printf("init translator failed, err:%v\n", err)
		return
	}
	// Register routes
	r := routers.SetupRoutes()
	err := r.Run(fmt.Sprintf(":%d", settings.Conf.Port))
	if err != nil {
		fmt.Printf("run server failed, err:%v\n", err)
		return
	}
}
