package main

import (
	"os"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func main() {
	//1
	// sugar := zap.NewExample().Sugar()
	// sugar.Infof("Hello Name: %s, age: %d", "HiuMX", 21)

	// logger := zap.NewExample()
	// logger.Info("Hello", zap.String("name", "HiuMx"), zap.Int("age", 21))

	//2
	// logger := zap.NewExample()
	// logger.Info("Hello Zap")

	// logger, _ = zap.NewDevelopment()
	// logger.Info("Hello Zap")

	// logger, _ = zap.NewProduction()
	// logger.Warn("Hello Zap")

	//3

	encoder := getEncoderLog()
	sync := getWriterSync()

	core := zapcore.NewCore(encoder, sync, zapcore.InfoLevel)

	logger := zap.New(core, zap.AddCaller())

	logger.Info("This is an info message")
	logger.Error("This is an error message")

}

func getEncoderLog() zapcore.Encoder {
	encoderConfig := zap.NewProductionEncoderConfig()

	encoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	encoderConfig.TimeKey = "time"
	encoderConfig.EncodeLevel = zapcore.CapitalLevelEncoder
	encoderConfig.EncodeCaller = zapcore.ShortCallerEncoder

	return zapcore.NewJSONEncoder(encoderConfig)
}

func getWriterSync() zapcore.WriteSyncer {
	file, err := os.OpenFile("./logs/log.txt", os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0666)

	if err != nil {
		panic(err)
	}

	syncFile := zapcore.AddSync(file)
	syncConsole := zapcore.AddSync(os.Stderr)

	return zapcore.NewMultiWriteSyncer(syncConsole, syncFile)

}
