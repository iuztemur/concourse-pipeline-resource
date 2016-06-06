// This file was generated by counterfeiter
package fakes

import (
	"io"
	"os"
	"sync"

	"github.com/concourse/atc"
	"github.com/concourse/atc/resource"
)

type FakeVersionedSource struct {
	RunStub        func(signals <-chan os.Signal, ready chan<- struct{}) error
	runMutex       sync.RWMutex
	runArgsForCall []struct {
		signals <-chan os.Signal
		ready   chan<- struct{}
	}
	runReturns struct {
		result1 error
	}
	VersionStub        func() atc.Version
	versionMutex       sync.RWMutex
	versionArgsForCall []struct{}
	versionReturns     struct {
		result1 atc.Version
	}
	MetadataStub        func() []atc.MetadataField
	metadataMutex       sync.RWMutex
	metadataArgsForCall []struct{}
	metadataReturns     struct {
		result1 []atc.MetadataField
	}
	StreamOutStub        func(string) (io.ReadCloser, error)
	streamOutMutex       sync.RWMutex
	streamOutArgsForCall []struct {
		arg1 string
	}
	streamOutReturns struct {
		result1 io.ReadCloser
		result2 error
	}
	StreamInStub        func(string, io.Reader) error
	streamInMutex       sync.RWMutex
	streamInArgsForCall []struct {
		arg1 string
		arg2 io.Reader
	}
	streamInReturns struct {
		result1 error
	}
}

func (fake *FakeVersionedSource) Run(signals <-chan os.Signal, ready chan<- struct{}) error {
	fake.runMutex.Lock()
	fake.runArgsForCall = append(fake.runArgsForCall, struct {
		signals <-chan os.Signal
		ready   chan<- struct{}
	}{signals, ready})
	fake.runMutex.Unlock()
	if fake.RunStub != nil {
		return fake.RunStub(signals, ready)
	} else {
		return fake.runReturns.result1
	}
}

func (fake *FakeVersionedSource) RunCallCount() int {
	fake.runMutex.RLock()
	defer fake.runMutex.RUnlock()
	return len(fake.runArgsForCall)
}

func (fake *FakeVersionedSource) RunArgsForCall(i int) (<-chan os.Signal, chan<- struct{}) {
	fake.runMutex.RLock()
	defer fake.runMutex.RUnlock()
	return fake.runArgsForCall[i].signals, fake.runArgsForCall[i].ready
}

func (fake *FakeVersionedSource) RunReturns(result1 error) {
	fake.RunStub = nil
	fake.runReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeVersionedSource) Version() atc.Version {
	fake.versionMutex.Lock()
	fake.versionArgsForCall = append(fake.versionArgsForCall, struct{}{})
	fake.versionMutex.Unlock()
	if fake.VersionStub != nil {
		return fake.VersionStub()
	} else {
		return fake.versionReturns.result1
	}
}

func (fake *FakeVersionedSource) VersionCallCount() int {
	fake.versionMutex.RLock()
	defer fake.versionMutex.RUnlock()
	return len(fake.versionArgsForCall)
}

func (fake *FakeVersionedSource) VersionReturns(result1 atc.Version) {
	fake.VersionStub = nil
	fake.versionReturns = struct {
		result1 atc.Version
	}{result1}
}

func (fake *FakeVersionedSource) Metadata() []atc.MetadataField {
	fake.metadataMutex.Lock()
	fake.metadataArgsForCall = append(fake.metadataArgsForCall, struct{}{})
	fake.metadataMutex.Unlock()
	if fake.MetadataStub != nil {
		return fake.MetadataStub()
	} else {
		return fake.metadataReturns.result1
	}
}

func (fake *FakeVersionedSource) MetadataCallCount() int {
	fake.metadataMutex.RLock()
	defer fake.metadataMutex.RUnlock()
	return len(fake.metadataArgsForCall)
}

func (fake *FakeVersionedSource) MetadataReturns(result1 []atc.MetadataField) {
	fake.MetadataStub = nil
	fake.metadataReturns = struct {
		result1 []atc.MetadataField
	}{result1}
}

func (fake *FakeVersionedSource) StreamOut(arg1 string) (io.ReadCloser, error) {
	fake.streamOutMutex.Lock()
	fake.streamOutArgsForCall = append(fake.streamOutArgsForCall, struct {
		arg1 string
	}{arg1})
	fake.streamOutMutex.Unlock()
	if fake.StreamOutStub != nil {
		return fake.StreamOutStub(arg1)
	} else {
		return fake.streamOutReturns.result1, fake.streamOutReturns.result2
	}
}

func (fake *FakeVersionedSource) StreamOutCallCount() int {
	fake.streamOutMutex.RLock()
	defer fake.streamOutMutex.RUnlock()
	return len(fake.streamOutArgsForCall)
}

func (fake *FakeVersionedSource) StreamOutArgsForCall(i int) string {
	fake.streamOutMutex.RLock()
	defer fake.streamOutMutex.RUnlock()
	return fake.streamOutArgsForCall[i].arg1
}

func (fake *FakeVersionedSource) StreamOutReturns(result1 io.ReadCloser, result2 error) {
	fake.StreamOutStub = nil
	fake.streamOutReturns = struct {
		result1 io.ReadCloser
		result2 error
	}{result1, result2}
}

func (fake *FakeVersionedSource) StreamIn(arg1 string, arg2 io.Reader) error {
	fake.streamInMutex.Lock()
	fake.streamInArgsForCall = append(fake.streamInArgsForCall, struct {
		arg1 string
		arg2 io.Reader
	}{arg1, arg2})
	fake.streamInMutex.Unlock()
	if fake.StreamInStub != nil {
		return fake.StreamInStub(arg1, arg2)
	} else {
		return fake.streamInReturns.result1
	}
}

func (fake *FakeVersionedSource) StreamInCallCount() int {
	fake.streamInMutex.RLock()
	defer fake.streamInMutex.RUnlock()
	return len(fake.streamInArgsForCall)
}

func (fake *FakeVersionedSource) StreamInArgsForCall(i int) (string, io.Reader) {
	fake.streamInMutex.RLock()
	defer fake.streamInMutex.RUnlock()
	return fake.streamInArgsForCall[i].arg1, fake.streamInArgsForCall[i].arg2
}

func (fake *FakeVersionedSource) StreamInReturns(result1 error) {
	fake.StreamInStub = nil
	fake.streamInReturns = struct {
		result1 error
	}{result1}
}

var _ resource.VersionedSource = new(FakeVersionedSource)