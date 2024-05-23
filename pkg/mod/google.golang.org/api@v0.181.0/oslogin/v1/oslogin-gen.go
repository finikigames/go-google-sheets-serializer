// Copyright 2024 Google LLC.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Code generated file. DO NOT EDIT.

// Package oslogin provides access to the Cloud OS Login API.
//
// This package is DEPRECATED. Use package cloud.google.com/go/oslogin/apiv1 instead.
//
// For product documentation, see: https://cloud.google.com/compute/docs/oslogin/
//
// # Library status
//
// These client libraries are officially supported by Google. However, this
// library is considered complete and is in maintenance mode. This means
// that we will address critical bugs and security issues but will not add
// any new features.
//
// When possible, we recommend using our newer
// [Cloud Client Libraries for Go](https://pkg.go.dev/cloud.google.com/go)
// that are still actively being worked and iterated on.
//
// # Creating a client
//
// Usage example:
//
//	import "google.golang.org/api/oslogin/v1"
//	...
//	ctx := context.Background()
//	osloginService, err := oslogin.NewService(ctx)
//
// In this example, Google Application Default Credentials are used for
// authentication. For information on how to create and obtain Application
// Default Credentials, see https://developers.google.com/identity/protocols/application-default-credentials.
//
// # Other authentication options
//
// By default, all available scopes (see "Constants") are used to authenticate.
// To restrict scopes, use [google.golang.org/api/option.WithScopes]:
//
//	osloginService, err := oslogin.NewService(ctx, option.WithScopes(oslogin.ComputeReadonlyScope))
//
// To use an API key for authentication (note: some APIs do not support API
// keys), use [google.golang.org/api/option.WithAPIKey]:
//
//	osloginService, err := oslogin.NewService(ctx, option.WithAPIKey("AIza..."))
//
// To use an OAuth token (e.g., a user token obtained via a three-legged OAuth
// flow, use [google.golang.org/api/option.WithTokenSource]:
//
//	config := &oauth2.Config{...}
//	// ...
//	token, err := config.Exchange(ctx, ...)
//	osloginService, err := oslogin.NewService(ctx, option.WithTokenSource(config.TokenSource(ctx, token)))
//
// See [google.golang.org/api/option.ClientOption] for details on options.
package oslogin // import "google.golang.org/api/oslogin/v1"

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strconv"
	"strings"

	googleapi "google.golang.org/api/googleapi"
	internal "google.golang.org/api/internal"
	gensupport "google.golang.org/api/internal/gensupport"
	option "google.golang.org/api/option"
	internaloption "google.golang.org/api/option/internaloption"
	htransport "google.golang.org/api/transport/http"
)

// Always reference these packages, just in case the auto-generated code
// below doesn't.
var _ = bytes.NewBuffer
var _ = strconv.Itoa
var _ = fmt.Sprintf
var _ = json.NewDecoder
var _ = io.Copy
var _ = url.Parse
var _ = gensupport.MarshalJSON
var _ = googleapi.Version
var _ = errors.New
var _ = strings.Replace
var _ = context.Canceled
var _ = internaloption.WithDefaultEndpoint
var _ = internal.Version

const apiId = "oslogin:v1"
const apiName = "oslogin"
const apiVersion = "v1"
const basePath = "https://oslogin.googleapis.com/"
const basePathTemplate = "https://oslogin.UNIVERSE_DOMAIN/"
const mtlsBasePath = "https://oslogin.mtls.googleapis.com/"

// OAuth2 scopes used by this API.
const (
	// See, edit, configure, and delete your Google Cloud data and see the email
	// address for your Google Account.
	CloudPlatformScope = "https://www.googleapis.com/auth/cloud-platform"

	// View your data across Google Cloud services and see the email address of
	// your Google Account
	CloudPlatformReadOnlyScope = "https://www.googleapis.com/auth/cloud-platform.read-only"

	// View and manage your Google Compute Engine resources
	ComputeScope = "https://www.googleapis.com/auth/compute"

	// View your Google Compute Engine resources
	ComputeReadonlyScope = "https://www.googleapis.com/auth/compute.readonly"
)

// NewService creates a new Service.
func NewService(ctx context.Context, opts ...option.ClientOption) (*Service, error) {
	scopesOption := internaloption.WithDefaultScopes(
		"https://www.googleapis.com/auth/cloud-platform",
		"https://www.googleapis.com/auth/cloud-platform.read-only",
		"https://www.googleapis.com/auth/compute",
		"https://www.googleapis.com/auth/compute.readonly",
	)
	// NOTE: prepend, so we don't override user-specified scopes.
	opts = append([]option.ClientOption{scopesOption}, opts...)
	opts = append(opts, internaloption.WithDefaultEndpoint(basePath))
	opts = append(opts, internaloption.WithDefaultEndpointTemplate(basePathTemplate))
	opts = append(opts, internaloption.WithDefaultMTLSEndpoint(mtlsBasePath))
	opts = append(opts, internaloption.EnableNewAuthLibrary())
	client, endpoint, err := htransport.NewClient(ctx, opts...)
	if err != nil {
		return nil, err
	}
	s, err := New(client)
	if err != nil {
		return nil, err
	}
	if endpoint != "" {
		s.BasePath = endpoint
	}
	return s, nil
}

