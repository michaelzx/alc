package alc_gen

import (
	"database/sql"
	"fmt"
	"github.com/michaelzx/alc/alc_config"
	"github.com/michaelzx/alc/alc_logger"
	"go.uber.org/zap"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Config struct {
	RootPath    string
	RootPackage string
	DbCfg       alc_config.MysqlConfig
	Tables      []string
	TablePrefix string
	GenModel    bool
}

type Gen struct {
	rootPath    string
	rootPackage string
	dbCfg       alc_config.MysqlConfig
	db          *sql.DB
	tables      []*Table
	tablePrefix string
	logger      *zap.Logger
	genModel    bool
}

func New(cfg Config) (*Gen, error) {

	logger, err := alc_logger.New(alc_config.LoggerConfig{Mode: "dev"}, 0)
	if cfg.RootPath == "" {
		logger.Fatal("请指定RootPath")
	}
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=true&loc=Local",
		cfg.DbCfg.Usr,
		cfg.DbCfg.Psw,
		cfg.DbCfg.Host,
		cfg.DbCfg.Port,
		cfg.DbCfg.DbName,
	)
	gormDB, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	mysqlDB, err := gormDB.DB()
	if err != nil {
		return nil, err
	}
	tables := make([]*Table, 0, 0)
	for _, tableName := range cfg.Tables {
		tables = append(tables, NewTable(tableName, mysqlDB))
	}

	return &Gen{
		rootPath:    cfg.RootPath,
		rootPackage: cfg.RootPackage,
		db:          mysqlDB,
		dbCfg:       cfg.DbCfg,
		tables:      tables,
		tablePrefix: cfg.TablePrefix,
		logger:      logger,
		genModel:    cfg.GenModel,
	}, nil
}

func (g *Gen) Run() {
	if len(g.tables) == 0 {
		g.logger.Fatal("未指定任何需要生成的表")
	}
	g.logger.Info("对需要生成表，进行验证，并获取表信息...")
	for _, tbl := range g.tables {
		err := tbl.Check(g)
		if err != nil {
			g.logger.Fatal(err.Error())
		}
	}
	g.logger.Info("根据数据表生成...")
	for _, tbl := range g.tables {
		err := tbl.gen(g)
		if err != nil {
			g.logger.Fatal(err.Error())
		}
	}
}
