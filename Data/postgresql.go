package Data

import (
	"XcessAlipay/Config"
	"fmt"
	"github.com/golang/glog"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
	"strings"
	"time"
)

var _db *gorm.DB
var Loc, _ = time.LoadLocation("Asia/Shanghai")

//包初始化函数，golang特性，每个包初始化的时候会自动执行init函数，这里用来初始化gorm。
func init() {

	var err error
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d sslmode=disable TimeZone=Asia/Shanghai", ServerConf.PGDBConf.Address, ServerConf.PGDBConf.User, ServerConf.PGDBConf.Password, ServerConf.PGDBConf.DatabaseName, ServerConf.PGDBConf.Port)
	_db, err = gorm.Open(postgres.New(postgres.Config{
		DSN:                  dsn,
		PreferSimpleProtocol: true, // disables implicit prepared statement usage
	}), &gorm.Config{
		PrepareStmt: true, //
		NamingStrategy: schema.NamingStrategy{
			TablePrefix:   "z_",                              // 表名前缀，`User`表为`t_users`
			SingularTable: true,                              // 使用单数表名，启用该选项后，`User` 表将是`user`
			NameReplacer:  strings.NewReplacer("CID", "Cid"), // 在转为数据库名称之前，使用NameReplacer更改结构/字段名称。

		},
		NowFunc: func() time.Time {
			return time.Now().In(Loc).Local()
		},
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		glog.Fatalf("client error is %v", err)
	} else {
		sqlDB, err := _db.DB()
		if err != nil {
			glog.Fatalf("error is %v", err)
		}
		// fmt.Println("PostgresqlDB", sqlDB)

		// SetMaxIdleConns 设置空闲连接池中连接的最大数量
		sqlDB.SetMaxIdleConns(20)

		// SetMaxOpenConns 设置打开数据库连接的最大数量。
		sqlDB.SetMaxOpenConns(10000)

		// SetConnMaxLifetime 设置了连接可复用的最大时间。
		sqlDB.SetConnMaxLifetime(time.Hour)
	}

}

//获取gorm db对象，其他包需要执行数据库查询的时候，只要通过tools.getDB()获取db对象即可。
//不用担心协程并发使用同样的db对象会共用同一个连接，db对象在调用他的方法的时候会从数据库连接池中获取新的连接

func GetDB() *gorm.DB {
	return _db
}