// New creates a new Service. It uses the provided http.Client for requests.
//
// Deprecated: please use NewService instead.
// To provide a custom HTTP client, use option.WithHTTPClient.
// If you are using google.golang.org/api/googleapis/transport.APIKey, use option.WithAPIKey with NewService instead.
func New(client *http.Client) (*Service, error) {
	if client == nil {
		return nil, errors.New("client is nil")
	}
	s := &Service{client: client, BasePath: basePath}
	s.Users = NewUsersService(s)
	return s, nil
}

type Service struct {
	client    *http.Client
	BasePath  string // API endpoint base URL
	UserAgent string // optional additional User-Agent fragment

	Users *UsersService
}

func (s *Service) userAgent() string {
	if s.UserAgent == "" {
		return googleapi.UserAgent
	}
	return googleapi.UserAgent + " " + s.UserAgent
}

func NewUsersService(s *Service) *UsersService {
	rs := &UsersService{s: s}
	rs.Projects = NewUsersProjectsService(s)
	rs.SshPublicKeys = NewUsersSshPublicKeysService(s)
	return rs
}

type UsersService struct {
	s *Service

	Projects *UsersProjectsService

	SshPublicKeys *UsersSshPublicKeysService
}

func NewUsersProjectsService(s *Service) *UsersProjectsService {
	rs := &UsersProjectsService{s: s}
	return rs
}

type UsersProjectsService struct {
	s *Service
}

func NewUsersSshPublicKeysService(s *Service) *UsersSshPublicKeysService {
	rs := &UsersSshPublicKeysService{s: s}
	return rs
}

type UsersSshPublicKeysService struct {
	s *Service
}

// Empty: A generic empty message that you can re-use to avoid defining
// duplicated empty messages in your APIs. A typical example is to use it as
// the request or the response type of an API method. For instance: service Foo
// { rpc Bar(google.protobuf.Empty) returns (google.protobuf.Empty); }
type Empty struct {
	// ServerResponse contains the HTTP response code and headers from the server.
	googleapi.ServerResponse `json:"-"`
}

// ImportSshPublicKeyResponse: A response message for importing an SSH public
// key.
type ImportSshPublicKeyResponse struct {
	// Details: Detailed information about import results.
	Details string `json:"details,omitempty"`
	// LoginProfile: The login profile information for the user.
	LoginProfile *LoginProfile `json:"loginProfile,omitempty"`

	// ServerResponse contains the HTTP response code and headers from the server.
	googleapi.ServerResponse `json:"-"`
	// ForceSendFields is a list of field names (e.g. "Details") to unconditionally
	// include in API requests. By default, fields with empty or default values are
	// omitted from API requests. See
	// https://pkg.go.dev/google.golang.org/api#hdr-ForceSendFields for more
	// details.
	ForceSendFields []string `json:"-"`
	// NullFields is a list of field names (e.g. "Details") to include in API
	// requests with the JSON null value. By default, fields with empty values are
	// omitted from API requests. See
	// https://pkg.go.dev/google.golang.org/api#hdr-NullFields for more details.
	NullFields []string `json:"-"`
}

func (s *ImportSshPublicKeyResponse) MarshalJSON() ([]byte, error) {
	type NoMethod ImportSshPublicKeyResponse
	return gensupport.MarshalJSON(NoMethod(*s), s.ForceSendFields, s.NullFields)
}

// LoginProfile: The user profile information used for logging in to a virtual
// machine on Google Compute Engine.
type LoginProfile struct {
	// Name: Required. A unique user ID.
	Name string `json:"name,omitempty"`
	// PosixAccounts: The list of POSIX accounts associated with the user.
	PosixAccounts []*PosixAccount `json:"posixAccounts,omitempty"`
	// SshPublicKeys: A map from SSH public key fingerprint to the associated key
	// object.
	SshPublicKeys map[string]SshPublicKey `json:"sshPublicKeys,omitempty"`

	// ServerResponse contains the HTTP response code and headers from the server.
	googleapi.ServerResponse `json:"-"`
	// ForceSendFields is a list of field names (e.g. "Name") to unconditionally
	// include in API requests. By default, fields with empty or default values are
	// omitted from API requests. See
	// https://pkg.go.dev/google.golang.org/api#hdr-ForceSendFields for more
	// details.
	ForceSendFields []string `json:"-"`
	// NullFields is a list of field names (e.g. "Name") to include in API requests
	// with the JSON null value. By default, fields with empty values are omitted
	// from API requests. See
	// https://pkg.go.dev/google.golang.org/api#hdr-NullFields for more details.
	NullFields []string `json:"-"`
}

func (s *LoginProfile) MarshalJSON() ([]byte, error) {
	type NoMethod LoginProfile
	return gensupport.MarshalJSON(NoMethod(*s), s.ForceSendFields, s.NullFields)
}

