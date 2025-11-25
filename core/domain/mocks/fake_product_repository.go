// Code generated manually (mockgen + ajustes para generics). DO NOT EDIT AUTOMATICALLY.

package mocks

import (
	"reflect"

	"github.com/brudrago/clean-go/core/domain"
	"github.com/brudrago/clean-go/core/dto"
	"github.com/golang/mock/gomock"
)

// MockProductRepository is a mock of ProductRepository interface.
type MockProductRepository struct {
	ctrl     *gomock.Controller
	recorder *MockProductRepositoryMockRecorder
}

// MockProductRepositoryMockRecorder is the mock recorder for MockProductRepository.
type MockProductRepositoryMockRecorder struct {
	mock *MockProductRepository
}

// NewMockProductRepository creates a new mock instance.
func NewMockProductRepository(ctrl *gomock.Controller) *MockProductRepository {
	mock := &MockProductRepository{ctrl: ctrl}
	mock.recorder = &MockProductRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockProductRepository) EXPECT() *MockProductRepositoryMockRecorder {
	return m.recorder
}

// Create mocks base method.
func (m *MockProductRepository) Create(arg0 *dto.CreateProductRequest) (*domain.Product, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Create", arg0)
	ret0, _ := ret[0].(*domain.Product)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Create indicates an expected call of Create.
func (mr *MockProductRepositoryMockRecorder) Create(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(
		mr.mock,
		"Create",
		reflect.TypeOf((*MockProductRepository)(nil).Create),
		arg0,
	)
}

// Fetch mocks base method.
func (m *MockProductRepository) Fetch(arg0 *dto.PaginationRequestParams) (*domain.Pagination[[]domain.Product], error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Fetch", arg0)
	ret0, _ := ret[0].(*domain.Pagination[[]domain.Product])
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Fetch indicates an expected call of Fetch.
func (mr *MockProductRepositoryMockRecorder) Fetch(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(
		mr.mock,
		"Fetch",
		reflect.TypeOf((*MockProductRepository)(nil).Fetch),
		arg0,
	)
}
