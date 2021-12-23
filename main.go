package main

import (
	"context"
	"fmt"
	"github.com/ZmaximillianZ/stskp_sport_api/internal/grpc/user_grpc/pb"
	"github.com/go-playground/validator"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
	"log"
	"net"
	"time"

	"github.com/ZmaximillianZ/stskp_sport_api/internal/controllers"
	"github.com/ZmaximillianZ/stskp_sport_api/internal/db"
	errorhandler "github.com/ZmaximillianZ/stskp_sport_api/internal/e"
	"github.com/ZmaximillianZ/stskp_sport_api/internal/logging"
	"github.com/ZmaximillianZ/stskp_sport_api/internal/repository"
	"github.com/ZmaximillianZ/stskp_sport_api/internal/routes"
	"github.com/ZmaximillianZ/stskp_sport_api/internal/setting"
	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
)

// init is invoked before main()
// @title Swagger Example API
// @version 1.0
// @description This is a sample user_protobuf Petstore user_protobuf.
// @termsOfService http://swagger.io/terms/

// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html
// @securityDefinitions.apikey JWT
// @in header
// @name X-AUTH-TOKEN
// @host localhost:8081
func main() {
	if err := godotenv.Load(); err != nil {
		log.Print("No .env file found")
	}
	settings := setting.LoadSetting()

	logging.Setup(&settings.App)
	defer func() {
		err := logging.Close()
		if err != nil {
			log.Print(err)
		}
	}()
	dbContext, err := db.CreateDatabaseContext(settings.DBConfig)
	if err != nil {
		log.Print(err)
		panic(err)
	}
	userRepo := repository.NewUserRepository(dbContext.Connection, dbContext.QueryBuilder)
	workoutRepo := repository.NewWorkoutRepository(dbContext.Connection, dbContext.QueryBuilder)
	workoutTypeRepo := repository.NewWorkoutTypeRepository(dbContext.Connection, dbContext.QueryBuilder)
	workoutValueRepo := repository.NewWorkoutValueRepository(dbContext.Connection, dbContext.QueryBuilder)
	eHandler := errorhandler.ErrorHandler{}
	userController := controllers.NewUserController(userRepo, eHandler, validator.New())
	workoutController := controllers.NewWorkoutController(workoutRepo, eHandler, validator.New())
	workoutTypeController := controllers.NewWorkoutTypeController(workoutTypeRepo, eHandler, validator.New())
	workoutValueController := controllers.NewWorkoutValueController(workoutValueRepo, eHandler, validator.New())
	eco := echo.New()
	router := eco.Group("/api/v1")
	router.Use(eHandler.Handle)
	routes.RegisterAPIV1(router, userController, workoutController, workoutTypeController, workoutValueController)

	lis, errLis := net.Listen("tcp", ":8081")
	if errLis != nil {
		fmt.Println(errLis.Error())
		log.Fatalln("can't listen port", errLis)
	}
	userGrpc := pb.NewGrpcServer(
		grpc.UnaryInterceptor(authInterceptor),
		//grpc.InTapHandle(rateLimiter),
	)
	userGrpc.Register()
	err = userGrpc.Server.Serve(lis)
	if err != nil {
		if err != nil {
			log.Print(err)
			fmt.Println(err.Error())
		}
	}

	eco.Logger.Fatal(eco.StartTLS(":1323", "cert.pem", "key.pem"))
}

func authInterceptor(
	ctx context.Context,
	req interface{},
	info *grpc.UnaryServerInfo,
	handler grpc.UnaryHandler,
) (interface{}, error) {
	start := time.Now()
	md, _ := metadata.FromIncomingContext(ctx)
	reply, err := handler(ctx, req)
	fmt.Printf(`--
	after incoming call=%v
	req=%#v
	reply=%#v
	time=%v
	md=%v
	err=%v
	`, info.FullMethod, req, reply, time.Since(start), md, err)
	return reply, err
}
