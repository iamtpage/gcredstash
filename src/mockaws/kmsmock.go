// Automatically generated by MockGen. DO NOT EDIT!
// Source: /Users/sugawara/.go/src/github.com/aws/aws-sdk-go/service/kms/kmsiface/interface.go

package mockaws

import (
	request "github.com/aws/aws-sdk-go/aws/request"
	kms "github.com/aws/aws-sdk-go/service/kms"
	gomock "github.com/golang/mock/gomock"
)

// Mock of KMSAPI interface
type MockKMSAPI struct {
	ctrl     *gomock.Controller
	recorder *_MockKMSAPIRecorder
}

// Recorder for MockKMSAPI (not exported)
type _MockKMSAPIRecorder struct {
	mock *MockKMSAPI
}

func NewMockKMSAPI(ctrl *gomock.Controller) *MockKMSAPI {
	mock := &MockKMSAPI{ctrl: ctrl}
	mock.recorder = &_MockKMSAPIRecorder{mock}
	return mock
}

func (_m *MockKMSAPI) EXPECT() *_MockKMSAPIRecorder {
	return _m.recorder
}

func (_m *MockKMSAPI) CancelKeyDeletionRequest(_param0 *kms.CancelKeyDeletionInput) (*request.Request, *kms.CancelKeyDeletionOutput) {
	ret := _m.ctrl.Call(_m, "CancelKeyDeletionRequest", _param0)
	ret0, _ := ret[0].(*request.Request)
	ret1, _ := ret[1].(*kms.CancelKeyDeletionOutput)
	return ret0, ret1
}

func (_mr *_MockKMSAPIRecorder) CancelKeyDeletionRequest(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "CancelKeyDeletionRequest", arg0)
}

func (_m *MockKMSAPI) CancelKeyDeletion(_param0 *kms.CancelKeyDeletionInput) (*kms.CancelKeyDeletionOutput, error) {
	ret := _m.ctrl.Call(_m, "CancelKeyDeletion", _param0)
	ret0, _ := ret[0].(*kms.CancelKeyDeletionOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (_mr *_MockKMSAPIRecorder) CancelKeyDeletion(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "CancelKeyDeletion", arg0)
}

func (_m *MockKMSAPI) CreateAliasRequest(_param0 *kms.CreateAliasInput) (*request.Request, *kms.CreateAliasOutput) {
	ret := _m.ctrl.Call(_m, "CreateAliasRequest", _param0)
	ret0, _ := ret[0].(*request.Request)
	ret1, _ := ret[1].(*kms.CreateAliasOutput)
	return ret0, ret1
}

func (_mr *_MockKMSAPIRecorder) CreateAliasRequest(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "CreateAliasRequest", arg0)
}

func (_m *MockKMSAPI) CreateAlias(_param0 *kms.CreateAliasInput) (*kms.CreateAliasOutput, error) {
	ret := _m.ctrl.Call(_m, "CreateAlias", _param0)
	ret0, _ := ret[0].(*kms.CreateAliasOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (_mr *_MockKMSAPIRecorder) CreateAlias(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "CreateAlias", arg0)
}

func (_m *MockKMSAPI) CreateGrantRequest(_param0 *kms.CreateGrantInput) (*request.Request, *kms.CreateGrantOutput) {
	ret := _m.ctrl.Call(_m, "CreateGrantRequest", _param0)
	ret0, _ := ret[0].(*request.Request)
	ret1, _ := ret[1].(*kms.CreateGrantOutput)
	return ret0, ret1
}

func (_mr *_MockKMSAPIRecorder) CreateGrantRequest(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "CreateGrantRequest", arg0)
}

