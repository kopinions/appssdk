// This file was generated by counterfeiter
package fakes

import (
	"sync"

	"github.com/sjkyspa/stacks/controller/api/api"
)

type FakeAuthRepository struct {
	CreateStub        func(params api.UserParams) (auth api.Auth, apiErr error)
	createMutex       sync.RWMutex
	createArgsForCall []struct {
		params api.UserParams
	}
	createReturns struct {
		result1 api.Auth
		result2 error
	}
	GetStub        func() (User user, apiErr error)
	getMutex       sync.RWMutex
	getArgsForCall []struct{}
	getReturns     struct {
		result1 user
		result2 error
	}
	DeleteStub        func(id string) error
	deleteMutex       sync.RWMutex
	deleteArgsForCall []struct {
		id string
	}
	deleteReturns struct {
		result1 error
	}
}

func (fake *FakeAuthRepository) Create(params api.UserParams) (auth api.Auth, apiErr error) {
	fake.createMutex.Lock()
	fake.createArgsForCall = append(fake.createArgsForCall, struct {
		params api.UserParams
	}{params})
	fake.createMutex.Unlock()
	if fake.CreateStub != nil {
		return fake.CreateStub(params)
	} else {
		return fake.createReturns.result1, fake.createReturns.result2
	}
}

func (fake *FakeAuthRepository) CreateCallCount() int {
	fake.createMutex.RLock()
	defer fake.createMutex.RUnlock()
	return len(fake.createArgsForCall)
}

func (fake *FakeAuthRepository) CreateArgsForCall(i int) api.UserParams {
	fake.createMutex.RLock()
	defer fake.createMutex.RUnlock()
	return fake.createArgsForCall[i].params
}

func (fake *FakeAuthRepository) CreateReturns(result1 api.Auth, result2 error) {
	fake.CreateStub = nil
	fake.createReturns = struct {
		result1 api.Auth
		result2 error
	}{result1, result2}
}

func (fake *FakeAuthRepository) Get() (User user, apiErr error) {
	fake.getMutex.Lock()
	fake.getArgsForCall = append(fake.getArgsForCall, struct{}{})
	fake.getMutex.Unlock()
	if fake.GetStub != nil {
		return fake.GetStub()
	} else {
		return fake.getReturns.result1, fake.getReturns.result2
	}
}

func (fake *FakeAuthRepository) GetCallCount() int {
	fake.getMutex.RLock()
	defer fake.getMutex.RUnlock()
	return len(fake.getArgsForCall)
}

func (fake *FakeAuthRepository) GetReturns(result1 user, result2 error) {
	fake.GetStub = nil
	fake.getReturns = struct {
		result1 user
		result2 error
	}{result1, result2}
}

func (fake *FakeAuthRepository) Delete(id string) error {
	fake.deleteMutex.Lock()
	fake.deleteArgsForCall = append(fake.deleteArgsForCall, struct {
		id string
	}{id})
	fake.deleteMutex.Unlock()
	if fake.DeleteStub != nil {
		return fake.DeleteStub(id)
	} else {
		return fake.deleteReturns.result1
	}
}

func (fake *FakeAuthRepository) DeleteCallCount() int {
	fake.deleteMutex.RLock()
	defer fake.deleteMutex.RUnlock()
	return len(fake.deleteArgsForCall)
}

func (fake *FakeAuthRepository) DeleteArgsForCall(i int) string {
	fake.deleteMutex.RLock()
	defer fake.deleteMutex.RUnlock()
	return fake.deleteArgsForCall[i].id
}

func (fake *FakeAuthRepository) DeleteReturns(result1 error) {
	fake.DeleteStub = nil
	fake.deleteReturns = struct {
		result1 error
	}{result1}
}

var _ api.AuthRepository = new(FakeAuthRepository)
