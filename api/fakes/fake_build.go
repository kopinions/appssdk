// Code generated by counterfeiter. DO NOT EDIT.
package fakes

import (
	"sync"

	"github.com/cnupp/appssdk/api"
)

type FakeBuild struct {
	IdStub        func() string
	idMutex       sync.RWMutex
	idArgsForCall []struct{}
	idReturns     struct {
		result1 string
	}
	idReturnsOnCall map[int]struct {
		result1 string
	}
	GitShaStub        func() string
	gitShaMutex       sync.RWMutex
	gitShaArgsForCall []struct{}
	gitShaReturns     struct {
		result1 string
	}
	gitShaReturnsOnCall map[int]struct {
		result1 string
	}
	StatusStub        func() string
	statusMutex       sync.RWMutex
	statusArgsForCall []struct{}
	statusReturns     struct {
		result1 string
	}
	statusReturnsOnCall map[int]struct {
		result1 string
	}
	VerifyStub        func() api.Verify
	verifyMutex       sync.RWMutex
	verifyArgsForCall []struct{}
	verifyReturns     struct {
		result1 api.Verify
	}
	verifyReturnsOnCall map[int]struct {
		result1 api.Verify
	}
	LinksStub        func() api.Links
	linksMutex       sync.RWMutex
	linksArgsForCall []struct{}
	linksReturns     struct {
		result1 api.Links
	}
	linksReturnsOnCall map[int]struct {
		result1 api.Links
	}
	GetAppStub        func() api.App
	getAppMutex       sync.RWMutex
	getAppArgsForCall []struct{}
	getAppReturns     struct {
		result1 api.App
	}
	getAppReturnsOnCall map[int]struct {
		result1 api.App
	}
	SuccessStub        func() error
	successMutex       sync.RWMutex
	successArgsForCall []struct{}
	successReturns     struct {
		result1 error
	}
	successReturnsOnCall map[int]struct {
		result1 error
	}
	FailStub        func() error
	failMutex       sync.RWMutex
	failArgsForCall []struct{}
	failReturns     struct {
		result1 error
	}
	failReturnsOnCall map[int]struct {
		result1 error
	}
	IsSuccessStub        func() bool
	isSuccessMutex       sync.RWMutex
	isSuccessArgsForCall []struct{}
	isSuccessReturns     struct {
		result1 bool
	}
	isSuccessReturnsOnCall map[int]struct {
		result1 bool
	}
	IsFailStub        func() bool
	isFailMutex       sync.RWMutex
	isFailArgsForCall []struct{}
	isFailReturns     struct {
		result1 bool
	}
	isFailReturnsOnCall map[int]struct {
		result1 bool
	}
	VerifySuccessStub        func() error
	verifySuccessMutex       sync.RWMutex
	verifySuccessArgsForCall []struct{}
	verifySuccessReturns     struct {
		result1 error
	}
	verifySuccessReturnsOnCall map[int]struct {
		result1 error
	}
	IsVerifySuccessStub        func() bool
	isVerifySuccessMutex       sync.RWMutex
	isVerifySuccessArgsForCall []struct{}
	isVerifySuccessReturns     struct {
		result1 bool
	}
	isVerifySuccessReturnsOnCall map[int]struct {
		result1 bool
	}
	VerifyFailStub        func() error
	verifyFailMutex       sync.RWMutex
	verifyFailArgsForCall []struct{}
	verifyFailReturns     struct {
		result1 error
	}
	verifyFailReturnsOnCall map[int]struct {
		result1 error
	}
	IsVerifyFailStub        func() bool
	isVerifyFailMutex       sync.RWMutex
	isVerifyFailArgsForCall []struct{}
	isVerifyFailReturns     struct {
		result1 bool
	}
	isVerifyFailReturnsOnCall map[int]struct {
		result1 bool
	}
	CreateVerifyStub        func(api.VerifyParams) (api.Verify, error)
	createVerifyMutex       sync.RWMutex
	createVerifyArgsForCall []struct {
		arg1 api.VerifyParams
	}
	createVerifyReturns struct {
		result1 api.Verify
		result2 error
	}
	createVerifyReturnsOnCall map[int]struct {
		result1 api.Verify
		result2 error
	}
	GetVerifyStub        func(id string) (api.Verify, error)
	getVerifyMutex       sync.RWMutex
	getVerifyArgsForCall []struct {
		id string
	}
	getVerifyReturns struct {
		result1 api.Verify
		result2 error
	}
	getVerifyReturnsOnCall map[int]struct {
		result1 api.Verify
		result2 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeBuild) Id() string {
	fake.idMutex.Lock()
	ret, specificReturn := fake.idReturnsOnCall[len(fake.idArgsForCall)]
	fake.idArgsForCall = append(fake.idArgsForCall, struct{}{})
	fake.recordInvocation("Id", []interface{}{})
	fake.idMutex.Unlock()
	if fake.IdStub != nil {
		return fake.IdStub()
	}
	if specificReturn {
		return ret.result1
	}
	return fake.idReturns.result1
}

func (fake *FakeBuild) IdCallCount() int {
	fake.idMutex.RLock()
	defer fake.idMutex.RUnlock()
	return len(fake.idArgsForCall)
}

func (fake *FakeBuild) IdReturns(result1 string) {
	fake.IdStub = nil
	fake.idReturns = struct {
		result1 string
	}{result1}
}

func (fake *FakeBuild) IdReturnsOnCall(i int, result1 string) {
	fake.IdStub = nil
	if fake.idReturnsOnCall == nil {
		fake.idReturnsOnCall = make(map[int]struct {
			result1 string
		})
	}
	fake.idReturnsOnCall[i] = struct {
		result1 string
	}{result1}
}

func (fake *FakeBuild) GitSha() string {
	fake.gitShaMutex.Lock()
	ret, specificReturn := fake.gitShaReturnsOnCall[len(fake.gitShaArgsForCall)]
	fake.gitShaArgsForCall = append(fake.gitShaArgsForCall, struct{}{})
	fake.recordInvocation("GitSha", []interface{}{})
	fake.gitShaMutex.Unlock()
	if fake.GitShaStub != nil {
		return fake.GitShaStub()
	}
	if specificReturn {
		return ret.result1
	}
	return fake.gitShaReturns.result1
}

func (fake *FakeBuild) GitShaCallCount() int {
	fake.gitShaMutex.RLock()
	defer fake.gitShaMutex.RUnlock()
	return len(fake.gitShaArgsForCall)
}

func (fake *FakeBuild) GitShaReturns(result1 string) {
	fake.GitShaStub = nil
	fake.gitShaReturns = struct {
		result1 string
	}{result1}
}

func (fake *FakeBuild) GitShaReturnsOnCall(i int, result1 string) {
	fake.GitShaStub = nil
	if fake.gitShaReturnsOnCall == nil {
		fake.gitShaReturnsOnCall = make(map[int]struct {
			result1 string
		})
	}
	fake.gitShaReturnsOnCall[i] = struct {
		result1 string
	}{result1}
}

func (fake *FakeBuild) Status() string {
	fake.statusMutex.Lock()
	ret, specificReturn := fake.statusReturnsOnCall[len(fake.statusArgsForCall)]
	fake.statusArgsForCall = append(fake.statusArgsForCall, struct{}{})
	fake.recordInvocation("Status", []interface{}{})
	fake.statusMutex.Unlock()
	if fake.StatusStub != nil {
		return fake.StatusStub()
	}
	if specificReturn {
		return ret.result1
	}
	return fake.statusReturns.result1
}

func (fake *FakeBuild) StatusCallCount() int {
	fake.statusMutex.RLock()
	defer fake.statusMutex.RUnlock()
	return len(fake.statusArgsForCall)
}

func (fake *FakeBuild) StatusReturns(result1 string) {
	fake.StatusStub = nil
	fake.statusReturns = struct {
		result1 string
	}{result1}
}

func (fake *FakeBuild) StatusReturnsOnCall(i int, result1 string) {
	fake.StatusStub = nil
	if fake.statusReturnsOnCall == nil {
		fake.statusReturnsOnCall = make(map[int]struct {
			result1 string
		})
	}
	fake.statusReturnsOnCall[i] = struct {
		result1 string
	}{result1}
}

func (fake *FakeBuild) Verify() api.Verify {
	fake.verifyMutex.Lock()
	ret, specificReturn := fake.verifyReturnsOnCall[len(fake.verifyArgsForCall)]
	fake.verifyArgsForCall = append(fake.verifyArgsForCall, struct{}{})
	fake.recordInvocation("Verify", []interface{}{})
	fake.verifyMutex.Unlock()
	if fake.VerifyStub != nil {
		return fake.VerifyStub()
	}
	if specificReturn {
		return ret.result1
	}
	return fake.verifyReturns.result1
}

func (fake *FakeBuild) VerifyCallCount() int {
	fake.verifyMutex.RLock()
	defer fake.verifyMutex.RUnlock()
	return len(fake.verifyArgsForCall)
}

func (fake *FakeBuild) VerifyReturns(result1 api.Verify) {
	fake.VerifyStub = nil
	fake.verifyReturns = struct {
		result1 api.Verify
	}{result1}
}

func (fake *FakeBuild) VerifyReturnsOnCall(i int, result1 api.Verify) {
	fake.VerifyStub = nil
	if fake.verifyReturnsOnCall == nil {
		fake.verifyReturnsOnCall = make(map[int]struct {
			result1 api.Verify
		})
	}
	fake.verifyReturnsOnCall[i] = struct {
		result1 api.Verify
	}{result1}
}

func (fake *FakeBuild) Links() api.Links {
	fake.linksMutex.Lock()
	ret, specificReturn := fake.linksReturnsOnCall[len(fake.linksArgsForCall)]
	fake.linksArgsForCall = append(fake.linksArgsForCall, struct{}{})
	fake.recordInvocation("Links", []interface{}{})
	fake.linksMutex.Unlock()
	if fake.LinksStub != nil {
		return fake.LinksStub()
	}
	if specificReturn {
		return ret.result1
	}
	return fake.linksReturns.result1
}

func (fake *FakeBuild) LinksCallCount() int {
	fake.linksMutex.RLock()
	defer fake.linksMutex.RUnlock()
	return len(fake.linksArgsForCall)
}

func (fake *FakeBuild) LinksReturns(result1 api.Links) {
	fake.LinksStub = nil
	fake.linksReturns = struct {
		result1 api.Links
	}{result1}
}

func (fake *FakeBuild) LinksReturnsOnCall(i int, result1 api.Links) {
	fake.LinksStub = nil
	if fake.linksReturnsOnCall == nil {
		fake.linksReturnsOnCall = make(map[int]struct {
			result1 api.Links
		})
	}
	fake.linksReturnsOnCall[i] = struct {
		result1 api.Links
	}{result1}
}

func (fake *FakeBuild) GetApp() api.App {
	fake.getAppMutex.Lock()
	ret, specificReturn := fake.getAppReturnsOnCall[len(fake.getAppArgsForCall)]
	fake.getAppArgsForCall = append(fake.getAppArgsForCall, struct{}{})
	fake.recordInvocation("GetApp", []interface{}{})
	fake.getAppMutex.Unlock()
	if fake.GetAppStub != nil {
		return fake.GetAppStub()
	}
	if specificReturn {
		return ret.result1
	}
	return fake.getAppReturns.result1
}

func (fake *FakeBuild) GetAppCallCount() int {
	fake.getAppMutex.RLock()
	defer fake.getAppMutex.RUnlock()
	return len(fake.getAppArgsForCall)
}

func (fake *FakeBuild) GetAppReturns(result1 api.App) {
	fake.GetAppStub = nil
	fake.getAppReturns = struct {
		result1 api.App
	}{result1}
}

func (fake *FakeBuild) GetAppReturnsOnCall(i int, result1 api.App) {
	fake.GetAppStub = nil
	if fake.getAppReturnsOnCall == nil {
		fake.getAppReturnsOnCall = make(map[int]struct {
			result1 api.App
		})
	}
	fake.getAppReturnsOnCall[i] = struct {
		result1 api.App
	}{result1}
}

func (fake *FakeBuild) Success() error {
	fake.successMutex.Lock()
	ret, specificReturn := fake.successReturnsOnCall[len(fake.successArgsForCall)]
	fake.successArgsForCall = append(fake.successArgsForCall, struct{}{})
	fake.recordInvocation("Success", []interface{}{})
	fake.successMutex.Unlock()
	if fake.SuccessStub != nil {
		return fake.SuccessStub()
	}
	if specificReturn {
		return ret.result1
	}
	return fake.successReturns.result1
}

func (fake *FakeBuild) SuccessCallCount() int {
	fake.successMutex.RLock()
	defer fake.successMutex.RUnlock()
	return len(fake.successArgsForCall)
}

func (fake *FakeBuild) SuccessReturns(result1 error) {
	fake.SuccessStub = nil
	fake.successReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeBuild) SuccessReturnsOnCall(i int, result1 error) {
	fake.SuccessStub = nil
	if fake.successReturnsOnCall == nil {
		fake.successReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.successReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakeBuild) Fail() error {
	fake.failMutex.Lock()
	ret, specificReturn := fake.failReturnsOnCall[len(fake.failArgsForCall)]
	fake.failArgsForCall = append(fake.failArgsForCall, struct{}{})
	fake.recordInvocation("Fail", []interface{}{})
	fake.failMutex.Unlock()
	if fake.FailStub != nil {
		return fake.FailStub()
	}
	if specificReturn {
		return ret.result1
	}
	return fake.failReturns.result1
}

func (fake *FakeBuild) FailCallCount() int {
	fake.failMutex.RLock()
	defer fake.failMutex.RUnlock()
	return len(fake.failArgsForCall)
}

func (fake *FakeBuild) FailReturns(result1 error) {
	fake.FailStub = nil
	fake.failReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeBuild) FailReturnsOnCall(i int, result1 error) {
	fake.FailStub = nil
	if fake.failReturnsOnCall == nil {
		fake.failReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.failReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakeBuild) IsSuccess() bool {
	fake.isSuccessMutex.Lock()
	ret, specificReturn := fake.isSuccessReturnsOnCall[len(fake.isSuccessArgsForCall)]
	fake.isSuccessArgsForCall = append(fake.isSuccessArgsForCall, struct{}{})
	fake.recordInvocation("IsSuccess", []interface{}{})
	fake.isSuccessMutex.Unlock()
	if fake.IsSuccessStub != nil {
		return fake.IsSuccessStub()
	}
	if specificReturn {
		return ret.result1
	}
	return fake.isSuccessReturns.result1
}

func (fake *FakeBuild) IsSuccessCallCount() int {
	fake.isSuccessMutex.RLock()
	defer fake.isSuccessMutex.RUnlock()
	return len(fake.isSuccessArgsForCall)
}

func (fake *FakeBuild) IsSuccessReturns(result1 bool) {
	fake.IsSuccessStub = nil
	fake.isSuccessReturns = struct {
		result1 bool
	}{result1}
}

func (fake *FakeBuild) IsSuccessReturnsOnCall(i int, result1 bool) {
	fake.IsSuccessStub = nil
	if fake.isSuccessReturnsOnCall == nil {
		fake.isSuccessReturnsOnCall = make(map[int]struct {
			result1 bool
		})
	}
	fake.isSuccessReturnsOnCall[i] = struct {
		result1 bool
	}{result1}
}

func (fake *FakeBuild) IsFail() bool {
	fake.isFailMutex.Lock()
	ret, specificReturn := fake.isFailReturnsOnCall[len(fake.isFailArgsForCall)]
	fake.isFailArgsForCall = append(fake.isFailArgsForCall, struct{}{})
	fake.recordInvocation("IsFail", []interface{}{})
	fake.isFailMutex.Unlock()
	if fake.IsFailStub != nil {
		return fake.IsFailStub()
	}
	if specificReturn {
		return ret.result1
	}
	return fake.isFailReturns.result1
}

func (fake *FakeBuild) IsFailCallCount() int {
	fake.isFailMutex.RLock()
	defer fake.isFailMutex.RUnlock()
	return len(fake.isFailArgsForCall)
}

func (fake *FakeBuild) IsFailReturns(result1 bool) {
	fake.IsFailStub = nil
	fake.isFailReturns = struct {
		result1 bool
	}{result1}
}

func (fake *FakeBuild) IsFailReturnsOnCall(i int, result1 bool) {
	fake.IsFailStub = nil
	if fake.isFailReturnsOnCall == nil {
		fake.isFailReturnsOnCall = make(map[int]struct {
			result1 bool
		})
	}
	fake.isFailReturnsOnCall[i] = struct {
		result1 bool
	}{result1}
}

func (fake *FakeBuild) VerifySuccess() error {
	fake.verifySuccessMutex.Lock()
	ret, specificReturn := fake.verifySuccessReturnsOnCall[len(fake.verifySuccessArgsForCall)]
	fake.verifySuccessArgsForCall = append(fake.verifySuccessArgsForCall, struct{}{})
	fake.recordInvocation("VerifySuccess", []interface{}{})
	fake.verifySuccessMutex.Unlock()
	if fake.VerifySuccessStub != nil {
		return fake.VerifySuccessStub()
	}
	if specificReturn {
		return ret.result1
	}
	return fake.verifySuccessReturns.result1
}

func (fake *FakeBuild) VerifySuccessCallCount() int {
	fake.verifySuccessMutex.RLock()
	defer fake.verifySuccessMutex.RUnlock()
	return len(fake.verifySuccessArgsForCall)
}

func (fake *FakeBuild) VerifySuccessReturns(result1 error) {
	fake.VerifySuccessStub = nil
	fake.verifySuccessReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeBuild) VerifySuccessReturnsOnCall(i int, result1 error) {
	fake.VerifySuccessStub = nil
	if fake.verifySuccessReturnsOnCall == nil {
		fake.verifySuccessReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.verifySuccessReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakeBuild) IsVerifySuccess() bool {
	fake.isVerifySuccessMutex.Lock()
	ret, specificReturn := fake.isVerifySuccessReturnsOnCall[len(fake.isVerifySuccessArgsForCall)]
	fake.isVerifySuccessArgsForCall = append(fake.isVerifySuccessArgsForCall, struct{}{})
	fake.recordInvocation("IsVerifySuccess", []interface{}{})
	fake.isVerifySuccessMutex.Unlock()
	if fake.IsVerifySuccessStub != nil {
		return fake.IsVerifySuccessStub()
	}
	if specificReturn {
		return ret.result1
	}
	return fake.isVerifySuccessReturns.result1
}

func (fake *FakeBuild) IsVerifySuccessCallCount() int {
	fake.isVerifySuccessMutex.RLock()
	defer fake.isVerifySuccessMutex.RUnlock()
	return len(fake.isVerifySuccessArgsForCall)
}

func (fake *FakeBuild) IsVerifySuccessReturns(result1 bool) {
	fake.IsVerifySuccessStub = nil
	fake.isVerifySuccessReturns = struct {
		result1 bool
	}{result1}
}

func (fake *FakeBuild) IsVerifySuccessReturnsOnCall(i int, result1 bool) {
	fake.IsVerifySuccessStub = nil
	if fake.isVerifySuccessReturnsOnCall == nil {
		fake.isVerifySuccessReturnsOnCall = make(map[int]struct {
			result1 bool
		})
	}
	fake.isVerifySuccessReturnsOnCall[i] = struct {
		result1 bool
	}{result1}
}

func (fake *FakeBuild) VerifyFail() error {
	fake.verifyFailMutex.Lock()
	ret, specificReturn := fake.verifyFailReturnsOnCall[len(fake.verifyFailArgsForCall)]
	fake.verifyFailArgsForCall = append(fake.verifyFailArgsForCall, struct{}{})
	fake.recordInvocation("VerifyFail", []interface{}{})
	fake.verifyFailMutex.Unlock()
	if fake.VerifyFailStub != nil {
		return fake.VerifyFailStub()
	}
	if specificReturn {
		return ret.result1
	}
	return fake.verifyFailReturns.result1
}

func (fake *FakeBuild) VerifyFailCallCount() int {
	fake.verifyFailMutex.RLock()
	defer fake.verifyFailMutex.RUnlock()
	return len(fake.verifyFailArgsForCall)
}

func (fake *FakeBuild) VerifyFailReturns(result1 error) {
	fake.VerifyFailStub = nil
	fake.verifyFailReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeBuild) VerifyFailReturnsOnCall(i int, result1 error) {
	fake.VerifyFailStub = nil
	if fake.verifyFailReturnsOnCall == nil {
		fake.verifyFailReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.verifyFailReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakeBuild) IsVerifyFail() bool {
	fake.isVerifyFailMutex.Lock()
	ret, specificReturn := fake.isVerifyFailReturnsOnCall[len(fake.isVerifyFailArgsForCall)]
	fake.isVerifyFailArgsForCall = append(fake.isVerifyFailArgsForCall, struct{}{})
	fake.recordInvocation("IsVerifyFail", []interface{}{})
	fake.isVerifyFailMutex.Unlock()
	if fake.IsVerifyFailStub != nil {
		return fake.IsVerifyFailStub()
	}
	if specificReturn {
		return ret.result1
	}
	return fake.isVerifyFailReturns.result1
}

func (fake *FakeBuild) IsVerifyFailCallCount() int {
	fake.isVerifyFailMutex.RLock()
	defer fake.isVerifyFailMutex.RUnlock()
	return len(fake.isVerifyFailArgsForCall)
}

func (fake *FakeBuild) IsVerifyFailReturns(result1 bool) {
	fake.IsVerifyFailStub = nil
	fake.isVerifyFailReturns = struct {
		result1 bool
	}{result1}
}

func (fake *FakeBuild) IsVerifyFailReturnsOnCall(i int, result1 bool) {
	fake.IsVerifyFailStub = nil
	if fake.isVerifyFailReturnsOnCall == nil {
		fake.isVerifyFailReturnsOnCall = make(map[int]struct {
			result1 bool
		})
	}
	fake.isVerifyFailReturnsOnCall[i] = struct {
		result1 bool
	}{result1}
}

func (fake *FakeBuild) CreateVerify(arg1 api.VerifyParams) (api.Verify, error) {
	fake.createVerifyMutex.Lock()
	ret, specificReturn := fake.createVerifyReturnsOnCall[len(fake.createVerifyArgsForCall)]
	fake.createVerifyArgsForCall = append(fake.createVerifyArgsForCall, struct {
		arg1 api.VerifyParams
	}{arg1})
	fake.recordInvocation("CreateVerify", []interface{}{arg1})
	fake.createVerifyMutex.Unlock()
	if fake.CreateVerifyStub != nil {
		return fake.CreateVerifyStub(arg1)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fake.createVerifyReturns.result1, fake.createVerifyReturns.result2
}

func (fake *FakeBuild) CreateVerifyCallCount() int {
	fake.createVerifyMutex.RLock()
	defer fake.createVerifyMutex.RUnlock()
	return len(fake.createVerifyArgsForCall)
}

func (fake *FakeBuild) CreateVerifyArgsForCall(i int) api.VerifyParams {
	fake.createVerifyMutex.RLock()
	defer fake.createVerifyMutex.RUnlock()
	return fake.createVerifyArgsForCall[i].arg1
}

func (fake *FakeBuild) CreateVerifyReturns(result1 api.Verify, result2 error) {
	fake.CreateVerifyStub = nil
	fake.createVerifyReturns = struct {
		result1 api.Verify
		result2 error
	}{result1, result2}
}

func (fake *FakeBuild) CreateVerifyReturnsOnCall(i int, result1 api.Verify, result2 error) {
	fake.CreateVerifyStub = nil
	if fake.createVerifyReturnsOnCall == nil {
		fake.createVerifyReturnsOnCall = make(map[int]struct {
			result1 api.Verify
			result2 error
		})
	}
	fake.createVerifyReturnsOnCall[i] = struct {
		result1 api.Verify
		result2 error
	}{result1, result2}
}

func (fake *FakeBuild) GetVerify(id string) (api.Verify, error) {
	fake.getVerifyMutex.Lock()
	ret, specificReturn := fake.getVerifyReturnsOnCall[len(fake.getVerifyArgsForCall)]
	fake.getVerifyArgsForCall = append(fake.getVerifyArgsForCall, struct {
		id string
	}{id})
	fake.recordInvocation("GetVerify", []interface{}{id})
	fake.getVerifyMutex.Unlock()
	if fake.GetVerifyStub != nil {
		return fake.GetVerifyStub(id)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fake.getVerifyReturns.result1, fake.getVerifyReturns.result2
}

func (fake *FakeBuild) GetVerifyCallCount() int {
	fake.getVerifyMutex.RLock()
	defer fake.getVerifyMutex.RUnlock()
	return len(fake.getVerifyArgsForCall)
}

func (fake *FakeBuild) GetVerifyArgsForCall(i int) string {
	fake.getVerifyMutex.RLock()
	defer fake.getVerifyMutex.RUnlock()
	return fake.getVerifyArgsForCall[i].id
}

func (fake *FakeBuild) GetVerifyReturns(result1 api.Verify, result2 error) {
	fake.GetVerifyStub = nil
	fake.getVerifyReturns = struct {
		result1 api.Verify
		result2 error
	}{result1, result2}
}

func (fake *FakeBuild) GetVerifyReturnsOnCall(i int, result1 api.Verify, result2 error) {
	fake.GetVerifyStub = nil
	if fake.getVerifyReturnsOnCall == nil {
		fake.getVerifyReturnsOnCall = make(map[int]struct {
			result1 api.Verify
			result2 error
		})
	}
	fake.getVerifyReturnsOnCall[i] = struct {
		result1 api.Verify
		result2 error
	}{result1, result2}
}

func (fake *FakeBuild) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.idMutex.RLock()
	defer fake.idMutex.RUnlock()
	fake.gitShaMutex.RLock()
	defer fake.gitShaMutex.RUnlock()
	fake.statusMutex.RLock()
	defer fake.statusMutex.RUnlock()
	fake.verifyMutex.RLock()
	defer fake.verifyMutex.RUnlock()
	fake.linksMutex.RLock()
	defer fake.linksMutex.RUnlock()
	fake.getAppMutex.RLock()
	defer fake.getAppMutex.RUnlock()
	fake.successMutex.RLock()
	defer fake.successMutex.RUnlock()
	fake.failMutex.RLock()
	defer fake.failMutex.RUnlock()
	fake.isSuccessMutex.RLock()
	defer fake.isSuccessMutex.RUnlock()
	fake.isFailMutex.RLock()
	defer fake.isFailMutex.RUnlock()
	fake.verifySuccessMutex.RLock()
	defer fake.verifySuccessMutex.RUnlock()
	fake.isVerifySuccessMutex.RLock()
	defer fake.isVerifySuccessMutex.RUnlock()
	fake.verifyFailMutex.RLock()
	defer fake.verifyFailMutex.RUnlock()
	fake.isVerifyFailMutex.RLock()
	defer fake.isVerifyFailMutex.RUnlock()
	fake.createVerifyMutex.RLock()
	defer fake.createVerifyMutex.RUnlock()
	fake.getVerifyMutex.RLock()
	defer fake.getVerifyMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakeBuild) recordInvocation(key string, args []interface{}) {
	fake.invocationsMutex.Lock()
	defer fake.invocationsMutex.Unlock()
	if fake.invocations == nil {
		fake.invocations = map[string][][]interface{}{}
	}
	if fake.invocations[key] == nil {
		fake.invocations[key] = [][]interface{}{}
	}
	fake.invocations[key] = append(fake.invocations[key], args)
}

var _ api.Build = new(FakeBuild)