func (_m *MockKMSAPI) CreateGrant(_param0 *kms.CreateGrantInput) (*kms.CreateGrantOutput, error) {
	ret := _m.ctrl.Call(_m, "CreateGrant", _param0)
	ret0, _ := ret[0].(*kms.CreateGrantOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (_mr *_MockKMSAPIRecorder) CreateGrant(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "CreateGrant", arg0)
}

func (_m *MockKMSAPI) CreateKeyRequest(_param0 *kms.CreateKeyInput) (*request.Request, *kms.CreateKeyOutput) {
	ret := _m.ctrl.Call(_m, "CreateKeyRequest", _param0)
	ret0, _ := ret[0].(*request.Request)
	ret1, _ := ret[1].(*kms.CreateKeyOutput)
	return ret0, ret1
}

func (_mr *_MockKMSAPIRecorder) CreateKeyRequest(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "CreateKeyRequest", arg0)
}

func (_m *MockKMSAPI) CreateKey(_param0 *kms.CreateKeyInput) (*kms.CreateKeyOutput, error) {
	ret := _m.ctrl.Call(_m, "CreateKey", _param0)
	ret0, _ := ret[0].(*kms.CreateKeyOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (_mr *_MockKMSAPIRecorder) CreateKey(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "CreateKey", arg0)
}

func (_m *MockKMSAPI) DecryptRequest(_param0 *kms.DecryptInput) (*request.Request, *kms.DecryptOutput) {
	ret := _m.ctrl.Call(_m, "DecryptRequest", _param0)
	ret0, _ := ret[0].(*request.Request)
	ret1, _ := ret[1].(*kms.DecryptOutput)
	return ret0, ret1
}

func (_mr *_MockKMSAPIRecorder) DecryptRequest(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "DecryptRequest", arg0)
}

func (_m *MockKMSAPI) Decrypt(_param0 *kms.DecryptInput) (*kms.DecryptOutput, error) {
	ret := _m.ctrl.Call(_m, "Decrypt", _param0)
	ret0, _ := ret[0].(*kms.DecryptOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (_mr *_MockKMSAPIRecorder) Decrypt(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "Decrypt", arg0)
}

func (_m *MockKMSAPI) DeleteAliasRequest(_param0 *kms.DeleteAliasInput) (*request.Request, *kms.DeleteAliasOutput) {
	ret := _m.ctrl.Call(_m, "DeleteAliasRequest", _param0)
	ret0, _ := ret[0].(*request.Request)
	ret1, _ := ret[1].(*kms.DeleteAliasOutput)
	return ret0, ret1
}

func (_mr *_MockKMSAPIRecorder) DeleteAliasRequest(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "DeleteAliasRequest", arg0)
}

func (_m *MockKMSAPI) DeleteAlias(_param0 *kms.DeleteAliasInput) (*kms.DeleteAliasOutput, error) {
	ret := _m.ctrl.Call(_m, "DeleteAlias", _param0)
	ret0, _ := ret[0].(*kms.DeleteAliasOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (_mr *_MockKMSAPIRecorder) DeleteAlias(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "DeleteAlias", arg0)
}

func (_m *MockKMSAPI) DescribeKeyRequest(_param0 *kms.DescribeKeyInput) (*request.Request, *kms.DescribeKeyOutput) {
	ret := _m.ctrl.Call(_m, "DescribeKeyRequest", _param0)
	ret0, _ := ret[0].(*request.Request)
	ret1, _ := ret[1].(*kms.DescribeKeyOutput)
	return ret0, ret1
}

func (_mr *_MockKMSAPIRecorder) DescribeKeyRequest(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "DescribeKeyRequest", arg0)
}

func (_m *MockKMSAPI) DescribeKey(_param0 *kms.DescribeKeyInput) (*kms.DescribeKeyOutput, error) {
	ret := _m.ctrl.Call(_m, "DescribeKey", _param0)
	ret0, _ := ret[0].(*kms.DescribeKeyOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (_mr *_MockKMSAPIRecorder) DescribeKey(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "DescribeKey", arg0)
}

func (_m *MockKMSAPI) DisableKeyRequest(_param0 *kms.DisableKeyInput) (*request.Request, *kms.DisableKeyOutput) {
	ret := _m.ctrl.Call(_m, "DisableKeyRequest", _param0)
	ret0, _ := ret[0].(*request.Request)
	ret1, _ := ret[1].(*kms.DisableKeyOutput)
	return ret0, ret1
}

func (_mr *_MockKMSAPIRecorder) DisableKeyRequest(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "DisableKeyRequest", arg0)
}

