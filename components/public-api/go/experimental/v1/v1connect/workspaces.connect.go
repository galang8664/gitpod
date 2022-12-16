// Copyright (c) 2022 Gitpod GmbH. All rights reserved.
// Licensed under the GNU Affero General Public License (AGPL).
// See License.AGPL.txt in the project root for license information.

// Code generated by protoc-gen-connect-go. DO NOT EDIT.
//
// Source: gitpod/experimental/v1/workspaces.proto

package v1connect

import (
	context "context"
	errors "errors"
	connect_go "github.com/bufbuild/connect-go"
	v1 "github.com/gitpod-io/gitpod/components/public-api/go/experimental/v1"
	http "net/http"
	strings "strings"
)

// This is a compile-time assertion to ensure that this generated file and the connect package are
// compatible. If you get a compiler error that this constant is not defined, this code was
// generated with a version of connect newer than the one compiled into your binary. You can fix the
// problem by either regenerating this code with an older version of connect or updating the connect
// version compiled into your binary.
const _ = connect_go.IsAtLeastVersion0_1_0

const (
	// WorkspacesServiceName is the fully-qualified name of the WorkspacesService service.
	WorkspacesServiceName = "gitpod.experimental.v1.WorkspacesService"
)

// WorkspacesServiceClient is a client for the gitpod.experimental.v1.WorkspacesService service.
type WorkspacesServiceClient interface {
	// ListWorkspaces enumerates all workspaces belonging to the authenticated user.
	ListWorkspaces(context.Context, *connect_go.Request[v1.ListWorkspacesRequest]) (*connect_go.Response[v1.ListWorkspacesResponse], error)
	// GetWorkspace returns a single workspace.
	GetWorkspace(context.Context, *connect_go.Request[v1.GetWorkspaceRequest]) (*connect_go.Response[v1.GetWorkspaceResponse], error)
	// GetOwnerToken returns an owner token.
	GetOwnerToken(context.Context, *connect_go.Request[v1.GetOwnerTokenRequest]) (*connect_go.Response[v1.GetOwnerTokenResponse], error)
	// CreateAndStartWorkspace creates a new workspace and starts it.
	CreateAndStartWorkspace(context.Context, *connect_go.Request[v1.CreateAndStartWorkspaceRequest]) (*connect_go.Response[v1.CreateAndStartWorkspaceResponse], error)
	// StopWorkspace stops a running workspace (instance).
	// Errors:
	//
	//	NOT_FOUND:           the workspace_id is unkown
	//	FAILED_PRECONDITION: if there's no running instance
	StopWorkspace(context.Context, *connect_go.Request[v1.StopWorkspaceRequest]) (*connect_go.ServerStreamForClient[v1.StopWorkspaceResponse], error)
	UpdatePort(context.Context, *connect_go.Request[v1.UpdatePortRequest]) (*connect_go.Response[v1.UpdatePortResponse], error)
}

// NewWorkspacesServiceClient constructs a client for the gitpod.experimental.v1.WorkspacesService
// service. By default, it uses the Connect protocol with the binary Protobuf Codec, asks for
// gzipped responses, and sends uncompressed requests. To use the gRPC or gRPC-Web protocols, supply
// the connect.WithGRPC() or connect.WithGRPCWeb() options.
//
// The URL supplied here should be the base URL for the Connect or gRPC server (for example,
// http://api.acme.com or https://acme.com/grpc).
func NewWorkspacesServiceClient(httpClient connect_go.HTTPClient, baseURL string, opts ...connect_go.ClientOption) WorkspacesServiceClient {
	baseURL = strings.TrimRight(baseURL, "/")
	return &workspacesServiceClient{
		listWorkspaces: connect_go.NewClient[v1.ListWorkspacesRequest, v1.ListWorkspacesResponse](
			httpClient,
			baseURL+"/gitpod.experimental.v1.WorkspacesService/ListWorkspaces",
			opts...,
		),
		getWorkspace: connect_go.NewClient[v1.GetWorkspaceRequest, v1.GetWorkspaceResponse](
			httpClient,
			baseURL+"/gitpod.experimental.v1.WorkspacesService/GetWorkspace",
			opts...,
		),
		getOwnerToken: connect_go.NewClient[v1.GetOwnerTokenRequest, v1.GetOwnerTokenResponse](
			httpClient,
			baseURL+"/gitpod.experimental.v1.WorkspacesService/GetOwnerToken",
			opts...,
		),
		createAndStartWorkspace: connect_go.NewClient[v1.CreateAndStartWorkspaceRequest, v1.CreateAndStartWorkspaceResponse](
			httpClient,
			baseURL+"/gitpod.experimental.v1.WorkspacesService/CreateAndStartWorkspace",
			opts...,
		),
		stopWorkspace: connect_go.NewClient[v1.StopWorkspaceRequest, v1.StopWorkspaceResponse](
			httpClient,
			baseURL+"/gitpod.experimental.v1.WorkspacesService/StopWorkspace",
			opts...,
		),
		updatePort: connect_go.NewClient[v1.UpdatePortRequest, v1.UpdatePortResponse](
			httpClient,
			baseURL+"/gitpod.experimental.v1.WorkspacesService/UpdatePort",
			opts...,
		),
	}
}

