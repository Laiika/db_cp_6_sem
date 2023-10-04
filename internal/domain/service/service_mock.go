// Code generated by MockGen. DO NOT EDIT.
// Source: db_cp_6_sem/internal/domain/service (interfaces: IRepo)

// Package service is a generated GoMock package.
package service

import (
	context "context"
	entity "db_cp_6_sem/internal/domain/entity"
	reflect "reflect"
	time "time"

	gomock "github.com/golang/mock/gomock"
)

// MockIRepo is a mock of IRepo interface.
type MockIRepo struct {
	ctrl     *gomock.Controller
	recorder *MockIRepoMockRecorder
}

// MockIRepoMockRecorder is the mock recorder for MockIRepo.
type MockIRepoMockRecorder struct {
	mock *MockIRepo
}

// NewMockIRepo creates a new mock instance.
func NewMockIRepo(ctrl *gomock.Controller) *MockIRepo {
	mock := &MockIRepo{ctrl: ctrl}
	mock.recorder = &MockIRepoMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockIRepo) EXPECT() *MockIRepoMockRecorder {
	return m.recorder
}

// CreateAnalyzer mocks base method.
func (m *MockIRepo) CreateAnalyzer(arg0 Client, arg1 context.Context, arg2 *entity.Analyzer) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateAnalyzer", arg0, arg1, arg2)
	ret0, _ := ret[0].(error)
	return ret0
}

// CreateAnalyzer indicates an expected call of CreateAnalyzer.
func (mr *MockIRepoMockRecorder) CreateAnalyzer(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateAnalyzer", reflect.TypeOf((*MockIRepo)(nil).CreateAnalyzer), arg0, arg1, arg2)
}

// CreateEvent mocks base method.
func (m *MockIRepo) CreateEvent(arg0 Client, arg1 context.Context, arg2 *entity.Event) (int, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateEvent", arg0, arg1, arg2)
	ret0, _ := ret[0].(int)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateEvent indicates an expected call of CreateEvent.
func (mr *MockIRepoMockRecorder) CreateEvent(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateEvent", reflect.TypeOf((*MockIRepo)(nil).CreateEvent), arg0, arg1, arg2)
}

// CreateGas mocks base method.
func (m *MockIRepo) CreateGas(arg0 Client, arg1 context.Context, arg2 *entity.Gas) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateGas", arg0, arg1, arg2)
	ret0, _ := ret[0].(error)
	return ret0
}

// CreateGas indicates an expected call of CreateGas.
func (mr *MockIRepoMockRecorder) CreateGas(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateGas", reflect.TypeOf((*MockIRepo)(nil).CreateGas), arg0, arg1, arg2)
}

// CreateSensor mocks base method.
func (m *MockIRepo) CreateSensor(arg0 Client, arg1 context.Context, arg2 *entity.Sensor) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateSensor", arg0, arg1, arg2)
	ret0, _ := ret[0].(error)
	return ret0
}

// CreateSensor indicates an expected call of CreateSensor.
func (mr *MockIRepoMockRecorder) CreateSensor(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateSensor", reflect.TypeOf((*MockIRepo)(nil).CreateSensor), arg0, arg1, arg2)
}

// CreateType mocks base method.
func (m *MockIRepo) CreateType(arg0 Client, arg1 context.Context, arg2 *entity.Type, arg3 []string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateType", arg0, arg1, arg2, arg3)
	ret0, _ := ret[0].(error)
	return ret0
}

// CreateType indicates an expected call of CreateType.
func (mr *MockIRepoMockRecorder) CreateType(arg0, arg1, arg2, arg3 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateType", reflect.TypeOf((*MockIRepo)(nil).CreateType), arg0, arg1, arg2, arg3)
}