func (_m *MockKMSAPI) DisableKey(_param0 *kms.DisableKeyInput) (*kms.DisableKeyOutput, error) {
	ret := _m.ctrl.Call(_m, "DisableKey", _param0)
	ret0, _ := ret[0].(*kms.DisableKeyOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (_mr *_MockKMSAPIRecorder) DisableKey(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "DisableKey", arg0)
}

func (_m *MockKMSAPI) DisableKeyRotationRequest(_param0 *kms.DisableKeyRotationInput) (*request.Request, *kms.DisableKeyRotationOutput) {
	ret := _m.ctrl.Call(_m, "DisableKeyRotationRequest", _param0)
	ret0, _ := ret[0].(*request.Request)
	ret1, _ := ret[1].(*kms.DisableKeyRotationOutput)
	return ret0, ret1
}

func (_mr *_MockKMSAPIRecorder) DisableKeyRotationRequest(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "DisableKeyRotationRequest", arg0)
}

func (_m *MockKMSAPI) DisableKeyRotation(_param0 *kms.DisableKeyRotationInput) (*kms.DisableKeyRotationOutput, error) {
	ret := _m.ctrl.Call(_m, "DisableKeyRotation", _param0)
	ret0, _ := ret[0].(*kms.DisableKeyRotationOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (_mr *_MockKMSAPIRecorder) DisableKeyRotation(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "DisableKeyRotation", arg0)
}

func (_m *MockKMSAPI) EnableKeyRequest(_param0 *kms.EnableKeyInput) (*request.Request, *kms.EnableKeyOutput) {
	ret := _m.ctrl.Call(_m, "EnableKeyRequest", _param0)
	ret0, _ := ret[0].(*request.Request)
	ret1, _ := ret[1].(*kms.EnableKeyOutput)
	return ret0, ret1
}

func (_mr *_MockKMSAPIRecorder) EnableKeyRequest(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "EnableKeyRequest", arg0)
}

func (_m *MockKMSAPI) EnableKey(_param0 *kms.EnableKeyInput) (*kms.EnableKeyOutput, error) {
	ret := _m.ctrl.Call(_m, "EnableKey", _param0)
	ret0, _ := ret[0].(*kms.EnableKeyOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (_mr *_MockKMSAPIRecorder) EnableKey(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "EnableKey", arg0)
}

func (_m *MockKMSAPI) EnableKeyRotationRequest(_param0 *kms.EnableKeyRotationInput) (*request.Request, *kms.EnableKeyRotationOutput) {
	ret := _m.ctrl.Call(_m, "EnableKeyRotationRequest", _param0)
	ret0, _ := ret[0].(*request.Request)
	ret1, _ := ret[1].(*kms.EnableKeyRotationOutput)
	return ret0, ret1
}

func (_mr *_MockKMSAPIRecorder) EnableKeyRotationRequest(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "EnableKeyRotationRequest", arg0)
}

func (_m *MockKMSAPI) EnableKeyRotation(_param0 *kms.EnableKeyRotationInput) (*kms.EnableKeyRotationOutput, error) {
	ret := _m.ctrl.Call(_m, "EnableKeyRotation", _param0)
	ret0, _ := ret[0].(*kms.EnableKeyRotationOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (_mr *_MockKMSAPIRecorder) EnableKeyRotation(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "EnableKeyRotation", arg0)
}

func (_m *MockKMSAPI) EncryptRequest(_param0 *kms.EncryptInput) (*request.Request, *kms.EncryptOutput) {
	ret := _m.ctrl.Call(_m, "EncryptRequest", _param0)
	ret0, _ := ret[0].(*request.Request)
	ret1, _ := ret[1].(*kms.EncryptOutput)
	return ret0, ret1
}

func (_mr *_MockKMSAPIRecorder) EncryptRequest(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "EncryptRequest", arg0)
}

