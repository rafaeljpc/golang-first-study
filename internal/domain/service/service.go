package service

import "fmt"

type Service struct {
}

func NewService() *Service {
	service := &Service{}
	return service
}

func (s *Service) DoSomething() {
	fmt.Println("Doing something...")
}