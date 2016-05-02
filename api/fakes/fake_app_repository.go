// This file was generated by counterfeiter
package fakes

import (
	"sync"

	"github.com/sjkyspa/stacks/controller/api/api"
)

type FakeAppRepository struct {
	CreateStub        func(params api.AppParams) (createdApp api.App, apiErr error)
	createMutex       sync.RWMutex
	createArgsForCall []struct {
		params api.AppParams
	}
	createReturns struct {
		result1 api.App
		result2 error
	}
	GetAppStub        func(id string) (api.App, error)
	getAppMutex       sync.RWMutex
	getAppArgsForCall []struct {
		id string
	}
	getAppReturns struct {
		result1 api.App
		result2 error
	}
	GetAppsStub        func() (api.Apps, error)
	getAppsMutex       sync.RWMutex
	getAppsArgsForCall []struct{}
	getAppsReturns     struct {
		result1 api.Apps
		result2 error
	}
	DeleteStub        func(id string) (apiErr error)
	deleteMutex       sync.RWMutex
	deleteArgsForCall []struct {
		id string
	}
	deleteReturns struct {
		result1 error
	}
	BindWithRouteStub        func(app api.App, params api.AppRouteParams) error
	bindWithRouteMutex       sync.RWMutex
	bindWithRouteArgsForCall []struct {
		app    api.App
		params api.AppRouteParams
	}
	bindWithRouteReturns struct {
		result1 error
	}
	UnbindRouteStub        func(app api.App, routeId string) error
	unbindRouteMutex       sync.RWMutex
	unbindRouteArgsForCall []struct {
		app     api.App
		routeId string
	}
	unbindRouteReturns struct {
		result1 error
	}
	GetRoutesStub        func(app api.App) (routes api.AppRoutes, apiErr error)
	getRoutesMutex       sync.RWMutex
	getRoutesArgsForCall []struct {
		app api.App
	}
	getRoutesReturns struct {
		result1 api.AppRoutes
		result2 error
	}
	GetRoutesByURIStub        func(uri string) (routes api.AppRoutes, apiErr error)
	getRoutesByURIMutex       sync.RWMutex
	getRoutesByURIArgsForCall []struct {
		uri string
	}
	getRoutesByURIReturns struct {
		result1 api.AppRoutes
		result2 error
	}
	SetEnvStub        func(app api.App, kvs map[string]interface{}) error
	setEnvMutex       sync.RWMutex
	setEnvArgsForCall []struct {
		app api.App
		kvs map[string]interface{}
	}
	setEnvReturns struct {
		result1 error
	}
	UnsetEnvStub        func(app api.App, keys []string) error
	unsetEnvMutex       sync.RWMutex
	unsetEnvArgsForCall []struct {
		app  api.App
		keys []string
	}
	unsetEnvReturns struct {
		result1 error
	}
	SwitchStackStub        func(id string, params api.UpdateStackParams) (apiErr error)
	switchStackMutex       sync.RWMutex
	switchStackArgsForCall []struct {
		id     string
		params api.UpdateStackParams
	}
	switchStackReturns struct {
		result1 error
	}
}

func (fake *FakeAppRepository) Create(params api.AppParams) (createdApp api.App, apiErr error) {
	fake.createMutex.Lock()
	fake.createArgsForCall = append(fake.createArgsForCall, struct {
		params api.AppParams
	}{params})
	fake.createMutex.Unlock()
	if fake.CreateStub != nil {
		return fake.CreateStub(params)
	} else {
		return fake.createReturns.result1, fake.createReturns.result2
	}
}

func (fake *FakeAppRepository) CreateCallCount() int {
	fake.createMutex.RLock()
	defer fake.createMutex.RUnlock()
	return len(fake.createArgsForCall)
}

func (fake *FakeAppRepository) CreateArgsForCall(i int) api.AppParams {
	fake.createMutex.RLock()
	defer fake.createMutex.RUnlock()
	return fake.createArgsForCall[i].params
}

