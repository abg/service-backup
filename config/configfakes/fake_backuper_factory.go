// This file was generated by counterfeiter
package configfakes

import (
	"sync"

	"github.com/pivotal-cf/service-backup/azure"
	"github.com/pivotal-cf/service-backup/config"
	"github.com/pivotal-cf/service-backup/gcp"
	"github.com/pivotal-cf/service-backup/s3"
	"github.com/pivotal-cf/service-backup/scp"
)

type FakeBackuperFactory struct {
	S3Stub        func(destination config.Destination, locator config.SystemTrustStoreLocator) (*s3.S3CliClient, error)
	s3Mutex       sync.RWMutex
	s3ArgsForCall []struct {
		destination config.Destination
		locator     config.SystemTrustStoreLocator
	}
	s3Returns struct {
		result1 *s3.S3CliClient
		result2 error
	}
	s3ReturnsOnCall map[int]struct {
		result1 *s3.S3CliClient
		result2 error
	}
	SCPStub        func(destination config.Destination) *scp.SCPClient
	sCPMutex       sync.RWMutex
	sCPArgsForCall []struct {
		destination config.Destination
	}
	sCPReturns struct {
		result1 *scp.SCPClient
	}
	sCPReturnsOnCall map[int]struct {
		result1 *scp.SCPClient
	}
	AzureStub        func(destination config.Destination) *azure.AzureClient
	azureMutex       sync.RWMutex
	azureArgsForCall []struct {
		destination config.Destination
	}
	azureReturns struct {
		result1 *azure.AzureClient
	}
	azureReturnsOnCall map[int]struct {
		result1 *azure.AzureClient
	}
	GCPStub        func(destination config.Destination) *gcp.StorageClient
	gCPMutex       sync.RWMutex
	gCPArgsForCall []struct {
		destination config.Destination
	}
	gCPReturns struct {
		result1 *gcp.StorageClient
	}
	gCPReturnsOnCall map[int]struct {
		result1 *gcp.StorageClient
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeBackuperFactory) S3(destination config.Destination, locator config.SystemTrustStoreLocator) (*s3.S3CliClient, error) {
	fake.s3Mutex.Lock()
	ret, specificReturn := fake.s3ReturnsOnCall[len(fake.s3ArgsForCall)]
	fake.s3ArgsForCall = append(fake.s3ArgsForCall, struct {
		destination config.Destination
		locator     config.SystemTrustStoreLocator
	}{destination, locator})
	fake.recordInvocation("S3", []interface{}{destination, locator})
	fake.s3Mutex.Unlock()
	if fake.S3Stub != nil {
		return fake.S3Stub(destination, locator)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fake.s3Returns.result1, fake.s3Returns.result2
}

func (fake *FakeBackuperFactory) S3CallCount() int {
	fake.s3Mutex.RLock()
	defer fake.s3Mutex.RUnlock()
	return len(fake.s3ArgsForCall)
}

func (fake *FakeBackuperFactory) S3ArgsForCall(i int) (config.Destination, config.SystemTrustStoreLocator) {
	fake.s3Mutex.RLock()
	defer fake.s3Mutex.RUnlock()
	return fake.s3ArgsForCall[i].destination, fake.s3ArgsForCall[i].locator
}

func (fake *FakeBackuperFactory) S3Returns(result1 *s3.S3CliClient, result2 error) {
	fake.S3Stub = nil
	fake.s3Returns = struct {
		result1 *s3.S3CliClient
		result2 error
	}{result1, result2}
}

func (fake *FakeBackuperFactory) S3ReturnsOnCall(i int, result1 *s3.S3CliClient, result2 error) {
	fake.S3Stub = nil
	if fake.s3ReturnsOnCall == nil {
		fake.s3ReturnsOnCall = make(map[int]struct {
			result1 *s3.S3CliClient
			result2 error
		})
	}
	fake.s3ReturnsOnCall[i] = struct {
		result1 *s3.S3CliClient
		result2 error
	}{result1, result2}
}

func (fake *FakeBackuperFactory) SCP(destination config.Destination) *scp.SCPClient {
	fake.sCPMutex.Lock()
	ret, specificReturn := fake.sCPReturnsOnCall[len(fake.sCPArgsForCall)]
	fake.sCPArgsForCall = append(fake.sCPArgsForCall, struct {
		destination config.Destination
	}{destination})
	fake.recordInvocation("SCP", []interface{}{destination})
	fake.sCPMutex.Unlock()
	if fake.SCPStub != nil {
		return fake.SCPStub(destination)
	}
	if specificReturn {
		return ret.result1
	}
	return fake.sCPReturns.result1
}

func (fake *FakeBackuperFactory) SCPCallCount() int {
	fake.sCPMutex.RLock()
	defer fake.sCPMutex.RUnlock()
	return len(fake.sCPArgsForCall)
}

func (fake *FakeBackuperFactory) SCPArgsForCall(i int) config.Destination {
	fake.sCPMutex.RLock()
	defer fake.sCPMutex.RUnlock()
	return fake.sCPArgsForCall[i].destination
}

func (fake *FakeBackuperFactory) SCPReturns(result1 *scp.SCPClient) {
	fake.SCPStub = nil
	fake.sCPReturns = struct {
		result1 *scp.SCPClient
	}{result1}
}

func (fake *FakeBackuperFactory) SCPReturnsOnCall(i int, result1 *scp.SCPClient) {
	fake.SCPStub = nil
	if fake.sCPReturnsOnCall == nil {
		fake.sCPReturnsOnCall = make(map[int]struct {
			result1 *scp.SCPClient
		})
	}
	fake.sCPReturnsOnCall[i] = struct {
		result1 *scp.SCPClient
	}{result1}
}

func (fake *FakeBackuperFactory) Azure(destination config.Destination) *azure.AzureClient {
	fake.azureMutex.Lock()
	ret, specificReturn := fake.azureReturnsOnCall[len(fake.azureArgsForCall)]
	fake.azureArgsForCall = append(fake.azureArgsForCall, struct {
		destination config.Destination
	}{destination})
	fake.recordInvocation("Azure", []interface{}{destination})
	fake.azureMutex.Unlock()
	if fake.AzureStub != nil {
		return fake.AzureStub(destination)
	}
	if specificReturn {
		return ret.result1
	}
	return fake.azureReturns.result1
}

func (fake *FakeBackuperFactory) AzureCallCount() int {
	fake.azureMutex.RLock()
	defer fake.azureMutex.RUnlock()
	return len(fake.azureArgsForCall)
}

func (fake *FakeBackuperFactory) AzureArgsForCall(i int) config.Destination {
	fake.azureMutex.RLock()
	defer fake.azureMutex.RUnlock()
	return fake.azureArgsForCall[i].destination
}

func (fake *FakeBackuperFactory) AzureReturns(result1 *azure.AzureClient) {
	fake.AzureStub = nil
	fake.azureReturns = struct {
		result1 *azure.AzureClient
	}{result1}
}

func (fake *FakeBackuperFactory) AzureReturnsOnCall(i int, result1 *azure.AzureClient) {
	fake.AzureStub = nil
	if fake.azureReturnsOnCall == nil {
		fake.azureReturnsOnCall = make(map[int]struct {
			result1 *azure.AzureClient
		})
	}
	fake.azureReturnsOnCall[i] = struct {
		result1 *azure.AzureClient
	}{result1}
}

func (fake *FakeBackuperFactory) GCP(destination config.Destination) *gcp.StorageClient {
	fake.gCPMutex.Lock()
	ret, specificReturn := fake.gCPReturnsOnCall[len(fake.gCPArgsForCall)]
	fake.gCPArgsForCall = append(fake.gCPArgsForCall, struct {
		destination config.Destination
	}{destination})
	fake.recordInvocation("GCP", []interface{}{destination})
	fake.gCPMutex.Unlock()
	if fake.GCPStub != nil {
		return fake.GCPStub(destination)
	}
	if specificReturn {
		return ret.result1
	}
	return fake.gCPReturns.result1
}

func (fake *FakeBackuperFactory) GCPCallCount() int {
	fake.gCPMutex.RLock()
	defer fake.gCPMutex.RUnlock()
	return len(fake.gCPArgsForCall)
}

func (fake *FakeBackuperFactory) GCPArgsForCall(i int) config.Destination {
	fake.gCPMutex.RLock()
	defer fake.gCPMutex.RUnlock()
	return fake.gCPArgsForCall[i].destination
}

func (fake *FakeBackuperFactory) GCPReturns(result1 *gcp.StorageClient) {
	fake.GCPStub = nil
	fake.gCPReturns = struct {
		result1 *gcp.StorageClient
	}{result1}
}

func (fake *FakeBackuperFactory) GCPReturnsOnCall(i int, result1 *gcp.StorageClient) {
	fake.GCPStub = nil
	if fake.gCPReturnsOnCall == nil {
		fake.gCPReturnsOnCall = make(map[int]struct {
			result1 *gcp.StorageClient
		})
	}
	fake.gCPReturnsOnCall[i] = struct {
		result1 *gcp.StorageClient
	}{result1}
}

func (fake *FakeBackuperFactory) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.s3Mutex.RLock()
	defer fake.s3Mutex.RUnlock()
	fake.sCPMutex.RLock()
	defer fake.sCPMutex.RUnlock()
	fake.azureMutex.RLock()
	defer fake.azureMutex.RUnlock()
	fake.gCPMutex.RLock()
	defer fake.gCPMutex.RUnlock()
	return fake.invocations
}

func (fake *FakeBackuperFactory) recordInvocation(key string, args []interface{}) {
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

var _ config.BackuperFactory = new(FakeBackuperFactory)