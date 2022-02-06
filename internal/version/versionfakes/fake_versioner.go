// Code generated by counterfeiter. DO NOT EDIT.
package versionfakes

import (
	"sync"

	"github.com/rleszilm/tag-version/internal/version"
)

type FakeVersioner struct {
	BranchStub        func() (string, error)
	branchMutex       sync.RWMutex
	branchArgsForCall []struct {
	}
	branchReturns struct {
		result1 string
		result2 error
	}
	branchReturnsOnCall map[int]struct {
		result1 string
		result2 error
	}
	CommitStub        func() (string, error)
	commitMutex       sync.RWMutex
	commitArgsForCall []struct {
	}
	commitReturns struct {
		result1 string
		result2 error
	}
	commitReturnsOnCall map[int]struct {
		result1 string
		result2 error
	}
	CommittishStub        func() (string, error)
	committishMutex       sync.RWMutex
	committishArgsForCall []struct {
	}
	committishReturns struct {
		result1 string
		result2 error
	}
	committishReturnsOnCall map[int]struct {
		result1 string
		result2 error
	}
	TagStub        func() (string, error)
	tagMutex       sync.RWMutex
	tagArgsForCall []struct {
	}
	tagReturns struct {
		result1 string
		result2 error
	}
	tagReturnsOnCall map[int]struct {
		result1 string
		result2 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeVersioner) Branch() (string, error) {
	fake.branchMutex.Lock()
	ret, specificReturn := fake.branchReturnsOnCall[len(fake.branchArgsForCall)]
	fake.branchArgsForCall = append(fake.branchArgsForCall, struct {
	}{})
	stub := fake.BranchStub
	fakeReturns := fake.branchReturns
	fake.recordInvocation("Branch", []interface{}{})
	fake.branchMutex.Unlock()
	if stub != nil {
		return stub()
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeVersioner) BranchCallCount() int {
	fake.branchMutex.RLock()
	defer fake.branchMutex.RUnlock()
	return len(fake.branchArgsForCall)
}

func (fake *FakeVersioner) BranchCalls(stub func() (string, error)) {
	fake.branchMutex.Lock()
	defer fake.branchMutex.Unlock()
	fake.BranchStub = stub
}

func (fake *FakeVersioner) BranchReturns(result1 string, result2 error) {
	fake.branchMutex.Lock()
	defer fake.branchMutex.Unlock()
	fake.BranchStub = nil
	fake.branchReturns = struct {
		result1 string
		result2 error
	}{result1, result2}
}

func (fake *FakeVersioner) BranchReturnsOnCall(i int, result1 string, result2 error) {
	fake.branchMutex.Lock()
	defer fake.branchMutex.Unlock()
	fake.BranchStub = nil
	if fake.branchReturnsOnCall == nil {
		fake.branchReturnsOnCall = make(map[int]struct {
			result1 string
			result2 error
		})
	}
	fake.branchReturnsOnCall[i] = struct {
		result1 string
		result2 error
	}{result1, result2}
}

func (fake *FakeVersioner) Commit() (string, error) {
	fake.commitMutex.Lock()
	ret, specificReturn := fake.commitReturnsOnCall[len(fake.commitArgsForCall)]
	fake.commitArgsForCall = append(fake.commitArgsForCall, struct {
	}{})
	stub := fake.CommitStub
	fakeReturns := fake.commitReturns
	fake.recordInvocation("Commit", []interface{}{})
	fake.commitMutex.Unlock()
	if stub != nil {
		return stub()
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeVersioner) CommitCallCount() int {
	fake.commitMutex.RLock()
	defer fake.commitMutex.RUnlock()
	return len(fake.commitArgsForCall)
}

func (fake *FakeVersioner) CommitCalls(stub func() (string, error)) {
	fake.commitMutex.Lock()
	defer fake.commitMutex.Unlock()
	fake.CommitStub = stub
}

func (fake *FakeVersioner) CommitReturns(result1 string, result2 error) {
	fake.commitMutex.Lock()
	defer fake.commitMutex.Unlock()
	fake.CommitStub = nil
	fake.commitReturns = struct {
		result1 string
		result2 error
	}{result1, result2}
}

func (fake *FakeVersioner) CommitReturnsOnCall(i int, result1 string, result2 error) {
	fake.commitMutex.Lock()
	defer fake.commitMutex.Unlock()
	fake.CommitStub = nil
	if fake.commitReturnsOnCall == nil {
		fake.commitReturnsOnCall = make(map[int]struct {
			result1 string
			result2 error
		})
	}
	fake.commitReturnsOnCall[i] = struct {
		result1 string
		result2 error
	}{result1, result2}
}

func (fake *FakeVersioner) Committish() (string, error) {
	fake.committishMutex.Lock()
	ret, specificReturn := fake.committishReturnsOnCall[len(fake.committishArgsForCall)]
	fake.committishArgsForCall = append(fake.committishArgsForCall, struct {
	}{})
	stub := fake.CommittishStub
	fakeReturns := fake.committishReturns
	fake.recordInvocation("Committish", []interface{}{})
	fake.committishMutex.Unlock()
	if stub != nil {
		return stub()
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeVersioner) CommittishCallCount() int {
	fake.committishMutex.RLock()
	defer fake.committishMutex.RUnlock()
	return len(fake.committishArgsForCall)
}

func (fake *FakeVersioner) CommittishCalls(stub func() (string, error)) {
	fake.committishMutex.Lock()
	defer fake.committishMutex.Unlock()
	fake.CommittishStub = stub
}

func (fake *FakeVersioner) CommittishReturns(result1 string, result2 error) {
	fake.committishMutex.Lock()
	defer fake.committishMutex.Unlock()
	fake.CommittishStub = nil
	fake.committishReturns = struct {
		result1 string
		result2 error
	}{result1, result2}
}

func (fake *FakeVersioner) CommittishReturnsOnCall(i int, result1 string, result2 error) {
	fake.committishMutex.Lock()
	defer fake.committishMutex.Unlock()
	fake.CommittishStub = nil
	if fake.committishReturnsOnCall == nil {
		fake.committishReturnsOnCall = make(map[int]struct {
			result1 string
			result2 error
		})
	}
	fake.committishReturnsOnCall[i] = struct {
		result1 string
		result2 error
	}{result1, result2}
}

func (fake *FakeVersioner) Tag() (string, error) {
	fake.tagMutex.Lock()
	ret, specificReturn := fake.tagReturnsOnCall[len(fake.tagArgsForCall)]
	fake.tagArgsForCall = append(fake.tagArgsForCall, struct {
	}{})
	stub := fake.TagStub
	fakeReturns := fake.tagReturns
	fake.recordInvocation("Tag", []interface{}{})
	fake.tagMutex.Unlock()
	if stub != nil {
		return stub()
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeVersioner) TagCallCount() int {
	fake.tagMutex.RLock()
	defer fake.tagMutex.RUnlock()
	return len(fake.tagArgsForCall)
}

func (fake *FakeVersioner) TagCalls(stub func() (string, error)) {
	fake.tagMutex.Lock()
	defer fake.tagMutex.Unlock()
	fake.TagStub = stub
}

func (fake *FakeVersioner) TagReturns(result1 string, result2 error) {
	fake.tagMutex.Lock()
	defer fake.tagMutex.Unlock()
	fake.TagStub = nil
	fake.tagReturns = struct {
		result1 string
		result2 error
	}{result1, result2}
}

func (fake *FakeVersioner) TagReturnsOnCall(i int, result1 string, result2 error) {
	fake.tagMutex.Lock()
	defer fake.tagMutex.Unlock()
	fake.TagStub = nil
	if fake.tagReturnsOnCall == nil {
		fake.tagReturnsOnCall = make(map[int]struct {
			result1 string
			result2 error
		})
	}
	fake.tagReturnsOnCall[i] = struct {
		result1 string
		result2 error
	}{result1, result2}
}

func (fake *FakeVersioner) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.branchMutex.RLock()
	defer fake.branchMutex.RUnlock()
	fake.commitMutex.RLock()
	defer fake.commitMutex.RUnlock()
	fake.committishMutex.RLock()
	defer fake.committishMutex.RUnlock()
	fake.tagMutex.RLock()
	defer fake.tagMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakeVersioner) recordInvocation(key string, args []interface{}) {
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

var _ version.Versioner = new(FakeVersioner)