func (_m *MockKMSAPI) Encrypt(_param0 *kms.EncryptInput) (*kms.EncryptOutput, error) {
	ret := _m.ctrl.Call(_m, "Encrypt", _param0)
	ret0, _ := ret[0].(*kms.EncryptOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (_mr *_MockKMSAPIRecorder) Encrypt(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "Encrypt", arg0)
}

func (_m *MockKMSAPI) GenerateDataKeyRequest(_param0 *kms.GenerateDataKeyInput) (*request.Request, *kms.GenerateDataKeyOutput) {
	ret := _m.ctrl.Call(_m, "GenerateDataKeyRequest", _param0)
	ret0, _ := ret[0].(*request.Request)
	ret1, _ := ret[1].(*kms.GenerateDataKeyOutput)
	return ret0, ret1
}

func (_mr *_MockKMSAPIRecorder) GenerateDataKeyRequest(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "GenerateDataKeyRequest", arg0)
}

func (_m *MockKMSAPI) GenerateDataKey(_param0 *kms.GenerateDataKeyInput) (*kms.GenerateDataKeyOutput, error) {
	ret := _m.ctrl.Call(_m, "GenerateDataKey", _param0)
	ret0, _ := ret[0].(*kms.GenerateDataKeyOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (_mr *_MockKMSAPIRecorder) GenerateDataKey(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "GenerateDataKey", arg0)
}

func (_m *MockKMSAPI) GenerateDataKeyWithoutPlaintextRequest(_param0 *kms.GenerateDataKeyWithoutPlaintextInput) (*request.Request, *kms.GenerateDataKeyWithoutPlaintextOutput) {
	ret := _m.ctrl.Call(_m, "GenerateDataKeyWithoutPlaintextRequest", _param0)
	ret0, _ := ret[0].(*request.Request)
	ret1, _ := ret[1].(*kms.GenerateDataKeyWithoutPlaintextOutput)
	return ret0, ret1
}

func (_mr *_MockKMSAPIRecorder) GenerateDataKeyWithoutPlaintextRequest(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "GenerateDataKeyWithoutPlaintextRequest", arg0)
}

func (_m *MockKMSAPI) GenerateDataKeyWithoutPlaintext(_param0 *kms.GenerateDataKeyWithoutPlaintextInput) (*kms.GenerateDataKeyWithoutPlaintextOutput, error) {
	ret := _m.ctrl.Call(_m, "GenerateDataKeyWithoutPlaintext", _param0)
	ret0, _ := ret[0].(*kms.GenerateDataKeyWithoutPlaintextOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (_mr *_MockKMSAPIRecorder) GenerateDataKeyWithoutPlaintext(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "GenerateDataKeyWithoutPlaintext", arg0)
}

func (_m *MockKMSAPI) GenerateRandomRequest(_param0 *kms.GenerateRandomInput) (*request.Request, *kms.GenerateRandomOutput) {
	ret := _m.ctrl.Call(_m, "GenerateRandomRequest", _param0)
	ret0, _ := ret[0].(*request.Request)
	ret1, _ := ret[1].(*kms.GenerateRandomOutput)
	return ret0, ret1
}

func (_mr *_MockKMSAPIRecorder) GenerateRandomRequest(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "GenerateRandomRequest", arg0)
}

