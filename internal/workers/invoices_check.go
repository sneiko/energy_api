package workers

import (
	"context"
	"log/slog"
	"time"

	"energy_tk/internal/domain"
)

const daemonDelay = 1 * time.Minute

type InvoicesStorage interface {
	GetList(ctx context.Context) (domain.InvoiceList, error)
	UpdateState(ctx context.Context, invoice *domain.Invoice) error
}

type InvoicesCheck struct {
	storage InvoicesStorage
}

func NewInvoicesCheck(storage InvoicesStorage) *InvoicesCheck {
	return &InvoicesCheck{
		storage: storage,
	}
}

func (i *InvoicesCheck) Do(ctx context.Context) error {
	logger := slog.With("worker", "invoices_check")
	for {
		list, err := i.storage.GetList(ctx)
		if err != nil {
			logger.Error("GetList", err)
		}

		for _, invoice := range list {
			if err := i.storage.UpdateState(ctx, &invoice); err != nil {
				logger.Error("UpdateState", err)
			}
		}
		time.Sleep(daemonDelay)
	}
}
