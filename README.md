# 更新记录

## 1.1.4

- [x] 完善alc_color
    - [x] 调整gorm的关联代码

## 1.0.8

- [x] 增加：i18n基础库`alc_i18n`
- [x] alc_logger：生产模式下，初始化时输出日志写入的路径
- [x] 增加：命令行颜色基础库 `alc_color`
- [x] alc_gorm：增加新方法 `NewDBWithLogger(appDbCfg alc_config.MysqlConfig, zapLogger *zap.Logger)` ，替换内置日志
- [x] alc_logger：优化日志输出，尽量做到对齐
- [X] alc_gin：增加中间件 ZapLogger、ZapRecovery，以便统一日志输出

