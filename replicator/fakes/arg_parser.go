// This file was generated by counterfeiter
package fakes

import (
	"sync"

	"github.com/pivotal-cf/replicator/replicator"
)

type ArgParser struct {
	ParseStub        func([]string) (replicator.ApplicationConfig, error)
	parseMutex       sync.RWMutex
	parseArgsForCall []struct {
		arg1 []string
	}
	parseReturns struct {
		result1 replicator.ApplicationConfig
		result2 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *ArgParser) Parse(arg1 []string) (replicator.ApplicationConfig, error) {
	var arg1Copy []string
	if arg1 != nil {
		arg1Copy = make([]string, len(arg1))
		copy(arg1Copy, arg1)
	}
	fake.parseMutex.Lock()
	fake.parseArgsForCall = append(fake.parseArgsForCall, struct {
		arg1 []string
	}{arg1Copy})
	fake.recordInvocation("Parse", []interface{}{arg1Copy})
	fake.parseMutex.Unlock()
	if fake.ParseStub != nil {
		return fake.ParseStub(arg1)
	} else {
		return fake.parseReturns.result1, fake.parseReturns.result2
	}
}

func (fake *ArgParser) ParseCallCount() int {
	fake.parseMutex.RLock()
	defer fake.parseMutex.RUnlock()
	return len(fake.parseArgsForCall)
}

func (fake *ArgParser) ParseArgsForCall(i int) []string {
	fake.parseMutex.RLock()
	defer fake.parseMutex.RUnlock()
	return fake.parseArgsForCall[i].arg1
}

func (fake *ArgParser) ParseReturns(result1 replicator.ApplicationConfig, result2 error) {
	fake.ParseStub = nil
	fake.parseReturns = struct {
		result1 replicator.ApplicationConfig
		result2 error
	}{result1, result2}
}

func (fake *ArgParser) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.parseMutex.RLock()
	defer fake.parseMutex.RUnlock()
	return fake.invocations
}

func (fake *ArgParser) recordInvocation(key string, args []interface{}) {
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
