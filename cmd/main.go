package main

import (
	"database/sql"
	"log"
	"net"

	_ "github.com/mattn/go-sqlite3"
	"github.com/paulovitor-dock/clean-architecture-go/adapter/api"
	"github.com/paulovitor-dock/clean-architecture-go/adapter/grpc/pb"
	"github.com/paulovitor-dock/clean-architecture-go/adapter/grpc/service"
	"github.com/paulovitor-dock/clean-architecture-go/adapter/repository"
	"github.com/paulovitor-dock/clean-architecture-go/entity"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func main() {
	db, err := sql.Open("sqlite3", "test.db")

	if err != nil {
		log.Fatal(err)
	}

	repo := repository.NewTransactionRepositoryDb(db)
	// usecase := process_transaction.NewProcessTransaction(repo)
	// input := process_transaction.TransactionDtoInput{
	// 	ID:        "2",
	// 	AccountID: "1",
	// 	Amount:    10,
	// }
	// output, err := usecase.Execute(input)
	// if err != nil {
	// 	fmt.Println(err.Error())
	// }
	// outputJson, _ := json.Marshal(output)
	// fmt.Println(string(outputJson))
	webServer := api.NewWebServer()
	webServer.Repository = repo
	go webServer.Serve()
	startGRPCServer(repo)
}

func startGRPCServer(repo entity.TransactionRepository) {
	lis, err := net.Listen("tcp", "localhost:50052")
	if err != nil {
		log.Fatal(err)
	}
	grpcServer := grpc.NewServer()
	service := service.NewProcessService()
	service.Repository = repo
	pb.RegisterTransactionServiceServer(grpcServer, service)
	reflection.Register(grpcServer)
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatal(err)
	}
}
