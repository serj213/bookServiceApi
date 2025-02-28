package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	bsv1 "github.com/serj213/bookService/pb/grpc/grpc"
	"github.com/serj213/bookServiceApi/internal/config"
	HTTPServer "github.com/serj213/bookServiceApi/internal/http"
	"github.com/serj213/bookServiceApi/internal/services"
	"go.uber.org/zap"
	"google.golang.org/grpc"
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

	// инициализация grpc клиента


	conn, err := grpc.Dial(fmt.Sprintf("localhost:%d", cfg.GRPC.Port), grpc.WithInsecure())
	if err != nil {
		logSugar.Infof("failed start grpc client: %w", err)
		panic(err)
	}
	
	logSugar.Info("grpc client started...")

	defer conn.Close()

	bookClient := bsv1.NewBookClient(conn)

	bookService := services.New(logSugar, bookClient)

	httpServer := HTTPServer.New(logSugar, bookService)

	router := mux.NewRouter()

	router.HandleFunc("/create", httpServer.Create).Methods(http.MethodPost)
	router.HandleFunc("/books", httpServer.GetBooks).Methods(http.MethodGet)

	srv := &http.Server{
		Handler: router,
		Addr: cfg.HTTP.Addr,
	}

	logSugar.Infof("http server started: %d...", cfg.HTTP.Addr)

	if err := srv.ListenAndServe(); err != nil {
		logSugar.Infof("failed http server: %w", err)
		panic(err)
	}
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