package config

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"time"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

// Logger ...
var Lg *zap.Logger

// InitLogger ...
func InitLogger() {
	lp := Conf.Common.LogPath
	lv := Conf.Common.LogLevel
	var isDebug bool
	isDebug = Conf.Common.IsDebug
	fmt.Println("=======isdebug:=======", isDebug)
	initLogger(lp, lv, isDebug)
}

func fileExist(filename string) bool {
	_, err := os.Stat(filename)
	return err == nil || os.IsExist(err)
}

func checkLog(lp string, lv string, isDebug bool) {
	if !fileExist(lp) {
		initLogger(lp, lv, isDebug)
		log.SetFlags(log.Lmicroseconds | log.Lshortfile | log.LstdFlags)
	}
}

func initLogger(lp string, lv string, isDebug bool) {

	if isDebug {
		// 进入开发模式，日志输出到终端
		consoleEncoder := zapcore.NewConsoleEncoder(zap.NewDevelopmentEncoderConfig())
		core := zapcore.NewTee(
			//zapcore.NewCore(encoder, writeSyncer, l),
			zapcore.NewCore(consoleEncoder, zapcore.Lock(os.Stdout), zapcore.DebugLevel),
			//zapcore.NewCore(encoder, zapcore.Lock(os.Stdout), zapcore.DebugLevel),
		)
		Lg = zap.New(core, zap.AddCaller())

		zap.ReplaceGlobals(Lg)
		zap.L().Info("init logger success")
	} else {
		js := fmt.Sprintf(`{
		"level": "%s",
		"encoding": "json",
		"outputPaths": ["stdout","%s"],
		"errorOutputPaths": ["stderr","%s"]
	}`, lv, lp, lp)

		var cfg zap.Config
		if err := json.Unmarshal([]byte(js), &cfg); err != nil {
			panic(err)
		}
		cfg.EncoderConfig = zap.NewProductionEncoderConfig()
		cfg.EncoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
		cfg.EncoderConfig.EncodeLevel = zapcore.LowercaseColorLevelEncoder
		var err error
		Lg, err = cfg.Build()
		if err != nil {
			log.Fatal("init logger error: ", err)
		}
		log.SetFlags(log.Lmicroseconds | log.Lshortfile | log.LstdFlags)
		go func() {
			for {
				time.Sleep(5 * time.Second)
				checkLog(lp, lv, isDebug)
			}
		}()
	}

}