func (_m *MockKMSAPI) GenerateRandom(_param0 *kms.GenerateRandomInput) (*kms.GenerateRandomOutput, error) {
	ret := _m.ctrl.Call(_m, "GenerateRandom", _param0)
	ret0, _ := ret[0].(*kms.GenerateRandomOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (_mr *_MockKMSAPIRecorder) GenerateRandom(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "GenerateRandom", arg0)
}

func (_m *MockKMSAPI) GetKeyPolicyRequest(_param0 *kms.GetKeyPolicyInput) (*request.Request, *kms.GetKeyPolicyOutput) {
	ret := _m.ctrl.Call(_m, "GetKeyPolicyRequest", _param0)
	ret0, _ := ret[0].(*request.Request)
	ret1, _ := ret[1].(*kms.GetKeyPolicyOutput)
	return ret0, ret1
}

func (_mr *_MockKMSAPIRecorder) GetKeyPolicyRequest(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "GetKeyPolicyRequest", arg0)
}

func (_m *MockKMSAPI) GetKeyPolicy(_param0 *kms.GetKeyPolicyInput) (*kms.GetKeyPolicyOutput, error) {
	ret := _m.ctrl.Call(_m, "GetKeyPolicy", _param0)
	ret0, _ := ret[0].(*kms.GetKeyPolicyOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (_mr *_MockKMSAPIRecorder) GetKeyPolicy(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "GetKeyPolicy", arg0)
}

func (_m *MockKMSAPI) GetKeyRotationStatusRequest(_param0 *kms.GetKeyRotationStatusInput) (*request.Request, *kms.GetKeyRotationStatusOutput) {
	ret := _m.ctrl.Call(_m, "GetKeyRotationStatusRequest", _param0)
	ret0, _ := ret[0].(*request.Request)
	ret1, _ := ret[1].(*kms.GetKeyRotationStatusOutput)
	return ret0, ret1
}

func (_mr *_MockKMSAPIRecorder) GetKeyRotationStatusRequest(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "GetKeyRotationStatusRequest", arg0)
}

func (_m *MockKMSAPI) GetKeyRotationStatus(_param0 *kms.GetKeyRotationStatusInput) (*kms.GetKeyRotationStatusOutput, error) {
	ret := _m.ctrl.Call(_m, "GetKeyRotationStatus", _param0)
	ret0, _ := ret[0].(*kms.GetKeyRotationStatusOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (_mr *_MockKMSAPIRecorder) GetKeyRotationStatus(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "GetKeyRotationStatus", arg0)
}

func (_m *MockKMSAPI) ListAliasesRequest(_param0 *kms.ListAliasesInput) (*request.Request, *kms.ListAliasesOutput) {
	ret := _m.ctrl.Call(_m, "ListAliasesRequest", _param0)
	ret0, _ := ret[0].(*request.Request)
	ret1, _ := ret[1].(*kms.ListAliasesOutput)
	return ret0, ret1
}

func (_mr *_MockKMSAPIRecorder) ListAliasesRequest(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "ListAliasesRequest", arg0)
}

func (_m *MockKMSAPI) ListAliases(_param0 *kms.ListAliasesInput) (*kms.ListAliasesOutput, error) {
	ret := _m.ctrl.Call(_m, "ListAliases", _param0)
	ret0, _ := ret[0].(*kms.ListAliasesOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (_mr *_MockKMSAPIRecorder) ListAliases(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "ListAliases", arg0)
}

func (_m *MockKMSAPI) ListAliasesPages(_param0 *kms.ListAliasesInput, _param1 func(*kms.ListAliasesOutput, bool) bool) error {
	ret := _m.ctrl.Call(_m, "ListAliasesPages", _param0, _param1)
	ret0, _ := ret[0].(error)
	return ret0
}

func (_mr *_MockKMSAPIRecorder) ListAliasesPages(arg0, arg1 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "ListAliasesPages", arg0, arg1)
}

func (_m *MockKMSAPI) ListGrantsRequest(_param0 *kms.ListGrantsInput) (*request.Request, *kms.ListGrantsResponse) {
	ret := _m.ctrl.Call(_m, "ListGrantsRequest", _param0)
	ret0, _ := ret[0].(*request.Request)
	ret1, _ := ret[1].(*kms.ListGrantsResponse)
	return ret0, ret1
}

func (_mr *_MockKMSAPIRecorder) ListGrantsRequest(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "ListGrantsRequest", arg0)
}

