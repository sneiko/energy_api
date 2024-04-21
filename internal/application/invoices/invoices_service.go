package invoices

import (
	"context"
	"fmt"
	"slices"

	"energy_tk/internal/domain"
	"energy_tk/pkg/clients/energytksite"
)

type InvoicesRepository interface {
	Create(ctx context.Context, invoice *domain.Invoice) (int, error)
	GetListByUserToken(ctx context.Context, token string) (domain.InvoiceList, error)
	GetDetails(ctx context.Context, id int) (*domain.Invoice, error)
	GetList(ctx context.Context) (domain.InvoiceList, error)
	AddState(ctx context.Context, invoiceID int, state *domain.InvoiceState) (int, error)
}

type UsersRepository interface {
	GetByToken(ctx context.Context, token string) (*domain.User, error)
}

type Service struct {
	usersRepository    UsersRepository
	invoicesRepository InvoicesRepository
	siteClient         *energytksite.Client
}

func New(
	usersRepository UsersRepository,
	invoicesRepository InvoicesRepository,
	siteClient *energytksite.Client,
) *Service {
	return &Service{
		usersRepository:    usersRepository,
		invoicesRepository: invoicesRepository,
		siteClient:         siteClient,
	}
}

func (s *Service) Create(ctx context.Context, token string, number string) error {
	sClient, err := s.siteClient.CheckInvoice(number)
	if err != nil {
		return fmt.Errorf("CheckInvoice - error: %w", err)
	}

	user, err := s.usersRepository.GetByToken(ctx, token)
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
		UserID:                    user.ID,
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

	_, err = s.invoicesRepository.Create(ctx, invoice)
	return err
}

func (s *Service) GetListByUserToken(ctx context.Context, token string) (domain.InvoiceList, error) {
	list, err := s.invoicesRepository.GetListByUserToken(ctx, token)
	return list, err
}

func (s *Service) GetDetails(ctx context.Context, id int) (*domain.Invoice, error) {
	data, err := s.invoicesRepository.GetDetails(ctx, id)
	return data, err
}

func (s *Service) GetList(ctx context.Context) (domain.InvoiceList, error) {
	return s.invoicesRepository.GetList(ctx)
}

func (s *Service) UpdateState(ctx context.Context, invoice *domain.Invoice) error {
	siteData, err := s.siteClient.CheckInvoice(invoice.InvoiceNumber)
	if err != nil {
		return fmt.Errorf("CheckInvoice - error: %w", err)
	}

	states := make([]domain.InvoiceState, len(siteData.States))
	for _, s := range siteData.States {
		if slices.ContainsFunc(invoice.States, func(state domain.InvoiceState) bool {
			return state.SiteID == s.IdState
		}) {
			continue
		}

		states = append(states, domain.InvoiceState{
			Title:               s.Title,
			MovingDate:          s.MovingDate,
			MovingDateFormatted: s.MovingDateFormatted,
			MovingFromCity:      s.StateInfo.Trip.CityFrom.Name,
			MovingToCity:        s.StateInfo.Trip.CityTo.Name,
		})
	}

	for _, state := range invoice.States {
		if _, err = s.invoicesRepository.AddState(ctx, invoice.ID, &state); err != nil {
			return fmt.Errorf("UpdateState - error: %w", err)
		}
	}

	return nil
}
