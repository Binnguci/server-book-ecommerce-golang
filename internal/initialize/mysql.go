package initialize

import (
	"fmt"
	"go.uber.org/zap"
	"gorm.io/driver/mysql"
	"gorm.io/gen"
	"gorm.io/gorm"
	"server-furniture-ecommerce-gin/global"
	"time"
)

func checkErrorPanic(err error, errStr string) {
	if err != nil {
		global.Logger.Error(errStr, zap.Error(err))
		panic(fmt.Sprintf("failed to connect database, err: %v", err))
	}
}

func InitMySQL() {
	m := global.Config.MySQL
	dsn := "%s:%s@tcp(%s:%v)/%s?charset=utf8mb4&parseTime=True&loc=Local"
	var s = fmt.Sprintf(dsn, m.Username, m.Password, m.Host, m.Port, m.Database)
	db, err := gorm.Open(mysql.Open(s), &gorm.Config{})

	checkErrorPanic(err, "gorm.Open failed")
	global.Mdb = db
	//genEntity()
	SetPool()
}

func SetPool() {
	m := global.Config.MySQL
	mysql, err := global.Mdb.DB()
	if err != nil {
		global.Logger.Error("SetPool failed", zap.Error(err))
	}
	mysql.SetMaxIdleConns(int(time.Duration(m.MaxIdleConns)))
	mysql.SetConnMaxLifetime(time.Duration(m.ConnMaxLifetime))
	mysql.SetMaxOpenConns(int(time.Duration(m.MaxOpenConns)))

}

func genEntity() {
	g := gen.NewGenerator(gen.Config{
		OutPath: "./internal/model",
		Mode:    gen.WithoutContext | gen.WithDefaultQuery | gen.WithQueryInterface,
	})
	g.UseDB(global.Mdb)
	g.GenerateAllTable()
	g.Execute()
}
