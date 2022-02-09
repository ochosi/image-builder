// Package v1 provides primitives to interact the openapi HTTP API.
//
// Code generated by github.com/deepmap/oapi-codegen DO NOT EDIT.
package v1

import (
	"bytes"
	"compress/gzip"
	"encoding/base64"
	"fmt"
	"github.com/deepmap/oapi-codegen/pkg/runtime"
	"github.com/getkin/kin-openapi/openapi3"
	"github.com/labstack/echo/v4"
	"net/http"
	"strings"
)

// AWSS3UploadRequestOptions defines model for AWSS3UploadRequestOptions.
type AWSS3UploadRequestOptions map[string]interface{}

// AWSS3UploadStatus defines model for AWSS3UploadStatus.
type AWSS3UploadStatus struct {
	Url string `json:"url"`
}

// AWSUploadRequestOptions defines model for AWSUploadRequestOptions.
type AWSUploadRequestOptions struct {
	ShareWithAccounts []string `json:"share_with_accounts"`
}

// AWSUploadStatus defines model for AWSUploadStatus.
type AWSUploadStatus struct {
	Ami    string `json:"ami"`
	Region string `json:"region"`
}

// ArchitectureItem defines model for ArchitectureItem.
type ArchitectureItem struct {
	Arch       string   `json:"arch"`
	ImageTypes []string `json:"image_types"`
}

// Architectures defines model for Architectures.
type Architectures []ArchitectureItem

// AzureUploadRequestOptions defines model for AzureUploadRequestOptions.
type AzureUploadRequestOptions struct {

	// Name of the resource group where the image should be uploaded.
	ResourceGroup string `json:"resource_group"`

	// ID of subscription where the image should be uploaded.
	SubscriptionId string `json:"subscription_id"`

	// ID of the tenant where the image should be uploaded. This link explains how
	// to find it in the Azure Portal:
	// https://docs.microsoft.com/en-us/azure/active-directory/fundamentals/active-directory-how-to-find-tenant
	TenantId string `json:"tenant_id"`
}

// AzureUploadStatus defines model for AzureUploadStatus.
type AzureUploadStatus struct {
	ImageName string `json:"image_name"`
}

// ComposeMetadata defines model for ComposeMetadata.
type ComposeMetadata struct {

	// ID (hash) of the built commit
	OstreeCommit *string `json:"ostree_commit,omitempty"`

	// Package list including NEVRA
	Packages *[]PackageMetadata `json:"packages,omitempty"`
}

// ComposeRequest defines model for ComposeRequest.
type ComposeRequest struct {
	Customizations *Customizations `json:"customizations,omitempty"`
	Distribution   string          `json:"distribution"`
	ImageRequests  []ImageRequest  `json:"image_requests"`
}

// ComposeResponse defines model for ComposeResponse.
type ComposeResponse struct {
	Id string `json:"id"`
}

// ComposeStatus defines model for ComposeStatus.
type ComposeStatus struct {
	ImageStatus ImageStatus `json:"image_status"`
}

// ComposesResponse defines model for ComposesResponse.
type ComposesResponse struct {
	Data  []ComposesResponseItem `json:"data"`
	Links struct {
		First string `json:"first"`
		Last  string `json:"last"`
	} `json:"links"`
	Meta struct {
		Count int `json:"count"`
	} `json:"meta"`
}

// ComposesResponseItem defines model for ComposesResponseItem.
type ComposesResponseItem struct {
	CreatedAt string      `json:"created_at"`
	Id        string      `json:"id"`
	Request   interface{} `json:"request"`
}

// Customizations defines model for Customizations.
type Customizations struct {
	Filesystem   *[]Filesystem `json:"filesystem,omitempty"`
	Packages     *[]string     `json:"packages,omitempty"`
	Subscription *Subscription `json:"subscription,omitempty"`
}

// DistributionItem defines model for DistributionItem.
type DistributionItem struct {
	Description string `json:"description"`
	Name        string `json:"name"`
}

// Distributions defines model for Distributions.
type Distributions []DistributionItem

// Filesystem defines model for Filesystem.
type Filesystem struct {
	MinSize    int    `json:"min_size"`
	Mountpoint string `json:"mountpoint"`
}

