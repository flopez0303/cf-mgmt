package quota_test

import (
	"errors"

	cfclient "github.com/cloudfoundry-community/go-cfclient"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/pivotalservices/cf-mgmt/config"
	configfakes "github.com/pivotalservices/cf-mgmt/config/fakes"
	orgfakes "github.com/pivotalservices/cf-mgmt/organization/fakes"
	"github.com/pivotalservices/cf-mgmt/quota"
	quotafakes "github.com/pivotalservices/cf-mgmt/quota/fakes"
	spacefakes "github.com/pivotalservices/cf-mgmt/space/fakes"
)

var _ = Describe("given QuotaManager", func() {
	var (
		fakeReader   *configfakes.FakeReader
		fakeOrgMgr   *orgfakes.FakeManager
		fakeClient   *quotafakes.FakeCFClient
		fakeSpaceMgr *spacefakes.FakeManager
		quotaMgr     quota.DefaultManager
	)

	BeforeEach(func() {
		fakeReader = new(configfakes.FakeReader)
		fakeOrgMgr = new(orgfakes.FakeManager)
		fakeSpaceMgr = new(spacefakes.FakeManager)
		fakeClient = new(quotafakes.FakeCFClient)
		quotaMgr = quota.DefaultManager{
			Cfg:      fakeReader,
			Client:   fakeClient,
			OrgMgr:   fakeOrgMgr,
			SpaceMgr: fakeSpaceMgr,
			Peek:     false,
		}
	})

	Context("ListAllSpaceQuotasForOrg()", func() {
		It("should return 2 quotas", func() {
			fakeClient.ListOrgSpaceQuotasReturns([]cfclient.SpaceQuota{
				cfclient.SpaceQuota{
					Name: "quota-1",
					Guid: "quota-1-guid",
				},
				cfclient.SpaceQuota{
					Name: "quota-2",
					Guid: "quota-2-guid",
				},
			}, nil)
			quotas, err := quotaMgr.ListAllSpaceQuotasForOrg("orgGUID")
			Expect(err).Should(BeNil())
			Expect(fakeClient.ListOrgSpaceQuotasCallCount()).Should(Equal(1))
			orgGUID := fakeClient.ListOrgSpaceQuotasArgsForCall(0)
			Expect(orgGUID).Should(Equal("orgGUID"))
			Expect(len(quotas)).Should(Equal(2))
			Expect(quotas).Should(HaveKey("quota-1"))
			Expect(quotas).Should(HaveKey("quota-2"))
		})
		It("should return an error", func() {
			fakeClient.ListOrgSpaceQuotasReturns(nil, errors.New("error"))
			_, err := quotaMgr.ListAllSpaceQuotasForOrg("orgGUID")
			Expect(err).ShouldNot(BeNil())
			Expect(fakeClient.ListOrgSpaceQuotasCallCount()).Should(Equal(1))
		})
	})

	Context("CreateSpaceQuotas()", func() {

		BeforeEach(func() {
			spaceConfigs := []config.SpaceConfig{
				config.SpaceConfig{
					EnableSpaceQuota: true,
					Space:            "space1",
					Org:              "org1",
				},
				config.SpaceConfig{
					EnableSpaceQuota: false,
					Space:            "space2",
					Org:              "org1",
				},
			}
			fakeReader.GetSpaceConfigsReturns(spaceConfigs, nil)
			fakeSpaceMgr.FindSpaceReturns(cfclient.Space{
				Name:             "space1",
				Guid:             "space1-guid",
				OrganizationGuid: "org1-guid",
			}, nil)
		})
		It("should create a quota and assign it", func() {
			fakeClient.CreateSpaceQuotaReturns(&cfclient.SpaceQuota{Name: "space1", Guid: "space-quota-guid"}, nil)
			err := quotaMgr.CreateSpaceQuotas()
			Expect(err).Should(BeNil())
			Expect(fakeClient.CreateSpaceQuotaCallCount()).Should(Equal(1))
			quotaRequest := fakeClient.CreateSpaceQuotaArgsForCall(0)
			Expect(quotaRequest.Name).Should(Equal("space1"))
			Expect(fakeClient.AssignSpaceQuotaCallCount()).Should(Equal(1))
			quotaGUID, spaceGUID := fakeClient.AssignSpaceQuotaArgsForCall(0)
			Expect(quotaGUID).Should(Equal("space-quota-guid"))
			Expect(spaceGUID).Should(Equal("space1-guid"))
		})

		It("should error creating a quota", func() {
			fakeClient.CreateSpaceQuotaReturns(nil, errors.New("error"))
			err := quotaMgr.CreateSpaceQuotas()
			Expect(err).ShouldNot(BeNil())
			Expect(fakeClient.CreateSpaceQuotaCallCount()).Should(Equal(1))
			quotaRequest := fakeClient.CreateSpaceQuotaArgsForCall(0)
			Expect(quotaRequest.Name).Should(Equal("space1"))
		})

		It("should update a quota and assign it", func() {
			fakeClient.ListOrgSpaceQuotasReturns([]cfclient.SpaceQuota{
				cfclient.SpaceQuota{
					Name: "space1",
					Guid: "space-quota-guid",
				},
			}, nil)
			fakeClient.UpdateSpaceQuotaReturns(nil, nil)
			err := quotaMgr.CreateSpaceQuotas()
			Expect(err).Should(BeNil())
			Expect(fakeClient.UpdateSpaceQuotaCallCount()).Should(Equal(1))
			quotaGUID, quotaRequest := fakeClient.UpdateSpaceQuotaArgsForCall(0)
			Expect(quotaGUID).Should(Equal("space-quota-guid"))
			Expect(quotaRequest.Name).Should(Equal("space1"))
			Expect(fakeClient.AssignSpaceQuotaCallCount()).Should(Equal(1))
			quotaGUID, spaceGUID := fakeClient.AssignSpaceQuotaArgsForCall(0)
			Expect(quotaGUID).Should(Equal("space-quota-guid"))
			Expect(spaceGUID).Should(Equal("space1-guid"))
		})

		It("should update a quota and not assign it", func() {
			fakeSpaceMgr.FindSpaceReturns(cfclient.Space{
				Name:                "space1",
				Guid:                "space1-guid",
				OrganizationGuid:    "org1-guid",
				QuotaDefinitionGuid: "space-quota-guid",
			}, nil)
			fakeClient.ListOrgSpaceQuotasReturns([]cfclient.SpaceQuota{
				cfclient.SpaceQuota{
					Name: "space1",
					Guid: "space-quota-guid",
				},
			}, nil)
			fakeClient.UpdateSpaceQuotaReturns(nil, nil)
			err := quotaMgr.CreateSpaceQuotas()
			Expect(err).Should(BeNil())
			Expect(fakeClient.UpdateSpaceQuotaCallCount()).Should(Equal(1))
			quotaGUID, quotaRequest := fakeClient.UpdateSpaceQuotaArgsForCall(0)
			Expect(quotaGUID).Should(Equal("space-quota-guid"))
			Expect(quotaRequest.Name).Should(Equal("space1"))
			Expect(fakeClient.AssignSpaceQuotaCallCount()).Should(Equal(0))
		})

		It("should not update a quota or assign it", func() {
			fakeSpaceMgr.FindSpaceReturns(cfclient.Space{
				Name:                "space1",
				Guid:                "space1-guid",
				OrganizationGuid:    "org1-guid",
				QuotaDefinitionGuid: "space-quota-guid",
			}, nil)
			fakeClient.ListOrgSpaceQuotasReturns([]cfclient.SpaceQuota{
				cfclient.SpaceQuota{
					Name:             "space1",
					Guid:             "space-quota-guid",
					OrganizationGuid: "org1-guid",
				},
			}, nil)
			fakeClient.UpdateSpaceQuotaReturns(nil, nil)
			err := quotaMgr.CreateSpaceQuotas()
			Expect(err).Should(BeNil())
			Expect(fakeClient.UpdateSpaceQuotaCallCount()).Should(Equal(0))
			Expect(fakeClient.AssignSpaceQuotaCallCount()).Should(Equal(0))
		})

		It("should error updating a quota", func() {
			fakeClient.ListOrgSpaceQuotasReturns([]cfclient.SpaceQuota{
				cfclient.SpaceQuota{
					Name: "space1",
					Guid: "space-quota-guid",
				},
			}, nil)
			fakeClient.UpdateSpaceQuotaReturns(nil, errors.New("error"))
			err := quotaMgr.CreateSpaceQuotas()
			Expect(err).ShouldNot(BeNil())
			Expect(fakeClient.UpdateSpaceQuotaCallCount()).Should(Equal(1))
			quotaGUID, quotaRequest := fakeClient.UpdateSpaceQuotaArgsForCall(0)
			Expect(quotaGUID).Should(Equal("space-quota-guid"))
			Expect(quotaRequest.Name).Should(Equal("space1"))
			Expect(fakeClient.AssignSpaceQuotaCallCount()).Should(Equal(0))
		})

		It("should create a quota and fail to assign it", func() {
			fakeClient.CreateSpaceQuotaReturns(&cfclient.SpaceQuota{Name: "space1", Guid: "space-quota-guid"}, nil)
			fakeClient.AssignSpaceQuotaReturns(errors.New("error"))
			err := quotaMgr.CreateSpaceQuotas()
			Expect(err).ShouldNot(BeNil())
			Expect(fakeClient.CreateSpaceQuotaCallCount()).Should(Equal(1))
			quotaRequest := fakeClient.CreateSpaceQuotaArgsForCall(0)
			Expect(quotaRequest.Name).Should(Equal("space1"))
			Expect(fakeClient.AssignSpaceQuotaCallCount()).Should(Equal(1))
			quotaGUID, spaceGUID := fakeClient.AssignSpaceQuotaArgsForCall(0)
			Expect(quotaGUID).Should(Equal("space-quota-guid"))
			Expect(spaceGUID).Should(Equal("space1-guid"))
		})

		It("should peek create a quota and peek assign it", func() {
			quotaMgr.Peek = true
			err := quotaMgr.CreateSpaceQuotas()
			Expect(err).Should(BeNil())
			Expect(fakeClient.CreateSpaceQuotaCallCount()).Should(Equal(0))
			Expect(fakeClient.AssignSpaceQuotaCallCount()).Should(Equal(0))
		})

		It("Should error getting configs", func() {
			fakeReader.GetSpaceConfigsReturns(nil, errors.New("error"))
			err := quotaMgr.CreateSpaceQuotas()
			Expect(err).ShouldNot(BeNil())
		})
		It("Should error finding space", func() {
			fakeSpaceMgr.FindSpaceReturns(cfclient.Space{}, errors.New("error"))
			err := quotaMgr.CreateSpaceQuotas()
			Expect(err).ShouldNot(BeNil())
		})
		It("Should error listing space quotas", func() {
			fakeClient.ListOrgSpaceQuotasReturns(nil, errors.New("error"))
			err := quotaMgr.CreateSpaceQuotas()
			Expect(err).ShouldNot(BeNil())
		})
	})

	Context("CreateOrgQuotas()", func() {

		BeforeEach(func() {
			orgConfigs := []config.OrgConfig{
				config.OrgConfig{
					EnableOrgQuota: true,
					Org:            "org1",
				},
				config.OrgConfig{
					EnableOrgQuota: false,
					Org:            "org2",
				},
			}
			fakeReader.GetOrgConfigsReturns(orgConfigs, nil)
			fakeOrgMgr.FindOrgReturns(cfclient.Org{Name: "org1", Guid: "org-guid"}, nil)
		})
		It("should create a quota and assign it", func() {
			fakeClient.CreateOrgQuotaReturns(&cfclient.OrgQuota{Name: "org1", Guid: "org-quota-guid"}, nil)
			err := quotaMgr.CreateOrgQuotas()
			Expect(err).Should(BeNil())
			Expect(fakeClient.CreateOrgQuotaCallCount()).Should(Equal(1))
			quotaRequest := fakeClient.CreateOrgQuotaArgsForCall(0)
			Expect(quotaRequest.Name).Should(Equal("org1"))
			Expect(fakeOrgMgr.UpdateOrgCallCount()).Should(Equal(1))
			orgGUID, orgRequest := fakeOrgMgr.UpdateOrgArgsForCall(0)
			Expect(orgGUID).Should(Equal("org-guid"))
			Expect(orgRequest.QuotaDefinitionGuid).Should(Equal("org-quota-guid"))
		})

		It("should error creating a quota", func() {
			fakeClient.CreateOrgQuotaReturns(nil, errors.New("error"))
			err := quotaMgr.CreateOrgQuotas()
			Expect(err).ShouldNot(BeNil())
			Expect(fakeClient.CreateOrgQuotaCallCount()).Should(Equal(1))
			quotaRequest := fakeClient.CreateOrgQuotaArgsForCall(0)
			Expect(quotaRequest.Name).Should(Equal("org1"))
		})

		It("should update a quota and assign it", func() {
			fakeClient.ListOrgQuotasReturns([]cfclient.OrgQuota{
				cfclient.OrgQuota{
					Name:        "org1",
					Guid:        "org-quota-guid",
					TotalRoutes: 100,
				},
			}, nil)
			fakeClient.UpdateOrgQuotaReturns(nil, nil)
			err := quotaMgr.CreateOrgQuotas()
			Expect(err).Should(BeNil())
			Expect(fakeClient.UpdateOrgQuotaCallCount()).Should(Equal(1))
			quotaGUID, quotaRequest := fakeClient.UpdateOrgQuotaArgsForCall(0)
			Expect(quotaGUID).Should(Equal("org-quota-guid"))
			Expect(quotaRequest.Name).Should(Equal("org1"))
			Expect(fakeOrgMgr.UpdateOrgCallCount()).Should(Equal(1))
			orgGUID, orgRequest := fakeOrgMgr.UpdateOrgArgsForCall(0)
			Expect(orgGUID).Should(Equal("org-guid"))
			Expect(orgRequest.QuotaDefinitionGuid).Should(Equal("org-quota-guid"))
		})

		It("should update a quota and not assign it", func() {
			fakeOrgMgr.FindOrgReturns(cfclient.Org{Name: "org1", Guid: "org-guid", QuotaDefinitionGuid: "org-quota-guid"}, nil)
			fakeClient.ListOrgQuotasReturns([]cfclient.OrgQuota{
				cfclient.OrgQuota{
					Name:        "org1",
					Guid:        "org-quota-guid",
					TotalRoutes: 100,
				},
			}, nil)
			fakeClient.UpdateOrgQuotaReturns(nil, nil)
			err := quotaMgr.CreateOrgQuotas()
			Expect(err).Should(BeNil())
			Expect(fakeClient.UpdateOrgQuotaCallCount()).Should(Equal(1))
			quotaGUID, quotaRequest := fakeClient.UpdateOrgQuotaArgsForCall(0)
			Expect(quotaGUID).Should(Equal("org-quota-guid"))
			Expect(quotaRequest.Name).Should(Equal("org1"))
			Expect(fakeOrgMgr.UpdateOrgCallCount()).Should(Equal(0))
		})

		It("should not update a quota or assign it", func() {
			fakeOrgMgr.FindOrgReturns(cfclient.Org{Name: "org1", Guid: "org-guid", QuotaDefinitionGuid: "org-quota-guid"}, nil)
			fakeClient.ListOrgQuotasReturns([]cfclient.OrgQuota{
				cfclient.OrgQuota{
					Name: "org1",
					Guid: "org-quota-guid",
				},
			}, nil)
			fakeClient.UpdateOrgQuotaReturns(nil, nil)
			err := quotaMgr.CreateOrgQuotas()
			Expect(err).Should(BeNil())
			Expect(fakeClient.UpdateOrgQuotaCallCount()).Should(Equal(0))
			Expect(fakeOrgMgr.UpdateOrgCallCount()).Should(Equal(0))
		})

		It("should error updating quota", func() {
			fakeOrgMgr.FindOrgReturns(cfclient.Org{Name: "org1", Guid: "org-guid", QuotaDefinitionGuid: "org-quota-guid"}, nil)
			fakeClient.ListOrgQuotasReturns([]cfclient.OrgQuota{
				cfclient.OrgQuota{
					Name:        "org1",
					Guid:        "org-quota-guid",
					TotalRoutes: 10,
				},
			}, nil)
			fakeClient.UpdateOrgQuotaReturns(nil, errors.New("error"))
			err := quotaMgr.CreateOrgQuotas()
			Expect(err).ShouldNot(BeNil())
			Expect(fakeClient.UpdateOrgQuotaCallCount()).Should(Equal(1))
			Expect(fakeOrgMgr.UpdateOrgCallCount()).Should(Equal(0))
		})

		It("should error assigning quota", func() {
			fakeOrgMgr.FindOrgReturns(cfclient.Org{Name: "org1", Guid: "org-guid", QuotaDefinitionGuid: "org-quota-guid"}, nil)
			fakeClient.ListOrgQuotasReturns([]cfclient.OrgQuota{
				cfclient.OrgQuota{
					Name:        "org1",
					Guid:        "org-quota-guid2",
					TotalRoutes: 100,
				},
			}, nil)
			fakeClient.UpdateOrgQuotaReturns(nil, nil)
			fakeOrgMgr.UpdateOrgReturns(cfclient.Org{}, errors.New("error"))
			err := quotaMgr.CreateOrgQuotas()
			Expect(err).ShouldNot(BeNil())
			Expect(fakeClient.UpdateOrgQuotaCallCount()).Should(Equal(1))
			quotaGUID, quotaRequest := fakeClient.UpdateOrgQuotaArgsForCall(0)
			Expect(quotaGUID).Should(Equal("org-quota-guid2"))
			Expect(quotaRequest.Name).Should(Equal("org1"))
			Expect(fakeOrgMgr.UpdateOrgCallCount()).Should(Equal(1))
		})
		It("should peek create a quota and peek assign it", func() {
			quotaMgr.Peek = true
			err := quotaMgr.CreateOrgQuotas()
			Expect(err).Should(BeNil())
			Expect(fakeClient.CreateOrgQuotaCallCount()).Should(Equal(0))
			Expect(fakeOrgMgr.UpdateOrgCallCount()).Should(Equal(0))
		})

		It("Should error getting configs", func() {
			fakeReader.GetOrgConfigsReturns(nil, errors.New("error"))
			err := quotaMgr.CreateOrgQuotas()
			Expect(err).ShouldNot(BeNil())
		})
		It("Should error finding org", func() {
			fakeOrgMgr.FindOrgReturns(cfclient.Org{}, errors.New("error"))
			err := quotaMgr.CreateOrgQuotas()
			Expect(err).ShouldNot(BeNil())
		})
		It("Should error listing org quotas", func() {
			fakeClient.ListOrgQuotasReturns(nil, errors.New("error"))
			err := quotaMgr.CreateOrgQuotas()
			Expect(err).ShouldNot(BeNil())
		})
	})

	Context("UpdateSpaceQuota()", func() {
		It("should update a quota", func() {
			fakeClient.UpdateSpaceQuotaReturns(nil, nil)

			err := quotaMgr.UpdateSpaceQuota("quotaGUID", cfclient.SpaceQuotaRequest{Name: "quota"})
			Expect(err).Should(BeNil())
			Expect(fakeClient.UpdateSpaceQuotaCallCount()).Should(Equal(1))
		})
		It("should peek and not update a quota", func() {
			quotaMgr.Peek = true
			fakeClient.UpdateSpaceQuotaReturns(nil, nil)

			err := quotaMgr.UpdateSpaceQuota("quotaGUID", cfclient.SpaceQuotaRequest{Name: "quota"})
			Expect(err).Should(BeNil())
			Expect(fakeClient.UpdateSpaceQuotaCallCount()).Should(Equal(0))
		})
		It("should return an error", func() {
			fakeClient.UpdateSpaceQuotaReturns(nil, errors.New("error"))

			err := quotaMgr.UpdateSpaceQuota("quotaGUID", cfclient.SpaceQuotaRequest{})
			Expect(err).ShouldNot(BeNil())
		})
	})

	Context("CreateSpaceQuota()", func() {
		It("should create a quota", func() {
			fakeClient.CreateSpaceQuotaReturns(nil, nil)

			_, err := quotaMgr.CreateSpaceQuota(cfclient.SpaceQuotaRequest{Name: "quota"})
			Expect(err).Should(BeNil())
			Expect(fakeClient.CreateSpaceQuotaCallCount()).Should(Equal(1))
		})
		It("should peek and not create a quota", func() {
			quotaMgr.Peek = true
			fakeClient.CreateSpaceQuotaReturns(nil, nil)

			_, err := quotaMgr.CreateSpaceQuota(cfclient.SpaceQuotaRequest{Name: "quota"})
			Expect(err).Should(BeNil())
			Expect(fakeClient.CreateSpaceQuotaCallCount()).Should(Equal(0))
		})
		It("should return an error", func() {
			fakeClient.CreateSpaceQuotaReturns(nil, errors.New("error"))

			_, err := quotaMgr.CreateSpaceQuota(cfclient.SpaceQuotaRequest{Name: "quota"})
			Expect(err).ShouldNot(BeNil())
		})
	})

	/*It("should update 2 quota", func() {
	  spaces := []cfclient.Space{
	    {
	      Name:             "space1",
	      OrganizationGuid: "testOrgGUID",
	      Guid:             "space1GUID",
	    },
	    {
	      Name:             "space2",
	      OrganizationGuid: "testOrgGUID",
	      Guid:             "space2GUID",
	    },
	  }
	  quotas := make(map[string]string)
	  quotas["space1"] = "space1QuotaGUID"
	  quotas["space2"] = "space2QuotaGUID"
	  mockOrgMgr.EXPECT().GetOrgGUID("test").Return("testOrgGUID", nil)
	  mockCloudController.EXPECT().ListSpaces("testOrgGUID").Return(spaces, nil)
	  mockCloudController.EXPECT().ListAllSpaceQuotasForOrg("testOrgGUID").Return(quotas, nil)
	  mockCloudController.EXPECT().UpdateSpaceQuota("space1QuotaGUID", cloudcontroller.SpaceQuotaEntity{
	    OrgGUID: "testOrgGUID",
	    QuotaEntity: cloudcontroller.QuotaEntity{
	      Name:                    "space1",
	      MemoryLimit:             10240,
	      InstanceMemoryLimit:     -1,
	      TotalRoutes:             10,
	      TotalServices:           -1,
	      PaidServicePlansAllowed: true,
	      AppInstanceLimit:        -1,
	      TotalReservedRoutePorts: 0,
	      TotalPrivateDomains:     -1,
	      TotalServiceKeys:        -1,
	    }}).Return(nil)
	  mockCloudController.EXPECT().AssignQuotaToSpace("space1GUID", "space1QuotaGUID")

	  mockOrgMgr.EXPECT().GetOrgGUID("test").Return("testOrgGUID", nil)
	  mockCloudController.EXPECT().ListSpaces("testOrgGUID").Return(spaces, nil)
	  mockCloudController.EXPECT().ListAllSpaceQuotasForOrg("testOrgGUID").Return(quotas, nil)
	  mockCloudController.EXPECT().UpdateSpaceQuota("space2QuotaGUID", cloudcontroller.SpaceQuotaEntity{
	    OrgGUID: "testOrgGUID",
	    QuotaEntity: cloudcontroller.QuotaEntity{
	      Name:                    "space2",
	      MemoryLimit:             10240,
	      InstanceMemoryLimit:     -1,
	      TotalRoutes:             10,
	      TotalServices:           -1,
	      PaidServicePlansAllowed: true,
	      AppInstanceLimit:        -1,
	      TotalReservedRoutePorts: 0,
	      TotalPrivateDomains:     -1,
	      TotalServiceKeys:        -1,
	    }}).Return(nil)
	  mockCloudController.EXPECT().AssignQuotaToSpace("space2GUID", "space2QuotaGUID")
	  err := spaceManager.CreateQuotas("./fixtures/config")
	  Ω(err).Should(BeNil())
	})*/

})
