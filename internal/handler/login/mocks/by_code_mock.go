// Code generated by MockGen. DO NOT EDIT.
// Source: by_code.go
//
// Generated by this command:
//
//	mockgen -destination=./mocks/by_code_mock.go -source=by_code.go
//

// Package mock_login is a generated GoMock package.
package mock_login

import (
	context "context"
	models "pets-backend/internal/models"
	reflect "reflect"

	gomock "go.uber.org/mock/gomock"
)

// MockByCodeOtpService is a mock of ByCodeOtpService interface.
type MockByCodeOtpService struct {
	ctrl     *gomock.Controller
	recorder *MockByCodeOtpServiceMockRecorder
}

// MockByCodeOtpServiceMockRecorder is the mock recorder for MockByCodeOtpService.
type MockByCodeOtpServiceMockRecorder struct {
	mock *MockByCodeOtpService
}

// NewMockByCodeOtpService creates a new mock instance.
func NewMockByCodeOtpService(ctrl *gomock.Controller) *MockByCodeOtpService {
	mock := &MockByCodeOtpService{ctrl: ctrl}
	mock.recorder = &MockByCodeOtpServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockByCodeOtpService) EXPECT() *MockByCodeOtpServiceMockRecorder {
	return m.recorder
}

// VerifyCode mocks base method.
func (m *MockByCodeOtpService) VerifyCode(ctx context.Context, email, code string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "VerifyCode", ctx, email, code)
	ret0, _ := ret[0].(error)
	return ret0
}

// VerifyCode indicates an expected call of VerifyCode.
func (mr *MockByCodeOtpServiceMockRecorder) VerifyCode(ctx, email, code any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "VerifyCode", reflect.TypeOf((*MockByCodeOtpService)(nil).VerifyCode), ctx, email, code)
}

// MockByCodeUserService is a mock of ByCodeUserService interface.
type MockByCodeUserService struct {
	ctrl     *gomock.Controller
	recorder *MockByCodeUserServiceMockRecorder
}

// MockByCodeUserServiceMockRecorder is the mock recorder for MockByCodeUserService.
type MockByCodeUserServiceMockRecorder struct {
	mock *MockByCodeUserService
}

// NewMockByCodeUserService creates a new mock instance.
func NewMockByCodeUserService(ctrl *gomock.Controller) *MockByCodeUserService {
	mock := &MockByCodeUserService{ctrl: ctrl}
	mock.recorder = &MockByCodeUserServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockByCodeUserService) EXPECT() *MockByCodeUserServiceMockRecorder {
	return m.recorder
}

// GetByEmail mocks base method.
func (m *MockByCodeUserService) GetByEmail(ctx context.Context, email string) (*models.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetByEmail", ctx, email)
	ret0, _ := ret[0].(*models.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetByEmail indicates an expected call of GetByEmail.
func (mr *MockByCodeUserServiceMockRecorder) GetByEmail(ctx, email any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetByEmail", reflect.TypeOf((*MockByCodeUserService)(nil).GetByEmail), ctx, email)
}

// MockByCodeSSOService is a mock of ByCodeSSOService interface.
type MockByCodeSSOService struct {
	ctrl     *gomock.Controller
	recorder *MockByCodeSSOServiceMockRecorder
}

// MockByCodeSSOServiceMockRecorder is the mock recorder for MockByCodeSSOService.
type MockByCodeSSOServiceMockRecorder struct {
	mock *MockByCodeSSOService
}

// NewMockByCodeSSOService creates a new mock instance.
func NewMockByCodeSSOService(ctrl *gomock.Controller) *MockByCodeSSOService {
	mock := &MockByCodeSSOService{ctrl: ctrl}
	mock.recorder = &MockByCodeSSOServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockByCodeSSOService) EXPECT() *MockByCodeSSOServiceMockRecorder {
	return m.recorder
}

// GenerateToken mocks base method.
func (m *MockByCodeSSOService) GenerateToken(userID string) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GenerateToken", userID)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GenerateToken indicates an expected call of GenerateToken.
func (mr *MockByCodeSSOServiceMockRecorder) GenerateToken(userID any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GenerateToken", reflect.TypeOf((*MockByCodeSSOService)(nil).GenerateToken), userID)
}
