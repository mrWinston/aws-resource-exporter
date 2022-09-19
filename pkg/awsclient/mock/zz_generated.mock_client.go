// Code generated by MockGen. DO NOT EDIT.
// Source: ./awsclient.go

// Package mock is a generated GoMock package.
package mock

import (
	reflect "reflect"

	aws "github.com/aws/aws-sdk-go/aws"
	request "github.com/aws/aws-sdk-go/aws/request"
	ec2 "github.com/aws/aws-sdk-go/service/ec2"
	servicequotas "github.com/aws/aws-sdk-go/service/servicequotas"
	gomock "github.com/golang/mock/gomock"
)

// MockClient is a mock of Client interface.
type MockClient struct {
	ctrl     *gomock.Controller
	recorder *MockClientMockRecorder
}

// MockClientMockRecorder is the mock recorder for MockClient.
type MockClientMockRecorder struct {
	mock *MockClient
}

// NewMockClient creates a new mock instance.
func NewMockClient(ctrl *gomock.Controller) *MockClient {
	mock := &MockClient{ctrl: ctrl}
	mock.recorder = &MockClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockClient) EXPECT() *MockClientMockRecorder {
	return m.recorder
}

// DescribeTransitGatewaysWithContext mocks base method.
func (m *MockClient) DescribeTransitGatewaysWithContext(ctx aws.Context, input *ec2.DescribeTransitGatewaysInput, opts ...request.Option) (*ec2.DescribeTransitGatewaysOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, input}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "DescribeTransitGatewaysWithContext", varargs...)
	ret0, _ := ret[0].(*ec2.DescribeTransitGatewaysOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DescribeTransitGatewaysWithContext indicates an expected call of DescribeTransitGatewaysWithContext.
func (mr *MockClientMockRecorder) DescribeTransitGatewaysWithContext(ctx, input interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, input}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DescribeTransitGatewaysWithContext", reflect.TypeOf((*MockClient)(nil).DescribeTransitGatewaysWithContext), varargs...)
}

// GetServiceQuota mocks base method.
func (m *MockClient) GetServiceQuota(arg0 *servicequotas.GetServiceQuotaInput) (*servicequotas.GetServiceQuotaOutput, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetServiceQuota", arg0)
	ret0, _ := ret[0].(*servicequotas.GetServiceQuotaOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetServiceQuota indicates an expected call of GetServiceQuota.
func (mr *MockClientMockRecorder) GetServiceQuota(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetServiceQuota", reflect.TypeOf((*MockClient)(nil).GetServiceQuota), arg0)
}

// GetServiceQuotaWithContext mocks base method.
func (m *MockClient) GetServiceQuotaWithContext(ctx aws.Context, input *servicequotas.GetServiceQuotaInput, opts ...request.Option) (*servicequotas.GetServiceQuotaOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, input}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetServiceQuotaWithContext", varargs...)
	ret0, _ := ret[0].(*servicequotas.GetServiceQuotaOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetServiceQuotaWithContext indicates an expected call of GetServiceQuotaWithContext.
func (mr *MockClientMockRecorder) GetServiceQuotaWithContext(ctx, input interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, input}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetServiceQuotaWithContext", reflect.TypeOf((*MockClient)(nil).GetServiceQuotaWithContext), varargs...)
}

// ListRequestedServiceQuotaChangeHistory mocks base method.
func (m *MockClient) ListRequestedServiceQuotaChangeHistory(arg0 *servicequotas.ListRequestedServiceQuotaChangeHistoryInput) (*servicequotas.ListRequestedServiceQuotaChangeHistoryOutput, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListRequestedServiceQuotaChangeHistory", arg0)
	ret0, _ := ret[0].(*servicequotas.ListRequestedServiceQuotaChangeHistoryOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListRequestedServiceQuotaChangeHistory indicates an expected call of ListRequestedServiceQuotaChangeHistory.
func (mr *MockClientMockRecorder) ListRequestedServiceQuotaChangeHistory(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListRequestedServiceQuotaChangeHistory", reflect.TypeOf((*MockClient)(nil).ListRequestedServiceQuotaChangeHistory), arg0)
}

// ListRequestedServiceQuotaChangeHistoryByQuota mocks base method.
func (m *MockClient) ListRequestedServiceQuotaChangeHistoryByQuota(arg0 *servicequotas.ListRequestedServiceQuotaChangeHistoryByQuotaInput) (*servicequotas.ListRequestedServiceQuotaChangeHistoryByQuotaOutput, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListRequestedServiceQuotaChangeHistoryByQuota", arg0)
	ret0, _ := ret[0].(*servicequotas.ListRequestedServiceQuotaChangeHistoryByQuotaOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListRequestedServiceQuotaChangeHistoryByQuota indicates an expected call of ListRequestedServiceQuotaChangeHistoryByQuota.
func (mr *MockClientMockRecorder) ListRequestedServiceQuotaChangeHistoryByQuota(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListRequestedServiceQuotaChangeHistoryByQuota", reflect.TypeOf((*MockClient)(nil).ListRequestedServiceQuotaChangeHistoryByQuota), arg0)
}

// RequestServiceQuotaIncrease mocks base method.
func (m *MockClient) RequestServiceQuotaIncrease(arg0 *servicequotas.RequestServiceQuotaIncreaseInput) (*servicequotas.RequestServiceQuotaIncreaseOutput, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RequestServiceQuotaIncrease", arg0)
	ret0, _ := ret[0].(*servicequotas.RequestServiceQuotaIncreaseOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// RequestServiceQuotaIncrease indicates an expected call of RequestServiceQuotaIncrease.
func (mr *MockClientMockRecorder) RequestServiceQuotaIncrease(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RequestServiceQuotaIncrease", reflect.TypeOf((*MockClient)(nil).RequestServiceQuotaIncrease), arg0)
}