// PosixAccount: The POSIX account information associated with a Google
// account.
type PosixAccount struct {
	// AccountId: Output only. A POSIX account identifier.
	AccountId string `json:"accountId,omitempty"`
	// Gecos: The GECOS (user information) entry for this account.
	Gecos string `json:"gecos,omitempty"`
	// Gid: The default group ID.
	Gid int64 `json:"gid,omitempty,string"`
	// HomeDirectory: The path to the home directory for this account.
	HomeDirectory string `json:"homeDirectory,omitempty"`
	// Name: Output only. The canonical resource name.
	Name string `json:"name,omitempty"`
	// OperatingSystemType: The operating system type where this account applies.
	//
	// Possible values:
	//   "OPERATING_SYSTEM_TYPE_UNSPECIFIED" - The operating system type associated
	// with the user account information is unspecified.
	//   "LINUX" - Linux user account information.
	//   "WINDOWS" - Windows user account information.
	OperatingSystemType string `json:"operatingSystemType,omitempty"`
	// Primary: Only one POSIX account can be marked as primary.
	Primary bool `json:"primary,omitempty"`
	// Shell: The path to the logic shell for this account.
	Shell string `json:"shell,omitempty"`
	// SystemId: System identifier for which account the username or uid applies
	// to. By default, the empty value is used.
	SystemId string `json:"systemId,omitempty"`
	// Uid: The user ID.
	Uid int64 `json:"uid,omitempty,string"`
	// Username: The username of the POSIX account.
	Username string `json:"username,omitempty"`
	// ForceSendFields is a list of field names (e.g. "AccountId") to
	// unconditionally include in API requests. By default, fields with empty or
	// default values are omitted from API requests. See
	// https://pkg.go.dev/google.golang.org/api#hdr-ForceSendFields for more
	// details.
	ForceSendFields []string `json:"-"`
	// NullFields is a list of field names (e.g. "AccountId") to include in API
	// requests with the JSON null value. By default, fields with empty values are
	// omitted from API requests. See
	// https://pkg.go.dev/google.golang.org/api#hdr-NullFields for more details.
	NullFields []string `json:"-"`
}

func (s *PosixAccount) MarshalJSON() ([]byte, error) {
	type NoMethod PosixAccount
	return gensupport.MarshalJSON(NoMethod(*s), s.ForceSendFields, s.NullFields)
}

// SshPublicKey: The SSH public key information associated with a Google
// account.
type SshPublicKey struct {
	// ExpirationTimeUsec: An expiration time in microseconds since epoch.
	ExpirationTimeUsec int64 `json:"expirationTimeUsec,omitempty,string"`
	// Fingerprint: Output only. The SHA-256 fingerprint of the SSH public key.
	Fingerprint string `json:"fingerprint,omitempty"`
	// Key: Public key text in SSH format, defined by RFC4253 section 6.6.
	Key string `json:"key,omitempty"`
	// Name: Output only. The canonical resource name.
	Name string `json:"name,omitempty"`

	// ServerResponse contains the HTTP response code and headers from the server.
	googleapi.ServerResponse `json:"-"`
	// ForceSendFields is a list of field names (e.g. "ExpirationTimeUsec") to
	// unconditionally include in API requests. By default, fields with empty or
	// default values are omitted from API requests. See
	// https://pkg.go.dev/google.golang.org/api#hdr-ForceSendFields for more
	// details.
	ForceSendFields []string `json:"-"`
	// NullFields is a list of field names (e.g. "ExpirationTimeUsec") to include
	// in API requests with the JSON null value. By default, fields with empty
	// values are omitted from API requests. See
	// https://pkg.go.dev/google.golang.org/api#hdr-NullFields for more details.
	NullFields []string `json:"-"`
}

func (s *SshPublicKey) MarshalJSON() ([]byte, error) {
	type NoMethod SshPublicKey
	return gensupport.MarshalJSON(NoMethod(*s), s.ForceSendFields, s.NullFields)
}

type UsersGetLoginProfileCall struct {
	s            *Service
	name         string
	urlParams_   gensupport.URLParams
	ifNoneMatch_ string
	ctx_         context.Context
	header_      http.Header
}

// GetLoginProfile: Retrieves the profile information used for logging in to a
// virtual machine on Google Compute Engine.
//
// - name: The unique ID for the user in format `users/{user}`.
func (r *UsersService) GetLoginProfile(name string) *UsersGetLoginProfileCall {
	c := &UsersGetLoginProfileCall{s: r.s, urlParams_: make(gensupport.URLParams)}
	c.name = name
	return c
}

// ProjectId sets the optional parameter "projectId": The project ID of the
// Google Cloud Platform project.
func (c *UsersGetLoginProfileCall) ProjectId(projectId string) *UsersGetLoginProfileCall {
	c.urlParams_.Set("projectId", projectId)
	return c
}

// SystemId sets the optional parameter "systemId": A system ID for filtering
// the results of the request.
func (c *UsersGetLoginProfileCall) SystemId(systemId string) *UsersGetLoginProfileCall {
	c.urlParams_.Set("systemId", systemId)
	return c
}

// Fields allows partial responses to be retrieved. See
// https://developers.google.com/gdata/docs/2.0/basics#PartialResponse for more
// details.
func (c *UsersGetLoginProfileCall) Fields(s ...googleapi.Field) *UsersGetLoginProfileCall {
	c.urlParams_.Set("fields", googleapi.CombineFields(s))
	return c
}

// IfNoneMatch sets an optional parameter which makes the operation fail if the
// object's ETag matches the given value. This is useful for getting updates
// only after the object has changed since the last request.
func (c *UsersGetLoginProfileCall) IfNoneMatch(entityTag string) *UsersGetLoginProfileCall {
	c.ifNoneMatch_ = entityTag
	return c
}

// Context sets the context to be used in this call's Do method.
func (c *UsersGetLoginProfileCall) Context(ctx context.Context) *UsersGetLoginProfileCall {
	c.ctx_ = ctx
	return c
}

