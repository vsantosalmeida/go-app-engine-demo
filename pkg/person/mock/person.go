// Code generated by MockGen. DO NOT EDIT.
// Source: interface.go

// Package mock_person is a generated GoMock package.
package mock_person

import (
	entity "go-app-engine-demo/pkg/entity"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockReader is a mock of Reader interface.
type MockReader struct {
	ctrl     *gomock.Controller
	recorder *MockReaderMockRecorder
}

// MockReaderMockRecorder is the mock recorder for MockReader.
type MockReaderMockRecorder struct {
	mock *MockReader
}

// NewMockReader creates a new mock instance.
func NewMockReader(ctrl *gomock.Controller) *MockReader {
	mock := &MockReader{ctrl: ctrl}
	mock.recorder = &MockReaderMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockReader) EXPECT() *MockReaderMockRecorder {
	return m.recorder
}

// FindAll mocks base method.
func (m *MockReader) FindAll() ([]*entity.Person, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindAll")
	ret0, _ := ret[0].([]*entity.Person)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindAll indicates an expected call of FindAll.
func (mr *MockReaderMockRecorder) FindAll() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindAll", reflect.TypeOf((*MockReader)(nil).FindAll))
}

// FindByKey mocks base method.
func (m *MockReader) FindByKey(k string) (*entity.Person, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindByKey", k)
	ret0, _ := ret[0].(*entity.Person)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindByKey indicates an expected call of FindByKey.
func (mr *MockReaderMockRecorder) FindByKey(k interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindByKey", reflect.TypeOf((*MockReader)(nil).FindByKey), k)
}

// MockWriter is a mock of Writer interface.
type MockWriter struct {
	ctrl     *gomock.Controller
	recorder *MockWriterMockRecorder
}

// MockWriterMockRecorder is the mock recorder for MockWriter.
type MockWriterMockRecorder struct {
	mock *MockWriter
}

// NewMockWriter creates a new mock instance.
func NewMockWriter(ctrl *gomock.Controller) *MockWriter {
	mock := &MockWriter{ctrl: ctrl}
	mock.recorder = &MockWriterMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockWriter) EXPECT() *MockWriterMockRecorder {
	return m.recorder
}

// Delete mocks base method.
func (m *MockWriter) Delete(k string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Delete", k)
	ret0, _ := ret[0].(error)
	return ret0
}

// Delete indicates an expected call of Delete.
func (mr *MockWriterMockRecorder) Delete(k interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockWriter)(nil).Delete), k)
}

// Store mocks base method.
func (m *MockWriter) Store(p *entity.Person) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Store", p)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Store indicates an expected call of Store.
func (mr *MockWriterMockRecorder) Store(p interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Store", reflect.TypeOf((*MockWriter)(nil).Store), p)
}

// StoreMulti mocks base method.
func (m *MockWriter) StoreMulti(p []*entity.Person) ([]string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "StoreMulti", p)
	ret0, _ := ret[0].([]string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// StoreMulti indicates an expected call of StoreMulti.
func (mr *MockWriterMockRecorder) StoreMulti(p interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "StoreMulti", reflect.TypeOf((*MockWriter)(nil).StoreMulti), p)
}

// MockRepository is a mock of Repository interface.
type MockRepository struct {
	ctrl     *gomock.Controller
	recorder *MockRepositoryMockRecorder
}

// MockRepositoryMockRecorder is the mock recorder for MockRepository.
type MockRepositoryMockRecorder struct {
	mock *MockRepository
}

// NewMockRepository creates a new mock instance.
func NewMockRepository(ctrl *gomock.Controller) *MockRepository {
	mock := &MockRepository{ctrl: ctrl}
	mock.recorder = &MockRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockRepository) EXPECT() *MockRepositoryMockRecorder {
	return m.recorder
}

// Delete mocks base method.
func (m *MockRepository) Delete(k string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Delete", k)
	ret0, _ := ret[0].(error)
	return ret0
}

