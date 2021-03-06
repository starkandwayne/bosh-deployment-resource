// This file was generated by counterfeiter
package licensefakes

import (
	"sync"

	"github.com/cloudfoundry/bosh-cli/release/license"
)

type FakeDirReader struct {
	ReadStub        func(string) (*license.License, error)
	readMutex       sync.RWMutex
	readArgsForCall []struct {
		arg1 string
	}
	readReturns struct {
		result1 *license.License
		result2 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeDirReader) Read(arg1 string) (*license.License, error) {
	fake.readMutex.Lock()
	fake.readArgsForCall = append(fake.readArgsForCall, struct {
		arg1 string
	}{arg1})
	fake.recordInvocation("Read", []interface{}{arg1})
	fake.readMutex.Unlock()
	if fake.ReadStub != nil {
		return fake.ReadStub(arg1)
	} else {
		return fake.readReturns.result1, fake.readReturns.result2
	}
}

func (fake *FakeDirReader) ReadCallCount() int {
	fake.readMutex.RLock()
	defer fake.readMutex.RUnlock()
	return len(fake.readArgsForCall)
}

func (fake *FakeDirReader) ReadArgsForCall(i int) string {
	fake.readMutex.RLock()
	defer fake.readMutex.RUnlock()
	return fake.readArgsForCall[i].arg1
}

func (fake *FakeDirReader) ReadReturns(result1 *license.License, result2 error) {
	fake.ReadStub = nil
	fake.readReturns = struct {
		result1 *license.License
		result2 error
	}{result1, result2}
}

func (fake *FakeDirReader) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.readMutex.RLock()
	defer fake.readMutex.RUnlock()
	return fake.invocations
}

func (fake *FakeDirReader) recordInvocation(key string, args []interface{}) {
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

var _ license.DirReader = new(FakeDirReader)
