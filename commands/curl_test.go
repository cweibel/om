package commands_test

import (
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"strings"

	"github.com/pivotal-cf/om/api"
	"github.com/pivotal-cf/om/commands"
	"github.com/pivotal-cf/om/commands/fakes"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

type errReader struct{}

func (er errReader) Read([]byte) (int, error) {
	return 0, errors.New("failed to read")
}

var _ = Describe("Curl", func() {
	stringCloser := func(s string) io.ReadCloser {
		return ioutil.NopCloser(strings.NewReader(s))
	}
	Describe("Execute", func() {
		var (
			command     *commands.Curl
			fakeService *fakes.CurlService
			stdout      *fakes.Logger
			stderr      *fakes.Logger
		)

		BeforeEach(func() {
			fakeService = &fakes.CurlService{}
			stdout = &fakes.Logger{}
			stderr = &fakes.Logger{}
			command = commands.NewCurl(fakeService, stdout, stderr)
		})

		It("executes the API call", func() {
			fakeService.CurlReturns(api.RequestServiceCurlOutput{
				StatusCode: http.StatusOK,
				Headers: http.Header{
					"Content-Length": []string{"33"},
					"Content-Type":   []string{"application/json"},
					"Accept":         []string{"text/plain"},
				},
				Body: stringCloser(`{"some-response-key": "%some-response-value"}`),
			}, nil)

			err := executeCommand(command, []string{
				"--path", "/api/v0/some/path",
				"--request", "POST",
				"--data", `{"some-key": "some-value"}`,
			}, nil)
			Expect(err).NotTo(HaveOccurred())

			input := fakeService.CurlArgsForCall(0)
			Expect(input.Path).To(Equal("/api/v0/some/path"))
			Expect(input.Method).To(Equal("POST"))
			Expect(input.Headers).To(HaveKeyWithValue("Content-Type", []string{"application/json"}))

			data, err := ioutil.ReadAll(input.Data)
			Expect(err).NotTo(HaveOccurred())
			Expect(string(data)).To(Equal(`{"some-key": "some-value"}`))
			content := stdout.PrintlnArgsForCall(0)
			Expect(fmt.Sprint(content...)).To(MatchJSON(`{"some-response-key": "%some-response-value"}`))

			format, content := stderr.PrintfArgsForCall(0)
			Expect(fmt.Sprintf(format, content...)).To(Equal("Status: 200 OK"))

			format, content = stderr.PrintfArgsForCall(1)
			Expect(fmt.Sprintf(format, content...)).To(Equal("Accept: text/plain\r\nContent-Length: 33\r\nContent-Type: application/json\r\n"))
		})

		Context("when --silent is specified", func() {
			It("does not write anything to stderr if the status is 200", func() {
				fakeService.CurlReturns(api.RequestServiceCurlOutput{
					StatusCode: http.StatusOK,
					Headers: http.Header{
						"Content-Type": []string{"application/json"},
					},
					Body: stringCloser("{}"),
				}, nil)

				err := executeCommand(command, []string{
					"--path", "/api/v0/some/path",
					"--request", "GET",
					"--silent",
				}, nil)
				Expect(err).NotTo(HaveOccurred())

				Expect(stderr.Invocations()).To(BeEmpty())
			})

			It("does not write anything to stderr if the status is 201", func() {
				fakeService.CurlReturns(api.RequestServiceCurlOutput{
					StatusCode: http.StatusCreated,
					Headers: http.Header{
						"Content-Type": []string{"application/json"},
					},
					Body: stringCloser("{}"),
				}, nil)

				err := executeCommand(command, []string{
					"--path", "/api/v0/some/path",
					"--request", "POST",
					"--silent",
				}, nil)
				Expect(err).NotTo(HaveOccurred())

				Expect(stderr.Invocations()).To(HaveLen(0))
			})

			It("still writes response headers to stderr if the status is 404", func() {
				fakeService.CurlReturns(api.RequestServiceCurlOutput{
					StatusCode: http.StatusNotFound,
					Headers: http.Header{
						"Content-Type": []string{"application/json"},
					},
					Body: stringCloser("{}"),
				}, nil)

				err := executeCommand(command, []string{
					"--path", "/api/v0/some/path",
					"--request", "GET",
					"--silent",
				}, nil)
				Expect(err).To(MatchError("server responded with an error"))

				format, content := stderr.PrintfArgsForCall(0)
				Expect(fmt.Sprintf(format, content...)).To(Equal("Status: 404 Not Found"))

				format, content = stderr.PrintfArgsForCall(1)
				Expect(fmt.Sprintf(format, content...)).To(Equal("Content-Type: application/json\r\n"))
			})
		})

		Context("When a custom content-type is passed in", func() {
			It("executes the API call with the given content type", func() {
				fakeService.CurlReturns(api.RequestServiceCurlOutput{
					Headers: http.Header{
						"Content-Length": []string{"33"},
						"Content-Type":   []string{"application/json"},
						"Accept":         []string{"text/plain"},
					},
					Body: stringCloser(`{"some-response-key": "%some-response-value"}`),
				}, nil)

				err := executeCommand(command, []string{
					"--path", "/api/v0/some/path",
					"--request", "POST",
					"--data", `some_key=some_value`,
					"--header", "Content-Type: application/x-www-form-urlencoded",
				}, nil)
				Expect(err).NotTo(HaveOccurred())

				input := fakeService.CurlArgsForCall(0)
				Expect(input.Path).To(Equal("/api/v0/some/path"))
				Expect(input.Method).To(Equal("POST"))
				Expect(input.Headers).To(HaveKeyWithValue("Content-Type", []string{"application/x-www-form-urlencoded"}))
			})
		})

		Describe("pretty printing", func() {
			Context("when the response type is JSON", func() {
				It("pretty prints the response body", func() {
					fakeService.CurlReturns(api.RequestServiceCurlOutput{
						Headers: http.Header{
							"Content-Length": []string{"33"},
							"Content-Type":   []string{"application/json; charset=utf-8"},
						},
						Body: stringCloser(`{"some-response-key": "some-response-value"}`),
					}, nil)

					err := executeCommand(command, []string{
						"--path", "/api/v0/some/path",
						"--request", "POST",
						"--data", `{"some-key": "some-value"}`,
					}, nil)
					Expect(err).NotTo(HaveOccurred())

					content := stdout.PrintlnArgsForCall(0)
					Expect(fmt.Sprint(content...)).To(Equal("{\n  \"some-response-key\": \"some-response-value\"\n}"))
				})
			})

			Context("when the response type is not JSON", func() {
				It("doesn't format the response body", func() {
					fakeService.CurlReturns(api.RequestServiceCurlOutput{
						Headers: http.Header{
							"Content-Length": []string{"33"},
							"Content-Type":   []string{"text/plain; charset=utf-8"},
						},
						Body: stringCloser(`{"some-response-key": "some-response-value"}`),
					}, nil)

					err := executeCommand(command, []string{
						"--path", "/api/v0/some/path",
						"--request", "POST",
						"--data", `{"some-key": "some-value"}`,
					}, nil)
					Expect(err).NotTo(HaveOccurred())

					content := stdout.PrintlnArgsForCall(0)
					Expect(fmt.Sprint(content...)).To(Equal(`{"some-response-key": "some-response-value"}`))
				})
			})
		})

		Context("failure cases", func() {
			Context("when the flags cannot be parsed", func() {
				It("returns an error", func() {
					err := executeCommand(command, []string{"--bad-flag", "some-value"}, nil)
					Expect(err).To(MatchError("unknown flag `bad-flag'"))
				})
			})

			Context("when the request path is not provided", func() {
				It("returns an error", func() {
					err := executeCommand(command, []string{
						"--request", "GET",
						"--data", `{"some-key": "some-value"}`,
					}, nil)
					Expect(err.Error()).To(MatchRegexp("the required flag.*--path"))
				})
			})

			Context("when the request service returns an error", func() {
				It("returns an error", func() {
					fakeService.CurlReturns(api.RequestServiceCurlOutput{}, errors.New("some request error"))
					err := executeCommand(command, []string{
						"--path", "/api/v0/some/path",
						"--request", "POST",
						"--data", `{"some-key": "some-value"}`,
					}, nil)
					Expect(err).To(MatchError("failed to make api request: some request error"))
				})
			})

			Context("when the response body cannot be read", func() {
				It("returns an error", func() {
					fakeService.CurlReturns(api.RequestServiceCurlOutput{
						Body: ioutil.NopCloser(errReader{}),
					}, nil)
					err := executeCommand(command, []string{
						"--path", "/api/v0/some/path",
						"--request", "POST",
						"--data", `{"some-key": "some-value"}`,
					}, nil)
					Expect(err).To(MatchError("failed to read api response body: failed to read"))
				})
			})

			Context("when the response code is 400 or higher", func() {
				It("returns an error", func() {
					fakeService.CurlReturns(api.RequestServiceCurlOutput{
						StatusCode: 401,
						Body:       stringCloser(`{"some-response-key": "some-response-value"}`),
					}, nil)

					err := executeCommand(command, []string{
						"--path", "/api/v0/some/path",
						"--request", "POST",
						"--data", `{"some-key": "some-value"}`,
					}, nil)
					Expect(err).To(MatchError("server responded with an error"))
				})
			})
		})
	})
})