// GCPUploadRequestOptions defines model for GCPUploadRequestOptions.
type GCPUploadRequestOptions struct {

	// List of valid Google accounts to share the imported Compute Node image with.
	// Each string must contain a specifier of the account type. Valid formats are:
	//   - 'user:{emailid}': An email address that represents a specific
	//     Google account. For example, 'alice@example.com'.
	//   - 'serviceAccount:{emailid}': An email address that represents a
	//     service account. For example, 'my-other-app@appspot.gserviceaccount.com'.
	//   - 'group:{emailid}': An email address that represents a Google group.
	//     For example, 'admins@example.com'.
	//   - 'domain:{domain}': The G Suite domain (primary) that represents all
	//     the users of that domain. For example, 'google.com' or 'example.com'.
	//     If not specified, the imported Compute Node image is not shared with any
	//     account.
	ShareWithAccounts []string `json:"share_with_accounts"`
}

// GCPUploadStatus defines model for GCPUploadStatus.
type GCPUploadStatus struct {
	ImageName string `json:"image_name"`
	ProjectId string `json:"project_id"`
}

// HTTPError defines model for HTTPError.
type HTTPError struct {
	Detail string `json:"detail"`
	Title  string `json:"title"`
}

// HTTPErrorList defines model for HTTPErrorList.
type HTTPErrorList struct {
	Errors []HTTPError `json:"errors"`
}

// ImageRequest defines model for ImageRequest.
type ImageRequest struct {
	Architecture  string        `json:"architecture"`
	ImageType     ImageTypes    `json:"image_type"`
	Ostree        *OSTree       `json:"ostree,omitempty"`
	UploadRequest UploadRequest `json:"upload_request"`
}

// ImageStatus defines model for ImageStatus.
type ImageStatus struct {
	Status       string        `json:"status"`
	UploadStatus *UploadStatus `json:"upload_status,omitempty"`
}

// ImageTypes defines model for ImageTypes.
type ImageTypes string

// List of ImageTypes
const (
	ImageTypes_ami                 ImageTypes = "ami"
	ImageTypes_aws                 ImageTypes = "aws"
	ImageTypes_azure               ImageTypes = "azure"
	ImageTypes_edge_commit         ImageTypes = "edge-commit"
	ImageTypes_edge_container      ImageTypes = "edge-container"
	ImageTypes_edge_installer      ImageTypes = "edge-installer"
	ImageTypes_gcp                 ImageTypes = "gcp"
	ImageTypes_guest_image         ImageTypes = "guest-image"
	ImageTypes_image_installer     ImageTypes = "image-installer"
	ImageTypes_rhel_edge_commit    ImageTypes = "rhel-edge-commit"
	ImageTypes_rhel_edge_installer ImageTypes = "rhel-edge-installer"
	ImageTypes_vhd                 ImageTypes = "vhd"
	ImageTypes_vsphere             ImageTypes = "vsphere"
)

// OSTree defines model for OSTree.
type OSTree struct {
	Ref *string `json:"ref,omitempty"`
	Url *string `json:"url,omitempty"`
}

// Package defines model for Package.
type Package struct {
	Name    string `json:"name"`
	Summary string `json:"summary"`
}

// PackageMetadata defines model for PackageMetadata.
type PackageMetadata struct {
	Arch      string  `json:"arch"`
	Epoch     *string `json:"epoch,omitempty"`
	Name      string  `json:"name"`
	Release   string  `json:"release"`
	Sigmd5    string  `json:"sigmd5"`
	Signature *string `json:"signature,omitempty"`
	Type      string  `json:"type"`
	Version   string  `json:"version"`
}

// PackagesResponse defines model for PackagesResponse.
type PackagesResponse struct {
	Data  []Package `json:"data"`
	Links struct {
		First string `json:"first"`
		Last  string `json:"last"`
	} `json:"links"`
	Meta struct {
		Count int `json:"count"`
	} `json:"meta"`
}

// Readiness defines model for Readiness.
type Readiness struct {
	Readiness string `json:"readiness"`
}

// Subscription defines model for Subscription.
type Subscription struct {
	ActivationKey string `json:"activation-key"`
	BaseUrl       string `json:"base-url"`
	Insights      bool   `json:"insights"`
	Organization  int    `json:"organization"`
	ServerUrl     string `json:"server-url"`
}

// UploadRequest defines model for UploadRequest.
type UploadRequest struct {
	Options interface{} `json:"options"`
	Type    UploadTypes `json:"type"`
}

// UploadStatus defines model for UploadStatus.
type UploadStatus struct {
	Options interface{} `json:"options"`
	Status  string      `json:"status"`
	Type    UploadTypes `json:"type"`
}

// UploadTypes defines model for UploadTypes.
type UploadTypes string