func (_m *MockKMSAPI) ListGrants(_param0 *kms.ListGrantsInput) (*kms.ListGrantsResponse, error) {
	ret := _m.ctrl.Call(_m, "ListGrants", _param0)
	ret0, _ := ret[0].(*kms.ListGrantsResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (_mr *_MockKMSAPIRecorder) ListGrants(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "ListGrants", arg0)
}

func (_m *MockKMSAPI) ListGrantsPages(_param0 *kms.ListGrantsInput, _param1 func(*kms.ListGrantsResponse, bool) bool) error {
	ret := _m.ctrl.Call(_m, "ListGrantsPages", _param0, _param1)
	ret0, _ := ret[0].(error)
	return ret0
}

func (_mr *_MockKMSAPIRecorder) ListGrantsPages(arg0, arg1 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "ListGrantsPages", arg0, arg1)
}

func (_m *MockKMSAPI) ListKeyPoliciesRequest(_param0 *kms.ListKeyPoliciesInput) (*request.Request, *kms.ListKeyPoliciesOutput) {
	ret := _m.ctrl.Call(_m, "ListKeyPoliciesRequest", _param0)
	ret0, _ := ret[0].(*request.Request)
	ret1, _ := ret[1].(*kms.ListKeyPoliciesOutput)
	return ret0, ret1
}

func (_mr *_MockKMSAPIRecorder) ListKeyPoliciesRequest(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "ListKeyPoliciesRequest", arg0)
}

func (_m *MockKMSAPI) ListKeyPolicies(_param0 *kms.ListKeyPoliciesInput) (*kms.ListKeyPoliciesOutput, error) {
	ret := _m.ctrl.Call(_m, "ListKeyPolicies", _param0)
	ret0, _ := ret[0].(*kms.ListKeyPoliciesOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (_mr *_MockKMSAPIRecorder) ListKeyPolicies(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "ListKeyPolicies", arg0)
}

func (_m *MockKMSAPI) ListKeyPoliciesPages(_param0 *kms.ListKeyPoliciesInput, _param1 func(*kms.ListKeyPoliciesOutput, bool) bool) error {
	ret := _m.ctrl.Call(_m, "ListKeyPoliciesPages", _param0, _param1)
	ret0, _ := ret[0].(error)
	return ret0
}

func (_mr *_MockKMSAPIRecorder) ListKeyPoliciesPages(arg0, arg1 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "ListKeyPoliciesPages", arg0, arg1)
}

func (_m *MockKMSAPI) ListKeysRequest(_param0 *kms.ListKeysInput) (*request.Request, *kms.ListKeysOutput) {
	ret := _m.ctrl.Call(_m, "ListKeysRequest", _param0)
	ret0, _ := ret[0].(*request.Request)
	ret1, _ := ret[1].(*kms.ListKeysOutput)
	return ret0, ret1
}

func (_mr *_MockKMSAPIRecorder) ListKeysRequest(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "ListKeysRequest", arg0)
}

func (_m *MockKMSAPI) ListKeys(_param0 *kms.ListKeysInput) (*kms.ListKeysOutput, error) {
	ret := _m.ctrl.Call(_m, "ListKeys", _param0)
	ret0, _ := ret[0].(*kms.ListKeysOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (_mr *_MockKMSAPIRecorder) ListKeys(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "ListKeys", arg0)
}

func (_m *MockKMSAPI) ListKeysPages(_param0 *kms.ListKeysInput, _param1 func(*kms.ListKeysOutput, bool) bool) error {
	ret := _m.ctrl.Call(_m, "ListKeysPages", _param0, _param1)
	ret0, _ := ret[0].(error)
	return ret0
}

func (_mr *_MockKMSAPIRecorder) ListKeysPages(arg0, arg1 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "ListKeysPages", arg0, arg1)
}

