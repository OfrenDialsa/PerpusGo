package service

import (
	"PerpusGo/domain"
	"PerpusGo/dto"
	"context"
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