// Header returns a http.Header that can be modified by the caller to add
// headers to the request.
func (c *UsersGetLoginProfileCall) Header() http.Header {
	if c.header_ == nil {
		c.header_ = make(http.Header)
	}
	return c.header_
}

func (c *UsersGetLoginProfileCall) doRequest(alt string) (*http.Response, error) {
	reqHeaders := gensupport.SetHeaders(c.s.userAgent(), "", c.header_)
	if c.ifNoneMatch_ != "" {
		reqHeaders.Set("If-None-Match", c.ifNoneMatch_)
	}
	var body io.Reader = nil
	c.urlParams_.Set("alt", alt)
	c.urlParams_.Set("prettyPrint", "false")
	urls := googleapi.ResolveRelative(c.s.BasePath, "v1/{+name}/loginProfile")
	urls += "?" + c.urlParams_.Encode()
	req, err := http.NewRequest("GET", urls, body)
	if err != nil {
		return nil, err
	}
	req.Header = reqHeaders
	googleapi.Expand(req.URL, map[string]string{
		"name": c.name,
	})
	return gensupport.SendRequest(c.ctx_, c.s.client, req)
}

// Do executes the "oslogin.users.getLoginProfile" call.
// Any non-2xx status code is an error. Response headers are in either
// *LoginProfile.ServerResponse.Header or (if a response was returned at all)
// in error.(*googleapi.Error).Header. Use googleapi.IsNotModified to check
// whether the returned error was because http.StatusNotModified was returned.
func (c *UsersGetLoginProfileCall) Do(opts ...googleapi.CallOption) (*LoginProfile, error) {
	gensupport.SetOptions(c.urlParams_, opts...)
	res, err := c.doRequest("json")
	if res != nil && res.StatusCode == http.StatusNotModified {
		if res.Body != nil {
			res.Body.Close()
		}
		return nil, gensupport.WrapError(&googleapi.Error{
			Code:   res.StatusCode,
			Header: res.Header,
		})
	}
	if err != nil {
		return nil, err
	}
	defer googleapi.CloseBody(res)
	if err := googleapi.CheckResponse(res); err != nil {
		return nil, gensupport.WrapError(err)
	}
	ret := &LoginProfile{
		ServerResponse: googleapi.ServerResponse{
			Header:         res.Header,
			HTTPStatusCode: res.StatusCode,
		},
	}
	target := &ret
	if err := gensupport.DecodeResponse(target, res); err != nil {
		return nil, err
	}
	return ret, nil
}

type UsersImportSshPublicKeyCall struct {
	s            *Service
	parent       string
	sshpublickey *SshPublicKey
	urlParams_   gensupport.URLParams
	ctx_         context.Context
	header_      http.Header
}

// ImportSshPublicKey: Adds an SSH public key and returns the profile
// information. Default POSIX account information is set when no username and
// UID exist as part of the login profile.
//
// - parent: The unique ID for the user in format `users/{user}`.
func (r *UsersService) ImportSshPublicKey(parent string, sshpublickey *SshPublicKey) *UsersImportSshPublicKeyCall {
	c := &UsersImportSshPublicKeyCall{s: r.s, urlParams_: make(gensupport.URLParams)}
	c.parent = parent
	c.sshpublickey = sshpublickey
	return c
}

// ProjectId sets the optional parameter "projectId": The project ID of the
// Google Cloud Platform project.
func (c *UsersImportSshPublicKeyCall) ProjectId(projectId string) *UsersImportSshPublicKeyCall {
	c.urlParams_.Set("projectId", projectId)
	return c
}

// Regions sets the optional parameter "regions": The regions to which to
// assert that the key was written. If unspecified, defaults to all regions.
// Regions are listed at https://cloud.google.com/about/locations#region.
func (c *UsersImportSshPublicKeyCall) Regions(regions ...string) *UsersImportSshPublicKeyCall {
	c.urlParams_.SetMulti("regions", append([]string{}, regions...))
	return c
}

// Fields allows partial responses to be retrieved. See
// https://developers.google.com/gdata/docs/2.0/basics#PartialResponse for more
// details.
func (c *UsersImportSshPublicKeyCall) Fields(s ...googleapi.Field) *UsersImportSshPublicKeyCall {
	c.urlParams_.Set("fields", googleapi.CombineFields(s))
	return c
}

// Context sets the context to be used in this call's Do method.
func (c *UsersImportSshPublicKeyCall) Context(ctx context.Context) *UsersImportSshPublicKeyCall {
	c.ctx_ = ctx
	return c
}

// Header returns a http.Header that can be modified by the caller to add
// headers to the request.
func (c *UsersImportSshPublicKeyCall) Header() http.Header {
	if c.header_ == nil {
		c.header_ = make(http.Header)
	}
	return c.header_
}

