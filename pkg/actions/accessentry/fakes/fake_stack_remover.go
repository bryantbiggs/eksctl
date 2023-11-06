// Code generated by counterfeiter. DO NOT EDIT.
package fakes

import (
	"context"
	"sync"

	"github.com/aws/aws-sdk-go-v2/service/cloudformation/types"
	"github.com/weaveworks/eksctl/pkg/actions/accessentry"
)

type FakeStackRemover struct {
	DeleteStackBySpecStub        func(context.Context, *types.Stack) (*types.Stack, error)
	deleteStackBySpecMutex       sync.RWMutex
	deleteStackBySpecArgsForCall []struct {
		arg1 context.Context
		arg2 *types.Stack
	}
	deleteStackBySpecReturns struct {
		result1 *types.Stack
		result2 error
	}
	deleteStackBySpecReturnsOnCall map[int]struct {
		result1 *types.Stack
		result2 error
	}
	DescribeStackStub        func(context.Context, *types.Stack) (*types.Stack, error)
	describeStackMutex       sync.RWMutex
	describeStackArgsForCall []struct {
		arg1 context.Context
		arg2 *types.Stack
	}
	describeStackReturns struct {
		result1 *types.Stack
		result2 error
	}
	describeStackReturnsOnCall map[int]struct {
		result1 *types.Stack
		result2 error
	}
	ListAccessEntryStackNamesStub        func(context.Context, string) ([]string, error)
	listAccessEntryStackNamesMutex       sync.RWMutex
	listAccessEntryStackNamesArgsForCall []struct {
		arg1 context.Context
		arg2 string
	}
	listAccessEntryStackNamesReturns struct {
		result1 []string
		result2 error
	}
	listAccessEntryStackNamesReturnsOnCall map[int]struct {
		result1 []string
		result2 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeStackRemover) DeleteStackBySpec(arg1 context.Context, arg2 *types.Stack) (*types.Stack, error) {
	fake.deleteStackBySpecMutex.Lock()
	ret, specificReturn := fake.deleteStackBySpecReturnsOnCall[len(fake.deleteStackBySpecArgsForCall)]
	fake.deleteStackBySpecArgsForCall = append(fake.deleteStackBySpecArgsForCall, struct {
		arg1 context.Context
		arg2 *types.Stack
	}{arg1, arg2})
	stub := fake.DeleteStackBySpecStub
	fakeReturns := fake.deleteStackBySpecReturns
	fake.recordInvocation("DeleteStackBySpec", []interface{}{arg1, arg2})
	fake.deleteStackBySpecMutex.Unlock()
	if stub != nil {
		return stub(arg1, arg2)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeStackRemover) DeleteStackBySpecCallCount() int {
	fake.deleteStackBySpecMutex.RLock()
	defer fake.deleteStackBySpecMutex.RUnlock()
	return len(fake.deleteStackBySpecArgsForCall)
}

func (fake *FakeStackRemover) DeleteStackBySpecCalls(stub func(context.Context, *types.Stack) (*types.Stack, error)) {
	fake.deleteStackBySpecMutex.Lock()
	defer fake.deleteStackBySpecMutex.Unlock()
	fake.DeleteStackBySpecStub = stub
}

func (fake *FakeStackRemover) DeleteStackBySpecArgsForCall(i int) (context.Context, *types.Stack) {
	fake.deleteStackBySpecMutex.RLock()
	defer fake.deleteStackBySpecMutex.RUnlock()
	argsForCall := fake.deleteStackBySpecArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2
}

func (fake *FakeStackRemover) DeleteStackBySpecReturns(result1 *types.Stack, result2 error) {
	fake.deleteStackBySpecMutex.Lock()
	defer fake.deleteStackBySpecMutex.Unlock()
	fake.DeleteStackBySpecStub = nil
	fake.deleteStackBySpecReturns = struct {
		result1 *types.Stack
		result2 error
	}{result1, result2}
}

func (fake *FakeStackRemover) DeleteStackBySpecReturnsOnCall(i int, result1 *types.Stack, result2 error) {
	fake.deleteStackBySpecMutex.Lock()
	defer fake.deleteStackBySpecMutex.Unlock()
	fake.DeleteStackBySpecStub = nil
	if fake.deleteStackBySpecReturnsOnCall == nil {
		fake.deleteStackBySpecReturnsOnCall = make(map[int]struct {
			result1 *types.Stack
			result2 error
		})
	}
	fake.deleteStackBySpecReturnsOnCall[i] = struct {
		result1 *types.Stack
		result2 error
	}{result1, result2}
}

func (fake *FakeStackRemover) DescribeStack(arg1 context.Context, arg2 *types.Stack) (*types.Stack, error) {
	fake.describeStackMutex.Lock()
	ret, specificReturn := fake.describeStackReturnsOnCall[len(fake.describeStackArgsForCall)]
	fake.describeStackArgsForCall = append(fake.describeStackArgsForCall, struct {
		arg1 context.Context
		arg2 *types.Stack
	}{arg1, arg2})
	stub := fake.DescribeStackStub
	fakeReturns := fake.describeStackReturns
	fake.recordInvocation("DescribeStack", []interface{}{arg1, arg2})
	fake.describeStackMutex.Unlock()
	if stub != nil {
		return stub(arg1, arg2)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeStackRemover) DescribeStackCallCount() int {
	fake.describeStackMutex.RLock()
	defer fake.describeStackMutex.RUnlock()
	return len(fake.describeStackArgsForCall)
}

func (fake *FakeStackRemover) DescribeStackCalls(stub func(context.Context, *types.Stack) (*types.Stack, error)) {
	fake.describeStackMutex.Lock()
	defer fake.describeStackMutex.Unlock()
	fake.DescribeStackStub = stub
}

func (fake *FakeStackRemover) DescribeStackArgsForCall(i int) (context.Context, *types.Stack) {
	fake.describeStackMutex.RLock()
	defer fake.describeStackMutex.RUnlock()
	argsForCall := fake.describeStackArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2
}

func (fake *FakeStackRemover) DescribeStackReturns(result1 *types.Stack, result2 error) {
	fake.describeStackMutex.Lock()
	defer fake.describeStackMutex.Unlock()
	fake.DescribeStackStub = nil
	fake.describeStackReturns = struct {
		result1 *types.Stack
		result2 error
	}{result1, result2}
}

func (fake *FakeStackRemover) DescribeStackReturnsOnCall(i int, result1 *types.Stack, result2 error) {
	fake.describeStackMutex.Lock()
	defer fake.describeStackMutex.Unlock()
	fake.DescribeStackStub = nil
	if fake.describeStackReturnsOnCall == nil {
		fake.describeStackReturnsOnCall = make(map[int]struct {
			result1 *types.Stack
			result2 error
		})
	}
	fake.describeStackReturnsOnCall[i] = struct {
		result1 *types.Stack
		result2 error
	}{result1, result2}
}

func (fake *FakeStackRemover) ListAccessEntryStackNames(arg1 context.Context, arg2 string) ([]string, error) {
	fake.listAccessEntryStackNamesMutex.Lock()
	ret, specificReturn := fake.listAccessEntryStackNamesReturnsOnCall[len(fake.listAccessEntryStackNamesArgsForCall)]
	fake.listAccessEntryStackNamesArgsForCall = append(fake.listAccessEntryStackNamesArgsForCall, struct {
		arg1 context.Context
		arg2 string
	}{arg1, arg2})
	stub := fake.ListAccessEntryStackNamesStub
	fakeReturns := fake.listAccessEntryStackNamesReturns
	fake.recordInvocation("ListAccessEntryStackNames", []interface{}{arg1, arg2})
	fake.listAccessEntryStackNamesMutex.Unlock()
	if stub != nil {
		return stub(arg1, arg2)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeStackRemover) ListAccessEntryStackNamesCallCount() int {
	fake.listAccessEntryStackNamesMutex.RLock()
	defer fake.listAccessEntryStackNamesMutex.RUnlock()
	return len(fake.listAccessEntryStackNamesArgsForCall)
}

func (fake *FakeStackRemover) ListAccessEntryStackNamesCalls(stub func(context.Context, string) ([]string, error)) {
	fake.listAccessEntryStackNamesMutex.Lock()
	defer fake.listAccessEntryStackNamesMutex.Unlock()
	fake.ListAccessEntryStackNamesStub = stub
}

func (fake *FakeStackRemover) ListAccessEntryStackNamesArgsForCall(i int) (context.Context, string) {
	fake.listAccessEntryStackNamesMutex.RLock()
	defer fake.listAccessEntryStackNamesMutex.RUnlock()
	argsForCall := fake.listAccessEntryStackNamesArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2
}

func (fake *FakeStackRemover) ListAccessEntryStackNamesReturns(result1 []string, result2 error) {
	fake.listAccessEntryStackNamesMutex.Lock()
	defer fake.listAccessEntryStackNamesMutex.Unlock()
	fake.ListAccessEntryStackNamesStub = nil
	fake.listAccessEntryStackNamesReturns = struct {
		result1 []string
		result2 error
	}{result1, result2}
}

func (fake *FakeStackRemover) ListAccessEntryStackNamesReturnsOnCall(i int, result1 []string, result2 error) {
	fake.listAccessEntryStackNamesMutex.Lock()
	defer fake.listAccessEntryStackNamesMutex.Unlock()
	fake.ListAccessEntryStackNamesStub = nil
	if fake.listAccessEntryStackNamesReturnsOnCall == nil {
		fake.listAccessEntryStackNamesReturnsOnCall = make(map[int]struct {
			result1 []string
			result2 error
		})
	}
	fake.listAccessEntryStackNamesReturnsOnCall[i] = struct {
		result1 []string
		result2 error
	}{result1, result2}
}

func (fake *FakeStackRemover) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.deleteStackBySpecMutex.RLock()
	defer fake.deleteStackBySpecMutex.RUnlock()
	fake.describeStackMutex.RLock()
	defer fake.describeStackMutex.RUnlock()
	fake.listAccessEntryStackNamesMutex.RLock()
	defer fake.listAccessEntryStackNamesMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakeStackRemover) recordInvocation(key string, args []interface{}) {
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

var _ accessentry.StackRemover = new(FakeStackRemover)
