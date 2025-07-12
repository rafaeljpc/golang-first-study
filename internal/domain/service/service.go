package service

import (
	"fmt"

	"github.com/rafaeljpc/golang-first-study/internal/domain/model"
)

type Service struct {
	repository Repository
}

func NewService(repository Repository) *Service {
	service := &Service{
		repository: repository,
	}
	return service
}

func (s *Service) ListProducts() []model.Product {
	fmt.Printf("ListProducts")
	response, err := s.repository.ListProducts()
	if err != nil {
		fmt.Printf("Error: %v", err)
		return []model.Product{}
	}

	return response
}