func (c *UsersImportSshPublicKeyCall) doRequest(alt string) (*http.Response, error) {
	reqHeaders := gensupport.SetHeaders(c.s.userAgent(), "application/json", c.header_)
	var body io.Reader = nil
	body, err := googleapi.WithoutDataWrapper.JSONReader(c.sshpublickey)
	if err != nil {
		return nil, err
	}
	c.urlParams_.Set("alt", alt)
	c.urlParams_.Set("prettyPrint", "false")
	urls := googleapi.ResolveRelative(c.s.BasePath, "v1/{+parent}:importSshPublicKey")
	urls += "?" + c.urlParams_.Encode()
	req, err := http.NewRequest("POST", urls, body)
	if err != nil {
		return nil, err
	}
	req.Header = reqHeaders
	googleapi.Expand(req.URL, map[string]string{
		"parent": c.parent,
	})
	return gensupport.SendRequest(c.ctx_, c.s.client, req)
}

// Do executes the "oslogin.users.importSshPublicKey" call.
// Any non-2xx status code is an error. Response headers are in either
// *ImportSshPublicKeyResponse.ServerResponse.Header or (if a response was
// returned at all) in error.(*googleapi.Error).Header. Use
// googleapi.IsNotModified to check whether the returned error was because
// http.StatusNotModified was returned.
func (c *UsersImportSshPublicKeyCall) Do(opts ...googleapi.CallOption) (*ImportSshPublicKeyResponse, error) {
	gensupport.SetOptions(c.urlParams_, opts...)
	res, err := c.doRequest("json")
	if res != nil && res.StatusCode == http.StatusNotModified {
		if res.Body != nil {
			res.Body.Close()
		}
		return nil, gensupport.WrapError(&googleapi.Error{
			Code:   res.StatusCode,
			Header: res.Header,
		})
	}
	if err != nil {
		return nil, err
	}
	defer googleapi.CloseBody(res)
	if err := googleapi.CheckResponse(res); err != nil {
		return nil, gensupport.WrapError(err)
	}
	ret := &ImportSshPublicKeyResponse{
		ServerResponse: googleapi.ServerResponse{
			Header:         res.Header,
			HTTPStatusCode: res.StatusCode,
		},
	}
	target := &ret
	if err := gensupport.DecodeResponse(target, res); err != nil {
		return nil, err
	}
	return ret, nil
}

type UsersProjectsDeleteCall struct {
	s          *Service
	name       string
	urlParams_ gensupport.URLParams
	ctx_       context.Context
	header_    http.Header
}

// Delete: Deletes a POSIX account.
//
//   - name: A reference to the POSIX account to update. POSIX accounts are
//     identified by the project ID they are associated with. A reference to the
//     POSIX account is in format `users/{user}/projects/{project}`.
func (r *UsersProjectsService) Delete(name string) *UsersProjectsDeleteCall {
	c := &UsersProjectsDeleteCall{s: r.s, urlParams_: make(gensupport.URLParams)}
	c.name = name
	return c
}

// Fields allows partial responses to be retrieved. See
// https://developers.google.com/gdata/docs/2.0/basics#PartialResponse for more
// details.
func (c *UsersProjectsDeleteCall) Fields(s ...googleapi.Field) *UsersProjectsDeleteCall {
	c.urlParams_.Set("fields", googleapi.CombineFields(s))
	return c
}

// Context sets the context to be used in this call's Do method.
func (c *UsersProjectsDeleteCall) Context(ctx context.Context) *UsersProjectsDeleteCall {
	c.ctx_ = ctx
	return c
}

// Header returns a http.Header that can be modified by the caller to add
// headers to the request.
func (c *UsersProjectsDeleteCall) Header() http.Header {
	if c.header_ == nil {
		c.header_ = make(http.Header)
	}
	return c.header_
}

func (c *UsersProjectsDeleteCall) doRequest(alt string) (*http.Response, error) {
	reqHeaders := gensupport.SetHeaders(c.s.userAgent(), "", c.header_)
	var body io.Reader = nil
	c.urlParams_.Set("alt", alt)
	c.urlParams_.Set("prettyPrint", "false")
	urls := googleapi.ResolveRelative(c.s.BasePath, "v1/{+name}")
	urls += "?" + c.urlParams_.Encode()
	req, err := http.NewRequest("DELETE", urls, body)
	if err != nil {
		return nil, err
	}
	req.Header = reqHeaders
	googleapi.Expand(req.URL, map[string]string{
		"name": c.name,
	})
	return gensupport.SendRequest(c.ctx_, c.s.client, req)
}

// Do executes the "oslogin.users.projects.delete" call.
// Any non-2xx status code is an error. Response headers are in either
// *Empty.ServerResponse.Header or (if a response was returned at all) in
// error.(*googleapi.Error).Header. Use googleapi.IsNotModified to check
// whether the returned error was because http.StatusNotModified was returned.
func (c *UsersProjectsDeleteCall) Do(opts ...googleapi.CallOption) (*Empty, error) {
	gensupport.SetOptions(c.urlParams_, opts...)
	res, err := c.doRequest("json")
	if res != nil && res.StatusCode == http.StatusNotModified {
		if res.Body != nil {
			res.Body.Close()
		}
		return nil, gensupport.WrapError(&googleapi.Error{
			Code:   res.StatusCode,
			Header: res.Header,
		})
	}
	if err != nil {
		return nil, err
	}
	defer googleapi.CloseBody(res)
	if err := googleapi.CheckResponse(res); err != nil {
		return nil, gensupport.WrapError(err)
	}
	ret := &Empty{
		ServerResponse: googleapi.ServerResponse{
			Header:         res.Header,
			HTTPStatusCode: res.StatusCode,
		},
	}
	target := &ret
	if err := gensupport.DecodeResponse(target, res); err != nil {
		return nil, err
	}
	return ret, nil
}

