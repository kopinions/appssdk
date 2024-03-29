// Code generated by counterfeiter. DO NOT EDIT.
package fakes

import (
	"sync"

	"github.com/cnupp/appssdk/api"
)

type FakeKeyRepository struct {
	UploadStub        func(user api.User, params api.KeyParams) (uploaded api.Key, apiErr error)
	uploadMutex       sync.RWMutex
	uploadArgsForCall []struct {
		user   api.User
		params api.KeyParams
	}
	uploadReturns struct {
		result1 api.Key
		result2 error
	}
	uploadReturnsOnCall map[int]struct {
		result1 api.Key
		result2 error
	}
	GetKeyStub        func(id string) (key api.Key, apiErr error)
	getKeyMutex       sync.RWMutex
	getKeyArgsForCall []struct {
		id string
	}
	getKeyReturns struct {
		result1 api.Key
		result2 error
	}
	getKeyReturnsOnCall map[int]struct {
		result1 api.Key
		result2 error
	}
	GetKeysStub        func() (keys api.Keys, apiErr error)
	getKeysMutex       sync.RWMutex
	getKeysArgsForCall []struct{}
	getKeysReturns     struct {
		result1 api.Keys
		result2 error
	}
	getKeysReturnsOnCall map[int]struct {
		result1 api.Keys
		result2 error
	}
	GetKeysForUserStub        func(user api.User) (keys api.Keys, apiErr error)
	getKeysForUserMutex       sync.RWMutex
	getKeysForUserArgsForCall []struct {
		user api.User
	}
	getKeysForUserReturns struct {
		result1 api.Keys
		result2 error
	}
	getKeysForUserReturnsOnCall map[int]struct {
		result1 api.Keys
		result2 error
	}
	DeleteStub        func(user api.User, id string) (apiErr error)
	deleteMutex       sync.RWMutex
	deleteArgsForCall []struct {
		user api.User
		id   string
	}
	deleteReturns struct {
		result1 error
	}
	deleteReturnsOnCall map[int]struct {
		result1 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeKeyRepository) Upload(user api.User, params api.KeyParams) (uploaded api.Key, apiErr error) {
	fake.uploadMutex.Lock()
	ret, specificReturn := fake.uploadReturnsOnCall[len(fake.uploadArgsForCall)]
	fake.uploadArgsForCall = append(fake.uploadArgsForCall, struct {
		user   api.User
		params api.KeyParams
	}{user, params})
	fake.recordInvocation("Upload", []interface{}{user, params})
	fake.uploadMutex.Unlock()
	if fake.UploadStub != nil {
		return fake.UploadStub(user, params)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fake.uploadReturns.result1, fake.uploadReturns.result2
}

func (fake *FakeKeyRepository) UploadCallCount() int {
	fake.uploadMutex.RLock()
	defer fake.uploadMutex.RUnlock()
	return len(fake.uploadArgsForCall)
}

func (fake *FakeKeyRepository) UploadArgsForCall(i int) (api.User, api.KeyParams) {
	fake.uploadMutex.RLock()
	defer fake.uploadMutex.RUnlock()
	return fake.uploadArgsForCall[i].user, fake.uploadArgsForCall[i].params
}

func (fake *FakeKeyRepository) UploadReturns(result1 api.Key, result2 error) {
	fake.UploadStub = nil
	fake.uploadReturns = struct {
		result1 api.Key
		result2 error
	}{result1, result2}
}

func (fake *FakeKeyRepository) UploadReturnsOnCall(i int, result1 api.Key, result2 error) {
	fake.UploadStub = nil
	if fake.uploadReturnsOnCall == nil {
		fake.uploadReturnsOnCall = make(map[int]struct {
			result1 api.Key
			result2 error
		})
	}
	fake.uploadReturnsOnCall[i] = struct {
		result1 api.Key
		result2 error
	}{result1, result2}
}

func (fake *FakeKeyRepository) GetKey(id string) (key api.Key, apiErr error) {
	fake.getKeyMutex.Lock()
	ret, specificReturn := fake.getKeyReturnsOnCall[len(fake.getKeyArgsForCall)]
	fake.getKeyArgsForCall = append(fake.getKeyArgsForCall, struct {
		id string
	}{id})
	fake.recordInvocation("GetKey", []interface{}{id})
	fake.getKeyMutex.Unlock()
	if fake.GetKeyStub != nil {
		return fake.GetKeyStub(id)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fake.getKeyReturns.result1, fake.getKeyReturns.result2
}

func (fake *FakeKeyRepository) GetKeyCallCount() int {
	fake.getKeyMutex.RLock()
	defer fake.getKeyMutex.RUnlock()
	return len(fake.getKeyArgsForCall)
}

func (fake *FakeKeyRepository) GetKeyArgsForCall(i int) string {
	fake.getKeyMutex.RLock()
	defer fake.getKeyMutex.RUnlock()
	return fake.getKeyArgsForCall[i].id
}

func (fake *FakeKeyRepository) GetKeyReturns(result1 api.Key, result2 error) {
	fake.GetKeyStub = nil
	fake.getKeyReturns = struct {
		result1 api.Key
		result2 error
	}{result1, result2}
}

func (fake *FakeKeyRepository) GetKeyReturnsOnCall(i int, result1 api.Key, result2 error) {
	fake.GetKeyStub = nil
	if fake.getKeyReturnsOnCall == nil {
		fake.getKeyReturnsOnCall = make(map[int]struct {
			result1 api.Key
			result2 error
		})
	}
	fake.getKeyReturnsOnCall[i] = struct {
		result1 api.Key
		result2 error
	}{result1, result2}
}

func (fake *FakeKeyRepository) GetKeys() (keys api.Keys, apiErr error) {
	fake.getKeysMutex.Lock()
	ret, specificReturn := fake.getKeysReturnsOnCall[len(fake.getKeysArgsForCall)]
	fake.getKeysArgsForCall = append(fake.getKeysArgsForCall, struct{}{})
	fake.recordInvocation("GetKeys", []interface{}{})
	fake.getKeysMutex.Unlock()
	if fake.GetKeysStub != nil {
		return fake.GetKeysStub()
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fake.getKeysReturns.result1, fake.getKeysReturns.result2
}

func (fake *FakeKeyRepository) GetKeysCallCount() int {
	fake.getKeysMutex.RLock()
	defer fake.getKeysMutex.RUnlock()
	return len(fake.getKeysArgsForCall)
}

func (fake *FakeKeyRepository) GetKeysReturns(result1 api.Keys, result2 error) {
	fake.GetKeysStub = nil
	fake.getKeysReturns = struct {
		result1 api.Keys
		result2 error
	}{result1, result2}
}

func (fake *FakeKeyRepository) GetKeysReturnsOnCall(i int, result1 api.Keys, result2 error) {
	fake.GetKeysStub = nil
	if fake.getKeysReturnsOnCall == nil {
		fake.getKeysReturnsOnCall = make(map[int]struct {
			result1 api.Keys
			result2 error
		})
	}
	fake.getKeysReturnsOnCall[i] = struct {
		result1 api.Keys
		result2 error
	}{result1, result2}
}

func (fake *FakeKeyRepository) GetKeysForUser(user api.User) (keys api.Keys, apiErr error) {
	fake.getKeysForUserMutex.Lock()
	ret, specificReturn := fake.getKeysForUserReturnsOnCall[len(fake.getKeysForUserArgsForCall)]
	fake.getKeysForUserArgsForCall = append(fake.getKeysForUserArgsForCall, struct {
		user api.User
	}{user})
	fake.recordInvocation("GetKeysForUser", []interface{}{user})
	fake.getKeysForUserMutex.Unlock()
	if fake.GetKeysForUserStub != nil {
		return fake.GetKeysForUserStub(user)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fake.getKeysForUserReturns.result1, fake.getKeysForUserReturns.result2
}

func (fake *FakeKeyRepository) GetKeysForUserCallCount() int {
	fake.getKeysForUserMutex.RLock()
	defer fake.getKeysForUserMutex.RUnlock()
	return len(fake.getKeysForUserArgsForCall)
}

func (fake *FakeKeyRepository) GetKeysForUserArgsForCall(i int) api.User {
	fake.getKeysForUserMutex.RLock()
	defer fake.getKeysForUserMutex.RUnlock()
	return fake.getKeysForUserArgsForCall[i].user
}

func (fake *FakeKeyRepository) GetKeysForUserReturns(result1 api.Keys, result2 error) {
	fake.GetKeysForUserStub = nil
	fake.getKeysForUserReturns = struct {
		result1 api.Keys
		result2 error
	}{result1, result2}
}

func (fake *FakeKeyRepository) GetKeysForUserReturnsOnCall(i int, result1 api.Keys, result2 error) {
	fake.GetKeysForUserStub = nil
	if fake.getKeysForUserReturnsOnCall == nil {
		fake.getKeysForUserReturnsOnCall = make(map[int]struct {
			result1 api.Keys
			result2 error
		})
	}
	fake.getKeysForUserReturnsOnCall[i] = struct {
		result1 api.Keys
		result2 error
	}{result1, result2}
}

func (fake *FakeKeyRepository) Delete(user api.User, id string) (apiErr error) {
	fake.deleteMutex.Lock()
	ret, specificReturn := fake.deleteReturnsOnCall[len(fake.deleteArgsForCall)]
	fake.deleteArgsForCall = append(fake.deleteArgsForCall, struct {
		user api.User
		id   string
	}{user, id})
	fake.recordInvocation("Delete", []interface{}{user, id})
	fake.deleteMutex.Unlock()
	if fake.DeleteStub != nil {
		return fake.DeleteStub(user, id)
	}
	if specificReturn {
		return ret.result1
	}
	return fake.deleteReturns.result1
}

func (fake *FakeKeyRepository) DeleteCallCount() int {
	fake.deleteMutex.RLock()
	defer fake.deleteMutex.RUnlock()
	return len(fake.deleteArgsForCall)
}

func (fake *FakeKeyRepository) DeleteArgsForCall(i int) (api.User, string) {
	fake.deleteMutex.RLock()
	defer fake.deleteMutex.RUnlock()
	return fake.deleteArgsForCall[i].user, fake.deleteArgsForCall[i].id
}

func (fake *FakeKeyRepository) DeleteReturns(result1 error) {
	fake.DeleteStub = nil
	fake.deleteReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeKeyRepository) DeleteReturnsOnCall(i int, result1 error) {
	fake.DeleteStub = nil
	if fake.deleteReturnsOnCall == nil {
		fake.deleteReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.deleteReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakeKeyRepository) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.uploadMutex.RLock()
	defer fake.uploadMutex.RUnlock()
	fake.getKeyMutex.RLock()
	defer fake.getKeyMutex.RUnlock()
	fake.getKeysMutex.RLock()
	defer fake.getKeysMutex.RUnlock()
	fake.getKeysForUserMutex.RLock()
	defer fake.getKeysForUserMutex.RUnlock()
	fake.deleteMutex.RLock()
	defer fake.deleteMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakeKeyRepository) recordInvocation(key string, args []interface{}) {
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

var _ api.KeyRepository = new(FakeKeyRepository)
