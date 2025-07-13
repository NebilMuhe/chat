package initiator

import (
	handler "chat_system/internal/handler/user"
	module "chat_system/internal/module/user"
	persistence "chat_system/internal/persistence/user"
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/spf13/viper"
)

func Initiator() {
	InitConfig()
	redisClient := InitRedis(context.Background(), viper.GetString("redis_url"),viper.GetString("password"))

	persistence := persistence.InitPersistence(redisClient)

	module := module.InitUserModule(persistence, viper.GetString("secret_key"))

	userHandler := handler.InitUserHandler(module)

	serverMux := &http.ServeMux{}

	serverMux.HandleFunc("POST /signup", func(w http.ResponseWriter, r *http.Request) {
		userHandler.CreateUser(w, r)
	})
	serverMux.HandleFunc("POST /login", func(w http.ResponseWriter, r *http.Request) {
		userHandler.LoginUser(w, r)
	})

	server := &http.Server{
		Addr:    ":" + viper.GetString("port"),
		Handler: serverMux,
	}

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, syscall.SIGTERM)

	go func() {
		if err := server.ListenAndServe(); err != nil {
			log.Printf("server stopped with error %s\n", err)
		}
	}()

	log.Println("server started")
	<-quit
}