func (_m *MockKMSAPI) ListRetirableGrantsRequest(_param0 *kms.ListRetirableGrantsInput) (*request.Request, *kms.ListGrantsResponse) {
	ret := _m.ctrl.Call(_m, "ListRetirableGrantsRequest", _param0)
	ret0, _ := ret[0].(*request.Request)
	ret1, _ := ret[1].(*kms.ListGrantsResponse)
	return ret0, ret1
}

func (_mr *_MockKMSAPIRecorder) ListRetirableGrantsRequest(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "ListRetirableGrantsRequest", arg0)
}

func (_m *MockKMSAPI) ListRetirableGrants(_param0 *kms.ListRetirableGrantsInput) (*kms.ListGrantsResponse, error) {
	ret := _m.ctrl.Call(_m, "ListRetirableGrants", _param0)
	ret0, _ := ret[0].(*kms.ListGrantsResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (_mr *_MockKMSAPIRecorder) ListRetirableGrants(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "ListRetirableGrants", arg0)
}

func (_m *MockKMSAPI) PutKeyPolicyRequest(_param0 *kms.PutKeyPolicyInput) (*request.Request, *kms.PutKeyPolicyOutput) {
	ret := _m.ctrl.Call(_m, "PutKeyPolicyRequest", _param0)
	ret0, _ := ret[0].(*request.Request)
	ret1, _ := ret[1].(*kms.PutKeyPolicyOutput)
	return ret0, ret1
}

func (_mr *_MockKMSAPIRecorder) PutKeyPolicyRequest(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "PutKeyPolicyRequest", arg0)
}

func (_m *MockKMSAPI) PutKeyPolicy(_param0 *kms.PutKeyPolicyInput) (*kms.PutKeyPolicyOutput, error) {
	ret := _m.ctrl.Call(_m, "PutKeyPolicy", _param0)
	ret0, _ := ret[0].(*kms.PutKeyPolicyOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (_mr *_MockKMSAPIRecorder) PutKeyPolicy(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "PutKeyPolicy", arg0)
}

func (_m *MockKMSAPI) ReEncryptRequest(_param0 *kms.ReEncryptInput) (*request.Request, *kms.ReEncryptOutput) {
	ret := _m.ctrl.Call(_m, "ReEncryptRequest", _param0)
	ret0, _ := ret[0].(*request.Request)
	ret1, _ := ret[1].(*kms.ReEncryptOutput)
	return ret0, ret1
}

func (_mr *_MockKMSAPIRecorder) ReEncryptRequest(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "ReEncryptRequest", arg0)
}

func (_m *MockKMSAPI) ReEncrypt(_param0 *kms.ReEncryptInput) (*kms.ReEncryptOutput, error) {
	ret := _m.ctrl.Call(_m, "ReEncrypt", _param0)
	ret0, _ := ret[0].(*kms.ReEncryptOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (_mr *_MockKMSAPIRecorder) ReEncrypt(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "ReEncrypt", arg0)
}

func (_m *MockKMSAPI) RetireGrantRequest(_param0 *kms.RetireGrantInput) (*request.Request, *kms.RetireGrantOutput) {
	ret := _m.ctrl.Call(_m, "RetireGrantRequest", _param0)
	ret0, _ := ret[0].(*request.Request)
	ret1, _ := ret[1].(*kms.RetireGrantOutput)
	return ret0, ret1
}

func (_mr *_MockKMSAPIRecorder) RetireGrantRequest(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "RetireGrantRequest", arg0)
}

func (_m *MockKMSAPI) RetireGrant(_param0 *kms.RetireGrantInput) (*kms.RetireGrantOutput, error) {
	ret := _m.ctrl.Call(_m, "RetireGrant", _param0)
	ret0, _ := ret[0].(*kms.RetireGrantOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (_mr *_MockKMSAPIRecorder) RetireGrant(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "RetireGrant", arg0)
}

func (_m *MockKMSAPI) RevokeGrantRequest(_param0 *kms.RevokeGrantInput) (*request.Request, *kms.RevokeGrantOutput) {
	ret := _m.ctrl.Call(_m, "RevokeGrantRequest", _param0)
	ret0, _ := ret[0].(*request.Request)
	ret1, _ := ret[1].(*kms.RevokeGrantOutput)
	return ret0, ret1
}

func (_mr *_MockKMSAPIRecorder) RevokeGrantRequest(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "RevokeGrantRequest", arg0)
}

