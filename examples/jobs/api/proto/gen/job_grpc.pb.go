// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             (unknown)
// source: job.proto

package gen

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

const (
	Jobs_ListCompanies_FullMethodName = "/Jobs/ListCompanies"
	Jobs_CreateCompany_FullMethodName = "/Jobs/CreateCompany"
	Jobs_DeleteCompany_FullMethodName = "/Jobs/DeleteCompany"
	Jobs_ListJobs_FullMethodName      = "/Jobs/ListJobs"
	Jobs_CreateJob_FullMethodName     = "/Jobs/CreateJob"
	Jobs_GetJob_FullMethodName        = "/Jobs/GetJob"
	Jobs_DeleteJob_FullMethodName     = "/Jobs/DeleteJob"
)

// JobsClient is the client API for Jobs service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type JobsClient interface {
	// List the companies
	ListCompanies(ctx context.Context, in *ListCompaniesRequest, opts ...grpc.CallOption) (*ListCompaniesResponse, error)
	// Creates a new company.
	CreateCompany(ctx context.Context, in *CreateCompanyRequest, opts ...grpc.CallOption) (*Company, error)
	// Deletes a company, including all jobs that are in the company.
	DeleteCompany(ctx context.Context, in *DeleteCompanyRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
	// Return a list of all jobs of a company
	ListJobs(ctx context.Context, in *ListJobsRequest, opts ...grpc.CallOption) (*ListJobsResponse, error)
	// Creates a new job of the company.
	CreateJob(ctx context.Context, in *CreateJobRequest, opts ...grpc.CallOption) (*Job, error)
	// Returns a specific job.
	GetJob(ctx context.Context, in *GetJobRequest, opts ...grpc.CallOption) (*Job, error)
	// Deletes a job from a company.
	DeleteJob(ctx context.Context, in *DeleteJobRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
}

type jobsClient struct {
	cc grpc.ClientConnInterface
}

func NewJobsClient(cc grpc.ClientConnInterface) JobsClient {
	return &jobsClient{cc}
}

func (c *jobsClient) ListCompanies(ctx context.Context, in *ListCompaniesRequest, opts ...grpc.CallOption) (*ListCompaniesResponse, error) {
	out := new(ListCompaniesResponse)
	err := c.cc.Invoke(ctx, Jobs_ListCompanies_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *jobsClient) CreateCompany(ctx context.Context, in *CreateCompanyRequest, opts ...grpc.CallOption) (*Company, error) {
	out := new(Company)
	err := c.cc.Invoke(ctx, Jobs_CreateCompany_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *jobsClient) DeleteCompany(ctx context.Context, in *DeleteCompanyRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, Jobs_DeleteCompany_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *jobsClient) ListJobs(ctx context.Context, in *ListJobsRequest, opts ...grpc.CallOption) (*ListJobsResponse, error) {
	out := new(ListJobsResponse)
	err := c.cc.Invoke(ctx, Jobs_ListJobs_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *jobsClient) CreateJob(ctx context.Context, in *CreateJobRequest, opts ...grpc.CallOption) (*Job, error) {
	out := new(Job)
	err := c.cc.Invoke(ctx, Jobs_CreateJob_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *jobsClient) GetJob(ctx context.Context, in *GetJobRequest, opts ...grpc.CallOption) (*Job, error) {
	out := new(Job)
	err := c.cc.Invoke(ctx, Jobs_GetJob_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *jobsClient) DeleteJob(ctx context.Context, in *DeleteJobRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, Jobs_DeleteJob_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// JobsServer is the server API for Jobs service.
// All implementations must embed UnimplementedJobsServer
// for forward compatibility
type JobsServer interface {
	// List the companies
	ListCompanies(context.Context, *ListCompaniesRequest) (*ListCompaniesResponse, error)
	// Creates a new company.
	CreateCompany(context.Context, *CreateCompanyRequest) (*Company, error)
	// Deletes a company, including all jobs that are in the company.
	DeleteCompany(context.Context, *DeleteCompanyRequest) (*emptypb.Empty, error)
	// Return a list of all jobs of a company
	ListJobs(context.Context, *ListJobsRequest) (*ListJobsResponse, error)
	// Creates a new job of the company.
	CreateJob(context.Context, *CreateJobRequest) (*Job, error)
	// Returns a specific job.
	GetJob(context.Context, *GetJobRequest) (*Job, error)
	// Deletes a job from a company.
	DeleteJob(context.Context, *DeleteJobRequest) (*emptypb.Empty, error)
	mustEmbedUnimplementedJobsServer()
}

// UnimplementedJobsServer must be embedded to have forward compatible implementations.
type UnimplementedJobsServer struct {
}

func (UnimplementedJobsServer) ListCompanies(context.Context, *ListCompaniesRequest) (*ListCompaniesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListCompanies not implemented")
}
func (UnimplementedJobsServer) CreateCompany(context.Context, *CreateCompanyRequest) (*Company, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateCompany not implemented")
}
func (UnimplementedJobsServer) DeleteCompany(context.Context, *DeleteCompanyRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteCompany not implemented")
}
func (UnimplementedJobsServer) ListJobs(context.Context, *ListJobsRequest) (*ListJobsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListJobs not implemented")
}
func (UnimplementedJobsServer) CreateJob(context.Context, *CreateJobRequest) (*Job, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateJob not implemented")
}
func (UnimplementedJobsServer) GetJob(context.Context, *GetJobRequest) (*Job, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetJob not implemented")
}
func (UnimplementedJobsServer) DeleteJob(context.Context, *DeleteJobRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteJob not implemented")
}
func (UnimplementedJobsServer) mustEmbedUnimplementedJobsServer() {}

// UnsafeJobsServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to JobsServer will
// result in compilation errors.
type UnsafeJobsServer interface {
	mustEmbedUnimplementedJobsServer()
}

func RegisterJobsServer(s grpc.ServiceRegistrar, srv JobsServer) {
	s.RegisterService(&Jobs_ServiceDesc, srv)
}

func _Jobs_ListCompanies_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListCompaniesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(JobsServer).ListCompanies(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Jobs_ListCompanies_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(JobsServer).ListCompanies(ctx, req.(*ListCompaniesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Jobs_CreateCompany_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateCompanyRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(JobsServer).CreateCompany(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Jobs_CreateCompany_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(JobsServer).CreateCompany(ctx, req.(*CreateCompanyRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Jobs_DeleteCompany_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteCompanyRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(JobsServer).DeleteCompany(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Jobs_DeleteCompany_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(JobsServer).DeleteCompany(ctx, req.(*DeleteCompanyRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Jobs_ListJobs_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListJobsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(JobsServer).ListJobs(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Jobs_ListJobs_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(JobsServer).ListJobs(ctx, req.(*ListJobsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Jobs_CreateJob_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateJobRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(JobsServer).CreateJob(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Jobs_CreateJob_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(JobsServer).CreateJob(ctx, req.(*CreateJobRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Jobs_GetJob_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetJobRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(JobsServer).GetJob(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Jobs_GetJob_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(JobsServer).GetJob(ctx, req.(*GetJobRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Jobs_DeleteJob_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteJobRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(JobsServer).DeleteJob(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Jobs_DeleteJob_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(JobsServer).DeleteJob(ctx, req.(*DeleteJobRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Jobs_ServiceDesc is the grpc.ServiceDesc for Jobs service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Jobs_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "Jobs",
	HandlerType: (*JobsServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ListCompanies",
			Handler:    _Jobs_ListCompanies_Handler,
		},
		{
			MethodName: "CreateCompany",
			Handler:    _Jobs_CreateCompany_Handler,
		},
		{
			MethodName: "DeleteCompany",
			Handler:    _Jobs_DeleteCompany_Handler,
		},
		{
			MethodName: "ListJobs",
			Handler:    _Jobs_ListJobs_Handler,
		},
		{
			MethodName: "CreateJob",
			Handler:    _Jobs_CreateJob_Handler,
		},
		{
			MethodName: "GetJob",
			Handler:    _Jobs_GetJob_Handler,
		},
		{
			MethodName: "DeleteJob",
			Handler:    _Jobs_DeleteJob_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "job.proto",
}