// workspacesServiceClient implements WorkspacesServiceClient.
type workspacesServiceClient struct {
	listWorkspaces          *connect_go.Client[v1.ListWorkspacesRequest, v1.ListWorkspacesResponse]
	getWorkspace            *connect_go.Client[v1.GetWorkspaceRequest, v1.GetWorkspaceResponse]
	getOwnerToken           *connect_go.Client[v1.GetOwnerTokenRequest, v1.GetOwnerTokenResponse]
	createAndStartWorkspace *connect_go.Client[v1.CreateAndStartWorkspaceRequest, v1.CreateAndStartWorkspaceResponse]
	stopWorkspace           *connect_go.Client[v1.StopWorkspaceRequest, v1.StopWorkspaceResponse]
	updatePort              *connect_go.Client[v1.UpdatePortRequest, v1.UpdatePortResponse]
}

// ListWorkspaces calls gitpod.experimental.v1.WorkspacesService.ListWorkspaces.
func (c *workspacesServiceClient) ListWorkspaces(ctx context.Context, req *connect_go.Request[v1.ListWorkspacesRequest]) (*connect_go.Response[v1.ListWorkspacesResponse], error) {
	return c.listWorkspaces.CallUnary(ctx, req)
}

// GetWorkspace calls gitpod.experimental.v1.WorkspacesService.GetWorkspace.
func (c *workspacesServiceClient) GetWorkspace(ctx context.Context, req *connect_go.Request[v1.GetWorkspaceRequest]) (*connect_go.Response[v1.GetWorkspaceResponse], error) {
	return c.getWorkspace.CallUnary(ctx, req)
}

// GetOwnerToken calls gitpod.experimental.v1.WorkspacesService.GetOwnerToken.
func (c *workspacesServiceClient) GetOwnerToken(ctx context.Context, req *connect_go.Request[v1.GetOwnerTokenRequest]) (*connect_go.Response[v1.GetOwnerTokenResponse], error) {
	return c.getOwnerToken.CallUnary(ctx, req)
}

// CreateAndStartWorkspace calls gitpod.experimental.v1.WorkspacesService.CreateAndStartWorkspace.
func (c *workspacesServiceClient) CreateAndStartWorkspace(ctx context.Context, req *connect_go.Request[v1.CreateAndStartWorkspaceRequest]) (*connect_go.Response[v1.CreateAndStartWorkspaceResponse], error) {
	return c.createAndStartWorkspace.CallUnary(ctx, req)
}

// StopWorkspace calls gitpod.experimental.v1.WorkspacesService.StopWorkspace.
func (c *workspacesServiceClient) StopWorkspace(ctx context.Context, req *connect_go.Request[v1.StopWorkspaceRequest]) (*connect_go.ServerStreamForClient[v1.StopWorkspaceResponse], error) {
	return c.stopWorkspace.CallServerStream(ctx, req)
}

// UpdatePort calls gitpod.experimental.v1.WorkspacesService.UpdatePort.
func (c *workspacesServiceClient) UpdatePort(ctx context.Context, req *connect_go.Request[v1.UpdatePortRequest]) (*connect_go.Response[v1.UpdatePortResponse], error) {
	return c.updatePort.CallUnary(ctx, req)
}

// WorkspacesServiceHandler is an implementation of the gitpod.experimental.v1.WorkspacesService
// service.
type WorkspacesServiceHandler interface {
	// ListWorkspaces enumerates all workspaces belonging to the authenticated user.
	ListWorkspaces(context.Context, *connect_go.Request[v1.ListWorkspacesRequest]) (*connect_go.Response[v1.ListWorkspacesResponse], error)
	// GetWorkspace returns a single workspace.
	GetWorkspace(context.Context, *connect_go.Request[v1.GetWorkspaceRequest]) (*connect_go.Response[v1.GetWorkspaceResponse], error)
	// GetOwnerToken returns an owner token.
	GetOwnerToken(context.Context, *connect_go.Request[v1.GetOwnerTokenRequest]) (*connect_go.Response[v1.GetOwnerTokenResponse], error)
	// CreateAndStartWorkspace creates a new workspace and starts it.
	CreateAndStartWorkspace(context.Context, *connect_go.Request[v1.CreateAndStartWorkspaceRequest]) (*connect_go.Response[v1.CreateAndStartWorkspaceResponse], error)
	// StopWorkspace stops a running workspace (instance).
	// Errors:
	//
	//	NOT_FOUND:           the workspace_id is unkown
	//	FAILED_PRECONDITION: if there's no running instance
	StopWorkspace(context.Context, *connect_go.Request[v1.StopWorkspaceRequest], *connect_go.ServerStream[v1.StopWorkspaceResponse]) error
	UpdatePort(context.Context, *connect_go.Request[v1.UpdatePortRequest]) (*connect_go.Response[v1.UpdatePortResponse], error)
}

