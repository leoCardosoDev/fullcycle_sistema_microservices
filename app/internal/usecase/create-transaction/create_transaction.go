package createtransaction

import (
	"github.com.br/devfullcycle/fc-ms-wallet/internal/entity"
	"github.com.br/devfullcycle/fc-ms-wallet/internal/gateway"
)

type CreateTransacionInputDTO struct {
	AccountIdFrom string
	AccountIdTo string
	Amount float64
}

type CreateTransacionOutputDTO struct {
	ID string
}

type CreateTransactionUseCase struct {
	TransactionGateway gateway.TransactionGateway
	AccountGateway gateway.AccountGateway
}

func NewCreateTransactionUseCase(transactionGateway gateway.TransactionGateway, accountGateway gateway.AccountGateway) *CreateTransactionUseCase {
	return &CreateTransactionUseCase {
		TransactionGateway: transactionGateway,
		AccountGateway: accountGateway,
	}
}

func (uc *CreateTransactionUseCase) Execute(input CreateTransacionInputDTO) (*CreateTransacionOutputDTO, error) {
	accountFrom, err := uc.AccountGateway.FindByID(input.AccountIdFrom)
	if err != nil {
		return nil, err
	}
	accountTo, err := uc.AccountGateway.FindByID(input.AccountIdTo)
	if err != nil {
		return nil, err
	}
	transaction, err := entity.NewTransaction(accountFrom, accountTo, input.Amount)
	if err != nil {
		return nil, err
	}
	err = uc.TransactionGateway.Create(transaction)
	if err != nil {
		return nil, err
	}
	return &CreateTransacionOutputDTO{ID: transaction.ID}, nil
}
