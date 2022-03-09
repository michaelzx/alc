package alc_mgo

import (
	"context"
	"fmt"
	"github.com/michaelzx/alc/alc_color"
	"github.com/michaelzx/alc/alc_config"
	"github.com/qiniu/qmgo"
	qmgoOpts "github.com/qiniu/qmgo/options"
	"github.com/shopspring/decimal"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/event"
	mgoOpts "go.mongodb.org/mongo-driver/mongo/options"
	"go.uber.org/zap"
	"reflect"
	"time"
)

// NewDBWithZapLogger 初始化数据库链接实例
func NewDBWithZapLogger(dbCfg alc_config.MongoDBConfig, zapLogger *zap.Logger) (*qmgo.Database, error) {
	zapLogger.Info("InitDB", zap.Any("dbCfg", dbCfg))
	// ------------------------------------------------------------------------------------------
	// 构建链接字符串
	// ------------------------------------------------------------------------------------------
	var err error
	connStr := fmt.Sprintf("mongodb://%s:%s/%s",
		dbCfg.DbHost,
		dbCfg.DbPort,
		dbCfg.DbName,
	)
	// ------------------------------------------------------------------------------------------
	// 构建clientOptions 主要是用来打印日志
	// ------------------------------------------------------------------------------------------

	monitor := &event.CommandMonitor{
		Started: func(_ context.Context, e *event.CommandStartedEvent) {
			if dbCfg.CmdLog {
				fmt.Printf(alc_color.LightYellowStr("qmgo[%d] cmd->")+" %s\n", e.RequestID, e.Command)
			}
		},
		Succeeded: func(_ context.Context, e *event.CommandSucceededEvent) {
			if dbCfg.SucceededLog {
				fmt.Printf(alc_color.LightGreenStr("qmgo[%d] succeeded in %fms\n"), e.RequestID, float64(e.DurationNanos)/float64(time.Millisecond))
				fmt.Printf(alc_color.LightGreenStr("qmgo[%d] reply-> %s\n"), e.RequestID, e.Reply)
			}
		},
		Failed: func(_ context.Context, e *event.CommandFailedEvent) {
			if dbCfg.FailedLog {
				fmt.Printf(alc_color.LightRedStr("qmgo[%d] failed in %fms\n"), e.RequestID, float64(e.DurationNanos)/float64(time.Millisecond))
				fmt.Printf(alc_color.LightRedStr("qmgo[%d] reply->")+" %s\n", e.Failure)
			}
		},
	}
	clientOptions := qmgoOpts.ClientOptions{ // <--注意：这个options.ClientOptions是qmgo自己封装的类型，里面继承了官方的
		ClientOptions: &mgoOpts.ClientOptions{ // 这个opt是mongoDrive官方的options，我给他起别名为opt
			Monitor: monitor,
			Registry: bson.NewRegistryBuilder().
				RegisterTypeDecoder(reflect.TypeOf(decimal.Decimal{}), DecimalString{}).
				RegisterTypeEncoder(reflect.TypeOf(decimal.Decimal{}), DecimalString{}).
				Build(),
		},
	}
	// ------------------------------------------------------------------------------------------
	// 创建 client & db
	// ------------------------------------------------------------------------------------------
	// After 10 seconds, this function will return a timeout error.
	var timeout int64 = 10
	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(timeout)*time.Second)
	defer cancel()

	cfg := &qmgo.Config{
		Uri: connStr,
	}
	if dbCfg.DbName != "" && dbCfg.DbUser != "" && dbCfg.DbPass != "" {
		credential := &qmgo.Credential{
			AuthSource: dbCfg.DbName,
			Username:   dbCfg.DbUser,
			Password:   dbCfg.DbPass,
		}
		cfg.Auth = credential
	}

	dbClient, err := qmgo.NewClient(ctx, cfg, clientOptions)
	if err != nil {
		return nil, err
	}
	if err = dbClient.Ping(timeout); err != nil {
		return nil, err
	}
	return dbClient.Database(dbCfg.DbName), nil
}
