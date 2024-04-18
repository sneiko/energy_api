package workers

import "context"

type InvoicesCheck struct {
}

func NewInvoicesCheck() *InvoicesCheck {
	return &InvoicesCheck{}
}

func (i *InvoicesCheck) Do(ctx context.Context) error {
	for {
		// todo: implement daemon
	}
	return nil
}
