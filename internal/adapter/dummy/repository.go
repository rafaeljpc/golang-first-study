package dummy

import (
	"github.com/rafaeljpc/golang-first-study/internal/domain/model"
	"github.com/rafaeljpc/golang-first-study/internal/util"
)

type DummyRepository struct {
	Products []model.Product
}

func NewDummyRepository() *DummyRepository {
	return &DummyRepository{}
}

func (r DummyRepository) ListProducts() []model.Product {
	return []model.Product{
		{
			ID:    util.GenerateUUID(),
			Name:  "Product 1",
			Price: 1.99,
		},
		{
			ID:    util.GenerateUUID(),
			Name:  "Product 2",
			Price: 2.99,
		},
		{
			ID:    util.GenerateUUID(),
			Name:  "Product 3",
			Price: 3.99,
		},
	}
}
