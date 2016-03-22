package api_test

import (
	. "github.com/sjkyspa/stacks/controller/api/api"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/sjkyspa/stacks/controller/api/net"
	testconfig "github.com/sjkyspa/stacks/controller/api/testhelpers/config"
	testnet "github.com/sjkyspa/stacks/controller/api/testhelpers/net"
	"net/http"
	"net/http/httptest"
)

var _ = Describe("Apps", func() {
	var createAppRequest = testnet.TestRequest{
		Method: "POST",
		Path:   "/apps",
		Response: testnet.TestResponse{
			Status: 201,
			Header: http.Header{
				"accept":   {"application/json"},
				"Location": {"/apps/ketsu"},
			},
		},
	}
	var getAppResponse = `
	{
	  "id": "b78dba51-8daf-4fe9-9345-c7ab582c3387",
	  "name": "ketsu",
	  "memory": 30,
	  "disk": 30,
	  "instances": 1,
	  "envs": {
	    "ENV": "PRODUCTION"
	  },
	  "links": [
		{
		  "rel": "self",
		  "uri": "/apps/ketsu"
		},
		{
		  "rel": "env",
		  "uri": "/apps/ketsu/env"
		},
		{
		  "rel": "routes",
		  "uri": "/apps/ketsu/routes"
		},
		{
		  "rel": "builds",
		  "uri": "/apps/ketsu/builds"
		},
		{
		  "rel": "releases",
		  "uri": "/apps/ketsu/releases"
		},
		{
		  "rel": "stack",
		  "uri": "/stacks/ketsu"
		}
	  ]
	}
	`

	var getAppRequest = testnet.TestRequest{
		Method: "GET",
		Path:   "/apps/ketsu",
		Response: testnet.TestResponse{
			Status: 200,
			Header: http.Header{
				"Content-Type": {"application/json"},
			},
			Body: getAppResponse,
		},
	}

	var getAppsResponse = `
	{
	  "count": 2,
	  "self": "/apps?page=1&per_page=30",
	  "first": "/apps?page=1&per_page=30",
	  "last": "/apps?page=1&per_page=30",
	  "prev": null,
	  "next": null,
	  "items": [
		{
		  "id": "b78dba51-8daf-4fe9-9345-c7ab582c3387",
		  "name": "ketsu",
		  "links": [
			{
			  "rel": "self",
			  "uri": "/apps/b78dba51-8daf-4fe9-9345-c7ab582c3387"
			},
			{
			  "rel": "env",
			  "uri": "/apps/b78dba51-8daf-4fe9-9345-c7ab582c3387/env"
			},
			{
			  "rel": "routes",
			  "uri": "/apps/b78dba51-8daf-4fe9-9345-c7ab582c3387/routes"
			},
			{
			  "rel": "builds",
			  "uri": "/apps/b78dba51-8daf-4fe9-9345-c7ab582c3387/builds"
			},
			{
			  "rel": "releases",
			  "uri": "/apps/b78dba51-8daf-4fe9-9345-c7ab582c3387/releases"
			},
			{
			  "rel": "stack",
			  "uri": "/stacks/74a052c9-76b3-44a1-ac0b-666faa1223b6"
			}
		  ]
		}
	  ]
	}
	`

	var getAppsRequest = testnet.TestRequest{
		Method: "GET",
		Path:   "/apps",
		Response: testnet.TestResponse{
			Status: 200,
			Header: http.Header{
				"Content-Type": {"application/json"},
			},
			Body: getAppsResponse,
		},
	}

	var destroyAppRequest = testnet.TestRequest{
		Method: "DELETE",
		Path:   "/apps/ketsu",
		Response: testnet.TestResponse{
			Status: 200,
			Header: http.Header{
				"Content-Type": {"application/json"},
			},
		},
	}

	var createAppRepository = func(requests []testnet.TestRequest) (ts *httptest.Server, handler *testnet.TestHandler, repo AppRepository) {
		ts, handler = testnet.NewServer(requests)
		configRepo := testconfig.NewRepositoryWithDefaults()
		configRepo.SetApiEndpoint(ts.URL)
		gateway := net.NewCloudControllerGateway(configRepo)
		repo = NewAppRepository(configRepo, gateway)
		return
	}

	var defaultAppParams = func() AppParams {
		name := "ketsu"

		return AppParams{
			Name:      name,
			Stack:     "/stacks/stackid",
		}
	}

	It("should able to create an app", func() {
		ts, _, repo := createAppRepository([]testnet.TestRequest{createAppRequest, getAppRequest})
		defer ts.Close()

		createdApp, err := repo.Create(defaultAppParams())
		Expect(err).To(BeNil())
		Expect(createdApp.Id()).To(Equal("ketsu"))
		Expect(createdApp.Links()).NotTo(BeNil())
	})

	It("should able to get an app", func() {
		ts, _, repo := createAppRepository([]testnet.TestRequest{getAppRequest})
		defer ts.Close()

		createdApp, err := repo.GetApp("ketsu")
		Expect(err).To(BeNil())
		Expect(createdApp.Id()).To(Equal("ketsu"))
		Expect(createdApp.GetEnvs()["ENV"]).To(Equal("PRODUCTION"))
		Expect(createdApp.Links()).NotTo(BeNil())
	})

	It("should able to get apps", func() {
		ts, _, repo := createAppRepository([]testnet.TestRequest{getAppsRequest})
		defer ts.Close()

		createdApps, err := repo.GetApps()
		Expect(err).To(BeNil())
		Expect(createdApps.Count()).To(Equal(2))
		Expect(createdApps.Items()[0].Id()).To(Equal("ketsu"))
		Expect(createdApps.Items()[0].Links()).NotTo(BeNil())
	})

	It("should able to delete apps", func() {
		ts, _, repo := createAppRepository([]testnet.TestRequest{destroyAppRequest})
		defer ts.Close()

		err := repo.Delete("ketsu")
		Expect(err).To(BeNil())
	})
})