func (_m *MockKMSAPI) RevokeGrant(_param0 *kms.RevokeGrantInput) (*kms.RevokeGrantOutput, error) {
	ret := _m.ctrl.Call(_m, "RevokeGrant", _param0)
	ret0, _ := ret[0].(*kms.RevokeGrantOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (_mr *_MockKMSAPIRecorder) RevokeGrant(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "RevokeGrant", arg0)
}

func (_m *MockKMSAPI) ScheduleKeyDeletionRequest(_param0 *kms.ScheduleKeyDeletionInput) (*request.Request, *kms.ScheduleKeyDeletionOutput) {
	ret := _m.ctrl.Call(_m, "ScheduleKeyDeletionRequest", _param0)
	ret0, _ := ret[0].(*request.Request)
	ret1, _ := ret[1].(*kms.ScheduleKeyDeletionOutput)
	return ret0, ret1
}

func (_mr *_MockKMSAPIRecorder) ScheduleKeyDeletionRequest(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "ScheduleKeyDeletionRequest", arg0)
}

func (_m *MockKMSAPI) ScheduleKeyDeletion(_param0 *kms.ScheduleKeyDeletionInput) (*kms.ScheduleKeyDeletionOutput, error) {
	ret := _m.ctrl.Call(_m, "ScheduleKeyDeletion", _param0)
	ret0, _ := ret[0].(*kms.ScheduleKeyDeletionOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (_mr *_MockKMSAPIRecorder) ScheduleKeyDeletion(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "ScheduleKeyDeletion", arg0)
}

func (_m *MockKMSAPI) UpdateAliasRequest(_param0 *kms.UpdateAliasInput) (*request.Request, *kms.UpdateAliasOutput) {
	ret := _m.ctrl.Call(_m, "UpdateAliasRequest", _param0)
	ret0, _ := ret[0].(*request.Request)
	ret1, _ := ret[1].(*kms.UpdateAliasOutput)
	return ret0, ret1
}

func (_mr *_MockKMSAPIRecorder) UpdateAliasRequest(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "UpdateAliasRequest", arg0)
}

func (_m *MockKMSAPI) UpdateAlias(_param0 *kms.UpdateAliasInput) (*kms.UpdateAliasOutput, error) {
	ret := _m.ctrl.Call(_m, "UpdateAlias", _param0)
	ret0, _ := ret[0].(*kms.UpdateAliasOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (_mr *_MockKMSAPIRecorder) UpdateAlias(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "UpdateAlias", arg0)
}

func (_m *MockKMSAPI) UpdateKeyDescriptionRequest(_param0 *kms.UpdateKeyDescriptionInput) (*request.Request, *kms.UpdateKeyDescriptionOutput) {
	ret := _m.ctrl.Call(_m, "UpdateKeyDescriptionRequest", _param0)
	ret0, _ := ret[0].(*request.Request)
	ret1, _ := ret[1].(*kms.UpdateKeyDescriptionOutput)
	return ret0, ret1
}

func (_mr *_MockKMSAPIRecorder) UpdateKeyDescriptionRequest(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "UpdateKeyDescriptionRequest", arg0)
}

func (_m *MockKMSAPI) UpdateKeyDescription(_param0 *kms.UpdateKeyDescriptionInput) (*kms.UpdateKeyDescriptionOutput, error) {
	ret := _m.ctrl.Call(_m, "UpdateKeyDescription", _param0)
	ret0, _ := ret[0].(*kms.UpdateKeyDescriptionOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (_mr *_MockKMSAPIRecorder) UpdateKeyDescription(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "UpdateKeyDescription", arg0)
}