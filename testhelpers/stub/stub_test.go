package stub_test

import (
	. "github.com/cnupp/appssdk/testhelpers/stub"

	"bytes"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"io/ioutil"
	"net/http"
)

var _ = Describe("Stub", func() {
	It("should able to start the stub server", func() {
		server, _ := NewStub([]TestRequest{
			TestRequest{
				Method: "GET",
				Path:   "/path",
				Response: TestResponse{
					Status: 200,
					Header: http.Header{
						"Content-Type": {"application/json"},
						"Set-Cookie":   {"cookie=cookie"},
					},
					Body: "content",
				},
			},
		})

		var body []byte
		req, err := http.NewRequest("GET", server.URL+"/path", bytes.NewBuffer(body))
		Expect(err).To(BeNil())

		client := http.Client{}

		res, err := client.Do(req)
		Expect(err).To(BeNil())

		Expect(res.StatusCode).To(Equal(http.StatusOK))
		Expect(res.Header.Get("Content-Type")).To(Equal("application/json"))
		Expect(res.Header.Get("Set-Cookie")).To(Equal("cookie=cookie"))

		bodyInBytes, err := ioutil.ReadAll(res.Body)
		Expect(err).To(BeNil())
		Expect(string(bodyInBytes)).To(Equal("content"))
	})

	It("should able to get the second request reponse when to stub defined", func() {
		server, _ := NewStub([]TestRequest{
			TestRequest{
				Method: "GET",
				Path:   "/path",
				Response: TestResponse{
					Status: 200,
					Header: http.Header{
						"Content-Type": {"application/json"},
						"Set-Cookie":   {"cookie=cookie"},
					},
					Body: "content",
				},
			},
			TestRequest{
				Method: "GET",
				Path:   "/another-path",
				Response: TestResponse{
					Status: 200,
					Header: http.Header{
						"Content-Type": {"application/json"},
						"Set-Cookie":   {"cookie=newcookie"},
					},
					Body: "anothercontent",
				},
			},
		})

		var body []byte
		req, err := http.NewRequest("GET", server.URL+"/another-path", bytes.NewBuffer(body))
		Expect(err).To(BeNil())

		client := http.Client{}

		res, err := client.Do(req)
		Expect(err).To(BeNil())

		Expect(res.StatusCode).To(Equal(http.StatusOK))
		Expect(res.Header.Get("Content-Type")).To(Equal("application/json"))
		Expect(res.Header.Get("Set-Cookie")).To(Equal("cookie=newcookie"))

		bodyInBytes, err := ioutil.ReadAll(res.Body)
		Expect(err).To(BeNil())
		Expect(string(bodyInBytes)).To(Equal("anothercontent"))
	})

	It("should able to 404 when no request stub defined", func() {
		server, _ := NewStub([]TestRequest{})

		var body []byte
		req, err := http.NewRequest("GET", server.URL+"/not-exists-path", bytes.NewBuffer(body))
		Expect(err).To(BeNil())

		client := http.Client{}

		res, err := client.Do(req)
		Expect(err).To(BeNil())

		Expect(res.StatusCode).To(Equal(http.StatusNotFound))
	})

	It("should able to check the spy data", func() {
		server, _ := NewStub([]TestRequest{
			TestRequest{
				Method: "PUT",
				Path:   "/path",
				Matcher: func(r *http.Request) {
					Expect(r.Header.Get("csrf")).To(Equal("csrf"))
				},
				Response: TestResponse{
					Status: 200,
					Header: http.Header{
						"Content-Type": {"application/json"},
						"Set-Cookie":   {"cookie=cookie"},
					},
					Body: "content",
				},
			},
		})

		var body []byte
		req, err := http.NewRequest("PUT", server.URL+"/path", bytes.NewBuffer(body))
		req.Header.Add("csrf", "csrf")
		Expect(err).To(BeNil())

		client := http.Client{}

		res, err := client.Do(req)
		Expect(err).To(BeNil())

		Expect(res.StatusCode).To(Equal(http.StatusOK))
	})
})
