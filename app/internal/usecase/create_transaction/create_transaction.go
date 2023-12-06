package create_transaction

import (
	"context"

	"github.com.br/devfullcycle/fc-ms-wallet/internal/entity"
	"github.com.br/devfullcycle/fc-ms-wallet/internal/gateway"
	"github.com.br/devfullcycle/fc-ms-wallet/pkg/events"
	"github.com.br/devfullcycle/fc-ms-wallet/pkg/uow"
)

type CreateTransacionInputDTO struct {
	AccountIdFrom string `json:"account_id_from"`
	AccountIdTo string `json:"account_id_to"`
	Amount float64 `json:"amount"`
}

type CreateTransacionOutputDTO struct {
	ID string `json:"id"`
	AccountIdFrom string `json:"account_id_from"`
	AccountIdTo string `json:"account_id_to"`
	Amount float64 `json:"amount"`
}

type CreateTransactionUseCase struct {
	Uow uow.UowInterface
	EventDispatcher events.EventDispatcherInterface
	TransactionCreated events.EventInterface
}

func NewCreateTransactionUseCase(
	Uow uow.UowInterface,
	eventDispatcher events.EventDispatcherInterface,
	transactionCreated events.EventInterface,
) *CreateTransactionUseCase {
	return &CreateTransactionUseCase {
		Uow: Uow,
		EventDispatcher: eventDispatcher,
		TransactionCreated: transactionCreated,
	}
}

func (uc *CreateTransactionUseCase) Execute(ctx context.Context, input CreateTransacionInputDTO) (*CreateTransacionOutputDTO, error) {
	output := &CreateTransacionOutputDTO{}
	err := uc.Uow.Do(ctx, func(_ *uow.Uow) error {
		accountRepository := uc.getAccountRepository(ctx)
		transactionRepository := uc.getTransactionRepository(ctx)
		accountFrom, err := accountRepository.FindByID(input.AccountIdFrom)
		if err != nil {
			return err
		}
		accountTo, err := accountRepository.FindByID(input.AccountIdTo)
		if err != nil {
			return err
		}
		transaction, err := entity.NewTransaction(accountFrom, accountTo, input.Amount)
		if err != nil {
			return err
		}
		err = accountRepository.UpdateBalance(accountFrom)
		if err!= nil {
			return err
		}
		err = accountRepository.UpdateBalance(accountTo)
		if err!= nil {
			return err
		}
		err = transactionRepository.Create(transaction)
		if err != nil {
			return err
		}
		output.ID = transaction.ID
		output.AccountIdFrom = input.AccountIdFrom
		output.AccountIdTo = input.AccountIdTo
		return nil
	})
	if err!= nil {
    return nil, err
  }
	uc.TransactionCreated.SetPayload(output)
	uc.EventDispatcher.Dispatch(uc.TransactionCreated)
	return output, nil
}


func (uc *CreateTransactionUseCase) getAccountRepository(ctx context.Context) gateway.AccountGateway {
	repo, err := uc.Uow.GetRepository(ctx, "AccountDB")
  if err!= nil {
    return nil
  }
  return repo.(gateway.AccountGateway)
}

func (uc *CreateTransactionUseCase) getTransactionRepository(ctx context.Context) gateway.TransactionGateway {
	repo, err := uc.Uow.GetRepository(ctx, "TransactionDB")
  if err!= nil {
    return nil
  }
  return repo.(gateway.TransactionGateway)
}
