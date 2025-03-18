package mocks_test

import (
	"context"
	"testing"

	"aws_challenge_pragma/infrastructure/mocks"

	"github.com/aws/aws-sdk-go-v2/service/sts"
	"github.com/stretchr/testify/assert"
)

func TestMockSTSClient_AssumeRole(t *testing.T) {
	mockSTS := new(mocks.MockSTSClient)
	ctx := context.TODO()
	input := &sts.AssumeRoleInput{}
	expectedOutput := &sts.AssumeRoleOutput{}

	mockSTS.On("AssumeRole", ctx, input).Return(expectedOutput, nil)
	result, err := mockSTS.AssumeRole(ctx, input)

	assert.NoError(t, err)
	assert.Equal(t, expectedOutput, result)
	mockSTS.AssertExpectations(t)
}

func TestMockSTSClient_AssumeRoleWithSAML(t *testing.T) {
	mockSTS := new(mocks.MockSTSClient)
	ctx := context.TODO()
	input := &sts.AssumeRoleWithSAMLInput{}
	expectedOutput := &sts.AssumeRoleWithSAMLOutput{}

	mockSTS.On("AssumeRoleWithSAML", ctx, input).Return(expectedOutput, nil)
	result, err := mockSTS.AssumeRoleWithSAML(ctx, input)

	assert.NoError(t, err)
	assert.Equal(t, expectedOutput, result)
	mockSTS.AssertExpectations(t)
}

func TestMockSTSClient_AssumeRoleWithWebIdentity(t *testing.T) {
	mockSTS := new(mocks.MockSTSClient)
	ctx := context.TODO()
	input := &sts.AssumeRoleWithWebIdentityInput{}
	expectedOutput := &sts.AssumeRoleWithWebIdentityOutput{}

	mockSTS.On("AssumeRoleWithWebIdentity", ctx, input).Return(expectedOutput, nil)
	result, err := mockSTS.AssumeRoleWithWebIdentity(ctx, input)

	assert.NoError(t, err)
	assert.Equal(t, expectedOutput, result)
	mockSTS.AssertExpectations(t)
}

func TestMockSTSClient_GetCallerIdentity(t *testing.T) {
	mockSTS := new(mocks.MockSTSClient)
	ctx := context.TODO()
	input := &sts.GetCallerIdentityInput{}
	expectedOutput := &sts.GetCallerIdentityOutput{}

	mockSTS.On("GetCallerIdentity", ctx, input).Return(expectedOutput, nil)
	result, err := mockSTS.GetCallerIdentity(ctx, input)

	assert.NoError(t, err)
	assert.Equal(t, expectedOutput, result)
	mockSTS.AssertExpectations(t)
}

func TestMockSTSClient_GetSessionToken(t *testing.T) {
	mockSTS := new(mocks.MockSTSClient)
	ctx := context.TODO()
	input := &sts.GetSessionTokenInput{}
	expectedOutput := &sts.GetSessionTokenOutput{}

	mockSTS.On("GetSessionToken", ctx, input).Return(expectedOutput, nil)
	result, err := mockSTS.GetSessionToken(ctx, input)

	assert.NoError(t, err)
	assert.Equal(t, expectedOutput, result)
	mockSTS.AssertExpectations(t)
}

func TestMockSTSClient_GetFederationToken(t *testing.T) {
	mockSTS := new(mocks.MockSTSClient)
	ctx := context.TODO()
	input := &sts.GetFederationTokenInput{}
	expectedOutput := &sts.GetFederationTokenOutput{}

	mockSTS.On("GetFederationToken", ctx, input).Return(expectedOutput, nil)
	result, err := mockSTS.GetFederationToken(ctx, input)

	assert.NoError(t, err)
	assert.Equal(t, expectedOutput, result)
	mockSTS.AssertExpectations(t)
}

func TestMockSTSClient_DecodeAuthorizationMessage(t *testing.T) {
	mockSTS := new(mocks.MockSTSClient)
	ctx := context.TODO()
	input := &sts.DecodeAuthorizationMessageInput{}
	expectedOutput := &sts.DecodeAuthorizationMessageOutput{}

	mockSTS.On("DecodeAuthorizationMessage", ctx, input).Return(expectedOutput, nil)
	result, err := mockSTS.DecodeAuthorizationMessage(ctx, input)

	assert.NoError(t, err)
	assert.Equal(t, expectedOutput, result)
	mockSTS.AssertExpectations(t)
}

func TestMockSTSClient_GetAccessKeyInfo(t *testing.T) {
	mockSTS := new(mocks.MockSTSClient)
	ctx := context.TODO()
	input := &sts.GetAccessKeyInfoInput{}
	expectedOutput := &sts.GetAccessKeyInfoOutput{}

	mockSTS.On("GetAccessKeyInfo", ctx, input).Return(expectedOutput, nil)
	result, err := mockSTS.GetAccessKeyInfo(ctx, input)

	assert.NoError(t, err)
	assert.Equal(t, expectedOutput, result)
	mockSTS.AssertExpectations(t)
}

func TestMockSTSClient_Options(t *testing.T) {
	mockSTS := new(mocks.MockSTSClient)
	expectedOptions := sts.Options{}

	mockSTS.On("Options").Return(expectedOptions)
	result := mockSTS.Options()

	assert.Equal(t, expectedOptions, result)
	mockSTS.AssertExpectations(t)
}

func TestMockSTSClient_AssumeRoot(t *testing.T) {
	mockSTS := new(mocks.MockSTSClient)
	ctx := context.TODO()
	input := &sts.AssumeRootInput{}
	expectedOutput := &sts.AssumeRootOutput{}

	mockSTS.On("AssumeRoot", ctx, input).Return(expectedOutput, nil)
	result, err := mockSTS.AssumeRoot(ctx, input)

	assert.NoError(t, err)
	assert.Equal(t, expectedOutput, result)
	mockSTS.AssertExpectations(t)
}
