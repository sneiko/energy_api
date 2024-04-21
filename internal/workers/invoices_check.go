package workers

import (
	"context"
	"log/slog"
	"time"

	"energy_tk/internal/domain"
)

const daemonDelay = 10 * time.Minute
const sleepBetweenQuery = 1 * time.Second

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
			tNow := time.Now()
			if err := i.storage.UpdateState(ctx, &invoice); err != nil {
				logger.Error("UpdateState",
					slog.Int("invoice_id", invoice.ID),
					slog.String("invoice_number", invoice.InvoiceNumber),
					slog.Duration("duration", time.Since(tNow)),
					slog.String("error", err.Error()),
				)
				continue
			}
			logger.Info("UpdateState",
				slog.Int("invoice_id", invoice.ID),
				slog.String("invoice_number", invoice.InvoiceNumber),
				slog.Duration("duration", time.Since(tNow)),
			)
			time.Sleep(sleepBetweenQuery)
		}
		time.Sleep(daemonDelay)
	}
}
