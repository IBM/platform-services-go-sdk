/**
 * (C) Copyright IBM Corp. 2023.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *      http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package main_test

import (
	"fmt"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	. "github.com/onsi/gomega/gexec"
	"net/http"
	"net/http/httptest"
	"os/exec"
)

const credentialErrMsg = "Could not authenticate the plugin.\nYou must either log in with 'ibmcloud login', export credentials as environment variables, or store them in a credentials file.\nFor more information, see https://github.com/IBM/ibm-cloud-sdk-common/blob/main/README.md#define-configuration-properties.\nError"

var testExecutable string

var _ = BeforeSuite(func() {
	var err error
	testExecutable, err = Build("./main.go")
	Expect(err).To(BeNil())
})

var _ = AfterSuite(func() {
	CleanupBuildArtifacts()
})

var _ = Describe("project CLI plugin", func() {
	var testServer *httptest.Server
	var testEnvironment []string

	runCmd := func(arguments []string) (*Session, error) {
		cmd := exec.Command(testExecutable, arguments...)
		cmd.Env = append(cmd.Env, testEnvironment...)

		session, err := Start(cmd, GinkgoWriter, GinkgoWriter)
		// Timeout after 10 seconds.
		session.Wait(10)

		return session, err
	}

	Describe("Run the `list` command", func() {
		operationPath := "/v1/projects"

		Context("successfully", func() {
			BeforeEach(func() {
				// Create the mock server.
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request.
					Expect(req.URL.EscapedPath()).To(Equal(operationPath))
					Expect(req.Method).To(Equal("GET"))
					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"limit": 1, "total_count": 0, "first": {"href": "testString", "start": "testString"}, "last": {"href": "testString", "start": "testString"}, "previous": {"href": "testString", "start": "testString"}, "next": {"href": "testString", "start": "testString"}, "projects": [{"id": "testString", "name": "testString", "description": "testString", "metadata": {"crn": "crn:v1:staging:public:project:global:a/4e1c48fcf8ac33c0a2441e4139f189ae:bf40ad13-b107-446a-8286-c6d576183bb1::", "created_at": "2019-01-01T12:00:00.000Z", "cumulative_needs_attention_view": {"event": "testString", "event_id": "testString", "config_id": "testString", "config_version": 38}, "cumulative_needs_attention_view_err": "testString", "location": "testString", "resource_group": "testString", "state": "testString"}}]}`)
				}))

				// Create the test environment.
				testEnvironment = append(testEnvironment, "PROJECTS_URL="+testServer.URL)
				testEnvironment = append(testEnvironment, "PROJECTS_AUTH_TYPE=noAuth")
			})

			AfterEach(func() {
				testServer.Close()
				testEnvironment = []string{}
			})

			It("with all flags", func() {
				// Put together mock arguments.
				args := []string{
					"list",
					`--start=testString`,
					`--limit=10`,
					`--complete=false`,
				}

				session, err := runCmd(args)

				Expect(err).To(BeNil())
				Expect(session).To(Exit(0))
				Expect(session.Out.Contents()).ToNot(BeNil())
			})

			It("with all flags (JSON output)", func() {
				// Put together mock arguments.
				args := []string{
					"list",
					`--start=testString`,
					`--limit=10`,
					`--complete=false`,
					`--output=json`,
				}

				session, err := runCmd(args)

				Expect(err).To(BeNil())
				Expect(session).To(Exit(0))

				out := session.Out.Contents()
				mockResponse := `{"limit": 1, "total_count": 0, "first": {"href": "testString", "start": "testString"}, "last": {"href": "testString", "start": "testString"}, "previous": {"href": "testString", "start": "testString"}, "next": {"href": "testString", "start": "testString"}, "projects": [{"id": "testString", "name": "testString", "description": "testString", "metadata": {"crn": "crn:v1:staging:public:project:global:a/4e1c48fcf8ac33c0a2441e4139f189ae:bf40ad13-b107-446a-8286-c6d576183bb1::", "created_at": "2019-01-01T12:00:00.000Z", "cumulative_needs_attention_view": {"event": "testString", "event_id": "testString", "config_id": "testString", "config_version": 38}, "cumulative_needs_attention_view_err": "testString", "location": "testString", "resource_group": "testString", "state": "testString"}}]}`
				Expect(string(out)).To(MatchJSON(mockResponse))
			})
		})

		Context("unsuccessfully", func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request.
					Expect(req.URL.EscapedPath()).To(Equal(operationPath))
					Expect(req.Method).To(Equal("GET"))
					// Set mock response
					res.WriteHeader(404)
				}))

				// Create the test environment.
				testEnvironment = append(testEnvironment, "PROJECTS_URL="+testServer.URL)
				testEnvironment = append(testEnvironment, "PROJECTS_AUTH_TYPE=noAuth")
			})

			AfterEach(func() {
				testServer.Close()
				testEnvironment = []string{}
			})

			It("because authentication error", func() {
				// Put together mock arguments.
				args := []string{
					"list",
					`--start=testString`,
					`--limit=10`,
					`--complete=false`,
				}

				// Clear the test environment that holds the credentials.
				testEnvironment = []string{}

				session, err := runCmd(args)

				Expect(err).To(BeNil())
				Expect(session).To(Exit(1))
				Expect(session.Err.Contents()).To(ContainSubstring(credentialErrMsg))
			})

			It("because response error", func() {
				args := []string{
					"list",
					`--start=testString`,
					`--limit=10`,
					`--complete=false`,
				}

				session, err := runCmd(args)

				Expect(err).To(BeNil())
				Expect(session).To(Exit(1))
				Expect(session.Err.Contents()).To(MatchRegexp("^FAILED\n.+\nNot Found\n\n$"))
			})
		})
	})

	Describe("Run the `get` command", func() {
		operationPath := "/v1/projects/testString"

		Context("successfully", func() {
			BeforeEach(func() {
				// Create the mock server.
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request.
					Expect(req.URL.EscapedPath()).To(Equal(operationPath))
					Expect(req.Method).To(Equal("GET"))
					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"name": "testString", "description": "testString", "id": "testString", "crn": "crn:v1:staging:public:project:global:a/4e1c48fcf8ac33c0a2441e4139f189ae:bf40ad13-b107-446a-8286-c6d576183bb1::", "configs": [{"id": "testString", "name": "testString", "labels": ["testString"], "description": "testString", "locator_id": "testString", "type": "terraform_template", "input": [{"name": "testString", "type": "array", "required": true}], "output": [{"name": "testString", "description": "testString", "value": ["testString"]}], "setting": [{"name": "testString", "value": "testString"}]}], "metadata": {"crn": "crn:v1:staging:public:project:global:a/4e1c48fcf8ac33c0a2441e4139f189ae:bf40ad13-b107-446a-8286-c6d576183bb1::", "created_at": "2019-01-01T12:00:00.000Z", "cumulative_needs_attention_view": {"event": "testString", "event_id": "testString", "config_id": "testString", "config_version": 38}, "cumulative_needs_attention_view_err": "testString", "location": "testString", "resource_group": "testString", "state": "testString", "event_notifications_crn": "testString"}}`)
				}))

				// Create the test environment.
				testEnvironment = append(testEnvironment, "PROJECTS_URL="+testServer.URL)
				testEnvironment = append(testEnvironment, "PROJECTS_AUTH_TYPE=noAuth")
			})

			AfterEach(func() {
				testServer.Close()
				testEnvironment = []string{}
			})

			It("with all flags", func() {
				// Put together mock arguments.
				args := []string{
					"get",
					`--id=testString`,
					`--exclude-configs=false`,
					`--complete=false`,
				}

				session, err := runCmd(args)

				Expect(err).To(BeNil())
				Expect(session).To(Exit(0))
				Expect(session.Out.Contents()).ToNot(BeNil())
			})

			It("with all flags (JSON output)", func() {
				// Put together mock arguments.
				args := []string{
					"get",
					`--id=testString`,
					`--exclude-configs=false`,
					`--complete=false`,
					`--output=json`,
				}

				session, err := runCmd(args)

				Expect(err).To(BeNil())
				Expect(session).To(Exit(0))

				out := session.Out.Contents()
				mockResponse := `{"name": "testString", "description": "testString", "id": "testString", "crn": "crn:v1:staging:public:project:global:a/4e1c48fcf8ac33c0a2441e4139f189ae:bf40ad13-b107-446a-8286-c6d576183bb1::", "configs": [{"id": "testString", "name": "testString", "labels": ["testString"], "description": "testString", "locator_id": "testString", "type": "terraform_template", "input": [{"name": "testString", "type": "array", "required": true}], "output": [{"name": "testString", "description": "testString", "value": ["testString"]}], "setting": [{"name": "testString", "value": "testString"}]}], "metadata": {"crn": "crn:v1:staging:public:project:global:a/4e1c48fcf8ac33c0a2441e4139f189ae:bf40ad13-b107-446a-8286-c6d576183bb1::", "created_at": "2019-01-01T12:00:00.000Z", "cumulative_needs_attention_view": {"event": "testString", "event_id": "testString", "config_id": "testString", "config_version": 38}, "cumulative_needs_attention_view_err": "testString", "location": "testString", "resource_group": "testString", "state": "testString", "event_notifications_crn": "testString"}}`
				Expect(string(out)).To(MatchJSON(mockResponse))
			})
		})

		Context("unsuccessfully", func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request.
					Expect(req.URL.EscapedPath()).To(Equal(operationPath))
					Expect(req.Method).To(Equal("GET"))
					// Set mock response
					res.WriteHeader(404)
				}))

				// Create the test environment.
				testEnvironment = append(testEnvironment, "PROJECTS_URL="+testServer.URL)
				testEnvironment = append(testEnvironment, "PROJECTS_AUTH_TYPE=noAuth")
			})

			AfterEach(func() {
				testServer.Close()
				testEnvironment = []string{}
			})

			It("because authentication error", func() {
				// Put together mock arguments.
				args := []string{
					"get",
					`--id=testString`,
					`--exclude-configs=false`,
					`--complete=false`,
				}

				// Clear the test environment that holds the credentials.
				testEnvironment = []string{}

				session, err := runCmd(args)

				Expect(err).To(BeNil())
				Expect(session).To(Exit(1))
				Expect(session.Err.Contents()).To(ContainSubstring(credentialErrMsg))
			})

			It("because response error", func() {
				args := []string{
					"get",
					`--id=testString`,
					`--exclude-configs=false`,
					`--complete=false`,
				}

				session, err := runCmd(args)

				Expect(err).To(BeNil())
				Expect(session).To(Exit(1))
				Expect(session.Err.Contents()).To(MatchRegexp("^FAILED\n.+\nNot Found\n\n$"))
			})

			It("because missing required flag", func() {
				args := []string{
					"get",
				}
				session, err := runCmd(args)

				Expect(err).To(BeNil())
				Expect(session).To(Exit(1))
				Expect(session.Err.Contents()).To(ContainSubstring("required flag(s)"))
				Expect(session.Err.Contents()).To(ContainSubstring("not set"))
			})
		})
	})

	Describe("Run the `delete` command", func() {
		operationPath := "/v1/projects/testString"

		Context("successfully", func() {
			BeforeEach(func() {
				// Create the mock server.
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request.
					Expect(req.URL.EscapedPath()).To(Equal(operationPath))
					Expect(req.Method).To(Equal("DELETE"))
					res.WriteHeader(204)
				}))

				// Create the test environment.
				testEnvironment = append(testEnvironment, "PROJECTS_URL="+testServer.URL)
				testEnvironment = append(testEnvironment, "PROJECTS_AUTH_TYPE=noAuth")
			})

			AfterEach(func() {
				testServer.Close()
				testEnvironment = []string{}
			})

			It("with all flags", func() {
				// Put together mock arguments.
				args := []string{
					"delete",
					`--id=testString`,
				}

				session, err := runCmd(args)

				Expect(err).To(BeNil())
				Expect(session).To(Exit(0))
				Expect(session.Out.Contents()).ToNot(BeNil())
			})

			It("with all flags (JSON output)", func() {
				// Put together mock arguments.
				args := []string{
					"delete",
					`--id=testString`,
					`--output=json`,
				}

				session, err := runCmd(args)

				Expect(err).To(BeNil())
				Expect(session).To(Exit(0))

				out := session.Out.Contents()
				Expect(out).ToNot(BeNil())
				Expect(string(out)).To(Equal("\"\"\n"))
			})
		})

		Context("unsuccessfully", func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request.
					Expect(req.URL.EscapedPath()).To(Equal(operationPath))
					Expect(req.Method).To(Equal("DELETE"))
					// Set mock response
					res.WriteHeader(404)
				}))

				// Create the test environment.
				testEnvironment = append(testEnvironment, "PROJECTS_URL="+testServer.URL)
				testEnvironment = append(testEnvironment, "PROJECTS_AUTH_TYPE=noAuth")
			})

			AfterEach(func() {
				testServer.Close()
				testEnvironment = []string{}
			})

			It("because authentication error", func() {
				// Put together mock arguments.
				args := []string{
					"delete",
					`--id=testString`,
				}

				// Clear the test environment that holds the credentials.
				testEnvironment = []string{}

				session, err := runCmd(args)

				Expect(err).To(BeNil())
				Expect(session).To(Exit(1))
				Expect(session.Err.Contents()).To(ContainSubstring(credentialErrMsg))
			})

			It("because response error", func() {
				args := []string{
					"delete",
					`--id=testString`,
				}

				session, err := runCmd(args)

				Expect(err).To(BeNil())
				Expect(session).To(Exit(1))
				Expect(session.Err.Contents()).To(MatchRegexp("^FAILED\n.+\nNot Found\n\n$"))
			})

			It("because missing required flag", func() {
				args := []string{
					"delete",
				}
				session, err := runCmd(args)

				Expect(err).To(BeNil())
				Expect(session).To(Exit(1))
				Expect(session.Err.Contents()).To(ContainSubstring("required flag(s)"))
				Expect(session.Err.Contents()).To(ContainSubstring("not set"))
			})
		})
	})

	Describe("Run the `configs` command", func() {
		operationPath := "/v1/projects/testString/configs"

		Context("successfully", func() {
			BeforeEach(func() {
				// Create the mock server.
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request.
					Expect(req.URL.EscapedPath()).To(Equal(operationPath))
					Expect(req.Method).To(Equal("GET"))
					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"configs": [{"id": "testString", "name": "testString", "labels": ["testString"], "description": "testString", "locator_id": "testString", "type": "terraform_template", "input": [{"name": "testString", "type": "array", "required": true}], "output": [{"name": "testString", "description": "testString", "value": ["testString"]}], "setting": [{"name": "testString", "value": "testString"}]}]}`)
				}))

				// Create the test environment.
				testEnvironment = append(testEnvironment, "PROJECTS_URL="+testServer.URL)
				testEnvironment = append(testEnvironment, "PROJECTS_AUTH_TYPE=noAuth")
			})

			AfterEach(func() {
				testServer.Close()
				testEnvironment = []string{}
			})

			It("with all flags", func() {
				// Put together mock arguments.
				args := []string{
					"configs",
					`--id=testString`,
					`--version=active`,
					`--complete=false`,
				}

				session, err := runCmd(args)

				Expect(err).To(BeNil())
				Expect(session).To(Exit(0))
				Expect(session.Out.Contents()).ToNot(BeNil())
			})

			It("with all flags (JSON output)", func() {
				// Put together mock arguments.
				args := []string{
					"configs",
					`--id=testString`,
					`--version=active`,
					`--complete=false`,
					`--output=json`,
				}

				session, err := runCmd(args)

				Expect(err).To(BeNil())
				Expect(session).To(Exit(0))

				out := session.Out.Contents()
				mockResponse := `{"configs": [{"id": "testString", "name": "testString", "labels": ["testString"], "description": "testString", "locator_id": "testString", "type": "terraform_template", "input": [{"name": "testString", "type": "array", "required": true}], "output": [{"name": "testString", "description": "testString", "value": ["testString"]}], "setting": [{"name": "testString", "value": "testString"}]}]}`
				Expect(string(out)).To(MatchJSON(mockResponse))
			})
		})

		Context("unsuccessfully", func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request.
					Expect(req.URL.EscapedPath()).To(Equal(operationPath))
					Expect(req.Method).To(Equal("GET"))
					// Set mock response
					res.WriteHeader(404)
				}))

				// Create the test environment.
				testEnvironment = append(testEnvironment, "PROJECTS_URL="+testServer.URL)
				testEnvironment = append(testEnvironment, "PROJECTS_AUTH_TYPE=noAuth")
			})

			AfterEach(func() {
				testServer.Close()
				testEnvironment = []string{}
			})

			It("because authentication error", func() {
				// Put together mock arguments.
				args := []string{
					"configs",
					`--id=testString`,
					`--version=active`,
					`--complete=false`,
				}

				// Clear the test environment that holds the credentials.
				testEnvironment = []string{}

				session, err := runCmd(args)

				Expect(err).To(BeNil())
				Expect(session).To(Exit(1))
				Expect(session.Err.Contents()).To(ContainSubstring(credentialErrMsg))
			})

			It("because response error", func() {
				args := []string{
					"configs",
					`--id=testString`,
					`--version=active`,
					`--complete=false`,
				}

				session, err := runCmd(args)

				Expect(err).To(BeNil())
				Expect(session).To(Exit(1))
				Expect(session.Err.Contents()).To(MatchRegexp("^FAILED\n.+\nNot Found\n\n$"))
			})

			It("because missing required flag", func() {
				args := []string{
					"configs",
				}
				session, err := runCmd(args)

				Expect(err).To(BeNil())
				Expect(session).To(Exit(1))
				Expect(session.Err.Contents()).To(ContainSubstring("required flag(s)"))
				Expect(session.Err.Contents()).To(ContainSubstring("not set"))
			})
		})
	})

	Describe("Run the `config-operation` command", func() {
		operationPath := "/v1/projects/testString/configs/testString"

		Context("successfully", func() {
			BeforeEach(func() {
				// Create the mock server.
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request.
					Expect(req.URL.EscapedPath()).To(Equal(operationPath))
					Expect(req.Method).To(Equal("GET"))
					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"id": "testString", "name": "testString", "labels": ["testString"], "description": "testString", "locator_id": "testString", "type": "terraform_template", "input": [{"name": "testString", "type": "array", "required": true}], "output": [{"name": "testString", "description": "testString", "value": ["testString"]}], "setting": [{"name": "testString", "value": "testString"}]}`)
				}))

				// Create the test environment.
				testEnvironment = append(testEnvironment, "PROJECTS_URL="+testServer.URL)
				testEnvironment = append(testEnvironment, "PROJECTS_AUTH_TYPE=noAuth")
			})

			AfterEach(func() {
				testServer.Close()
				testEnvironment = []string{}
			})

			It("with all flags", func() {
				// Put together mock arguments.
				args := []string{
					"config-operation",
					`--id=testString`,
					`--config-id=testString`,
					`--version=active`,
					`--complete=false`,
				}

				session, err := runCmd(args)

				Expect(err).To(BeNil())
				Expect(session).To(Exit(0))
				Expect(session.Out.Contents()).ToNot(BeNil())
			})

			It("with all flags (JSON output)", func() {
				// Put together mock arguments.
				args := []string{
					"config-operation",
					`--id=testString`,
					`--config-id=testString`,
					`--version=active`,
					`--complete=false`,
					`--output=json`,
				}

				session, err := runCmd(args)

				Expect(err).To(BeNil())
				Expect(session).To(Exit(0))

				out := session.Out.Contents()
				mockResponse := `{"id": "testString", "name": "testString", "labels": ["testString"], "description": "testString", "locator_id": "testString", "type": "terraform_template", "input": [{"name": "testString", "type": "array", "required": true}], "output": [{"name": "testString", "description": "testString", "value": ["testString"]}], "setting": [{"name": "testString", "value": "testString"}]}`
				Expect(string(out)).To(MatchJSON(mockResponse))
			})
		})

		Context("unsuccessfully", func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request.
					Expect(req.URL.EscapedPath()).To(Equal(operationPath))
					Expect(req.Method).To(Equal("GET"))
					// Set mock response
					res.WriteHeader(404)
				}))

				// Create the test environment.
				testEnvironment = append(testEnvironment, "PROJECTS_URL="+testServer.URL)
				testEnvironment = append(testEnvironment, "PROJECTS_AUTH_TYPE=noAuth")
			})

			AfterEach(func() {
				testServer.Close()
				testEnvironment = []string{}
			})

			It("because authentication error", func() {
				// Put together mock arguments.
				args := []string{
					"config-operation",
					`--id=testString`,
					`--config-id=testString`,
					`--version=active`,
					`--complete=false`,
				}

				// Clear the test environment that holds the credentials.
				testEnvironment = []string{}

				session, err := runCmd(args)

				Expect(err).To(BeNil())
				Expect(session).To(Exit(1))
				Expect(session.Err.Contents()).To(ContainSubstring(credentialErrMsg))
			})

			It("because response error", func() {
				args := []string{
					"config-operation",
					`--id=testString`,
					`--config-id=testString`,
					`--version=active`,
					`--complete=false`,
				}

				session, err := runCmd(args)

				Expect(err).To(BeNil())
				Expect(session).To(Exit(1))
				Expect(session.Err.Contents()).To(MatchRegexp("^FAILED\n.+\nNot Found\n\n$"))
			})

			It("because missing required flag", func() {
				args := []string{
					"config-operation",
				}
				session, err := runCmd(args)

				Expect(err).To(BeNil())
				Expect(session).To(Exit(1))
				Expect(session.Err.Contents()).To(ContainSubstring("required flag(s)"))
				Expect(session.Err.Contents()).To(ContainSubstring("not set"))
			})
		})
	})

	Describe("Run the `delete_config` command", func() {
		operationPath := "/v1/projects/testString/configs/testString"

		Context("successfully", func() {
			BeforeEach(func() {
				// Create the mock server.
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request.
					Expect(req.URL.EscapedPath()).To(Equal(operationPath))
					Expect(req.Method).To(Equal("DELETE"))
					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"id": "testString", "name": "testString"}`)
				}))

				// Create the test environment.
				testEnvironment = append(testEnvironment, "PROJECTS_URL="+testServer.URL)
				testEnvironment = append(testEnvironment, "PROJECTS_AUTH_TYPE=noAuth")
			})

			AfterEach(func() {
				testServer.Close()
				testEnvironment = []string{}
			})

			It("with all flags", func() {
				// Put together mock arguments.
				args := []string{
					"delete_config",
					`--id=testString`,
					`--config-id=testString`,
				}

				session, err := runCmd(args)

				Expect(err).To(BeNil())
				Expect(session).To(Exit(0))
				Expect(session.Out.Contents()).ToNot(BeNil())
			})

			It("with all flags (JSON output)", func() {
				// Put together mock arguments.
				args := []string{
					"delete_config",
					`--id=testString`,
					`--config-id=testString`,
					`--output=json`,
				}

				session, err := runCmd(args)

				Expect(err).To(BeNil())
				Expect(session).To(Exit(0))

				out := session.Out.Contents()
				mockResponse := `{"id": "testString", "name": "testString"}`
				Expect(string(out)).To(MatchJSON(mockResponse))
			})
		})

		Context("unsuccessfully", func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request.
					Expect(req.URL.EscapedPath()).To(Equal(operationPath))
					Expect(req.Method).To(Equal("DELETE"))
					// Set mock response
					res.WriteHeader(404)
				}))

				// Create the test environment.
				testEnvironment = append(testEnvironment, "PROJECTS_URL="+testServer.URL)
				testEnvironment = append(testEnvironment, "PROJECTS_AUTH_TYPE=noAuth")
			})

			AfterEach(func() {
				testServer.Close()
				testEnvironment = []string{}
			})

			It("because authentication error", func() {
				// Put together mock arguments.
				args := []string{
					"delete_config",
					`--id=testString`,
					`--config-id=testString`,
				}

				// Clear the test environment that holds the credentials.
				testEnvironment = []string{}

				session, err := runCmd(args)

				Expect(err).To(BeNil())
				Expect(session).To(Exit(1))
				Expect(session.Err.Contents()).To(ContainSubstring(credentialErrMsg))
			})

			It("because response error", func() {
				args := []string{
					"delete_config",
					`--id=testString`,
					`--config-id=testString`,
				}

				session, err := runCmd(args)

				Expect(err).To(BeNil())
				Expect(session).To(Exit(1))
				Expect(session.Err.Contents()).To(MatchRegexp("^FAILED\n.+\nNot Found\n\n$"))
			})

			It("because missing required flag", func() {
				args := []string{
					"delete_config",
				}
				session, err := runCmd(args)

				Expect(err).To(BeNil())
				Expect(session).To(Exit(1))
				Expect(session.Err.Contents()).To(ContainSubstring("required flag(s)"))
				Expect(session.Err.Contents()).To(ContainSubstring("not set"))
			})
		})
	})

	Describe("Run the `config-diff` command", func() {
		operationPath := "/v1/projects/testString/configs/testString/diff"

		Context("successfully", func() {
			BeforeEach(func() {
				// Create the mock server.
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request.
					Expect(req.URL.EscapedPath()).To(Equal(operationPath))
					Expect(req.Method).To(Equal("GET"))
					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"added": {"input": [{"name": "testString", "type": "array"}]}, "changed": {"input": [{"name": "testString", "type": "array"}]}, "removed": {"input": [{"name": "testString", "type": "array"}]}}`)
				}))

				// Create the test environment.
				testEnvironment = append(testEnvironment, "PROJECTS_URL="+testServer.URL)
				testEnvironment = append(testEnvironment, "PROJECTS_AUTH_TYPE=noAuth")
			})

			AfterEach(func() {
				testServer.Close()
				testEnvironment = []string{}
			})

			It("with all flags", func() {
				// Put together mock arguments.
				args := []string{
					"config-diff",
					`--id=testString`,
					`--config-id=testString`,
				}

				session, err := runCmd(args)

				Expect(err).To(BeNil())
				Expect(session).To(Exit(0))
				Expect(session.Out.Contents()).ToNot(BeNil())
			})

			It("with all flags (JSON output)", func() {
				// Put together mock arguments.
				args := []string{
					"config-diff",
					`--id=testString`,
					`--config-id=testString`,
					`--output=json`,
				}

				session, err := runCmd(args)

				Expect(err).To(BeNil())
				Expect(session).To(Exit(0))

				out := session.Out.Contents()
				mockResponse := `{"added": {"input": [{"name": "testString", "type": "array"}]}, "changed": {"input": [{"name": "testString", "type": "array"}]}, "removed": {"input": [{"name": "testString", "type": "array"}]}}`
				Expect(string(out)).To(MatchJSON(mockResponse))
			})
		})

		Context("unsuccessfully", func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request.
					Expect(req.URL.EscapedPath()).To(Equal(operationPath))
					Expect(req.Method).To(Equal("GET"))
					// Set mock response
					res.WriteHeader(404)
				}))

				// Create the test environment.
				testEnvironment = append(testEnvironment, "PROJECTS_URL="+testServer.URL)
				testEnvironment = append(testEnvironment, "PROJECTS_AUTH_TYPE=noAuth")
			})

			AfterEach(func() {
				testServer.Close()
				testEnvironment = []string{}
			})

			It("because authentication error", func() {
				// Put together mock arguments.
				args := []string{
					"config-diff",
					`--id=testString`,
					`--config-id=testString`,
				}

				// Clear the test environment that holds the credentials.
				testEnvironment = []string{}

				session, err := runCmd(args)

				Expect(err).To(BeNil())
				Expect(session).To(Exit(1))
				Expect(session.Err.Contents()).To(ContainSubstring(credentialErrMsg))
			})

			It("because response error", func() {
				args := []string{
					"config-diff",
					`--id=testString`,
					`--config-id=testString`,
				}

				session, err := runCmd(args)

				Expect(err).To(BeNil())
				Expect(session).To(Exit(1))
				Expect(session.Err.Contents()).To(MatchRegexp("^FAILED\n.+\nNot Found\n\n$"))
			})

			It("because missing required flag", func() {
				args := []string{
					"config-diff",
				}
				session, err := runCmd(args)

				Expect(err).To(BeNil())
				Expect(session).To(Exit(1))
				Expect(session.Err.Contents()).To(ContainSubstring("required flag(s)"))
				Expect(session.Err.Contents()).To(ContainSubstring("not set"))
			})
		})
	})

	Describe("Run the `install` command", func() {
		operationPath := "/v1/projects/testString/configs/testString/install"

		Context("successfully", func() {
			BeforeEach(func() {
				// Create the mock server.
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request.
					Expect(req.URL.EscapedPath()).To(Equal(operationPath))
					Expect(req.Method).To(Equal("POST"))
					res.WriteHeader(204)
				}))

				// Create the test environment.
				testEnvironment = append(testEnvironment, "PROJECTS_URL="+testServer.URL)
				testEnvironment = append(testEnvironment, "PROJECTS_AUTH_TYPE=noAuth")
			})

			AfterEach(func() {
				testServer.Close()
				testEnvironment = []string{}
			})

			It("with all flags", func() {
				// Put together mock arguments.
				args := []string{
					"install",
					`--id=testString`,
					`--config-id=testString`,
				}

				session, err := runCmd(args)

				Expect(err).To(BeNil())
				Expect(session).To(Exit(0))
				Expect(session.Out.Contents()).ToNot(BeNil())
			})

			It("with all flags (JSON output)", func() {
				// Put together mock arguments.
				args := []string{
					"install",
					`--id=testString`,
					`--config-id=testString`,
					`--output=json`,
				}

				session, err := runCmd(args)

				Expect(err).To(BeNil())
				Expect(session).To(Exit(0))

				out := session.Out.Contents()
				Expect(out).ToNot(BeNil())
				Expect(string(out)).To(Equal("\"\"\n"))
			})
		})

		Context("unsuccessfully", func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request.
					Expect(req.URL.EscapedPath()).To(Equal(operationPath))
					Expect(req.Method).To(Equal("POST"))
					// Set mock response
					res.WriteHeader(404)
				}))

				// Create the test environment.
				testEnvironment = append(testEnvironment, "PROJECTS_URL="+testServer.URL)
				testEnvironment = append(testEnvironment, "PROJECTS_AUTH_TYPE=noAuth")
			})

			AfterEach(func() {
				testServer.Close()
				testEnvironment = []string{}
			})

			It("because authentication error", func() {
				// Put together mock arguments.
				args := []string{
					"install",
					`--id=testString`,
					`--config-id=testString`,
				}

				// Clear the test environment that holds the credentials.
				testEnvironment = []string{}

				session, err := runCmd(args)

				Expect(err).To(BeNil())
				Expect(session).To(Exit(1))
				Expect(session.Err.Contents()).To(ContainSubstring(credentialErrMsg))
			})

			It("because response error", func() {
				args := []string{
					"install",
					`--id=testString`,
					`--config-id=testString`,
				}

				session, err := runCmd(args)

				Expect(err).To(BeNil())
				Expect(session).To(Exit(1))
				Expect(session.Err.Contents()).To(MatchRegexp("^FAILED\n.+\nNot Found\n\n$"))
			})

			It("because missing required flag", func() {
				args := []string{
					"install",
				}
				session, err := runCmd(args)

				Expect(err).To(BeNil())
				Expect(session).To(Exit(1))
				Expect(session.Err.Contents()).To(ContainSubstring("required flag(s)"))
				Expect(session.Err.Contents()).To(ContainSubstring("not set"))
			})
		})
	})

	Describe("Run the `uninstall` command", func() {
		operationPath := "/v1/projects/testString/configs/testString/uninstall"

		Context("successfully", func() {
			BeforeEach(func() {
				// Create the mock server.
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request.
					Expect(req.URL.EscapedPath()).To(Equal(operationPath))
					Expect(req.Method).To(Equal("POST"))
					res.WriteHeader(204)
				}))

				// Create the test environment.
				testEnvironment = append(testEnvironment, "PROJECTS_URL="+testServer.URL)
				testEnvironment = append(testEnvironment, "PROJECTS_AUTH_TYPE=noAuth")
			})

			AfterEach(func() {
				testServer.Close()
				testEnvironment = []string{}
			})

			It("with all flags", func() {
				// Put together mock arguments.
				args := []string{
					"uninstall",
					`--id=testString`,
					`--config-id=testString`,
				}

				session, err := runCmd(args)

				Expect(err).To(BeNil())
				Expect(session).To(Exit(0))
				Expect(session.Out.Contents()).ToNot(BeNil())
			})

			It("with all flags (JSON output)", func() {
				// Put together mock arguments.
				args := []string{
					"uninstall",
					`--id=testString`,
					`--config-id=testString`,
					`--output=json`,
				}

				session, err := runCmd(args)

				Expect(err).To(BeNil())
				Expect(session).To(Exit(0))

				out := session.Out.Contents()
				Expect(out).ToNot(BeNil())
				Expect(string(out)).To(Equal("\"\"\n"))
			})
		})

		Context("unsuccessfully", func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request.
					Expect(req.URL.EscapedPath()).To(Equal(operationPath))
					Expect(req.Method).To(Equal("POST"))
					// Set mock response
					res.WriteHeader(404)
				}))

				// Create the test environment.
				testEnvironment = append(testEnvironment, "PROJECTS_URL="+testServer.URL)
				testEnvironment = append(testEnvironment, "PROJECTS_AUTH_TYPE=noAuth")
			})

			AfterEach(func() {
				testServer.Close()
				testEnvironment = []string{}
			})

			It("because authentication error", func() {
				// Put together mock arguments.
				args := []string{
					"uninstall",
					`--id=testString`,
					`--config-id=testString`,
				}

				// Clear the test environment that holds the credentials.
				testEnvironment = []string{}

				session, err := runCmd(args)

				Expect(err).To(BeNil())
				Expect(session).To(Exit(1))
				Expect(session.Err.Contents()).To(ContainSubstring(credentialErrMsg))
			})

			It("because response error", func() {
				args := []string{
					"uninstall",
					`--id=testString`,
					`--config-id=testString`,
				}

				session, err := runCmd(args)

				Expect(err).To(BeNil())
				Expect(session).To(Exit(1))
				Expect(session.Err.Contents()).To(MatchRegexp("^FAILED\n.+\nNot Found\n\n$"))
			})

			It("because missing required flag", func() {
				args := []string{
					"uninstall",
				}
				session, err := runCmd(args)

				Expect(err).To(BeNil())
				Expect(session).To(Exit(1))
				Expect(session.Err.Contents()).To(ContainSubstring("required flag(s)"))
				Expect(session.Err.Contents()).To(ContainSubstring("not set"))
			})
		})
	})

	Describe("Run the `get_schematics_job` command", func() {
		operationPath := "/v1/projects/testString/configs/testString/plan/job"

		Context("successfully", func() {
			BeforeEach(func() {
				// Create the mock server.
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request.
					Expect(req.URL.EscapedPath()).To(Equal(operationPath))
					Expect(req.Method).To(Equal("GET"))
					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"id": "testString"}`)
				}))

				// Create the test environment.
				testEnvironment = append(testEnvironment, "PROJECTS_URL="+testServer.URL)
				testEnvironment = append(testEnvironment, "PROJECTS_AUTH_TYPE=noAuth")
			})

			AfterEach(func() {
				testServer.Close()
				testEnvironment = []string{}
			})

			It("with all flags", func() {
				// Put together mock arguments.
				args := []string{
					"get_schematics_job",
					`--id=testString`,
					`--config-id=testString`,
					`--action=plan`,
					`--since=38`,
				}

				session, err := runCmd(args)

				Expect(err).To(BeNil())
				Expect(session).To(Exit(0))
				Expect(session.Out.Contents()).ToNot(BeNil())
			})

			It("with all flags (JSON output)", func() {
				// Put together mock arguments.
				args := []string{
					"get_schematics_job",
					`--id=testString`,
					`--config-id=testString`,
					`--action=plan`,
					`--since=38`,
					`--output=json`,
				}

				session, err := runCmd(args)

				Expect(err).To(BeNil())
				Expect(session).To(Exit(0))

				out := session.Out.Contents()
				mockResponse := `{"id": "testString"}`
				Expect(string(out)).To(MatchJSON(mockResponse))
			})
		})

		Context("unsuccessfully", func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request.
					Expect(req.URL.EscapedPath()).To(Equal(operationPath))
					Expect(req.Method).To(Equal("GET"))
					// Set mock response
					res.WriteHeader(404)
				}))

				// Create the test environment.
				testEnvironment = append(testEnvironment, "PROJECTS_URL="+testServer.URL)
				testEnvironment = append(testEnvironment, "PROJECTS_AUTH_TYPE=noAuth")
			})

			AfterEach(func() {
				testServer.Close()
				testEnvironment = []string{}
			})

			It("because authentication error", func() {
				// Put together mock arguments.
				args := []string{
					"get_schematics_job",
					`--id=testString`,
					`--config-id=testString`,
					`--action=plan`,
					`--since=38`,
				}

				// Clear the test environment that holds the credentials.
				testEnvironment = []string{}

				session, err := runCmd(args)

				Expect(err).To(BeNil())
				Expect(session).To(Exit(1))
				Expect(session.Err.Contents()).To(ContainSubstring(credentialErrMsg))
			})

			It("because response error", func() {
				args := []string{
					"get_schematics_job",
					`--id=testString`,
					`--config-id=testString`,
					`--action=plan`,
					`--since=38`,
				}

				session, err := runCmd(args)

				Expect(err).To(BeNil())
				Expect(session).To(Exit(1))
				Expect(session.Err.Contents()).To(MatchRegexp("^FAILED\n.+\nNot Found\n\n$"))
			})

			It("because missing required flag", func() {
				args := []string{
					"get_schematics_job",
				}
				session, err := runCmd(args)

				Expect(err).To(BeNil())
				Expect(session).To(Exit(1))
				Expect(session.Err.Contents()).To(ContainSubstring("required flag(s)"))
				Expect(session.Err.Contents()).To(ContainSubstring("not set"))
			})
		})
	})

	Describe("Run the `get_cost_estimate` command", func() {
		operationPath := "/v1/projects/testString/configs/testString/cost_estimate"

		Context("successfully", func() {
			BeforeEach(func() {
				// Create the mock server.
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request.
					Expect(req.URL.EscapedPath()).To(Equal(operationPath))
					Expect(req.Method).To(Equal("GET"))
					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"anyKey": "anyValue"}`)
				}))

				// Create the test environment.
				testEnvironment = append(testEnvironment, "PROJECTS_URL="+testServer.URL)
				testEnvironment = append(testEnvironment, "PROJECTS_AUTH_TYPE=noAuth")
			})

			AfterEach(func() {
				testServer.Close()
				testEnvironment = []string{}
			})

			It("with all flags", func() {
				// Put together mock arguments.
				args := []string{
					"get_cost_estimate",
					`--id=testString`,
					`--config-id=testString`,
					`--version=active`,
				}

				session, err := runCmd(args)

				Expect(err).To(BeNil())
				Expect(session).To(Exit(0))
				Expect(session.Out.Contents()).ToNot(BeNil())
			})

			It("with all flags (JSON output)", func() {
				// Put together mock arguments.
				args := []string{
					"get_cost_estimate",
					`--id=testString`,
					`--config-id=testString`,
					`--version=active`,
					`--output=json`,
				}

				session, err := runCmd(args)

				Expect(err).To(BeNil())
				Expect(session).To(Exit(0))

				out := session.Out.Contents()
				mockResponse := `{"anyKey": "anyValue"}`
				Expect(string(out)).To(MatchJSON(mockResponse))
			})
		})

		Context("unsuccessfully", func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request.
					Expect(req.URL.EscapedPath()).To(Equal(operationPath))
					Expect(req.Method).To(Equal("GET"))
					// Set mock response
					res.WriteHeader(404)
				}))

				// Create the test environment.
				testEnvironment = append(testEnvironment, "PROJECTS_URL="+testServer.URL)
				testEnvironment = append(testEnvironment, "PROJECTS_AUTH_TYPE=noAuth")
			})

			AfterEach(func() {
				testServer.Close()
				testEnvironment = []string{}
			})

			It("because authentication error", func() {
				// Put together mock arguments.
				args := []string{
					"get_cost_estimate",
					`--id=testString`,
					`--config-id=testString`,
					`--version=active`,
				}

				// Clear the test environment that holds the credentials.
				testEnvironment = []string{}

				session, err := runCmd(args)

				Expect(err).To(BeNil())
				Expect(session).To(Exit(1))
				Expect(session.Err.Contents()).To(ContainSubstring(credentialErrMsg))
			})

			It("because response error", func() {
				args := []string{
					"get_cost_estimate",
					`--id=testString`,
					`--config-id=testString`,
					`--version=active`,
				}

				session, err := runCmd(args)

				Expect(err).To(BeNil())
				Expect(session).To(Exit(1))
				Expect(session.Err.Contents()).To(MatchRegexp("^FAILED\n.+\nNot Found\n\n$"))
			})

			It("because missing required flag", func() {
				args := []string{
					"get_cost_estimate",
				}
				session, err := runCmd(args)

				Expect(err).To(BeNil())
				Expect(session).To(Exit(1))
				Expect(session.Err.Contents()).To(ContainSubstring("required flag(s)"))
				Expect(session.Err.Contents()).To(ContainSubstring("not set"))
			})
		})
	})

	Describe("Run the `post-notification` command", func() {
		operationPath := "/v1/projects/testString/event"

		Context("successfully", func() {
			BeforeEach(func() {
				// Create the mock server.
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request.
					Expect(req.URL.EscapedPath()).To(Equal(operationPath))
					Expect(req.Method).To(Equal("POST"))
					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"notifications": [{"event": "testString", "target": "testString", "source": "testString", "action_url": "testString", "data": {"anyKey": "anyValue"}, "_id": "testString", "status": "testString", "reasons": [{"anyKey": "anyValue"}]}]}`)
				}))

				// Create the test environment.
				testEnvironment = append(testEnvironment, "PROJECTS_URL="+testServer.URL)
				testEnvironment = append(testEnvironment, "PROJECTS_AUTH_TYPE=noAuth")
			})

			AfterEach(func() {
				testServer.Close()
				testEnvironment = []string{}
			})

			It("with all flags", func() {
				// Put together mock arguments.
				args := []string{
					"post-notification",
					`--id=testString`,
					`--notifications=[{"event": "project.create.failed", "target": "234234324-3444-4556-224232432", "source": "id.of.project.service.instance", "action_url": "url.for.project.documentation", "data": {"anyKey": "anyValue"}}]`,
				}

				session, err := runCmd(args)

				Expect(err).To(BeNil())
				Expect(session).To(Exit(0))
				Expect(session.Out.Contents()).ToNot(BeNil())
			})

			It("with all flags (JSON output)", func() {
				// Put together mock arguments.
				args := []string{
					"post-notification",
					`--id=testString`,
					`--notifications=[{"event": "project.create.failed", "target": "234234324-3444-4556-224232432", "source": "id.of.project.service.instance", "action_url": "url.for.project.documentation", "data": {"anyKey": "anyValue"}}]`,
					`--output=json`,
				}

				session, err := runCmd(args)

				Expect(err).To(BeNil())
				Expect(session).To(Exit(0))

				out := session.Out.Contents()
				mockResponse := `{"notifications": [{"event": "testString", "target": "testString", "source": "testString", "action_url": "testString", "data": {"anyKey": "anyValue"}, "_id": "testString", "status": "testString", "reasons": [{"anyKey": "anyValue"}]}]}`
				Expect(string(out)).To(MatchJSON(mockResponse))
			})
		})

		Context("unsuccessfully", func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request.
					Expect(req.URL.EscapedPath()).To(Equal(operationPath))
					Expect(req.Method).To(Equal("POST"))
					// Set mock response
					res.WriteHeader(404)
				}))

				// Create the test environment.
				testEnvironment = append(testEnvironment, "PROJECTS_URL="+testServer.URL)
				testEnvironment = append(testEnvironment, "PROJECTS_AUTH_TYPE=noAuth")
			})

			AfterEach(func() {
				testServer.Close()
				testEnvironment = []string{}
			})

			It("because authentication error", func() {
				// Put together mock arguments.
				args := []string{
					"post-notification",
					`--id=testString`,
					`--notifications=[{"event": "project.create.failed", "target": "234234324-3444-4556-224232432", "source": "id.of.project.service.instance", "action_url": "url.for.project.documentation", "data": {"anyKey": "anyValue"}}]`,
				}

				// Clear the test environment that holds the credentials.
				testEnvironment = []string{}

				session, err := runCmd(args)

				Expect(err).To(BeNil())
				Expect(session).To(Exit(1))
				Expect(session.Err.Contents()).To(ContainSubstring(credentialErrMsg))
			})

			It("because response error", func() {
				args := []string{
					"post-notification",
					`--id=testString`,
					`--notifications=[{"event": "project.create.failed", "target": "234234324-3444-4556-224232432", "source": "id.of.project.service.instance", "action_url": "url.for.project.documentation", "data": {"anyKey": "anyValue"}}]`,
				}

				session, err := runCmd(args)

				Expect(err).To(BeNil())
				Expect(session).To(Exit(1))
				Expect(session.Err.Contents()).To(MatchRegexp("^FAILED\n.+\nNot Found\n\n$"))
			})

			It("because missing required flag", func() {
				args := []string{
					"post-notification",
				}
				session, err := runCmd(args)

				Expect(err).To(BeNil())
				Expect(session).To(Exit(1))
				Expect(session.Err.Contents()).To(ContainSubstring("required flag(s)"))
				Expect(session.Err.Contents()).To(ContainSubstring("not set"))
			})
		})
	})

	Describe("Run the `health` command", func() {
		operationPath := "/v1/health"

		Context("successfully", func() {
			BeforeEach(func() {
				// Create the mock server.
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request.
					Expect(req.URL.EscapedPath()).To(Equal(operationPath))
					Expect(req.Method).To(Equal("GET"))
					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"name": "testString", "version": "testString", "dependencies": {"anyKey": "anyValue"}}`)
				}))

				// Create the test environment.
				testEnvironment = append(testEnvironment, "PROJECTS_URL="+testServer.URL)
				testEnvironment = append(testEnvironment, "PROJECTS_AUTH_TYPE=noAuth")
			})

			AfterEach(func() {
				testServer.Close()
				testEnvironment = []string{}
			})

			It("with all flags", func() {
				// Put together mock arguments.
				args := []string{
					"health",
					`--info=false`,
				}

				session, err := runCmd(args)

				Expect(err).To(BeNil())
				Expect(session).To(Exit(0))
				Expect(session.Out.Contents()).ToNot(BeNil())
			})

			It("with all flags (JSON output)", func() {
				// Put together mock arguments.
				args := []string{
					"health",
					`--info=false`,
					`--output=json`,
				}

				session, err := runCmd(args)

				Expect(err).To(BeNil())
				Expect(session).To(Exit(0))

				out := session.Out.Contents()
				mockResponse := `{"name": "testString", "version": "testString", "dependencies": {"anyKey": "anyValue"}}`
				Expect(string(out)).To(MatchJSON(mockResponse))
			})
		})

		Context("unsuccessfully", func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request.
					Expect(req.URL.EscapedPath()).To(Equal(operationPath))
					Expect(req.Method).To(Equal("GET"))
					// Set mock response
					res.WriteHeader(404)
				}))

				// Create the test environment.
				testEnvironment = append(testEnvironment, "PROJECTS_URL="+testServer.URL)
				testEnvironment = append(testEnvironment, "PROJECTS_AUTH_TYPE=noAuth")
			})

			AfterEach(func() {
				testServer.Close()
				testEnvironment = []string{}
			})

			It("because authentication error", func() {
				// Put together mock arguments.
				args := []string{
					"health",
					`--info=false`,
				}

				// Clear the test environment that holds the credentials.
				testEnvironment = []string{}

				session, err := runCmd(args)

				Expect(err).To(BeNil())
				Expect(session).To(Exit(1))
				Expect(session.Err.Contents()).To(ContainSubstring(credentialErrMsg))
			})

			It("because response error", func() {
				args := []string{
					"health",
					`--info=false`,
				}

				session, err := runCmd(args)

				Expect(err).To(BeNil())
				Expect(session).To(Exit(1))
				Expect(session.Err.Contents()).To(MatchRegexp("^FAILED\n.+\nNot Found\n\n$"))
			})
		})
	})

	Describe("Run the `post-event-notifications-integration` command", func() {
		operationPath := "/v1/projects/testString/integrations/event_notifications"

		Context("successfully", func() {
			BeforeEach(func() {
				// Create the mock server.
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request.
					Expect(req.URL.EscapedPath()).To(Equal(operationPath))
					Expect(req.Method).To(Equal("POST"))
					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"description": "testString", "name": "testString", "enabled": true, "id": "testString", "type": "testString", "created_at": "2019-01-01T12:00:00.000Z"}`)
				}))

				// Create the test environment.
				testEnvironment = append(testEnvironment, "PROJECTS_URL="+testServer.URL)
				testEnvironment = append(testEnvironment, "PROJECTS_AUTH_TYPE=noAuth")
			})

			AfterEach(func() {
				testServer.Close()
				testEnvironment = []string{}
			})

			It("with all flags", func() {
				// Put together mock arguments.
				args := []string{
					"post-event-notifications-integration",
					`--id=testString`,
					`--instance-crn=crn of event notifications instance`,
					`--description=A sample project source`,
					`--name=Project name`,
					`--enabled=true`,
					`--source=CRN of the project instance`,
				}

				session, err := runCmd(args)

				Expect(err).To(BeNil())
				Expect(session).To(Exit(0))
				Expect(session.Out.Contents()).ToNot(BeNil())
			})

			It("with all flags (JSON output)", func() {
				// Put together mock arguments.
				args := []string{
					"post-event-notifications-integration",
					`--id=testString`,
					`--instance-crn=crn of event notifications instance`,
					`--description=A sample project source`,
					`--name=Project name`,
					`--enabled=true`,
					`--source=CRN of the project instance`,
					`--output=json`,
				}

				session, err := runCmd(args)

				Expect(err).To(BeNil())
				Expect(session).To(Exit(0))

				out := session.Out.Contents()
				mockResponse := `{"description": "testString", "name": "testString", "enabled": true, "id": "testString", "type": "testString", "created_at": "2019-01-01T12:00:00.000Z"}`
				Expect(string(out)).To(MatchJSON(mockResponse))
			})
		})

		Context("unsuccessfully", func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request.
					Expect(req.URL.EscapedPath()).To(Equal(operationPath))
					Expect(req.Method).To(Equal("POST"))
					// Set mock response
					res.WriteHeader(404)
				}))

				// Create the test environment.
				testEnvironment = append(testEnvironment, "PROJECTS_URL="+testServer.URL)
				testEnvironment = append(testEnvironment, "PROJECTS_AUTH_TYPE=noAuth")
			})

			AfterEach(func() {
				testServer.Close()
				testEnvironment = []string{}
			})

			It("because authentication error", func() {
				// Put together mock arguments.
				args := []string{
					"post-event-notifications-integration",
					`--id=testString`,
					`--instance-crn=crn of event notifications instance`,
					`--description=A sample project source`,
					`--name=Project name`,
					`--enabled=true`,
					`--source=CRN of the project instance`,
				}

				// Clear the test environment that holds the credentials.
				testEnvironment = []string{}

				session, err := runCmd(args)

				Expect(err).To(BeNil())
				Expect(session).To(Exit(1))
				Expect(session.Err.Contents()).To(ContainSubstring(credentialErrMsg))
			})

			It("because response error", func() {
				args := []string{
					"post-event-notifications-integration",
					`--id=testString`,
					`--instance-crn=crn of event notifications instance`,
					`--description=A sample project source`,
					`--name=Project name`,
					`--enabled=true`,
					`--source=CRN of the project instance`,
				}

				session, err := runCmd(args)

				Expect(err).To(BeNil())
				Expect(session).To(Exit(1))
				Expect(session.Err.Contents()).To(MatchRegexp("^FAILED\n.+\nNot Found\n\n$"))
			})

			It("because missing required flag", func() {
				args := []string{
					"post-event-notifications-integration",
				}
				session, err := runCmd(args)

				Expect(err).To(BeNil())
				Expect(session).To(Exit(1))
				Expect(session.Err.Contents()).To(ContainSubstring("required flag(s)"))
				Expect(session.Err.Contents()).To(ContainSubstring("not set"))
			})
		})
	})

	Describe("Run the `post-event-notification` command", func() {
		operationPath := "/v1/projects/testString/integrations/event_notifications/notifications"

		Context("successfully", func() {
			BeforeEach(func() {
				// Create the mock server.
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request.
					Expect(req.URL.EscapedPath()).To(Equal(operationPath))
					Expect(req.Method).To(Equal("POST"))
					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"datacontenttype": "testString", "ibmendefaultlong": "testString", "ibmendefaultshort": "testString", "ibmensourceid": "testString", "id": "testString", "source": "testString", "specversion": "testString", "type": "testString"}`)
				}))

				// Create the test environment.
				testEnvironment = append(testEnvironment, "PROJECTS_URL="+testServer.URL)
				testEnvironment = append(testEnvironment, "PROJECTS_AUTH_TYPE=noAuth")
			})

			AfterEach(func() {
				testServer.Close()
				testEnvironment = []string{}
			})

			It("with all flags", func() {
				// Put together mock arguments.
				args := []string{
					"post-event-notification",
					`--id=testString`,
					`--new-id=5f208fef-6b64-413c-aa07-dfed0b46abc1236`,
					`--new-source=crn of project`,
					`--new-datacontenttype=application/json`,
					`--new-ibmendefaultlong=long test notification message`,
					`--new-ibmendefaultshort=Test notification`,
					`--new-ibmensourceid=crn of project`,
					`--new-specversion=1.0`,
					`--new-type=com.ibm.cloud.project.project.test_notification`,
				}

				session, err := runCmd(args)

				Expect(err).To(BeNil())
				Expect(session).To(Exit(0))
				Expect(session.Out.Contents()).ToNot(BeNil())
			})

			It("with all flags (JSON output)", func() {
				// Put together mock arguments.
				args := []string{
					"post-event-notification",
					`--id=testString`,
					`--new-id=5f208fef-6b64-413c-aa07-dfed0b46abc1236`,
					`--new-source=crn of project`,
					`--new-datacontenttype=application/json`,
					`--new-ibmendefaultlong=long test notification message`,
					`--new-ibmendefaultshort=Test notification`,
					`--new-ibmensourceid=crn of project`,
					`--new-specversion=1.0`,
					`--new-type=com.ibm.cloud.project.project.test_notification`,
					`--output=json`,
				}

				session, err := runCmd(args)

				Expect(err).To(BeNil())
				Expect(session).To(Exit(0))

				out := session.Out.Contents()
				mockResponse := `{"datacontenttype": "testString", "ibmendefaultlong": "testString", "ibmendefaultshort": "testString", "ibmensourceid": "testString", "id": "testString", "source": "testString", "specversion": "testString", "type": "testString"}`
				Expect(string(out)).To(MatchJSON(mockResponse))
			})
		})

		Context("unsuccessfully", func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request.
					Expect(req.URL.EscapedPath()).To(Equal(operationPath))
					Expect(req.Method).To(Equal("POST"))
					// Set mock response
					res.WriteHeader(404)
				}))

				// Create the test environment.
				testEnvironment = append(testEnvironment, "PROJECTS_URL="+testServer.URL)
				testEnvironment = append(testEnvironment, "PROJECTS_AUTH_TYPE=noAuth")
			})

			AfterEach(func() {
				testServer.Close()
				testEnvironment = []string{}
			})

			It("because authentication error", func() {
				// Put together mock arguments.
				args := []string{
					"post-event-notification",
					`--id=testString`,
					`--new-id=5f208fef-6b64-413c-aa07-dfed0b46abc1236`,
					`--new-source=crn of project`,
					`--new-datacontenttype=application/json`,
					`--new-ibmendefaultlong=long test notification message`,
					`--new-ibmendefaultshort=Test notification`,
					`--new-ibmensourceid=crn of project`,
					`--new-specversion=1.0`,
					`--new-type=com.ibm.cloud.project.project.test_notification`,
				}

				// Clear the test environment that holds the credentials.
				testEnvironment = []string{}

				session, err := runCmd(args)

				Expect(err).To(BeNil())
				Expect(session).To(Exit(1))
				Expect(session.Err.Contents()).To(ContainSubstring(credentialErrMsg))
			})

			It("because response error", func() {
				args := []string{
					"post-event-notification",
					`--id=testString`,
					`--new-id=5f208fef-6b64-413c-aa07-dfed0b46abc1236`,
					`--new-source=crn of project`,
					`--new-datacontenttype=application/json`,
					`--new-ibmendefaultlong=long test notification message`,
					`--new-ibmendefaultshort=Test notification`,
					`--new-ibmensourceid=crn of project`,
					`--new-specversion=1.0`,
					`--new-type=com.ibm.cloud.project.project.test_notification`,
				}

				session, err := runCmd(args)

				Expect(err).To(BeNil())
				Expect(session).To(Exit(1))
				Expect(session.Err.Contents()).To(MatchRegexp("^FAILED\n.+\nNot Found\n\n$"))
			})

			It("because missing required flag", func() {
				args := []string{
					"post-event-notification",
				}
				session, err := runCmd(args)

				Expect(err).To(BeNil())
				Expect(session).To(Exit(1))
				Expect(session.Err.Contents()).To(ContainSubstring("required flag(s)"))
				Expect(session.Err.Contents()).To(ContainSubstring("not set"))
			})
		})
	})
})