// List of UploadTypes
const (
	UploadTypes_aws    UploadTypes = "aws"
	UploadTypes_aws_s3 UploadTypes = "aws.s3"
	UploadTypes_azure  UploadTypes = "azure"
	UploadTypes_gcp    UploadTypes = "gcp"
)

// Version defines model for Version.
type Version struct {
	Version string `json:"version"`
}

// ComposeImageJSONBody defines parameters for ComposeImage.
type ComposeImageJSONBody ComposeRequest

// GetComposesParams defines parameters for GetComposes.
type GetComposesParams struct {

	// max amount of composes, default 100
	Limit *int `json:"limit,omitempty"`

	// composes page offset, default 0
	Offset *int `json:"offset,omitempty"`
}

// GetPackagesParams defines parameters for GetPackages.
type GetPackagesParams struct {

	// distribution to look up packages for
	Distribution string `json:"distribution"`

	// architecture to look up packages for
	Architecture string `json:"architecture"`

	// packages to look for
	Search string `json:"search"`

	// max amount of packages, default 100
	Limit *int `json:"limit,omitempty"`

	// packages page offset, default 0
	Offset *int `json:"offset,omitempty"`
}

// ComposeImageRequestBody defines body for ComposeImage for application/json ContentType.
type ComposeImageJSONRequestBody ComposeImageJSONBody

// ServerInterface represents all server handlers.
type ServerInterface interface {
	// get the architectures and their image types available for a given distribution
	// (GET /architectures/{distribution})
	GetArchitectures(ctx echo.Context, distribution string) error
	// compose image
	// (POST /compose)
	ComposeImage(ctx echo.Context) error
	// get a collection of previous compose requests for the logged in user
	// (GET /composes)
	GetComposes(ctx echo.Context, params GetComposesParams) error
	// get status of an image compose
	// (GET /composes/{composeId})
	GetComposeStatus(ctx echo.Context, composeId string) error
	// get metadata of an image compose
	// (GET /composes/{composeId}/metadata)
	GetComposeMetadata(ctx echo.Context, composeId string) error
	// get the available distributions
	// (GET /distributions)
	GetDistributions(ctx echo.Context) error
	// get the openapi json specification
	// (GET /openapi.json)
	GetOpenapiJson(ctx echo.Context) error

	// (GET /packages)
	GetPackages(ctx echo.Context, params GetPackagesParams) error
	// return the readiness
	// (GET /ready)
	GetReadiness(ctx echo.Context) error
	// get the service version
	// (GET /version)
	GetVersion(ctx echo.Context) error
}

// ServerInterfaceWrapper converts echo contexts to parameters.
type ServerInterfaceWrapper struct {
	Handler ServerInterface
}

// GetArchitectures converts echo context to params.
func (w *ServerInterfaceWrapper) GetArchitectures(ctx echo.Context) error {
	var err error
	// ------------- Path parameter "distribution" -------------
	var distribution string

	err = runtime.BindStyledParameter("simple", false, "distribution", ctx.Param("distribution"), &distribution)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter distribution: %s", err))
	}

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.GetArchitectures(ctx, distribution)
	return err
}

// ComposeImage converts echo context to params.
func (w *ServerInterfaceWrapper) ComposeImage(ctx echo.Context) error {
	var err error

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.ComposeImage(ctx)
	return err
}

// GetComposes converts echo context to params.
func (w *ServerInterfaceWrapper) GetComposes(ctx echo.Context) error {
	var err error

	// Parameter object where we will unmarshal all parameters from the context
	var params GetComposesParams
	// ------------- Optional query parameter "limit" -------------

	err = runtime.BindQueryParameter("form", true, false, "limit", ctx.QueryParams(), &params.Limit)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter limit: %s", err))
	}

	// ------------- Optional query parameter "offset" -------------

	err = runtime.BindQueryParameter("form", true, false, "offset", ctx.QueryParams(), &params.Offset)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter offset: %s", err))
	}

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.GetComposes(ctx, params)
	return err
}

// GetComposeStatus converts echo context to params.
func (w *ServerInterfaceWrapper) GetComposeStatus(ctx echo.Context) error {
	var err error
	// ------------- Path parameter "composeId" -------------
	var composeId string

	err = runtime.BindStyledParameter("simple", false, "composeId", ctx.Param("composeId"), &composeId)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter composeId: %s", err))
	}

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.GetComposeStatus(ctx, composeId)
	return err
}

// GetComposeMetadata converts echo context to params.
func (w *ServerInterfaceWrapper) GetComposeMetadata(ctx echo.Context) error {
	var err error
	// ------------- Path parameter "composeId" -------------
	var composeId string

	err = runtime.BindStyledParameter("simple", false, "composeId", ctx.Param("composeId"), &composeId)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter composeId: %s", err))
	}

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.GetComposeMetadata(ctx, composeId)
	return err
}