// NewWorkspacesServiceHandler builds an HTTP handler from the service implementation. It returns
// the path on which to mount the handler and the handler itself.
//
// By default, handlers support the Connect, gRPC, and gRPC-Web protocols with the binary Protobuf
// and JSON codecs. They also support gzip compression.
func NewWorkspacesServiceHandler(svc WorkspacesServiceHandler, opts ...connect_go.HandlerOption) (string, http.Handler) {
	mux := http.NewServeMux()
	mux.Handle("/gitpod.experimental.v1.WorkspacesService/ListWorkspaces", connect_go.NewUnaryHandler(
		"/gitpod.experimental.v1.WorkspacesService/ListWorkspaces",
		svc.ListWorkspaces,
		opts...,
	))
	mux.Handle("/gitpod.experimental.v1.WorkspacesService/GetWorkspace", connect_go.NewUnaryHandler(
		"/gitpod.experimental.v1.WorkspacesService/GetWorkspace",
		svc.GetWorkspace,
		opts...,
	))
	mux.Handle("/gitpod.experimental.v1.WorkspacesService/GetOwnerToken", connect_go.NewUnaryHandler(
		"/gitpod.experimental.v1.WorkspacesService/GetOwnerToken",
		svc.GetOwnerToken,
		opts...,
	))
	mux.Handle("/gitpod.experimental.v1.WorkspacesService/CreateAndStartWorkspace", connect_go.NewUnaryHandler(
		"/gitpod.experimental.v1.WorkspacesService/CreateAndStartWorkspace",
		svc.CreateAndStartWorkspace,
		opts...,
	))
	mux.Handle("/gitpod.experimental.v1.WorkspacesService/StopWorkspace", connect_go.NewServerStreamHandler(
		"/gitpod.experimental.v1.WorkspacesService/StopWorkspace",
		svc.StopWorkspace,
		opts...,
	))
	mux.Handle("/gitpod.experimental.v1.WorkspacesService/UpdatePort", connect_go.NewUnaryHandler(
		"/gitpod.experimental.v1.WorkspacesService/UpdatePort",
		svc.UpdatePort,
		opts...,
	))
	return "/gitpod.experimental.v1.WorkspacesService/", mux
}

// UnimplementedWorkspacesServiceHandler returns CodeUnimplemented from all methods.
type UnimplementedWorkspacesServiceHandler struct{}

func (UnimplementedWorkspacesServiceHandler) ListWorkspaces(context.Context, *connect_go.Request[v1.ListWorkspacesRequest]) (*connect_go.Response[v1.ListWorkspacesResponse], error) {
	return nil, connect_go.NewError(connect_go.CodeUnimplemented, errors.New("gitpod.experimental.v1.WorkspacesService.ListWorkspaces is not implemented"))
}

func (UnimplementedWorkspacesServiceHandler) GetWorkspace(context.Context, *connect_go.Request[v1.GetWorkspaceRequest]) (*connect_go.Response[v1.GetWorkspaceResponse], error) {
	return nil, connect_go.NewError(connect_go.CodeUnimplemented, errors.New("gitpod.experimental.v1.WorkspacesService.GetWorkspace is not implemented"))
}

func (UnimplementedWorkspacesServiceHandler) GetOwnerToken(context.Context, *connect_go.Request[v1.GetOwnerTokenRequest]) (*connect_go.Response[v1.GetOwnerTokenResponse], error) {
	return nil, connect_go.NewError(connect_go.CodeUnimplemented, errors.New("gitpod.experimental.v1.WorkspacesService.GetOwnerToken is not implemented"))
}

func (UnimplementedWorkspacesServiceHandler) CreateAndStartWorkspace(context.Context, *connect_go.Request[v1.CreateAndStartWorkspaceRequest]) (*connect_go.Response[v1.CreateAndStartWorkspaceResponse], error) {
	return nil, connect_go.NewError(connect_go.CodeUnimplemented, errors.New("gitpod.experimental.v1.WorkspacesService.CreateAndStartWorkspace is not implemented"))
}

func (UnimplementedWorkspacesServiceHandler) StopWorkspace(context.Context, *connect_go.Request[v1.StopWorkspaceRequest], *connect_go.ServerStream[v1.StopWorkspaceResponse]) error {
	return connect_go.NewError(connect_go.CodeUnimplemented, errors.New("gitpod.experimental.v1.WorkspacesService.StopWorkspace is not implemented"))
}

func (UnimplementedWorkspacesServiceHandler) UpdatePort(context.Context, *connect_go.Request[v1.UpdatePortRequest]) (*connect_go.Response[v1.UpdatePortResponse], error) {
	return nil, connect_go.NewError(connect_go.CodeUnimplemented, errors.New("gitpod.experimental.v1.WorkspacesService.UpdatePort is not implemented"))
}
