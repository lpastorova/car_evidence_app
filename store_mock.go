package main

import (
	"github.com/stretchr/testify/mock"
)

// The mock store contains additonal methods for inspection
type MockStore struct {
	mock.Mock
}

func (m *MockStore) CreateCar(bird *Car) error {
	/*
		When this method is called, `m.Called` records the call, and also
		returns the result that we pass to it (which you will see in the
		handler tests)
	*/
	rets := m.Called(bird)
	return rets.Error(0)
}

func (m *MockStore) GetCar() ([]*Car, error) {
	rets := m.Called()
	/*
		Since `rets.Get()` is a generic method, that returns whatever we pass to it,
		we need to typecast it to the type we expect, which in this case is []*Car
	*/
	return rets.Get(0).([]*Car), rets.Error(1)
}
