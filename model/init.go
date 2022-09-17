package model

import (
	"fmt"
	"time"

	"github.com/AH-dark/bing-wallpaper/pkg/conf"
	"github.com/AH-dark/bing-wallpaper/pkg/util"
	"github.com/AH-dark/logger"
	"github.com/AH-dark/logger/gorm_logger"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/driver/postgres"
	"gorm.io/driver/sqlite"
	"gorm.io/driver/sqlserver"
	"gorm.io/gorm"
	gormlogger "gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
)

var DB *gorm.DB

func Init() {
	logger.Log().Info("Init database connection...")

	var dialector gorm.Dialector

	if gin.Mode() == gin.TestMode {
		dialector = sqlite.Open("file::memory:?cache=shared")
	} else {
		switch conf.DatabaseConfig.Type {
		case "sqlite", "sqlite3":
			dialector = sqlite.Open(util.AbsolutePath(conf.DatabaseConfig.DBFile))
		case "mysql":
			dialector = mysql.Open(fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=%s&parseTime=True&loc=Local",
				conf.DatabaseConfig.Username,
				conf.DatabaseConfig.Password,
				conf.DatabaseConfig.Host,
				conf.DatabaseConfig.Port,
				conf.DatabaseConfig.Database,
				conf.DatabaseConfig.Charset,
			))
		case "postgres", "postgresql":
			dialector = postgres.Open(fmt.Sprintf("host=%s port=%d user=%s dbname=%s password=%s sslmode=%s TimeZone=Asia/Shanghai",
				conf.DatabaseConfig.Host,
				conf.DatabaseConfig.Port,
				conf.DatabaseConfig.Username,
				conf.DatabaseConfig.Database,
				conf.DatabaseConfig.Password,
				conf.DatabaseConfig.SSLMode,
			))
		case "mssql", "sqlserver":
			dialector = sqlserver.Open(fmt.Sprintf("sqlserver://%s:%s@%s:%d?database=%s",
				conf.DatabaseConfig.Username,
				conf.DatabaseConfig.Password,
				conf.DatabaseConfig.Host,
				conf.DatabaseConfig.Port,
				conf.DatabaseConfig.Database,
			))
		default:
			logger.Log().Panicf("Database type error: %s", conf.DatabaseConfig.Type)
			return
		}
	}

	logLevel := gormlogger.Silent

	// Debug模式下，输出所有 SQL 日志
	if conf.SystemConfig.Debug {
		logLevel = gormlogger.Info
	}

	db, err := gorm.Open(dialector, &gorm.Config{
		PrepareStmt: true,
		Logger: gormlogger.New(gorm_logger.NewGormLogger(logger.GlobalLogger), gormlogger.Config{
			SlowThreshold:             200 * time.Millisecond,
			LogLevel:                  logLevel,
			IgnoreRecordNotFoundError: false,
			Colorful:                  true,
		}),
		NamingStrategy: schema.NamingStrategy{
			TablePrefix:   conf.DatabaseConfig.TablePrefix,
			SingularTable: true,
		},
	})
	if err != nil {
		logger.Log().Panicf("Database connection error: %s", err)
		return
	}

	logger.Log().Info("Database initial successful")

	DB = db
}
