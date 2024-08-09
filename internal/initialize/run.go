package initialize

import (
	"fmt"

	"github.com/hiumx/go-ecommerce-backend-api/global"
)

func Run() {
	LoadConfig()
	mySqlConfig := global.Config.Mysql
	fmt.Println("Load config from initialize", mySqlConfig.Host, mySqlConfig.Port)

	InitLogger()

	global.Logger.Info("Log success!")

	InitMySQL()
	InitRedis()
	r := InitRouter()

	r.Run(":8008")
}
