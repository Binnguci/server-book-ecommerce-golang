package initialize

import (
	"server-furniture-ecommerce-gin/global"
	"strconv"
)

func Run() {
	LoadConfig()
	InitLogger()
	InitMySQL()
	InitRedis()
	InitRouter()
	//InitElasticsearch()

	r := InitRouter()
	port := global.Config.Server.Port
	r.Run(":" + strconv.Itoa(port))
}
