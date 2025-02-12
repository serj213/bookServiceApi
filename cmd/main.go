package main

import (
	"github.com/serj213/bookServiceApi/internal/config"
	"go.uber.org/zap"
)

const (
	local = "local"
	dev = "develop"
)

func main(){
	cfg, err := config.Deal()
	if err != nil {
		panic(err)
	}


	log := setupLogger(cfg.Env)
	logSugar := log.Sugar()
	logSugar = logSugar.With(zap.String("env", cfg.Env))

	logSugar.Info("logger is enabled")

}

func setupLogger(env string) *zap.Logger{
	var log *zap.Logger

	switch(env){
	case local:
		log = zap.Must(zap.NewDevelopment())
	case dev:
		log = zap.Must(zap.NewProduction())
	default:
		log = zap.Must(zap.NewProduction())
	}
	return log
}