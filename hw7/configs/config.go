package config

import (
	"log"
	"os"
	"time"

	"github.com/kkyr/fig"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

// Config structure for all settings of application
type Config struct {
	App struct {
		Logger        *zap.Logger // logger for use, don't load from configuration file
		LogLevel      int         `fig:"logLevel" default:"0"`                                 // 0-info, 1-warn, -1-debug, 2-error, 4-panic, 5-fatal
		ServerAddress string      `fig:"serverAddress" default:"localhost"`                    // server address
		ServerPort    int         `fig:"serverPort" default:"9000"`                            // server port
		StoragePath   string      `fig:"storagePath" default:"../internal/storage/storage.db"` // storage db filepath
	} `fig:"app"`
}

// InitConfig function for initialize Config structure
func InitConfig(useConfig *string) (*Config, error) {
	var cfg = Config{}
	err := fig.Load(&cfg, fig.File(*useConfig))
	if err != nil {
		err = fig.Load(&cfg, fig.File("config.yml"))
		if err != nil {
			log.Fatalf("can't load configuration file: %s", err)
			return nil, err
		}
	}

	//Set log level
	atomicLevel := zap.NewAtomicLevel()
	switch cfg.App.LogLevel {
	case 0:
		{
			atomicLevel.SetLevel(zap.InfoLevel)
		}
	case 1:
		{
			atomicLevel.SetLevel(zap.WarnLevel)
		}
	case -1:
		{
			atomicLevel.SetLevel(zap.DebugLevel)
		}
	case 2:
		{
			atomicLevel.SetLevel(zap.ErrorLevel)
		}
	case 4:
		{
			atomicLevel.SetLevel(zap.PanicLevel)
		}
	case 5:
		{
			atomicLevel.SetLevel(zap.FatalLevel)
		}
	}

	encoderCfg := zap.NewProductionEncoderConfig()
	encoderCfg.EncodeTime = zapcore.TimeEncoderOfLayout(time.RFC3339)
	encoderCfg.EncodeLevel = zapcore.CapitalColorLevelEncoder
	encoderCfg.EncodeCaller = zapcore.ShortCallerEncoder

	logger := zap.New(zapcore.NewCore(
		zapcore.NewConsoleEncoder(encoderCfg),
		zapcore.Lock(os.Stdout),
		atomicLevel,
	), zap.AddCaller())

	cfg.App.Logger = logger

	return &cfg, err
}