// GetDistributions converts echo context to params.
func (w *ServerInterfaceWrapper) GetDistributions(ctx echo.Context) error {
	var err error

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.GetDistributions(ctx)
	return err
}

// GetOpenapiJson converts echo context to params.
func (w *ServerInterfaceWrapper) GetOpenapiJson(ctx echo.Context) error {
	var err error

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.GetOpenapiJson(ctx)
	return err
}

// GetPackages converts echo context to params.
func (w *ServerInterfaceWrapper) GetPackages(ctx echo.Context) error {
	var err error

	// Parameter object where we will unmarshal all parameters from the context
	var params GetPackagesParams
	// ------------- Required query parameter "distribution" -------------

	err = runtime.BindQueryParameter("form", true, true, "distribution", ctx.QueryParams(), &params.Distribution)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter distribution: %s", err))
	}

	// ------------- Required query parameter "architecture" -------------

	err = runtime.BindQueryParameter("form", true, true, "architecture", ctx.QueryParams(), &params.Architecture)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter architecture: %s", err))
	}

	// ------------- Required query parameter "search" -------------

	err = runtime.BindQueryParameter("form", true, true, "search", ctx.QueryParams(), &params.Search)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter search: %s", err))
	}

	// ------------- Optional query parameter "limit" -------------

	err = runtime.BindQueryParameter("form", true, false, "limit", ctx.QueryParams(), &params.Limit)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter limit: %s", err))
	}

	// ------------- Optional query parameter "offset" -------------

	err = runtime.BindQueryParameter("form", true, false, "offset", ctx.QueryParams(), &params.Offset)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter offset: %s", err))
	}

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.GetPackages(ctx, params)
	return err
}

// GetReadiness converts echo context to params.
func (w *ServerInterfaceWrapper) GetReadiness(ctx echo.Context) error {
	var err error

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.GetReadiness(ctx)
	return err
}

// GetVersion converts echo context to params.
func (w *ServerInterfaceWrapper) GetVersion(ctx echo.Context) error {
	var err error

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.GetVersion(ctx)
	return err
}

// This is a simple interface which specifies echo.Route addition functions which
// are present on both echo.Echo and echo.Group, since we want to allow using
// either of them for path registration
type EchoRouter interface {
	CONNECT(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	DELETE(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	GET(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	HEAD(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	OPTIONS(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	PATCH(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	POST(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	PUT(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	TRACE(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
}

// RegisterHandlers adds each server route to the EchoRouter.
func RegisterHandlers(router EchoRouter, si ServerInterface) {

	wrapper := ServerInterfaceWrapper{
		Handler: si,
	}

	router.GET("/architectures/:distribution", wrapper.GetArchitectures)
	router.POST("/compose", wrapper.ComposeImage)
	router.GET("/composes", wrapper.GetComposes)
	router.GET("/composes/:composeId", wrapper.GetComposeStatus)
	router.GET("/composes/:composeId/metadata", wrapper.GetComposeMetadata)
	router.GET("/distributions", wrapper.GetDistributions)
	router.GET("/openapi.json", wrapper.GetOpenapiJson)
	router.GET("/packages", wrapper.GetPackages)
	router.GET("/ready", wrapper.GetReadiness)
	router.GET("/version", wrapper.GetVersion)

}

// Base64 encoded, gzipped, json marshaled Swagger object
var swaggerSpec = []string{

	"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",
	"/2aqoE+WNAAA",
}

// GetSwagger returns the Swagger specification corresponding to the generated code
// in this file.
func GetSwagger() (*openapi3.Swagger, error) {
	zipped, err := base64.StdEncoding.DecodeString(strings.Join(swaggerSpec, ""))
	if err != nil {
		return nil, fmt.Errorf("error base64 decoding spec: %s", err)
	}
	zr, err := gzip.NewReader(bytes.NewReader(zipped))
	if err != nil {
		return nil, fmt.Errorf("error decompressing spec: %s", err)
	}
	var buf bytes.Buffer
	_, err = buf.ReadFrom(zr)
	if err != nil {
		return nil, fmt.Errorf("error decompressing spec: %s", err)
	}

	swagger, err := openapi3.NewSwaggerLoader().LoadSwaggerFromData(buf.Bytes())
	if err != nil {
		return nil, fmt.Errorf("error loading Swagger: %s", err)
	}
	return swagger, nil
}