// CreateUser mocks base method.
func (m *MockIRepo) CreateUser(arg0 Client, arg1 context.Context, arg2 *entity.User) (int, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateUser", arg0, arg1, arg2)
	ret0, _ := ret[0].(int)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateUser indicates an expected call of CreateUser.
func (mr *MockIRepoMockRecorder) CreateUser(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateUser", reflect.TypeOf((*MockIRepo)(nil).CreateUser), arg0, arg1, arg2)
}

// DeleteAnalyzer mocks base method.
func (m *MockIRepo) DeleteAnalyzer(arg0 Client, arg1 context.Context, arg2 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteAnalyzer", arg0, arg1, arg2)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteAnalyzer indicates an expected call of DeleteAnalyzer.
func (mr *MockIRepoMockRecorder) DeleteAnalyzer(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteAnalyzer", reflect.TypeOf((*MockIRepo)(nil).DeleteAnalyzer), arg0, arg1, arg2)
}

// DeleteEvent mocks base method.
func (m *MockIRepo) DeleteEvent(arg0 Client, arg1 context.Context, arg2 int) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteEvent", arg0, arg1, arg2)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteEvent indicates an expected call of DeleteEvent.
func (mr *MockIRepoMockRecorder) DeleteEvent(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteEvent", reflect.TypeOf((*MockIRepo)(nil).DeleteEvent), arg0, arg1, arg2)
}

// DeleteSensor mocks base method.
func (m *MockIRepo) DeleteSensor(arg0 Client, arg1 context.Context, arg2 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteSensor", arg0, arg1, arg2)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteSensor indicates an expected call of DeleteSensor.
func (mr *MockIRepoMockRecorder) DeleteSensor(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteSensor", reflect.TypeOf((*MockIRepo)(nil).DeleteSensor), arg0, arg1, arg2)
}

// DeleteType mocks base method.
func (m *MockIRepo) DeleteType(arg0 Client, arg1 context.Context, arg2 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteType", arg0, arg1, arg2)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteType indicates an expected call of DeleteType.
func (mr *MockIRepoMockRecorder) DeleteType(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteType", reflect.TypeOf((*MockIRepo)(nil).DeleteType), arg0, arg1, arg2)
}

// DeleteUser mocks base method.
func (m *MockIRepo) DeleteUser(arg0 Client, arg1 context.Context, arg2 int) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteUser", arg0, arg1, arg2)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteUser indicates an expected call of DeleteUser.
func (mr *MockIRepoMockRecorder) DeleteUser(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteUser", reflect.TypeOf((*MockIRepo)(nil).DeleteUser), arg0, arg1, arg2)
}

// GetAllAnalyzers mocks base method.
func (m *MockIRepo) GetAllAnalyzers(arg0 Client, arg1 context.Context) (entity.Analyzers, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAllAnalyzers", arg0, arg1)
	ret0, _ := ret[0].(entity.Analyzers)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAllAnalyzers indicates an expected call of GetAllAnalyzers.
func (mr *MockIRepoMockRecorder) GetAllAnalyzers(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAllAnalyzers", reflect.TypeOf((*MockIRepo)(nil).GetAllAnalyzers), arg0, arg1)
}

// GetAllGases mocks base method.
func (m *MockIRepo) GetAllGases(arg0 Client, arg1 context.Context) (entity.Gases, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAllGases", arg0, arg1)
	ret0, _ := ret[0].(entity.Gases)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAllGases indicates an expected call of GetAllGases.
func (mr *MockIRepoMockRecorder) GetAllGases(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAllGases", reflect.TypeOf((*MockIRepo)(nil).GetAllGases), arg0, arg1)
}

// GetAllSensors mocks base method.
func (m *MockIRepo) GetAllSensors(arg0 Client, arg1 context.Context) (entity.Sensors, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAllSensors", arg0, arg1)
	ret0, _ := ret[0].(entity.Sensors)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAllSensors indicates an expected call of GetAllSensors.
func (mr *MockIRepoMockRecorder) GetAllSensors(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAllSensors", reflect.TypeOf((*MockIRepo)(nil).GetAllSensors), arg0, arg1)
}

// GetAllTypes mocks base method.
func (m *MockIRepo) GetAllTypes(arg0 Client, arg1 context.Context) (entity.Types, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAllTypes", arg0, arg1)
	ret0, _ := ret[0].(entity.Types)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAllTypes indicates an expected call of GetAllTypes.
func (mr *MockIRepoMockRecorder) GetAllTypes(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAllTypes", reflect.TypeOf((*MockIRepo)(nil).GetAllTypes), arg0, arg1)
}

// GetAllUsers mocks base method.
func (m *MockIRepo) GetAllUsers(arg0 Client, arg1 context.Context) (entity.Users, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAllUsers", arg0, arg1)
	ret0, _ := ret[0].(entity.Users)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAllUsers indicates an expected call of GetAllUsers.
func (mr *MockIRepoMockRecorder) GetAllUsers(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAllUsers", reflect.TypeOf((*MockIRepo)(nil).GetAllUsers), arg0, arg1)
}

// GetAnalyzerById mocks base method.
func (m *MockIRepo) GetAnalyzerById(arg0 Client, arg1 context.Context, arg2 string) (*entity.Analyzer, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAnalyzerById", arg0, arg1, arg2)
	ret0, _ := ret[0].(*entity.Analyzer)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAnalyzerById indicates an expected call of GetAnalyzerById.
func (mr *MockIRepoMockRecorder) GetAnalyzerById(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAnalyzerById", reflect.TypeOf((*MockIRepo)(nil).GetAnalyzerById), arg0, arg1, arg2)
}

// GetAnalyzerSensors mocks base method.
func (m *MockIRepo) GetAnalyzerSensors(arg0 Client, arg1 context.Context, arg2 string) (entity.Sensors, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAnalyzerSensors", arg0, arg1, arg2)
	ret0, _ := ret[0].(entity.Sensors)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAnalyzerSensors indicates an expected call of GetAnalyzerSensors.
func (mr *MockIRepoMockRecorder) GetAnalyzerSensors(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAnalyzerSensors", reflect.TypeOf((*MockIRepo)(nil).GetAnalyzerSensors), arg0, arg1, arg2)
}

// GetEventsBySignalTime mocks base method.
func (m *MockIRepo) GetEventsBySignalTime(arg0 Client, arg1 context.Context, arg2, arg3 time.Time) (entity.Events, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetEventsBySignalTime", arg0, arg1, arg2, arg3)
	ret0, _ := ret[0].(entity.Events)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetEventsBySignalTime indicates an expected call of GetEventsBySignalTime.
func (mr *MockIRepoMockRecorder) GetEventsBySignalTime(arg0, arg1, arg2, arg3 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetEventsBySignalTime", reflect.TypeOf((*MockIRepo)(nil).GetEventsBySignalTime), arg0, arg1, arg2, arg3)
}

// GetGasByName mocks base method.
func (m *MockIRepo) GetGasByName(arg0 Client, arg1 context.Context, arg2 string) (*entity.Gas, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetGasByName", arg0, arg1, arg2)
	ret0, _ := ret[0].(*entity.Gas)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetGasByName indicates an expected call of GetGasByName.
func (mr *MockIRepoMockRecorder) GetGasByName(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetGasByName", reflect.TypeOf((*MockIRepo)(nil).GetGasByName), arg0, arg1, arg2)
}

// GetSensorById mocks base method.
func (m *MockIRepo) GetSensorById(arg0 Client, arg1 context.Context, arg2 string) (*entity.Sensor, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetSensorById", arg0, arg1, arg2)
	ret0, _ := ret[0].(*entity.Sensor)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetSensorById indicates an expected call of GetSensorById.
func (mr *MockIRepoMockRecorder) GetSensorById(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetSensorById", reflect.TypeOf((*MockIRepo)(nil).GetSensorById), arg0, arg1, arg2)
}

// GetSensorEvents mocks base method.
func (m *MockIRepo) GetSensorEvents(arg0 Client, arg1 context.Context, arg2 string) (entity.Events, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetSensorEvents", arg0, arg1, arg2)
	ret0, _ := ret[0].(entity.Events)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetSensorEvents indicates an expected call of GetSensorEvents.
func (mr *MockIRepoMockRecorder) GetSensorEvents(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetSensorEvents", reflect.TypeOf((*MockIRepo)(nil).GetSensorEvents), arg0, arg1, arg2)
}

// GetTypeAnalyzers mocks base method.
func (m *MockIRepo) GetTypeAnalyzers(arg0 Client, arg1 context.Context, arg2 string) (entity.Analyzers, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetTypeAnalyzers", arg0, arg1, arg2)
	ret0, _ := ret[0].(entity.Analyzers)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetTypeAnalyzers indicates an expected call of GetTypeAnalyzers.
func (mr *MockIRepoMockRecorder) GetTypeAnalyzers(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetTypeAnalyzers", reflect.TypeOf((*MockIRepo)(nil).GetTypeAnalyzers), arg0, arg1, arg2)
}

// GetTypeByName mocks base method.
func (m *MockIRepo) GetTypeByName(arg0 Client, arg1 context.Context, arg2 string) (*entity.Type, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetTypeByName", arg0, arg1, arg2)
	ret0, _ := ret[0].(*entity.Type)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetTypeByName indicates an expected call of GetTypeByName.
func (mr *MockIRepoMockRecorder) GetTypeByName(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetTypeByName", reflect.TypeOf((*MockIRepo)(nil).GetTypeByName), arg0, arg1, arg2)
}

// GetTypeGases mocks base method.
func (m *MockIRepo) GetTypeGases(arg0 Client, arg1 context.Context, arg2 string) (entity.Gases, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetTypeGases", arg0, arg1, arg2)
	ret0, _ := ret[0].(entity.Gases)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetTypeGases indicates an expected call of GetTypeGases.
func (mr *MockIRepoMockRecorder) GetTypeGases(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetTypeGases", reflect.TypeOf((*MockIRepo)(nil).GetTypeGases), arg0, arg1, arg2)
}

// GetUserById mocks base method.
func (m *MockIRepo) GetUserById(arg0 Client, arg1 context.Context, arg2 int) (*entity.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetUserById", arg0, arg1, arg2)
	ret0, _ := ret[0].(*entity.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetUserById indicates an expected call of GetUserById.
func (mr *MockIRepoMockRecorder) GetUserById(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetUserById", reflect.TypeOf((*MockIRepo)(nil).GetUserById), arg0, arg1, arg2)
}

// GetUserByName mocks base method.
func (m *MockIRepo) GetUserByName(arg0 Client, arg1 context.Context, arg2 string) (*entity.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetUserByName", arg0, arg1, arg2)
	ret0, _ := ret[0].(*entity.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetUserByName indicates an expected call of GetUserByName.
func (mr *MockIRepoMockRecorder) GetUserByName(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetUserByName", reflect.TypeOf((*MockIRepo)(nil).GetUserByName), arg0, arg1, arg2)
}

// UpdateSensorAnalyzer mocks base method.
func (m *MockIRepo) UpdateSensorAnalyzer(arg0 Client, arg1 context.Context, arg2, arg3 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateSensorAnalyzer", arg0, arg1, arg2, arg3)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateSensorAnalyzer indicates an expected call of UpdateSensorAnalyzer.
func (mr *MockIRepoMockRecorder) UpdateSensorAnalyzer(arg0, arg1, arg2, arg3 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateSensorAnalyzer", reflect.TypeOf((*MockIRepo)(nil).UpdateSensorAnalyzer), arg0, arg1, arg2, arg3)
}

// UpdateUserRole mocks base method.
func (m *MockIRepo) UpdateUserRole(arg0 Client, arg1 context.Context, arg2 int, arg3 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateUserRole", arg0, arg1, arg2, arg3)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateUserRole indicates an expected call of UpdateUserRole.
func (mr *MockIRepoMockRecorder) UpdateUserRole(arg0, arg1, arg2, arg3 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateUserRole", reflect.TypeOf((*MockIRepo)(nil).UpdateUserRole), arg0, arg1, arg2, arg3)
}