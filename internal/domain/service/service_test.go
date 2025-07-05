package service

import (
	"reflect"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/jaswdr/faker/v2"
	"github.com/rafaeljpc/golang-first-study/internal/domain/model"
	"github.com/rafaeljpc/golang-first-study/internal/util"
	"github.com/stretchr/testify/assert"
)

func TestService_should_ListProducts(t *testing.T) {
	// Given
	faker := faker.New()
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	expectedProducts := []model.Product{
		{ID: util.GenerateUUID(), Name: faker.Pokemon().English(), Price: faker.Float64(2, 1, 100)},
		{ID: util.GenerateUUID(), Name: faker.Pokemon().English(), Price: faker.Float64(2, 1, 100)},
	}

	mockRepo := NewMockRepository(ctrl)
	mockRepo.EXPECT().ListProducts().Return(expectedProducts)

	// Cria o serviço usando o repositório mockado
	service := NewService(mockRepo)

	// When
	actualProducts := service.ListProducts()

	// Then
	assert.True(t, reflect.DeepEqual(expectedProducts, actualProducts))
}