type UsersSshPublicKeysCreateCall struct {
	s            *Service
	parent       string
	sshpublickey *SshPublicKey
	urlParams_   gensupport.URLParams
	ctx_         context.Context
	header_      http.Header
}

// Create: Create an SSH public key
//
// - parent: The unique ID for the user in format `users/{user}`.
func (r *UsersSshPublicKeysService) Create(parent string, sshpublickey *SshPublicKey) *UsersSshPublicKeysCreateCall {
	c := &UsersSshPublicKeysCreateCall{s: r.s, urlParams_: make(gensupport.URLParams)}
	c.parent = parent
	c.sshpublickey = sshpublickey
	return c
}

// Fields allows partial responses to be retrieved. See
// https://developers.google.com/gdata/docs/2.0/basics#PartialResponse for more
// details.
func (c *UsersSshPublicKeysCreateCall) Fields(s ...googleapi.Field) *UsersSshPublicKeysCreateCall {
	c.urlParams_.Set("fields", googleapi.CombineFields(s))
	return c
}

// Context sets the context to be used in this call's Do method.
func (c *UsersSshPublicKeysCreateCall) Context(ctx context.Context) *UsersSshPublicKeysCreateCall {
	c.ctx_ = ctx
	return c
}

// Header returns a http.Header that can be modified by the caller to add
// headers to the request.
func (c *UsersSshPublicKeysCreateCall) Header() http.Header {
	if c.header_ == nil {
		c.header_ = make(http.Header)
	}
	return c.header_
}

func (c *UsersSshPublicKeysCreateCall) doRequest(alt string) (*http.Response, error) {
	reqHeaders := gensupport.SetHeaders(c.s.userAgent(), "application/json", c.header_)
	var body io.Reader = nil
	body, err := googleapi.WithoutDataWrapper.JSONReader(c.sshpublickey)
	if err != nil {
		return nil, err
	}
	c.urlParams_.Set("alt", alt)
	c.urlParams_.Set("prettyPrint", "false")
	urls := googleapi.ResolveRelative(c.s.BasePath, "v1/{+parent}/sshPublicKeys")
	urls += "?" + c.urlParams_.Encode()
	req, err := http.NewRequest("POST", urls, body)
	if err != nil {
		return nil, err
	}
	req.Header = reqHeaders
	googleapi.Expand(req.URL, map[string]string{
		"parent": c.parent,
	})
	return gensupport.SendRequest(c.ctx_, c.s.client, req)
}

// Do executes the "oslogin.users.sshPublicKeys.create" call.
// Any non-2xx status code is an error. Response headers are in either
// *SshPublicKey.ServerResponse.Header or (if a response was returned at all)
// in error.(*googleapi.Error).Header. Use googleapi.IsNotModified to check
// whether the returned error was because http.StatusNotModified was returned.
func (c *UsersSshPublicKeysCreateCall) Do(opts ...googleapi.CallOption) (*SshPublicKey, error) {
	gensupport.SetOptions(c.urlParams_, opts...)
	res, err := c.doRequest("json")
	if res != nil && res.StatusCode == http.StatusNotModified {
		if res.Body != nil {
			res.Body.Close()
		}
		return nil, gensupport.WrapError(&googleapi.Error{
			Code:   res.StatusCode,
			Header: res.Header,
		})
	}
	if err != nil {
		return nil, err
	}
	defer googleapi.CloseBody(res)
	if err := googleapi.CheckResponse(res); err != nil {
		return nil, gensupport.WrapError(err)
	}
	ret := &SshPublicKey{
		ServerResponse: googleapi.ServerResponse{
			Header:         res.Header,
			HTTPStatusCode: res.StatusCode,
		},
	}
	target := &ret
	if err := gensupport.DecodeResponse(target, res); err != nil {
		return nil, err
	}
	return ret, nil
}

type UsersSshPublicKeysDeleteCall struct {
	s          *Service
	name       string
	urlParams_ gensupport.URLParams
	ctx_       context.Context
	header_    http.Header
}

// Delete: Deletes an SSH public key.
//
//   - name: The fingerprint of the public key to update. Public keys are
//     identified by their SHA-256 fingerprint. The fingerprint of the public key
//     is in format `users/{user}/sshPublicKeys/{fingerprint}`.
func (r *UsersSshPublicKeysService) Delete(name string) *UsersSshPublicKeysDeleteCall {
	c := &UsersSshPublicKeysDeleteCall{s: r.s, urlParams_: make(gensupport.URLParams)}
	c.name = name
	return c
}

// Fields allows partial responses to be retrieved. See
// https://developers.google.com/gdata/docs/2.0/basics#PartialResponse for more
// details.
func (c *UsersSshPublicKeysDeleteCall) Fields(s ...googleapi.Field) *UsersSshPublicKeysDeleteCall {
	c.urlParams_.Set("fields", googleapi.CombineFields(s))
	return c
}

// Context sets the context to be used in this call's Do method.
func (c *UsersSshPublicKeysDeleteCall) Context(ctx context.Context) *UsersSshPublicKeysDeleteCall {
	c.ctx_ = ctx
	return c
}

// Header returns a http.Header that can be modified by the caller to add
// headers to the request.
func (c *UsersSshPublicKeysDeleteCall) Header() http.Header {
	if c.header_ == nil {
		c.header_ = make(http.Header)
	}
	return c.header_
}