func (fake *FakeAppRepository) CreateReturns(result1 api.App, result2 error) {
	fake.CreateStub = nil
	fake.createReturns = struct {
		result1 api.App
		result2 error
	}{result1, result2}
}

func (fake *FakeAppRepository) GetApp(id string) (api.App, error) {
	fake.getAppMutex.Lock()
	fake.getAppArgsForCall = append(fake.getAppArgsForCall, struct {
		id string
	}{id})
	fake.getAppMutex.Unlock()
	if fake.GetAppStub != nil {
		return fake.GetAppStub(id)
	} else {
		return fake.getAppReturns.result1, fake.getAppReturns.result2
	}
}

func (fake *FakeAppRepository) GetAppCallCount() int {
	fake.getAppMutex.RLock()
	defer fake.getAppMutex.RUnlock()
	return len(fake.getAppArgsForCall)
}

func (fake *FakeAppRepository) GetAppArgsForCall(i int) string {
	fake.getAppMutex.RLock()
	defer fake.getAppMutex.RUnlock()
	return fake.getAppArgsForCall[i].id
}

func (fake *FakeAppRepository) GetAppReturns(result1 api.App, result2 error) {
	fake.GetAppStub = nil
	fake.getAppReturns = struct {
		result1 api.App
		result2 error
	}{result1, result2}
}

func (fake *FakeAppRepository) GetApps() (api.Apps, error) {
	fake.getAppsMutex.Lock()
	fake.getAppsArgsForCall = append(fake.getAppsArgsForCall, struct{}{})
	fake.getAppsMutex.Unlock()
	if fake.GetAppsStub != nil {
		return fake.GetAppsStub()
	} else {
		return fake.getAppsReturns.result1, fake.getAppsReturns.result2
	}
}

func (fake *FakeAppRepository) GetAppsCallCount() int {
	fake.getAppsMutex.RLock()
	defer fake.getAppsMutex.RUnlock()
	return len(fake.getAppsArgsForCall)
}

func (fake *FakeAppRepository) GetAppsReturns(result1 api.Apps, result2 error) {
	fake.GetAppsStub = nil
	fake.getAppsReturns = struct {
		result1 api.Apps
		result2 error
	}{result1, result2}
}

