package invoices

import (
	"context"
	"fmt"

	"energy_tk/internal/domain"
	"energy_tk/pkg/clients/energytksite"
)

type InvoicesRepository interface {
	Create(ctx context.Context, token string, invoice *domain.Invoice) (int64, error)
	GetListByUserToken(ctx context.Context, token string) ([]domain.Invoice, error)
	GetDetails(ctx context.Context, id int) (*domain.Invoice, error)
}

type Service struct {
	invoicesRepository InvoicesRepository
	siteClient         *energytksite.Client
}

func New(invoicesRepository InvoicesRepository, client *energytksite.Client) *Service {
	return &Service{
		invoicesRepository: invoicesRepository,
	}
}

func (s *Service) Create(ctx context.Context, token string, number string) error {
	sClient, err := s.siteClient.CheckInvoice(number)
	if err != nil {
		return fmt.Errorf("CheckInvoice - error: %w", err)
	}

	states := make([]domain.InvoiceState, len(sClient.States))
	for _, s := range sClient.States {
		states = append(states, domain.InvoiceState{
			Title:               s.Title,
			MovingDate:          s.MovingDate,
			MovingDateFormatted: s.MovingDateFormatted,
			MovingFromCity:      s.StateInfo.Trip.CityFrom.Name,
			MovingToCity:        s.StateInfo.Trip.CityTo.Name,
		})
	}

	invoice := &domain.Invoice{
		InvoiceNumber:             number,
		FromCity:                  sClient.CityFrom.Name,
		ToCity:                    sClient.CityTo.Name,
		Places:                    sClient.Places,
		Weight:                    sClient.Weight,
		Volume:                    sClient.Volume,
		SenderIsPaid:              sClient.SenderIsPaid,
		RecipientIsPaid:           sClient.RecipientIsPaid,
		DeliveryDateFrom:          sClient.DeliveryDateFrom,
		DeliveryDateFromFormatted: sClient.DeliveryDateFromFormatted,
		DeliveryDateTo:            sClient.DeliveryDateTo,
		DeliveryDateToFormatted:   sClient.DeliveryDateToFormatted,
		SenderTotalPrice:          sClient.SenderTotalPrice,
		RecipientTotalPrice:       sClient.RecipientTotalPrice,
		States:                    states,
	}

	_, err = s.invoicesRepository.Create(ctx, token, invoice)
	return err
}

func (s *Service) GetListByUserToken(ctx context.Context, token string) ([]domain.Invoice, error) {
	list, err := s.invoicesRepository.GetListByUserToken(ctx, token)
	return list, err
}

func (s *Service) GetDetails(ctx context.Context, id int) (*domain.Invoice, error) {
	data, err := s.invoicesRepository.GetDetails(ctx, id)
	return data, err
}
