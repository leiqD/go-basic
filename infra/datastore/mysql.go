package datastore

import (
	"fmt"
	"github.com/leiqD/go-socket5/infra/conf"
	"gorm.io/gorm/logger"
	"log"
	"os"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

const (
	// DSN data source name
	// DSN := `gorm:gorm@tcp(127.0.0.1:3306)/gorm?charset=utf8&parseTime=True&loc=Local`
	DSNFormat string = `%s:%s@%s(%s)/%s?charset=utf8mb4&parseTime=True&loc=Local`
)

type MYSQLconfig interface {
	DataBase() *conf.DataStoreInfo
}

func NewMYSQLDB(cfg MYSQLconfig) (*gorm.DB, error) {
	fmt.Println(cfg.DataBase())
	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
		logger.Config{
			SlowThreshold: time.Second, // 慢 SQL 阈值
			LogLevel:      logger.Warn, // Log level
			Colorful:      false,       // 禁用彩色打印
		},
	)
	db, err := gorm.Open(
		mysql.New(mysql.Config{
			DSN: fmt.Sprintf(DSNFormat, cfg.DataBase().User, cfg.DataBase().Paswd, cfg.DataBase().Net, cfg.DataBase().Host, cfg.DataBase().Db),
		}),
		&gorm.Config{
			NamingStrategy: schema.NamingStrategy{
				SingularTable: true, // 使用单数表名，启用该选项，此时，`User` 的表名应该是 `User`
			},
			Logger: newLogger,
		})

	if err != nil {
		fmt.Printf("gorm open datastore failed: %s", err)
		return nil, err
	}

	sqlDB, err := db.DB()
	if err != nil {
		fmt.Printf("gorm get datastore failed: %s", err)
		return nil, err
	}
	// SetMaxIdleConns 设置空闲连接池中连接的最大数量
	sqlDB.SetMaxIdleConns(cfg.DataBase().Params.MaxIdleConns)

	// SetMaxOpenConns 设置打开数据库连接的最大数量
	sqlDB.SetMaxOpenConns(cfg.DataBase().Params.MaxOpenConns)

	// SetConnMaxLifetime 设置了连接可复用的最大时间
	sqlDB.SetConnMaxLifetime(time.Hour * time.Duration(cfg.DataBase().Params.ConnMaxLifetime))

	return db, nil
}
