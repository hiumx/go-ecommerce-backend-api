package initialize

import (
	"fmt"
	"time"

	"github.com/hiumx/go-ecommerce-backend-api/global"
	"github.com/hiumx/go-ecommerce-backend-api/internal/model"
	"go.uber.org/zap"
	"gorm.io/driver/mysql"
	"gorm.io/gen"
	"gorm.io/gorm"
)

func checkPanicError(err error, errorString string) {
	if err != nil {
		global.Logger.Error(errorString, zap.Error(err))
		panic(err)
	}
}

func InitMySQL() {
	m := global.Config.Mysql
	dsn := "%s:%s@tcp(%s:%v)/%s?charset=utf8mb4&parseTime=True&loc=Local"
	s := fmt.Sprintf(dsn, m.Username, m.Password, m.Host, m.Port, m.DbName)
	db, err := gorm.Open(mysql.Open(s), &gorm.Config{
		SkipDefaultTransaction: false,
	})
	checkPanicError(err, "Initialize MySQL failed!")
	global.Logger.Info("Initialize MySQL successfully")
	global.Mdb = db

	SetPool()
	migrateTable()
	generateTableDAO()
}

func SetPool() {
	m := global.Config.Mysql

	sqlDB, err := global.Mdb.DB()
	if err != nil {
		fmt.Printf("Mysql error: %s", err)
	}

	sqlDB.SetConnMaxIdleTime(time.Duration(m.MaxIdleConns))
	sqlDB.SetMaxOpenConns(m.MaxOpenConns)
	sqlDB.SetConnMaxLifetime(time.Duration(m.ConnMaxLifetime))

}

func generateTableDAO() {
	g := gen.NewGenerator(gen.Config{
		OutPath: "./internal/models",
		Mode:    gen.WithoutContext | gen.WithDefaultQuery | gen.WithQueryInterface, // generate mode
	})

	// gormdb, _ := gorm.Open(mysql.Open("root:@(127.0.0.1:3306)/demo?charset=utf8mb4&parseTime=True&loc=Local"))
	g.UseDB(global.Mdb) // reuse your gorm db

	// Generate basic type-safe DAO API for struct `model.User` following conventions
	//   g.ApplyBasic(model.User{})
	g.GenerateModel("go_crm_user")

	// Generate Type Safe API with Dynamic SQL defined on Querier interface for `model.User` and `model.Company`
	//   g.ApplyInterface(func(Querier){}, model.User{}, model.Company{})

	// Generate the code
	g.Execute()
}

func migrateTable() {
	err := global.Mdb.AutoMigrate(
		// &po.User{},
		// &po.Role{},
		&model.GoCrmUser{},
	)

	if err != nil {
		fmt.Println("Migration tables err: ", err)
	}
}
