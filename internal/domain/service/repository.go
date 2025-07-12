package service

import "github.com/rafaeljpc/golang-first-study/internal/domain/model"

type Repository interface {
	ListProducts() ([]model.Product, error)
}