package alc_gorm

import (
	"alchemy/alc/alc_config"
	"alchemy/alc/alc_errs"
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"time"
)

func NewDB(appDbCfg alc_config.MysqlConfig) (db *gorm.DB, err error) {
	// loc=Local,标识跟随系统
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=true&loc=Local",
		appDbCfg.Usr,
		appDbCfg.Psw,
		appDbCfg.Host,
		appDbCfg.Port,
		appDbCfg.DbName,
	)
	conn := mysql.Open(dsn)
	gormCfg := &gorm.Config{
		SkipDefaultTransaction:                   false,
		NamingStrategy:                           nil,
		FullSaveAssociations:                     false,
		Logger:                                   nil,
		NowFunc:                                  nil,
		DryRun:                                   false,
		PrepareStmt:                              false,
		DisableAutomaticPing:                     false,
		DisableForeignKeyConstraintWhenMigrating: false,
		DisableNestedTransaction:                 false,
		AllowGlobalUpdate:                        false,
		QueryFields:                              false,
		CreateBatchSize:                          0,
		ClauseBuilders:                           nil,
		ConnPool:                                 nil,
		Dialector:                                nil,
		Plugins:                                  nil,
	}
	if appDbCfg.Debug {
		gormCfg.Logger = logger.Default.LogMode(logger.Info)
	} else {
		gormCfg.Logger = logger.Default.LogMode(logger.Error)
	}
	db, err = gorm.Open(conn, gormCfg)
	if err != nil {
		err = alc_errs.Wrap(err, "数据库初始化失败"+err.Error())
		return
	}
	sqlDB, err := db.DB()
	if err != nil {
		err = alc_errs.Wrap(err, "数据库初始化失败"+err.Error())
		return
	}

	// db.SingularTable(true)

	// 连接最长存活期，超过这个时间连接将不再被复用
	// db.DB().SetConnMaxLifetime(1 * time.Second)
	// 最大空闲连接数
	// db.DB().SetMaxIdleConns(-1)
	// 数据库最大连接数
	// db.DB().SetMaxOpenConns(120)
	if appDbCfg.MaxLifetime > 0 {
		sqlDB.SetConnMaxLifetime(appDbCfg.MaxLifetime * time.Second)
	}
	if appDbCfg.MaxIdleConns > 0 {
		sqlDB.SetMaxIdleConns(appDbCfg.MaxIdleConns)
	}

	if appDbCfg.MaxOpenConns > 0 {
		sqlDB.SetMaxOpenConns(appDbCfg.MaxOpenConns)
	}
	return
}