func (c *UsersSshPublicKeysDeleteCall) doRequest(alt string) (*http.Response, error) {
	reqHeaders := gensupport.SetHeaders(c.s.userAgent(), "", c.header_)
	var body io.Reader = nil
	c.urlParams_.Set("alt", alt)
	c.urlParams_.Set("prettyPrint", "false")
	urls := googleapi.ResolveRelative(c.s.BasePath, "v1/{+name}")
	urls += "?" + c.urlParams_.Encode()
	req, err := http.NewRequest("DELETE", urls, body)
	if err != nil {
		return nil, err
	}
	req.Header = reqHeaders
	googleapi.Expand(req.URL, map[string]string{
		"name": c.name,
	})
	return gensupport.SendRequest(c.ctx_, c.s.client, req)
}

// Do executes the "oslogin.users.sshPublicKeys.delete" call.
// Any non-2xx status code is an error. Response headers are in either
// *Empty.ServerResponse.Header or (if a response was returned at all) in
// error.(*googleapi.Error).Header. Use googleapi.IsNotModified to check
// whether the returned error was because http.StatusNotModified was returned.
func (c *UsersSshPublicKeysDeleteCall) Do(opts ...googleapi.CallOption) (*Empty, error) {
	gensupport.SetOptions(c.urlParams_, opts...)
	res, err := c.doRequest("json")
	if res != nil && res.StatusCode == http.StatusNotModified {
		if res.Body != nil {
			res.Body.Close()
		}
		return nil, gensupport.WrapError(&googleapi.Error{
			Code:   res.StatusCode,
			Header: res.Header,
		})
	}
	if err != nil {
		return nil, err
	}
	defer googleapi.CloseBody(res)
	if err := googleapi.CheckResponse(res); err != nil {
		return nil, gensupport.WrapError(err)
	}
	ret := &Empty{
		ServerResponse: googleapi.ServerResponse{
			Header:         res.Header,
			HTTPStatusCode: res.StatusCode,
		},
	}
	target := &ret
	if err := gensupport.DecodeResponse(target, res); err != nil {
		return nil, err
	}
	return ret, nil
}

type UsersSshPublicKeysGetCall struct {
	s            *Service
	name         string
	urlParams_   gensupport.URLParams
	ifNoneMatch_ string
	ctx_         context.Context
	header_      http.Header
}

// Get: Retrieves an SSH public key.
//
//   - name: The fingerprint of the public key to retrieve. Public keys are
//     identified by their SHA-256 fingerprint. The fingerprint of the public key
//     is in format `users/{user}/sshPublicKeys/{fingerprint}`.
func (r *UsersSshPublicKeysService) Get(name string) *UsersSshPublicKeysGetCall {
	c := &UsersSshPublicKeysGetCall{s: r.s, urlParams_: make(gensupport.URLParams)}
	c.name = name
	return c
}

// Fields allows partial responses to be retrieved. See
// https://developers.google.com/gdata/docs/2.0/basics#PartialResponse for more
// details.
func (c *UsersSshPublicKeysGetCall) Fields(s ...googleapi.Field) *UsersSshPublicKeysGetCall {
	c.urlParams_.Set("fields", googleapi.CombineFields(s))
	return c
}

// IfNoneMatch sets an optional parameter which makes the operation fail if the
// object's ETag matches the given value. This is useful for getting updates
// only after the object has changed since the last request.
func (c *UsersSshPublicKeysGetCall) IfNoneMatch(entityTag string) *UsersSshPublicKeysGetCall {
	c.ifNoneMatch_ = entityTag
	return c
}

// Context sets the context to be used in this call's Do method.
func (c *UsersSshPublicKeysGetCall) Context(ctx context.Context) *UsersSshPublicKeysGetCall {
	c.ctx_ = ctx
	return c
}

// Header returns a http.Header that can be modified by the caller to add
// headers to the request.
func (c *UsersSshPublicKeysGetCall) Header() http.Header {
	if c.header_ == nil {
		c.header_ = make(http.Header)
	}
	return c.header_
}

func (c *UsersSshPublicKeysGetCall) doRequest(alt string) (*http.Response, error) {
	reqHeaders := gensupport.SetHeaders(c.s.userAgent(), "", c.header_)
	if c.ifNoneMatch_ != "" {
		reqHeaders.Set("If-None-Match", c.ifNoneMatch_)
	}
	var body io.Reader = nil
	c.urlParams_.Set("alt", alt)
	c.urlParams_.Set("prettyPrint", "false")
	urls := googleapi.ResolveRelative(c.s.BasePath, "v1/{+name}")
	urls += "?" + c.urlParams_.Encode()
	req, err := http.NewRequest("GET", urls, body)
	if err != nil {
		return nil, err
	}
	req.Header = reqHeaders
	googleapi.Expand(req.URL, map[string]string{
		"name": c.name,
	})
	return gensupport.SendRequest(c.ctx_, c.s.client, req)
}

