package server

import (
	"github.com/dumunari/codebank/infrastructure/grpc/pb"
	"github.com/dumunari/codebank/infrastructure/grpc/service"
	"github.com/dumunari/codebank/usecase"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"log"
	"net"
)

type GRPCServer struct {
	ProcessTransactionUseCase usecase.UseCaseTransaction
}

func NewGRPCServer() *GRPCServer {
	return &GRPCServer{}
}

func (g *GRPCServer) Serve() {
	listener, err := net.Listen("tcp", "0.0.0.0:50052")
	if err != nil {
		log.Fatal("could not listen tcp port")
	}
	transactionService := service.NewTransactionService()
	transactionService.ProcessTransactionUseCase = g.ProcessTransactionUseCase
	grpcServer := grpc.NewServer()
	reflection.Register(grpcServer)
	pb.RegisterPaymentServiceServer(grpcServer, transactionService)
	grpcServer.Serve(listener)
}