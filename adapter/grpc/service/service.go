package service

import (
	"context"

	"github.com/paulovitor-dock/clean-architecture-go/adapter/grpc/pb"
	"github.com/paulovitor-dock/clean-architecture-go/entity"
	"github.com/paulovitor-dock/clean-architecture-go/usecase/process_transaction"
)

type ProcessService struct {
	Repository entity.TransactionRepository
	pb.UnimplementedTransactionServiceServer
}

func NewProcessService() *ProcessService {
	return &ProcessService{}
}

func (p *ProcessService) Process(ctx context.Context, in *pb.ProcessRequest) (*pb.ProcessResponse, error) {
	dto := process_transaction.TransactionDtoInput{
		ID:        in.Id,
		AccountID: in.AccountId,
		Amount:    float64(in.Amount),
	}
	usecase := process_transaction.NewProcessTransaction(p.Repository)
	output, _ := usecase.Execute(dto)
	return &pb.ProcessResponse{
		Id:           output.ID,
		Status:       output.Status,
		ErrorMessage: output.ErrorMessage,
	}, nil
}
