// Code generated by protoc-gen-go-http. DO NOT EDIT.
// versions:
// - protoc-gen-go-http v2.7.3
// - protoc             v3.21.6
// source: ccnu_service/v1/ccnu_service.proto

package v1

import (
	context "context"
	http "github.com/go-kratos/kratos/v2/transport/http"
	binding "github.com/go-kratos/kratos/v2/transport/http/binding"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the kratos package it is being compiled against.
var _ = new(context.Context)
var _ = binding.EncodeURL

const _ = http.SupportPackageIsVersion1

const OperationCCNUServiceGetCookie = "/api.ccnu_service.v1.CCNUService/GetCookie"
const OperationCCNUServiceSaveUser = "/api.ccnu_service.v1.CCNUService/SaveUser"

type CCNUServiceHTTPServer interface {
	GetCookie(context.Context, *GetCookieRequest) (*GetCookieResponse, error)
	SaveUser(context.Context, *SaveUserRequest) (*SaveUserResponse, error)
}

func RegisterCCNUServiceHTTPServer(s *http.Server, srv CCNUServiceHTTPServer) {
	r := s.Route("/")
	r.POST("/v1/save_user", _CCNUService_SaveUser0_HTTP_Handler(srv))
	r.GET("/v1/login/{userid}", _CCNUService_GetCookie0_HTTP_Handler(srv))
}

func _CCNUService_SaveUser0_HTTP_Handler(srv CCNUServiceHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in SaveUserRequest
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationCCNUServiceSaveUser)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.SaveUser(ctx, req.(*SaveUserRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*SaveUserResponse)
		return ctx.Result(200, reply)
	}
}

func _CCNUService_GetCookie0_HTTP_Handler(srv CCNUServiceHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in GetCookieRequest
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		if err := ctx.BindVars(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationCCNUServiceGetCookie)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.GetCookie(ctx, req.(*GetCookieRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*GetCookieResponse)
		return ctx.Result(200, reply)
	}
}

type CCNUServiceHTTPClient interface {
	GetCookie(ctx context.Context, req *GetCookieRequest, opts ...http.CallOption) (rsp *GetCookieResponse, err error)
	SaveUser(ctx context.Context, req *SaveUserRequest, opts ...http.CallOption) (rsp *SaveUserResponse, err error)
}

type CCNUServiceHTTPClientImpl struct {
	cc *http.Client
}

func NewCCNUServiceHTTPClient(client *http.Client) CCNUServiceHTTPClient {
	return &CCNUServiceHTTPClientImpl{client}
}

func (c *CCNUServiceHTTPClientImpl) GetCookie(ctx context.Context, in *GetCookieRequest, opts ...http.CallOption) (*GetCookieResponse, error) {
	var out GetCookieResponse
	pattern := "/v1/login/{userid}"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation(OperationCCNUServiceGetCookie))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "GET", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, nil
}

func (c *CCNUServiceHTTPClientImpl) SaveUser(ctx context.Context, in *SaveUserRequest, opts ...http.CallOption) (*SaveUserResponse, error) {
	var out SaveUserResponse
	pattern := "/v1/save_user"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation(OperationCCNUServiceSaveUser))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, nil
}