// Do executes the "oslogin.users.sshPublicKeys.get" call.
// Any non-2xx status code is an error. Response headers are in either
// *SshPublicKey.ServerResponse.Header or (if a response was returned at all)
// in error.(*googleapi.Error).Header. Use googleapi.IsNotModified to check
// whether the returned error was because http.StatusNotModified was returned.
func (c *UsersSshPublicKeysGetCall) Do(opts ...googleapi.CallOption) (*SshPublicKey, error) {
	gensupport.SetOptions(c.urlParams_, opts...)
	res, err := c.doRequest("json")
	if res != nil && res.StatusCode == http.StatusNotModified {
		if res.Body != nil {
			res.Body.Close()
		}
		return nil, gensupport.WrapError(&googleapi.Error{
			Code:   res.StatusCode,
			Header: res.Header,
		})
	}
	if err != nil {
		return nil, err
	}
	defer googleapi.CloseBody(res)
	if err := googleapi.CheckResponse(res); err != nil {
		return nil, gensupport.WrapError(err)
	}
	ret := &SshPublicKey{
		ServerResponse: googleapi.ServerResponse{
			Header:         res.Header,
			HTTPStatusCode: res.StatusCode,
		},
	}
	target := &ret
	if err := gensupport.DecodeResponse(target, res); err != nil {
		return nil, err
	}
	return ret, nil
}

type UsersSshPublicKeysPatchCall struct {
	s            *Service
	name         string
	sshpublickey *SshPublicKey
	urlParams_   gensupport.URLParams
	ctx_         context.Context
	header_      http.Header
}

// Patch: Updates an SSH public key and returns the profile information. This
// method supports patch semantics.
//
//   - name: The fingerprint of the public key to update. Public keys are
//     identified by their SHA-256 fingerprint. The fingerprint of the public key
//     is in format `users/{user}/sshPublicKeys/{fingerprint}`.
func (r *UsersSshPublicKeysService) Patch(name string, sshpublickey *SshPublicKey) *UsersSshPublicKeysPatchCall {
	c := &UsersSshPublicKeysPatchCall{s: r.s, urlParams_: make(gensupport.URLParams)}
	c.name = name
	c.sshpublickey = sshpublickey
	return c
}

// UpdateMask sets the optional parameter "updateMask": Mask to control which
// fields get updated. Updates all if not present.
func (c *UsersSshPublicKeysPatchCall) UpdateMask(updateMask string) *UsersSshPublicKeysPatchCall {
	c.urlParams_.Set("updateMask", updateMask)
	return c
}

// Fields allows partial responses to be retrieved. See
// https://developers.google.com/gdata/docs/2.0/basics#PartialResponse for more
// details.
func (c *UsersSshPublicKeysPatchCall) Fields(s ...googleapi.Field) *UsersSshPublicKeysPatchCall {
	c.urlParams_.Set("fields", googleapi.CombineFields(s))
	return c
}

// Context sets the context to be used in this call's Do method.
func (c *UsersSshPublicKeysPatchCall) Context(ctx context.Context) *UsersSshPublicKeysPatchCall {
	c.ctx_ = ctx
	return c
}

// Header returns a http.Header that can be modified by the caller to add
// headers to the request.
func (c *UsersSshPublicKeysPatchCall) Header() http.Header {
	if c.header_ == nil {
		c.header_ = make(http.Header)
	}
	return c.header_
}

func (c *UsersSshPublicKeysPatchCall) doRequest(alt string) (*http.Response, error) {
	reqHeaders := gensupport.SetHeaders(c.s.userAgent(), "application/json", c.header_)
	var body io.Reader = nil
	body, err := googleapi.WithoutDataWrapper.JSONReader(c.sshpublickey)
	if err != nil {
		return nil, err
	}
	c.urlParams_.Set("alt", alt)
	c.urlParams_.Set("prettyPrint", "false")
	urls := googleapi.ResolveRelative(c.s.BasePath, "v1/{+name}")
	urls += "?" + c.urlParams_.Encode()
	req, err := http.NewRequest("PATCH", urls, body)
	if err != nil {
		return nil, err
	}
	req.Header = reqHeaders
	googleapi.Expand(req.URL, map[string]string{
		"name": c.name,
	})
	return gensupport.SendRequest(c.ctx_, c.s.client, req)
}

// Do executes the "oslogin.users.sshPublicKeys.patch" call.
// Any non-2xx status code is an error. Response headers are in either
// *SshPublicKey.ServerResponse.Header or (if a response was returned at all)
// in error.(*googleapi.Error).Header. Use googleapi.IsNotModified to check
// whether the returned error was because http.StatusNotModified was returned.
func (c *UsersSshPublicKeysPatchCall) Do(opts ...googleapi.CallOption) (*SshPublicKey, error) {
	gensupport.SetOptions(c.urlParams_, opts...)
	res, err := c.doRequest("json")
	if res != nil && res.StatusCode == http.StatusNotModified {
		if res.Body != nil {
			res.Body.Close()
		}
		return nil, gensupport.WrapError(&googleapi.Error{
			Code:   res.StatusCode,
			Header: res.Header,
		})
	}
	if err != nil {
		return nil, err
	}
	defer googleapi.CloseBody(res)
	if err := googleapi.CheckResponse(res); err != nil {
		return nil, gensupport.WrapError(err)
	}
	ret := &SshPublicKey{
		ServerResponse: googleapi.ServerResponse{
			Header:         res.Header,
			HTTPStatusCode: res.StatusCode,
		},
	}
	target := &ret
	if err := gensupport.DecodeResponse(target, res); err != nil {
		return nil, err
	}
	return ret, nil
}
