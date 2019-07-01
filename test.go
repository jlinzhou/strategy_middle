// package main
 
// import "github.com/wonderivan/logger"
 
// func main() {

// 	logger.SetLogger(`{"Console": {"level": "DEBG"}}`)
//     // 通过配置文件配置
//     logger.SetLogger("logging/log.json")
// 	logger.Trace("this is Trace") // 由于默认输出，只会在控制台输出Debug及其以上日志，所以该条不会输出
// 	logger.Debug("this is Debug")
// 	logger.Info("this is Info")
// 	logger.Warn("this is Warn")
// 	logger.Error("this is Error")
// 	logger.Crit("this is Critical")
// 	logger.Alert("this is Alert")
// 	logger.Emer("this is Emergency")
// }