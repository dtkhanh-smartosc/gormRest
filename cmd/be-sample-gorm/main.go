package main

import (
	"fmt"
	"github.com/HiBang15/sample-gorm.git/internal/database"
	userGrpc "github.com/HiBang15/sample-gorm.git/internal/module/user/grpc"
	pb "github.com/HiBang15/sample-gorm.git/proto/pb"
	"github.com/joho/godotenv"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"log"
	"net"
	"os"
	"strconv"
)

func init() {
	// load config from .env
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	databaseInfo := &database.DatabaseInfo{
		Name:     os.Getenv("DB_NAME"),
		Driver:   os.Getenv("DB_DRIVER"),
		Username: os.Getenv("DB_USER"),
		Password: os.Getenv("DB_PASSWORD"),
		SSLMode:  os.Getenv("DB_SSL_MODE"),
		Host:     os.Getenv("DB_HOST"),
		Port:     os.Getenv("DB_PORT"),
	}

	//connect database
	err = database.ConnectToDB(databaseInfo)
	if err != nil {
		log.Fatalf("Connect database fail with error: %v", err.Error())
	}
}

// @securityDefinitions.apikey bearerAuth
// @in header
// @name Authorization
//func main() {
//	//run Rest API
//	fmt.Println("Start run REST API OF TWC ADMIN API .....")
//	settingTwcAminApi := rest.SettingRestApi{
//		Environment: os.Getenv("ENVIRONMENT"),
//		Host:        os.Getenv("HOST"),
//		Port:        os.Getenv("PORT"),
//	}
//
//	rest.Load(settingTwcAminApi, public.SetRouter)
//}

func runGRPCServer(
	userServer pb.UserGRPCServiceServer,
	listener net.Listener,
	enableTLS bool,
) error {

	serverOptions := []grpc.ServerOption{}

	//if enableTLS {
	//	tlsCredentials, err := loadTLSCredentials()
	//	if err != nil {
	//		return fmt.Errorf("cannot load TLS credentials: %w", err)
	//	}
	//
	//	serverOptions = append(serverOptions, grpc.Creds(tlsCredentials))
	//}

	grpcServer := grpc.NewServer(serverOptions...)

	//------START register service Server------
	pb.RegisterUserGRPCServiceServer(grpcServer, userServer)
	//------END register service Server------
	reflection.Register(grpcServer)

	log.Printf("Start GRPC server at %s, TLS = %t", listener.Addr().String(), enableTLS)
	return grpcServer.Serve(listener)
}

func main() {
	fmt.Println("Starting GRPC Backend SmartContract Wallet!")

	host := os.Getenv("HOST")
	runGRPC, _ := strconv.ParseBool(os.Getenv("RUN_GRPC_ENDPOINT"))
	if runGRPC {
		//run grpc serve
		port := os.Getenv("LISTEN_GRPC_PORT")
		//var address string
		address := fmt.Sprintf("%s:%s", host, port)
		listener, err := net.Listen("tcp", address)
		if err != nil {
			log.Fatalf("[Backend SmartContract Wallet] Cannot start server at address: %s", address)
		}

		userServer := userGrpc.NewUserGrpcServer()

		err = runGRPCServer(
			userServer,
			listener,
			false,
		)
		if err != nil {
			log.Fatalf("[UserService][Main] Cannot start event with err: %s", err)
		}
	}
}
