package gateway

import "github.com/prtercio/fs-ws-wallet/internal/entity"

type TransactionGateway interface {
	Create(transaction *entity.Transaction) error
}
