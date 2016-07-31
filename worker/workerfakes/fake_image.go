// This file was generated by counterfeiter
package workerfakes

import (
	"io"
	"sync"

	"github.com/concourse/atc"
	"github.com/concourse/atc/worker"
)

type FakeImage struct {
	FetchStub        func() (worker.Volume, io.ReadCloser, atc.Version, error)
	fetchMutex       sync.RWMutex
	fetchArgsForCall []struct{}
	fetchReturns     struct {
		result1 worker.Volume
		result2 io.ReadCloser
		result3 atc.Version
		result4 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeImage) Fetch() (worker.Volume, io.ReadCloser, atc.Version, error) {
	fake.fetchMutex.Lock()
	fake.fetchArgsForCall = append(fake.fetchArgsForCall, struct{}{})
	fake.recordInvocation("Fetch", []interface{}{})
	fake.fetchMutex.Unlock()
	if fake.FetchStub != nil {
		return fake.FetchStub()
	} else {
		return fake.fetchReturns.result1, fake.fetchReturns.result2, fake.fetchReturns.result3, fake.fetchReturns.result4
	}
}

func (fake *FakeImage) FetchCallCount() int {
	fake.fetchMutex.RLock()
	defer fake.fetchMutex.RUnlock()
	return len(fake.fetchArgsForCall)
}

func (fake *FakeImage) FetchReturns(result1 worker.Volume, result2 io.ReadCloser, result3 atc.Version, result4 error) {
	fake.FetchStub = nil
	fake.fetchReturns = struct {
		result1 worker.Volume
		result2 io.ReadCloser
		result3 atc.Version
		result4 error
	}{result1, result2, result3, result4}
}

func (fake *FakeImage) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.fetchMutex.RLock()
	defer fake.fetchMutex.RUnlock()
	return fake.invocations
}

func (fake *FakeImage) recordInvocation(key string, args []interface{}) {
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

var _ worker.Image = new(FakeImage)