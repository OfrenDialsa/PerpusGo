package service

import (
	"PerpusGo/domain"
	"PerpusGo/dto"
	"context"
	"database/sql"
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
