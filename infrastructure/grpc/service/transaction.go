package service

import (
	"context"
	"github.com/dumunari/codebank/dto"
	"github.com/dumunari/codebank/infrastructure/grpc/pb"
	"github.com/dumunari/codebank/usecase"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
)

type TransactionService struct {
	ProcessTransactionUseCase usecase.UseCaseTransaction
	pb.UnimplementedPaymentServiceServer
}

func NewTransactionService() *TransactionService {
	return &TransactionService{}
}

func (t *TransactionService) Payment(ctx context.Context, in *pb.PaymentRequest) (*emptypb.Empty, error) {
	transactionDto := dto.Transaction{
		Name: in.CreditCard.Name,
		Number: in.CreditCard.Number,
		ExpirationMonth: in.CreditCard.ExpirationMonth,
		ExpirationYear: in.CreditCard.ExpirationYear,
		CVV: in.CreditCard.Cvv,
		Amount: in.GetAmount(),
		Store: in.GetStore(),
		Description: in.GetDescription(),
	}
	transaction, err := t.ProcessTransactionUseCase.ProcessTransaction(transactionDto)
	if err != nil {
		return &emptypb.Empty{}, status.Error(codes.FailedPrecondition, err.Error())
	}
	if transaction.Status != "approved" {
		return &emptypb.Empty{}, status.Error(codes.FailedPrecondition, "transaction " +
			"rejected by the bank")
	}
	return &emptypb.Empty{}, nil
}
