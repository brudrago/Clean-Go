package mocks

import (
	"reflect"

	"github.com/brudrago/clean-go/core/domain"
	"github.com/brudrago/clean-go/core/dto"
	"github.com/golang/mock/gomock"
)

// MockProductUseCase is a mock of ProductUseCase interface.
type MockProductUseCase struct {
	ctrl     *gomock.Controller
	recorder *MockProductUseCaseMockRecorder
}

// MockProductUseCaseMockRecorder is the mock recorder for MockProductUseCase.
type MockProductUseCaseMockRecorder struct {
	mock *MockProductUseCase
}

// NewMockProductUseCase creates a new mock instance.
func NewMockProductUseCase(ctrl *gomock.Controller) *MockProductUseCase {
	mock := &MockProductUseCase{ctrl: ctrl}
	mock.recorder = &MockProductUseCaseMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockProductUseCase) EXPECT() *MockProductUseCaseMockRecorder {
	return m.recorder
}

// Create mocks base method.
func (m *MockProductUseCase) Create(arg0 *dto.CreateProductRequest) (*domain.Product, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Create", arg0)
	ret0, _ := ret[0].(*domain.Product)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Create indicates an expected call of Create.
func (mr *MockProductUseCaseMockRecorder) Create(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(
		mr.mock,
		"Create",
		reflect.TypeOf((*MockProductUseCase)(nil).Create),
		arg0,
	)
}

// Fetch mocks base method.
func (m *MockProductUseCase) Fetch(arg0 *dto.PaginationRequestParams) (*domain.Pagination[[]domain.Product], error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Fetch", arg0)
	ret0, _ := ret[0].(*domain.Pagination[[]domain.Product])
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Fetch indicates an expected call of Fetch.
func (mr *MockProductUseCaseMockRecorder) Fetch(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(
		mr.mock,
		"Fetch",
		reflect.TypeOf((*MockProductUseCase)(nil).Fetch),
		arg0,
	)
}
