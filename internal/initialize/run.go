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
	InitMySQL()
	InitRedis()

	r := InitRouter()

	r.Run(":8008")
}
