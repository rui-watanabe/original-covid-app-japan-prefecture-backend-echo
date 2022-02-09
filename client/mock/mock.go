// Code generated by MockGen. DO NOT EDIT.
// Source: client/client.go

// Package mock_client is a generated GoMock package.
package mock_client

import (
	client "original-covid-app-japan-prefecture-backend-echo/client"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockApiClient is a mock of ApiClient interface.
type MockApiClient struct {
	ctrl     *gomock.Controller
	recorder *MockApiClientMockRecorder
}

// MockApiClientMockRecorder is the mock recorder for MockApiClient.
type MockApiClientMockRecorder struct {
	mock *MockApiClient
}

// NewMockApiClient creates a new mock instance.
func NewMockApiClient(ctrl *gomock.Controller) *MockApiClient {
	mock := &MockApiClient{ctrl: ctrl}
	mock.recorder = &MockApiClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockApiClient) EXPECT() *MockApiClientMockRecorder {
	return m.recorder
}

// Request mocks base method.
func (m *MockApiClient) Request() (client.Prefectures, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Request")
	ret0, _ := ret[0].(client.Prefectures)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Request indicates an expected call of Request.
func (mr *MockApiClientMockRecorder) Request() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Request", reflect.TypeOf((*MockApiClient)(nil).Request))
}
