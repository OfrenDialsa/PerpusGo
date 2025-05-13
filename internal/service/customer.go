package service

import (
	"PerpusGo/domain"
	"PerpusGo/dto"
	"context"
	"database/sql"
	"errors"
	"time"

	"github.com/google/uuid"
)

type customerService struct {
	customerRepository domain.CustomerRepository
}

func NewCustomer(customerRepository domain.CustomerRepository) domain.CustomerService {
	return &customerService{
		customerRepository: customerRepository,
	}
}

// index implements domain.CustomerService.
func (cs *customerService) Index(ctx context.Context) ([]dto.CustomerData, error) {
	customer, err := cs.customerRepository.FindAll(ctx)
	if err != nil {
		return nil, err
	}

	var customerData []dto.CustomerData
	for _, c := range customer {
		customerData = append(customerData, dto.CustomerData{
			ID:   c.ID,
			Code: c.Code,
			Name: c.Name,
		})
	}
	return customerData, nil
}

// Create implements domain.CustomerService.
func (cs *customerService) Create(ctx context.Context, req dto.CreateCustomerRequest) error {
	customer := domain.Customer{
		ID:         uuid.NewString(),
		Code:       req.Code,
		Name:       req.Name,
		Created_at: sql.NullTime{Time: time.Now(), Valid: true},
	}
	return cs.customerRepository.Save(ctx, &customer)
}

// Update implements domain.CustomerService.
func (cs *customerService) Update(ctx context.Context, req dto.UpdateCustomerRequest) error {
	persisted, err := cs.customerRepository.FindByID(ctx, req.ID)
	if err != nil {
		return err
	}

	if persisted.ID == "" {
		return errors.New("customer not found")
	}

	persisted.Code = req.Code
	persisted.Name = req.Name
	persisted.Updated_at = sql.NullTime{Time: time.Now(), Valid: true}
	return cs.customerRepository.Update(ctx, &persisted)
}

// Delete implements domain.CustomerService.
func (cs *customerService) Delete(ctx context.Context, id string) error {
	exist, err := cs.customerRepository.FindByID(ctx, id)

	if err != nil {
		return err
	}

	if exist.ID == "" {
		return errors.New("customer data not found")
	}

	return cs.customerRepository.Delete(ctx, id)
}

// Show implements domain.CustomerService.
func (cs *customerService) Show(ctx context.Context, id string) (dto.CustomerData, error) {
	persisted, err := cs.customerRepository.FindByID(ctx, id)

	if err != nil {
		return dto.CustomerData{}, err
	}

	if persisted.ID == "" {
		return dto.CustomerData{}, errors.New("customer not found")
	}

	return dto.CustomerData{
		ID:   persisted.ID,
		Code: persisted.Code,
		Name: persisted.Name,
	}, nil
}
