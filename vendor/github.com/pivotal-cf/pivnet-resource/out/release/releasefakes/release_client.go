// This file was generated by counterfeiter
package releasefakes

import (
	"sync"

	go_pivnet "github.com/pivotal-cf/go-pivnet"
)

type ReleaseClient struct {
	EULAsStub        func() ([]go_pivnet.EULA, error)
	eULAsMutex       sync.RWMutex
	eULAsArgsForCall []struct{}
	eULAsReturns     struct {
		result1 []go_pivnet.EULA
		result2 error
	}
	ReleaseTypesStub        func() ([]go_pivnet.ReleaseType, error)
	releaseTypesMutex       sync.RWMutex
	releaseTypesArgsForCall []struct{}
	releaseTypesReturns     struct {
		result1 []go_pivnet.ReleaseType
		result2 error
	}
	ReleasesForProductSlugStub        func(string) ([]go_pivnet.Release, error)
	releasesForProductSlugMutex       sync.RWMutex
	releasesForProductSlugArgsForCall []struct {
		arg1 string
	}
	releasesForProductSlugReturns struct {
		result1 []go_pivnet.Release
		result2 error
	}
	CreateReleaseStub        func(go_pivnet.CreateReleaseConfig) (go_pivnet.Release, error)
	createReleaseMutex       sync.RWMutex
	createReleaseArgsForCall []struct {
		arg1 go_pivnet.CreateReleaseConfig
	}
	createReleaseReturns struct {
		result1 go_pivnet.Release
		result2 error
	}
	DeleteReleaseStub        func(productSlug string, release go_pivnet.Release) error
	deleteReleaseMutex       sync.RWMutex
	deleteReleaseArgsForCall []struct {
		productSlug string
		release     go_pivnet.Release
	}
	deleteReleaseReturns struct {
		result1 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *ReleaseClient) EULAs() ([]go_pivnet.EULA, error) {
	fake.eULAsMutex.Lock()
	fake.eULAsArgsForCall = append(fake.eULAsArgsForCall, struct{}{})
	fake.recordInvocation("EULAs", []interface{}{})
	fake.eULAsMutex.Unlock()
	if fake.EULAsStub != nil {
		return fake.EULAsStub()
	} else {
		return fake.eULAsReturns.result1, fake.eULAsReturns.result2
	}
}

func (fake *ReleaseClient) EULAsCallCount() int {
	fake.eULAsMutex.RLock()
	defer fake.eULAsMutex.RUnlock()
	return len(fake.eULAsArgsForCall)
}

func (fake *ReleaseClient) EULAsReturns(result1 []go_pivnet.EULA, result2 error) {
	fake.EULAsStub = nil
	fake.eULAsReturns = struct {
		result1 []go_pivnet.EULA
		result2 error
	}{result1, result2}
}

func (fake *ReleaseClient) ReleaseTypes() ([]go_pivnet.ReleaseType, error) {
	fake.releaseTypesMutex.Lock()
	fake.releaseTypesArgsForCall = append(fake.releaseTypesArgsForCall, struct{}{})
	fake.recordInvocation("ReleaseTypes", []interface{}{})
	fake.releaseTypesMutex.Unlock()
	if fake.ReleaseTypesStub != nil {
		return fake.ReleaseTypesStub()
	} else {
		return fake.releaseTypesReturns.result1, fake.releaseTypesReturns.result2
	}
}

func (fake *ReleaseClient) ReleaseTypesCallCount() int {
	fake.releaseTypesMutex.RLock()
	defer fake.releaseTypesMutex.RUnlock()
	return len(fake.releaseTypesArgsForCall)
}

func (fake *ReleaseClient) ReleaseTypesReturns(result1 []go_pivnet.ReleaseType, result2 error) {
	fake.ReleaseTypesStub = nil
	fake.releaseTypesReturns = struct {
		result1 []go_pivnet.ReleaseType
		result2 error
	}{result1, result2}
}

func (fake *ReleaseClient) ReleasesForProductSlug(arg1 string) ([]go_pivnet.Release, error) {
	fake.releasesForProductSlugMutex.Lock()
	fake.releasesForProductSlugArgsForCall = append(fake.releasesForProductSlugArgsForCall, struct {
		arg1 string
	}{arg1})
	fake.recordInvocation("ReleasesForProductSlug", []interface{}{arg1})
	fake.releasesForProductSlugMutex.Unlock()
	if fake.ReleasesForProductSlugStub != nil {
		return fake.ReleasesForProductSlugStub(arg1)
	} else {
		return fake.releasesForProductSlugReturns.result1, fake.releasesForProductSlugReturns.result2
	}
}

func (fake *ReleaseClient) ReleasesForProductSlugCallCount() int {
	fake.releasesForProductSlugMutex.RLock()
	defer fake.releasesForProductSlugMutex.RUnlock()
	return len(fake.releasesForProductSlugArgsForCall)
}

func (fake *ReleaseClient) ReleasesForProductSlugArgsForCall(i int) string {
	fake.releasesForProductSlugMutex.RLock()
	defer fake.releasesForProductSlugMutex.RUnlock()
	return fake.releasesForProductSlugArgsForCall[i].arg1
}

func (fake *ReleaseClient) ReleasesForProductSlugReturns(result1 []go_pivnet.Release, result2 error) {
	fake.ReleasesForProductSlugStub = nil
	fake.releasesForProductSlugReturns = struct {
		result1 []go_pivnet.Release
		result2 error
	}{result1, result2}
}

func (fake *ReleaseClient) CreateRelease(arg1 go_pivnet.CreateReleaseConfig) (go_pivnet.Release, error) {
	fake.createReleaseMutex.Lock()
	fake.createReleaseArgsForCall = append(fake.createReleaseArgsForCall, struct {
		arg1 go_pivnet.CreateReleaseConfig
	}{arg1})
	fake.recordInvocation("CreateRelease", []interface{}{arg1})
	fake.createReleaseMutex.Unlock()
	if fake.CreateReleaseStub != nil {
		return fake.CreateReleaseStub(arg1)
	} else {
		return fake.createReleaseReturns.result1, fake.createReleaseReturns.result2
	}
}

func (fake *ReleaseClient) CreateReleaseCallCount() int {
	fake.createReleaseMutex.RLock()
	defer fake.createReleaseMutex.RUnlock()
	return len(fake.createReleaseArgsForCall)
}

func (fake *ReleaseClient) CreateReleaseArgsForCall(i int) go_pivnet.CreateReleaseConfig {
	fake.createReleaseMutex.RLock()
	defer fake.createReleaseMutex.RUnlock()
	return fake.createReleaseArgsForCall[i].arg1
}

func (fake *ReleaseClient) CreateReleaseReturns(result1 go_pivnet.Release, result2 error) {
	fake.CreateReleaseStub = nil
	fake.createReleaseReturns = struct {
		result1 go_pivnet.Release
		result2 error
	}{result1, result2}
}

func (fake *ReleaseClient) DeleteRelease(productSlug string, release go_pivnet.Release) error {
	fake.deleteReleaseMutex.Lock()
	fake.deleteReleaseArgsForCall = append(fake.deleteReleaseArgsForCall, struct {
		productSlug string
		release     go_pivnet.Release
	}{productSlug, release})
	fake.recordInvocation("DeleteRelease", []interface{}{productSlug, release})
	fake.deleteReleaseMutex.Unlock()
	if fake.DeleteReleaseStub != nil {
		return fake.DeleteReleaseStub(productSlug, release)
	} else {
		return fake.deleteReleaseReturns.result1
	}
}

func (fake *ReleaseClient) DeleteReleaseCallCount() int {
	fake.deleteReleaseMutex.RLock()
	defer fake.deleteReleaseMutex.RUnlock()
	return len(fake.deleteReleaseArgsForCall)
}

func (fake *ReleaseClient) DeleteReleaseArgsForCall(i int) (string, go_pivnet.Release) {
	fake.deleteReleaseMutex.RLock()
	defer fake.deleteReleaseMutex.RUnlock()
	return fake.deleteReleaseArgsForCall[i].productSlug, fake.deleteReleaseArgsForCall[i].release
}

func (fake *ReleaseClient) DeleteReleaseReturns(result1 error) {
	fake.DeleteReleaseStub = nil
	fake.deleteReleaseReturns = struct {
		result1 error
	}{result1}
}

func (fake *ReleaseClient) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.eULAsMutex.RLock()
	defer fake.eULAsMutex.RUnlock()
	fake.releaseTypesMutex.RLock()
	defer fake.releaseTypesMutex.RUnlock()
	fake.releasesForProductSlugMutex.RLock()
	defer fake.releasesForProductSlugMutex.RUnlock()
	fake.createReleaseMutex.RLock()
	defer fake.createReleaseMutex.RUnlock()
	fake.deleteReleaseMutex.RLock()
	defer fake.deleteReleaseMutex.RUnlock()
	return fake.invocations
}

func (fake *ReleaseClient) recordInvocation(key string, args []interface{}) {
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
