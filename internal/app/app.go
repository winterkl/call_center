package app

import (
	"contact_center/internal/config"
	v1 "contact_center/internal/controller/http/v1"
	auth_usecase "contact_center/internal/domain/auth/usecase"
	call_usecase "contact_center/internal/domain/call/usecase"
	auth_api "contact_center/internal/infrastructure/api/auth"
	call_repo "contact_center/internal/infrastructure/repo/call"
	"contact_center/pkg/http_server"
	"contact_center/pkg/postgres"
	"github.com/gin-gonic/gin"
	authv1 "github.com/winterkl/auth_protobuf/gen/go/proto/auth"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
)

type App struct {
	db         *postgres.Postgres
	handler    *gin.Engine
	httpServer *http_server.Server
	gRPCConn   *grpc.ClientConn
}

func NewApp(config *config.Config) *App {
	// Подключение к PSql

	psql, err := postgres.New(config.DataBase.User, config.DataBase.Password, config.DataBase.Host, config.DataBase.DB, config.DataBase.SSLMode, config.DataBase.Port)
	if err != nil {
		log.Fatalf("Не удалось подключиться к базе данных: %s", err.Error())
	}
	// Маршрутизатор
	handler := gin.New()

	// Создание gRPC клиента
	target := config.GRPC.Host + ":" + config.GRPC.Port
	conn, err := grpc.NewClient(target, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatal(err)
	}

	return &App{
		db:         psql,
		handler:    handler,
		httpServer: http_server.New(config.HTTP.Host, config.HTTP.Port, handler),
		gRPCConn:   conn,
	}
}

func (app *App) Run() {
	// Инициализация Repositories
	callRepo := call_repo.New(app.db)

	authClient := authv1.NewAuthClient(app.gRPCConn)

	// Инициализация Api
	authApi := auth_api.New(authClient)

	// Инициализация UseCases
	useCases := v1.UC{
		CallUC: call_usecase.New(callRepo),
		AuthUC: auth_usecase.New(authApi),
	}

	// Инициализация Router
	v1.New(app.handler, useCases, authApi)

	// Запуск http-сервера
	if err := app.httpServer.Start(); err != nil {
		log.Fatal(err)
	}
}
