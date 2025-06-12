package utilities

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var Logger *zap.Logger

// InitializeLogger menginisialisasi logger dengan konfigurasi yang sesuai
func InitializeLogger(isDevelopment bool) error {
	var config zap.Config

	if isDevelopment {
		config = zap.NewDevelopmentConfig()
		config.EncoderConfig.EncodeLevel = zapcore.CapitalColorLevelEncoder
	} else {
		config = zap.NewProductionConfig()
		config.EncoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	}

	var err error
	Logger, err = config.Build()
	if err != nil {
		return err
	}

	return nil
}

// GetLogger mengembalikan instance logger yang sudah diinisialisasi
func GetLogger() *zap.Logger {
	if Logger == nil {
		// Fallback ke development logger jika belum diinisialisasi
		Logger = zap.Must(zap.NewDevelopment())
	}
	return Logger
}