// Delete indicates an expected call of Delete.
func (mr *MockRepositoryMockRecorder) Delete(k interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockRepository)(nil).Delete), k)
}

// FindAll mocks base method.
func (m *MockRepository) FindAll() ([]*entity.Person, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindAll")
	ret0, _ := ret[0].([]*entity.Person)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindAll indicates an expected call of FindAll.
func (mr *MockRepositoryMockRecorder) FindAll() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindAll", reflect.TypeOf((*MockRepository)(nil).FindAll))
}

// FindByKey mocks base method.
func (m *MockRepository) FindByKey(k string) (*entity.Person, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindByKey", k)
	ret0, _ := ret[0].(*entity.Person)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindByKey indicates an expected call of FindByKey.
func (mr *MockRepositoryMockRecorder) FindByKey(k interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindByKey", reflect.TypeOf((*MockRepository)(nil).FindByKey), k)
}

// Store mocks base method.
func (m *MockRepository) Store(p *entity.Person) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Store", p)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Store indicates an expected call of Store.
func (mr *MockRepositoryMockRecorder) Store(p interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Store", reflect.TypeOf((*MockRepository)(nil).Store), p)
}

// StoreMulti mocks base method.
func (m *MockRepository) StoreMulti(p []*entity.Person) ([]string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "StoreMulti", p)
	ret0, _ := ret[0].([]string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// StoreMulti indicates an expected call of StoreMulti.
func (mr *MockRepositoryMockRecorder) StoreMulti(p interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "StoreMulti", reflect.TypeOf((*MockRepository)(nil).StoreMulti), p)
}

// MockUseCase is a mock of UseCase interface.
type MockUseCase struct {
	ctrl     *gomock.Controller
	recorder *MockUseCaseMockRecorder
}

// MockUseCaseMockRecorder is the mock recorder for MockUseCase.
type MockUseCaseMockRecorder struct {
	mock *MockUseCase
}

// NewMockUseCase creates a new mock instance.
func NewMockUseCase(ctrl *gomock.Controller) *MockUseCase {
	mock := &MockUseCase{ctrl: ctrl}
	mock.recorder = &MockUseCaseMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockUseCase) EXPECT() *MockUseCaseMockRecorder {
	return m.recorder
}

// Delete mocks base method.
func (m *MockUseCase) Delete(k string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Delete", k)
	ret0, _ := ret[0].(error)
	return ret0
}

// Delete indicates an expected call of Delete.
func (mr *MockUseCaseMockRecorder) Delete(k interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockUseCase)(nil).Delete), k)
}

// FindAll mocks base method.
func (m *MockUseCase) FindAll() ([]*entity.Person, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindAll")
	ret0, _ := ret[0].([]*entity.Person)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindAll indicates an expected call of FindAll.
func (mr *MockUseCaseMockRecorder) FindAll() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindAll", reflect.TypeOf((*MockUseCase)(nil).FindAll))
}

// FindByKey mocks base method.
func (m *MockUseCase) FindByKey(k string) (*entity.Person, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindByKey", k)
	ret0, _ := ret[0].(*entity.Person)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindByKey indicates an expected call of FindByKey.
func (mr *MockUseCaseMockRecorder) FindByKey(k interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindByKey", reflect.TypeOf((*MockUseCase)(nil).FindByKey), k)
}

// Store mocks base method.
func (m *MockUseCase) Store(p *entity.Person) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Store", p)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Store indicates an expected call of Store.
func (mr *MockUseCaseMockRecorder) Store(p interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Store", reflect.TypeOf((*MockUseCase)(nil).Store), p)
}

// StoreMulti mocks base method.
func (m *MockUseCase) StoreMulti(p []*entity.Person) ([]string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "StoreMulti", p)
	ret0, _ := ret[0].([]string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// StoreMulti indicates an expected call of StoreMulti.
func (mr *MockUseCaseMockRecorder) StoreMulti(p interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "StoreMulti", reflect.TypeOf((*MockUseCase)(nil).StoreMulti), p)
}