func (fake *FakeAppRepository) Delete(id string) (apiErr error) {
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

func (fake *FakeAppRepository) DeleteCallCount() int {
	fake.deleteMutex.RLock()
	defer fake.deleteMutex.RUnlock()
	return len(fake.deleteArgsForCall)
}

func (fake *FakeAppRepository) DeleteArgsForCall(i int) string {
	fake.deleteMutex.RLock()
	defer fake.deleteMutex.RUnlock()
	return fake.deleteArgsForCall[i].id
}

func (fake *FakeAppRepository) DeleteReturns(result1 error) {
	fake.DeleteStub = nil
	fake.deleteReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeAppRepository) BindWithRoute(app api.App, params api.AppRouteParams) error {
	fake.bindWithRouteMutex.Lock()
	fake.bindWithRouteArgsForCall = append(fake.bindWithRouteArgsForCall, struct {
		app    api.App
		params api.AppRouteParams
	}{app, params})
	fake.bindWithRouteMutex.Unlock()
	if fake.BindWithRouteStub != nil {
		return fake.BindWithRouteStub(app, params)
	} else {
		return fake.bindWithRouteReturns.result1
	}
}

func (fake *FakeAppRepository) BindWithRouteCallCount() int {
	fake.bindWithRouteMutex.RLock()
	defer fake.bindWithRouteMutex.RUnlock()
	return len(fake.bindWithRouteArgsForCall)
}

func (fake *FakeAppRepository) BindWithRouteArgsForCall(i int) (api.App, api.AppRouteParams) {
	fake.bindWithRouteMutex.RLock()
	defer fake.bindWithRouteMutex.RUnlock()
	return fake.bindWithRouteArgsForCall[i].app, fake.bindWithRouteArgsForCall[i].params
}

func (fake *FakeAppRepository) BindWithRouteReturns(result1 error) {
	fake.BindWithRouteStub = nil
	fake.bindWithRouteReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeAppRepository) UnbindRoute(app api.App, routeId string) error {
	fake.unbindRouteMutex.Lock()
	fake.unbindRouteArgsForCall = append(fake.unbindRouteArgsForCall, struct {
		app     api.App
		routeId string
	}{app, routeId})
	fake.unbindRouteMutex.Unlock()
	if fake.UnbindRouteStub != nil {
		return fake.UnbindRouteStub(app, routeId)
	} else {
		return fake.unbindRouteReturns.result1
	}
}

func (fake *FakeAppRepository) UnbindRouteCallCount() int {
	fake.unbindRouteMutex.RLock()
	defer fake.unbindRouteMutex.RUnlock()
	return len(fake.unbindRouteArgsForCall)
}

func (fake *FakeAppRepository) UnbindRouteArgsForCall(i int) (api.App, string) {
	fake.unbindRouteMutex.RLock()
	defer fake.unbindRouteMutex.RUnlock()
	return fake.unbindRouteArgsForCall[i].app, fake.unbindRouteArgsForCall[i].routeId
}

func (fake *FakeAppRepository) UnbindRouteReturns(result1 error) {
	fake.UnbindRouteStub = nil
	fake.unbindRouteReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeAppRepository) GetRoutes(app api.App) (routes api.AppRoutes, apiErr error) {
	fake.getRoutesMutex.Lock()
	fake.getRoutesArgsForCall = append(fake.getRoutesArgsForCall, struct {
		app api.App
	}{app})
	fake.getRoutesMutex.Unlock()
	if fake.GetRoutesStub != nil {
		return fake.GetRoutesStub(app)
	} else {
		return fake.getRoutesReturns.result1, fake.getRoutesReturns.result2
	}
}

func (fake *FakeAppRepository) GetRoutesCallCount() int {
	fake.getRoutesMutex.RLock()
	defer fake.getRoutesMutex.RUnlock()
	return len(fake.getRoutesArgsForCall)
}

func (fake *FakeAppRepository) GetRoutesArgsForCall(i int) api.App {
	fake.getRoutesMutex.RLock()
	defer fake.getRoutesMutex.RUnlock()
	return fake.getRoutesArgsForCall[i].app
}

func (fake *FakeAppRepository) GetRoutesReturns(result1 api.AppRoutes, result2 error) {
	fake.GetRoutesStub = nil
	fake.getRoutesReturns = struct {
		result1 api.AppRoutes
		result2 error
	}{result1, result2}
}

func (fake *FakeAppRepository) GetRoutesByURI(uri string) (routes api.AppRoutes, apiErr error) {
	fake.getRoutesByURIMutex.Lock()
	fake.getRoutesByURIArgsForCall = append(fake.getRoutesByURIArgsForCall, struct {
		uri string
	}{uri})
	fake.getRoutesByURIMutex.Unlock()
	if fake.GetRoutesByURIStub != nil {
		return fake.GetRoutesByURIStub(uri)
	} else {
		return fake.getRoutesByURIReturns.result1, fake.getRoutesByURIReturns.result2
	}
}

func (fake *FakeAppRepository) GetRoutesByURICallCount() int {
	fake.getRoutesByURIMutex.RLock()
	defer fake.getRoutesByURIMutex.RUnlock()
	return len(fake.getRoutesByURIArgsForCall)
}

func (fake *FakeAppRepository) GetRoutesByURIArgsForCall(i int) string {
	fake.getRoutesByURIMutex.RLock()
	defer fake.getRoutesByURIMutex.RUnlock()
	return fake.getRoutesByURIArgsForCall[i].uri
}

func (fake *FakeAppRepository) GetRoutesByURIReturns(result1 api.AppRoutes, result2 error) {
	fake.GetRoutesByURIStub = nil
	fake.getRoutesByURIReturns = struct {
		result1 api.AppRoutes
		result2 error
	}{result1, result2}
}

func (fake *FakeAppRepository) SetEnv(app api.App, kvs map[string]interface{}) error {
	fake.setEnvMutex.Lock()
	fake.setEnvArgsForCall = append(fake.setEnvArgsForCall, struct {
		app api.App
		kvs map[string]interface{}
	}{app, kvs})
	fake.setEnvMutex.Unlock()
	if fake.SetEnvStub != nil {
		return fake.SetEnvStub(app, kvs)
	} else {
		return fake.setEnvReturns.result1
	}
}

func (fake *FakeAppRepository) SetEnvCallCount() int {
	fake.setEnvMutex.RLock()
	defer fake.setEnvMutex.RUnlock()
	return len(fake.setEnvArgsForCall)
}

func (fake *FakeAppRepository) SetEnvArgsForCall(i int) (api.App, map[string]interface{}) {
	fake.setEnvMutex.RLock()
	defer fake.setEnvMutex.RUnlock()
	return fake.setEnvArgsForCall[i].app, fake.setEnvArgsForCall[i].kvs
}

func (fake *FakeAppRepository) SetEnvReturns(result1 error) {
	fake.SetEnvStub = nil
	fake.setEnvReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeAppRepository) UnsetEnv(app api.App, keys []string) error {
	fake.unsetEnvMutex.Lock()
	fake.unsetEnvArgsForCall = append(fake.unsetEnvArgsForCall, struct {
		app  api.App
		keys []string
	}{app, keys})
	fake.unsetEnvMutex.Unlock()
	if fake.UnsetEnvStub != nil {
		return fake.UnsetEnvStub(app, keys)
	} else {
		return fake.unsetEnvReturns.result1
	}
}

func (fake *FakeAppRepository) UnsetEnvCallCount() int {
	fake.unsetEnvMutex.RLock()
	defer fake.unsetEnvMutex.RUnlock()
	return len(fake.unsetEnvArgsForCall)
}

func (fake *FakeAppRepository) UnsetEnvArgsForCall(i int) (api.App, []string) {
	fake.unsetEnvMutex.RLock()
	defer fake.unsetEnvMutex.RUnlock()
	return fake.unsetEnvArgsForCall[i].app, fake.unsetEnvArgsForCall[i].keys
}

func (fake *FakeAppRepository) UnsetEnvReturns(result1 error) {
	fake.UnsetEnvStub = nil
	fake.unsetEnvReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeAppRepository) SwitchStack(id string, params api.UpdateStackParams) (apiErr error) {
	fake.switchStackMutex.Lock()
	fake.switchStackArgsForCall = append(fake.switchStackArgsForCall, struct {
		id     string
		params api.UpdateStackParams
	}{id, params})
	fake.switchStackMutex.Unlock()
	if fake.SwitchStackStub != nil {
		return fake.SwitchStackStub(id, params)
	} else {
		return fake.switchStackReturns.result1
	}
}

func (fake *FakeAppRepository) SwitchStackCallCount() int {
	fake.switchStackMutex.RLock()
	defer fake.switchStackMutex.RUnlock()
	return len(fake.switchStackArgsForCall)
}

func (fake *FakeAppRepository) SwitchStackArgsForCall(i int) (string, api.UpdateStackParams) {
	fake.switchStackMutex.RLock()
	defer fake.switchStackMutex.RUnlock()
	return fake.switchStackArgsForCall[i].id, fake.switchStackArgsForCall[i].params
}

func (fake *FakeAppRepository) SwitchStackReturns(result1 error) {
	fake.SwitchStackStub = nil
	fake.switchStackReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeAppRepository) GetLog(appId, buildId, logType string, lines, offset int64) (api.LogsModel, error) {
	return api.LogsModel{}, nil
}

var _ api.AppRepository = new(FakeAppRepository)
