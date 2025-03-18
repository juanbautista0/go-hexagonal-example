package mocks

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/service/sts"
	"github.com/stretchr/testify/mock"
)

type MockSTSClient struct {
	mock.Mock
}

func (m *MockSTSClient) AssumeRole(ctx context.Context, params *sts.AssumeRoleInput, optFns ...func(*sts.Options)) (*sts.AssumeRoleOutput, error) {
	args := m.Called(ctx, params)
	return args.Get(0).(*sts.AssumeRoleOutput), args.Error(1)
}

func (m *MockSTSClient) AssumeRoleWithSAML(ctx context.Context, params *sts.AssumeRoleWithSAMLInput, optFns ...func(*sts.Options)) (*sts.AssumeRoleWithSAMLOutput, error) {
	args := m.Called(ctx, params)
	return args.Get(0).(*sts.AssumeRoleWithSAMLOutput), args.Error(1)
}

func (m *MockSTSClient) AssumeRoleWithWebIdentity(ctx context.Context, params *sts.AssumeRoleWithWebIdentityInput, optFns ...func(*sts.Options)) (*sts.AssumeRoleWithWebIdentityOutput, error) {
	args := m.Called(ctx, params)
	return args.Get(0).(*sts.AssumeRoleWithWebIdentityOutput), args.Error(1)
}

func (m *MockSTSClient) AssumeRoot(ctx context.Context, params *sts.AssumeRootInput, optFns ...func(*sts.Options)) (*sts.AssumeRootOutput, error) {
	args := m.Called(ctx, params)
	return args.Get(0).(*sts.AssumeRootOutput), args.Error(1)
}

func (m *MockSTSClient) DecodeAuthorizationMessage(ctx context.Context, params *sts.DecodeAuthorizationMessageInput, optFns ...func(*sts.Options)) (*sts.DecodeAuthorizationMessageOutput, error) {
	args := m.Called(ctx, params)
	return args.Get(0).(*sts.DecodeAuthorizationMessageOutput), args.Error(1)
}

func (m *MockSTSClient) GetAccessKeyInfo(ctx context.Context, params *sts.GetAccessKeyInfoInput, optFns ...func(*sts.Options)) (*sts.GetAccessKeyInfoOutput, error) {
	args := m.Called(ctx, params)
	return args.Get(0).(*sts.GetAccessKeyInfoOutput), args.Error(1)
}

func (m *MockSTSClient) GetCallerIdentity(ctx context.Context, params *sts.GetCallerIdentityInput, optFns ...func(*sts.Options)) (*sts.GetCallerIdentityOutput, error) {
	args := m.Called(ctx, params)
	return args.Get(0).(*sts.GetCallerIdentityOutput), args.Error(1)
}

func (m *MockSTSClient) GetFederationToken(ctx context.Context, params *sts.GetFederationTokenInput, optFns ...func(*sts.Options)) (*sts.GetFederationTokenOutput, error) {
	args := m.Called(ctx, params)
	return args.Get(0).(*sts.GetFederationTokenOutput), args.Error(1)
}

func (m *MockSTSClient) GetSessionToken(ctx context.Context, params *sts.GetSessionTokenInput, optFns ...func(*sts.Options)) (*sts.GetSessionTokenOutput, error) {
	args := m.Called(ctx, params)
	return args.Get(0).(*sts.GetSessionTokenOutput), args.Error(1)
}

func (m *MockSTSClient) Options() sts.Options {
	args := m.Called()
	return args.Get(0).(sts.Options)
}
