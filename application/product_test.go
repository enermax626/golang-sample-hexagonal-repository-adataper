package application_test

import (
	"testing"

	"github.com/codeedu/go-hexagonal/application"
	"github.com/dgryski/trifles/uuid"
	"github.com/stretchr/testify/require"
)

func TestProductEnable(t *testing.T) {
	product := application.Product{}
	product.Name = "Hello"
	product.Status = application.DISABLED
	product.Price = 10

	err := product.Enable()
	require.Nil(t, err)
}

func TestProductEnableError(t *testing.T) {
	product := application.Product{}
	product.Name = "Hello"
	product.Status = application.DISABLED
	product.Price = -1

	err := product.Enable()
	require.Equal(t, "price must be greater than zero", err.Error())
}

func TestProductDisable(t *testing.T) {
	product := application.Product{}
	product.Name = "Mesa"
	product.Price = 0

	err := product.Disable()
	require.Nil(t, err)

}

func TestProductDisableError(t *testing.T) {
	product := application.Product{}
	product.Name = "Mouse"
	product.Price = 10

	err := product.Disable()

	require.Equal(t, "price must be zero", err.Error())
}

func TestProductIsValid(t *testing.T) {
	product := application.Product{}
	product.Id = uuid.UUIDv4()
	product.Name = "Mouse"
	product.Price = 10

	isValid, err := product.IsValid()

	require.True(t, isValid)
	require.Nil(t, err)
}

func TestProductIsValidThrowIdError(t *testing.T) {
	product := application.Product{}
	product.Id = "12345"
	product.Name = "Mouse"
	product.Price = 10

	isValid, err := product.IsValid()

	require.False(t, isValid)
	require.Equal(t, "Id: 12345 does not validate as uuidv4", err.Error())
}
