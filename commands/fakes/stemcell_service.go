// Code generated by counterfeiter. DO NOT EDIT.
package fakes

import (
	"sync"

	"github.com/pivotal-cf/om/api"
)

type StemcellService struct {
	UploadStub        func(api.StemcellUploadInput) (api.StemcellUploadOutput, error)
	uploadMutex       sync.RWMutex
	uploadArgsForCall []struct {
		arg1 api.StemcellUploadInput
	}
	uploadReturns struct {
		result1 api.StemcellUploadOutput
		result2 error
	}
	uploadReturnsOnCall map[int]struct {
		result1 api.StemcellUploadOutput
		result2 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *StemcellService) Upload(arg1 api.StemcellUploadInput) (api.StemcellUploadOutput, error) {
	fake.uploadMutex.Lock()
	ret, specificReturn := fake.uploadReturnsOnCall[len(fake.uploadArgsForCall)]
	fake.uploadArgsForCall = append(fake.uploadArgsForCall, struct {
		arg1 api.StemcellUploadInput
	}{arg1})
	fake.recordInvocation("Upload", []interface{}{arg1})
	fake.uploadMutex.Unlock()
	if fake.UploadStub != nil {
		return fake.UploadStub(arg1)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fake.uploadReturns.result1, fake.uploadReturns.result2
}

func (fake *StemcellService) UploadCallCount() int {
	fake.uploadMutex.RLock()
	defer fake.uploadMutex.RUnlock()
	return len(fake.uploadArgsForCall)
}

func (fake *StemcellService) UploadArgsForCall(i int) api.StemcellUploadInput {
	fake.uploadMutex.RLock()
	defer fake.uploadMutex.RUnlock()
	return fake.uploadArgsForCall[i].arg1
}

func (fake *StemcellService) UploadReturns(result1 api.StemcellUploadOutput, result2 error) {
	fake.UploadStub = nil
	fake.uploadReturns = struct {
		result1 api.StemcellUploadOutput
		result2 error
	}{result1, result2}
}

func (fake *StemcellService) UploadReturnsOnCall(i int, result1 api.StemcellUploadOutput, result2 error) {
	fake.UploadStub = nil
	if fake.uploadReturnsOnCall == nil {
		fake.uploadReturnsOnCall = make(map[int]struct {
			result1 api.StemcellUploadOutput
			result2 error
		})
	}
	fake.uploadReturnsOnCall[i] = struct {
		result1 api.StemcellUploadOutput
		result2 error
	}{result1, result2}
}

func (fake *StemcellService) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.uploadMutex.RLock()
	defer fake.uploadMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *StemcellService) recordInvocation(key string, args []interface{}) {
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
