package global

import "go.uber.org/zap"

var (
	SugarLog *zap.SugaredLogger
)

//InitResource 初始化
func InitResource() {
	logger, _ := initLogger()
	SugarLog = logger.Sugar()

}

