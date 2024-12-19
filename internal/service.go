package internal

import (
	"fmt"
	"github.com/antlko/golite/internal/db"
	"github.com/antlko/golite/internal/logger"
	oauthgoogle "github.com/antlko/golite/internal/oauth2/google"
	"github.com/antlko/golite/internal/server"
)

type AppConfig struct {
	Hostname        string `env:"HOSTNAME"`
	ApplicationName string `env:"APPLICATION_NAME"`

	Server       server.Config
	Ui           UIConfig
	DB           db.Config
	GoogleOauth2 oauthgoogle.Config
}

type UIConfig struct {
	ServerPort string `env:"VITE_SERVER_HOST,default=http://localhost:4000"`
	VuePort    string `env:"UI_PORT,default=5173"`
}

func InitService(cfg AppConfig) error {
	logger.InitLogger(logger.Config{
		AppName:  cfg.ApplicationName,
		Hostname: cfg.Hostname,
	})

	dbInst, err := db.NewDB(cfg.DB)
	if err != nil {
		return fmt.Errorf("new db: %w", err)
	}

	googleConfig := oauthgoogle.InitConfig(cfg.GoogleOauth2)

	if err := server.InitServer(cfg.Server, dbInst, googleConfig); err != nil {
		return fmt.Errorf("init server: %w", err)
	}
	return nil
}
