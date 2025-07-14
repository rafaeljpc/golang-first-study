package service

import (
	"log"

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
	log.Default().Printf("ListProducts")
	response, err := s.repository.ListProducts()
	if err != nil {
		log.Fatalf("Error: %v", err)
		return []model.Product{}
	}

	return response
}
