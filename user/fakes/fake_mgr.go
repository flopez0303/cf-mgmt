// This file was generated by counterfeiter
package fakes

import (
	"sync"

	"github.com/pivotalservices/cf-mgmt/uaa"
	"github.com/pivotalservices/cf-mgmt/user"
)

type FakeManager struct {
	InitializeLdapStub        func(ldapBindPassword string) error
	initializeLdapMutex       sync.RWMutex
	initializeLdapArgsForCall []struct {
		ldapBindPassword string
	}
	initializeLdapReturns struct {
		result1 error
	}
	DeinitializeLdapStub        func() error
	deinitializeLdapMutex       sync.RWMutex
	deinitializeLdapArgsForCall []struct{}
	deinitializeLdapReturns     struct {
		result1 error
	}
	UpdateSpaceUsersStub        func() error
	updateSpaceUsersMutex       sync.RWMutex
	updateSpaceUsersArgsForCall []struct{}
	updateSpaceUsersReturns     struct {
		result1 error
	}
	UpdateOrgUsersStub        func() error
	updateOrgUsersMutex       sync.RWMutex
	updateOrgUsersArgsForCall []struct{}
	updateOrgUsersReturns     struct {
		result1 error
	}
	CleanupOrgUsersStub        func() error
	cleanupOrgUsersMutex       sync.RWMutex
	cleanupOrgUsersArgsForCall []struct{}
	cleanupOrgUsersReturns     struct {
		result1 error
	}
	ListSpaceAuditorsStub        func(spaceGUID string, uaaUsers *uaa.Users) (*user.RoleUsers, error)
	listSpaceAuditorsMutex       sync.RWMutex
	listSpaceAuditorsArgsForCall []struct {
		spaceGUID string
		uaaUsers  *uaa.Users
	}
	listSpaceAuditorsReturns struct {
		result1 *user.RoleUsers
		result2 error
	}
	ListSpaceDevelopersStub        func(spaceGUID string, uaaUsers *uaa.Users) (*user.RoleUsers, error)
	listSpaceDevelopersMutex       sync.RWMutex
	listSpaceDevelopersArgsForCall []struct {
		spaceGUID string
		uaaUsers  *uaa.Users
	}
	listSpaceDevelopersReturns struct {
		result1 *user.RoleUsers
		result2 error
	}
	ListSpaceManagersStub        func(spaceGUID string, uaaUsers *uaa.Users) (*user.RoleUsers, error)
	listSpaceManagersMutex       sync.RWMutex
	listSpaceManagersArgsForCall []struct {
		spaceGUID string
		uaaUsers  *uaa.Users
	}
	listSpaceManagersReturns struct {
		result1 *user.RoleUsers
		result2 error
	}
	ListOrgAuditorsStub        func(orgGUID string, uaaUsers *uaa.Users) (*user.RoleUsers, error)
	listOrgAuditorsMutex       sync.RWMutex
	listOrgAuditorsArgsForCall []struct {
		orgGUID  string
		uaaUsers *uaa.Users
	}
	listOrgAuditorsReturns struct {
		result1 *user.RoleUsers
		result2 error
	}
	ListOrgBillingManagersStub        func(orgGUID string, uaaUsers *uaa.Users) (*user.RoleUsers, error)
	listOrgBillingManagersMutex       sync.RWMutex
	listOrgBillingManagersArgsForCall []struct {
		orgGUID  string
		uaaUsers *uaa.Users
	}
	listOrgBillingManagersReturns struct {
		result1 *user.RoleUsers
		result2 error
	}
	ListOrgManagersStub        func(orgGUID string, uaaUsers *uaa.Users) (*user.RoleUsers, error)
	listOrgManagersMutex       sync.RWMutex
	listOrgManagersArgsForCall []struct {
		orgGUID  string
		uaaUsers *uaa.Users
	}
	listOrgManagersReturns struct {
		result1 *user.RoleUsers
		result2 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeManager) InitializeLdap(ldapBindPassword string) error {
	fake.initializeLdapMutex.Lock()
	fake.initializeLdapArgsForCall = append(fake.initializeLdapArgsForCall, struct {
		ldapBindPassword string
	}{ldapBindPassword})
	fake.recordInvocation("InitializeLdap", []interface{}{ldapBindPassword})
	fake.initializeLdapMutex.Unlock()
	if fake.InitializeLdapStub != nil {
		return fake.InitializeLdapStub(ldapBindPassword)
	} else {
		return fake.initializeLdapReturns.result1
	}
}

func (fake *FakeManager) InitializeLdapCallCount() int {
	fake.initializeLdapMutex.RLock()
	defer fake.initializeLdapMutex.RUnlock()
	return len(fake.initializeLdapArgsForCall)
}

func (fake *FakeManager) InitializeLdapArgsForCall(i int) string {
	fake.initializeLdapMutex.RLock()
	defer fake.initializeLdapMutex.RUnlock()
	return fake.initializeLdapArgsForCall[i].ldapBindPassword
}

func (fake *FakeManager) InitializeLdapReturns(result1 error) {
	fake.InitializeLdapStub = nil
	fake.initializeLdapReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeManager) DeinitializeLdap() error {
	fake.deinitializeLdapMutex.Lock()
	fake.deinitializeLdapArgsForCall = append(fake.deinitializeLdapArgsForCall, struct{}{})
	fake.recordInvocation("DeinitializeLdap", []interface{}{})
	fake.deinitializeLdapMutex.Unlock()
	if fake.DeinitializeLdapStub != nil {
		return fake.DeinitializeLdapStub()
	} else {
		return fake.deinitializeLdapReturns.result1
	}
}

func (fake *FakeManager) DeinitializeLdapCallCount() int {
	fake.deinitializeLdapMutex.RLock()
	defer fake.deinitializeLdapMutex.RUnlock()
	return len(fake.deinitializeLdapArgsForCall)
}

func (fake *FakeManager) DeinitializeLdapReturns(result1 error) {
	fake.DeinitializeLdapStub = nil
	fake.deinitializeLdapReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeManager) UpdateSpaceUsers() error {
	fake.updateSpaceUsersMutex.Lock()
	fake.updateSpaceUsersArgsForCall = append(fake.updateSpaceUsersArgsForCall, struct{}{})
	fake.recordInvocation("UpdateSpaceUsers", []interface{}{})
	fake.updateSpaceUsersMutex.Unlock()
	if fake.UpdateSpaceUsersStub != nil {
		return fake.UpdateSpaceUsersStub()
	} else {
		return fake.updateSpaceUsersReturns.result1
	}
}

func (fake *FakeManager) UpdateSpaceUsersCallCount() int {
	fake.updateSpaceUsersMutex.RLock()
	defer fake.updateSpaceUsersMutex.RUnlock()
	return len(fake.updateSpaceUsersArgsForCall)
}

func (fake *FakeManager) UpdateSpaceUsersReturns(result1 error) {
	fake.UpdateSpaceUsersStub = nil
	fake.updateSpaceUsersReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeManager) UpdateOrgUsers() error {
	fake.updateOrgUsersMutex.Lock()
	fake.updateOrgUsersArgsForCall = append(fake.updateOrgUsersArgsForCall, struct{}{})
	fake.recordInvocation("UpdateOrgUsers", []interface{}{})
	fake.updateOrgUsersMutex.Unlock()
	if fake.UpdateOrgUsersStub != nil {
		return fake.UpdateOrgUsersStub()
	} else {
		return fake.updateOrgUsersReturns.result1
	}
}

func (fake *FakeManager) UpdateOrgUsersCallCount() int {
	fake.updateOrgUsersMutex.RLock()
	defer fake.updateOrgUsersMutex.RUnlock()
	return len(fake.updateOrgUsersArgsForCall)
}

func (fake *FakeManager) UpdateOrgUsersReturns(result1 error) {
	fake.UpdateOrgUsersStub = nil
	fake.updateOrgUsersReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeManager) CleanupOrgUsers() error {
	fake.cleanupOrgUsersMutex.Lock()
	fake.cleanupOrgUsersArgsForCall = append(fake.cleanupOrgUsersArgsForCall, struct{}{})
	fake.recordInvocation("CleanupOrgUsers", []interface{}{})
	fake.cleanupOrgUsersMutex.Unlock()
	if fake.CleanupOrgUsersStub != nil {
		return fake.CleanupOrgUsersStub()
	} else {
		return fake.cleanupOrgUsersReturns.result1
	}
}

func (fake *FakeManager) CleanupOrgUsersCallCount() int {
	fake.cleanupOrgUsersMutex.RLock()
	defer fake.cleanupOrgUsersMutex.RUnlock()
	return len(fake.cleanupOrgUsersArgsForCall)
}

func (fake *FakeManager) CleanupOrgUsersReturns(result1 error) {
	fake.CleanupOrgUsersStub = nil
	fake.cleanupOrgUsersReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeManager) ListSpaceAuditors(spaceGUID string, uaaUsers *uaa.Users) (*user.RoleUsers, error) {
	fake.listSpaceAuditorsMutex.Lock()
	fake.listSpaceAuditorsArgsForCall = append(fake.listSpaceAuditorsArgsForCall, struct {
		spaceGUID string
		uaaUsers  *uaa.Users
	}{spaceGUID, uaaUsers})
	fake.recordInvocation("ListSpaceAuditors", []interface{}{spaceGUID, uaaUsers})
	fake.listSpaceAuditorsMutex.Unlock()
	if fake.ListSpaceAuditorsStub != nil {
		return fake.ListSpaceAuditorsStub(spaceGUID, uaaUsers)
	} else {
		return fake.listSpaceAuditorsReturns.result1, fake.listSpaceAuditorsReturns.result2
	}
}

func (fake *FakeManager) ListSpaceAuditorsCallCount() int {
	fake.listSpaceAuditorsMutex.RLock()
	defer fake.listSpaceAuditorsMutex.RUnlock()
	return len(fake.listSpaceAuditorsArgsForCall)
}

func (fake *FakeManager) ListSpaceAuditorsArgsForCall(i int) (string, *uaa.Users) {
	fake.listSpaceAuditorsMutex.RLock()
	defer fake.listSpaceAuditorsMutex.RUnlock()
	return fake.listSpaceAuditorsArgsForCall[i].spaceGUID, fake.listSpaceAuditorsArgsForCall[i].uaaUsers
}

func (fake *FakeManager) ListSpaceAuditorsReturns(result1 *user.RoleUsers, result2 error) {
	fake.ListSpaceAuditorsStub = nil
	fake.listSpaceAuditorsReturns = struct {
		result1 *user.RoleUsers
		result2 error
	}{result1, result2}
}

func (fake *FakeManager) ListSpaceDevelopers(spaceGUID string, uaaUsers *uaa.Users) (*user.RoleUsers, error) {
	fake.listSpaceDevelopersMutex.Lock()
	fake.listSpaceDevelopersArgsForCall = append(fake.listSpaceDevelopersArgsForCall, struct {
		spaceGUID string
		uaaUsers  *uaa.Users
	}{spaceGUID, uaaUsers})
	fake.recordInvocation("ListSpaceDevelopers", []interface{}{spaceGUID, uaaUsers})
	fake.listSpaceDevelopersMutex.Unlock()
	if fake.ListSpaceDevelopersStub != nil {
		return fake.ListSpaceDevelopersStub(spaceGUID, uaaUsers)
	} else {
		return fake.listSpaceDevelopersReturns.result1, fake.listSpaceDevelopersReturns.result2
	}
}

func (fake *FakeManager) ListSpaceDevelopersCallCount() int {
	fake.listSpaceDevelopersMutex.RLock()
	defer fake.listSpaceDevelopersMutex.RUnlock()
	return len(fake.listSpaceDevelopersArgsForCall)
}

func (fake *FakeManager) ListSpaceDevelopersArgsForCall(i int) (string, *uaa.Users) {
	fake.listSpaceDevelopersMutex.RLock()
	defer fake.listSpaceDevelopersMutex.RUnlock()
	return fake.listSpaceDevelopersArgsForCall[i].spaceGUID, fake.listSpaceDevelopersArgsForCall[i].uaaUsers
}

func (fake *FakeManager) ListSpaceDevelopersReturns(result1 *user.RoleUsers, result2 error) {
	fake.ListSpaceDevelopersStub = nil
	fake.listSpaceDevelopersReturns = struct {
		result1 *user.RoleUsers
		result2 error
	}{result1, result2}
}

func (fake *FakeManager) ListSpaceManagers(spaceGUID string, uaaUsers *uaa.Users) (*user.RoleUsers, error) {
	fake.listSpaceManagersMutex.Lock()
	fake.listSpaceManagersArgsForCall = append(fake.listSpaceManagersArgsForCall, struct {
		spaceGUID string
		uaaUsers  *uaa.Users
	}{spaceGUID, uaaUsers})
	fake.recordInvocation("ListSpaceManagers", []interface{}{spaceGUID, uaaUsers})
	fake.listSpaceManagersMutex.Unlock()
	if fake.ListSpaceManagersStub != nil {
		return fake.ListSpaceManagersStub(spaceGUID, uaaUsers)
	} else {
		return fake.listSpaceManagersReturns.result1, fake.listSpaceManagersReturns.result2
	}
}

func (fake *FakeManager) ListSpaceManagersCallCount() int {
	fake.listSpaceManagersMutex.RLock()
	defer fake.listSpaceManagersMutex.RUnlock()
	return len(fake.listSpaceManagersArgsForCall)
}

func (fake *FakeManager) ListSpaceManagersArgsForCall(i int) (string, *uaa.Users) {
	fake.listSpaceManagersMutex.RLock()
	defer fake.listSpaceManagersMutex.RUnlock()
	return fake.listSpaceManagersArgsForCall[i].spaceGUID, fake.listSpaceManagersArgsForCall[i].uaaUsers
}

func (fake *FakeManager) ListSpaceManagersReturns(result1 *user.RoleUsers, result2 error) {
	fake.ListSpaceManagersStub = nil
	fake.listSpaceManagersReturns = struct {
		result1 *user.RoleUsers
		result2 error
	}{result1, result2}
}

func (fake *FakeManager) ListOrgAuditors(orgGUID string, uaaUsers *uaa.Users) (*user.RoleUsers, error) {
	fake.listOrgAuditorsMutex.Lock()
	fake.listOrgAuditorsArgsForCall = append(fake.listOrgAuditorsArgsForCall, struct {
		orgGUID  string
		uaaUsers *uaa.Users
	}{orgGUID, uaaUsers})
	fake.recordInvocation("ListOrgAuditors", []interface{}{orgGUID, uaaUsers})
	fake.listOrgAuditorsMutex.Unlock()
	if fake.ListOrgAuditorsStub != nil {
		return fake.ListOrgAuditorsStub(orgGUID, uaaUsers)
	} else {
		return fake.listOrgAuditorsReturns.result1, fake.listOrgAuditorsReturns.result2
	}
}

func (fake *FakeManager) ListOrgAuditorsCallCount() int {
	fake.listOrgAuditorsMutex.RLock()
	defer fake.listOrgAuditorsMutex.RUnlock()
	return len(fake.listOrgAuditorsArgsForCall)
}

func (fake *FakeManager) ListOrgAuditorsArgsForCall(i int) (string, *uaa.Users) {
	fake.listOrgAuditorsMutex.RLock()
	defer fake.listOrgAuditorsMutex.RUnlock()
	return fake.listOrgAuditorsArgsForCall[i].orgGUID, fake.listOrgAuditorsArgsForCall[i].uaaUsers
}

func (fake *FakeManager) ListOrgAuditorsReturns(result1 *user.RoleUsers, result2 error) {
	fake.ListOrgAuditorsStub = nil
	fake.listOrgAuditorsReturns = struct {
		result1 *user.RoleUsers
		result2 error
	}{result1, result2}
}

func (fake *FakeManager) ListOrgBillingManagers(orgGUID string, uaaUsers *uaa.Users) (*user.RoleUsers, error) {
	fake.listOrgBillingManagersMutex.Lock()
	fake.listOrgBillingManagersArgsForCall = append(fake.listOrgBillingManagersArgsForCall, struct {
		orgGUID  string
		uaaUsers *uaa.Users
	}{orgGUID, uaaUsers})
	fake.recordInvocation("ListOrgBillingManagers", []interface{}{orgGUID, uaaUsers})
	fake.listOrgBillingManagersMutex.Unlock()
	if fake.ListOrgBillingManagersStub != nil {
		return fake.ListOrgBillingManagersStub(orgGUID, uaaUsers)
	} else {
		return fake.listOrgBillingManagersReturns.result1, fake.listOrgBillingManagersReturns.result2
	}
}

func (fake *FakeManager) ListOrgBillingManagersCallCount() int {
	fake.listOrgBillingManagersMutex.RLock()
	defer fake.listOrgBillingManagersMutex.RUnlock()
	return len(fake.listOrgBillingManagersArgsForCall)
}

func (fake *FakeManager) ListOrgBillingManagersArgsForCall(i int) (string, *uaa.Users) {
	fake.listOrgBillingManagersMutex.RLock()
	defer fake.listOrgBillingManagersMutex.RUnlock()
	return fake.listOrgBillingManagersArgsForCall[i].orgGUID, fake.listOrgBillingManagersArgsForCall[i].uaaUsers
}

func (fake *FakeManager) ListOrgBillingManagersReturns(result1 *user.RoleUsers, result2 error) {
	fake.ListOrgBillingManagersStub = nil
	fake.listOrgBillingManagersReturns = struct {
		result1 *user.RoleUsers
		result2 error
	}{result1, result2}
}

func (fake *FakeManager) ListOrgManagers(orgGUID string, uaaUsers *uaa.Users) (*user.RoleUsers, error) {
	fake.listOrgManagersMutex.Lock()
	fake.listOrgManagersArgsForCall = append(fake.listOrgManagersArgsForCall, struct {
		orgGUID  string
		uaaUsers *uaa.Users
	}{orgGUID, uaaUsers})
	fake.recordInvocation("ListOrgManagers", []interface{}{orgGUID, uaaUsers})
	fake.listOrgManagersMutex.Unlock()
	if fake.ListOrgManagersStub != nil {
		return fake.ListOrgManagersStub(orgGUID, uaaUsers)
	} else {
		return fake.listOrgManagersReturns.result1, fake.listOrgManagersReturns.result2
	}
}

func (fake *FakeManager) ListOrgManagersCallCount() int {
	fake.listOrgManagersMutex.RLock()
	defer fake.listOrgManagersMutex.RUnlock()
	return len(fake.listOrgManagersArgsForCall)
}

func (fake *FakeManager) ListOrgManagersArgsForCall(i int) (string, *uaa.Users) {
	fake.listOrgManagersMutex.RLock()
	defer fake.listOrgManagersMutex.RUnlock()
	return fake.listOrgManagersArgsForCall[i].orgGUID, fake.listOrgManagersArgsForCall[i].uaaUsers
}

func (fake *FakeManager) ListOrgManagersReturns(result1 *user.RoleUsers, result2 error) {
	fake.ListOrgManagersStub = nil
	fake.listOrgManagersReturns = struct {
		result1 *user.RoleUsers
		result2 error
	}{result1, result2}
}

func (fake *FakeManager) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.initializeLdapMutex.RLock()
	defer fake.initializeLdapMutex.RUnlock()
	fake.deinitializeLdapMutex.RLock()
	defer fake.deinitializeLdapMutex.RUnlock()
	fake.updateSpaceUsersMutex.RLock()
	defer fake.updateSpaceUsersMutex.RUnlock()
	fake.updateOrgUsersMutex.RLock()
	defer fake.updateOrgUsersMutex.RUnlock()
	fake.cleanupOrgUsersMutex.RLock()
	defer fake.cleanupOrgUsersMutex.RUnlock()
	fake.listSpaceAuditorsMutex.RLock()
	defer fake.listSpaceAuditorsMutex.RUnlock()
	fake.listSpaceDevelopersMutex.RLock()
	defer fake.listSpaceDevelopersMutex.RUnlock()
	fake.listSpaceManagersMutex.RLock()
	defer fake.listSpaceManagersMutex.RUnlock()
	fake.listOrgAuditorsMutex.RLock()
	defer fake.listOrgAuditorsMutex.RUnlock()
	fake.listOrgBillingManagersMutex.RLock()
	defer fake.listOrgBillingManagersMutex.RUnlock()
	fake.listOrgManagersMutex.RLock()
	defer fake.listOrgManagersMutex.RUnlock()
	return fake.invocations
}

func (fake *FakeManager) recordInvocation(key string, args []interface{}) {
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

var _ user.Manager = new(FakeManager)
