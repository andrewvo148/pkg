package grpc

import (
	"context"

	"github.com/andrewvo148/pkg/examples/jobs/api/gen"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
)

type server struct {
	gen.UnimplementedLibraryServer
}

var _ gen.JobsServer = (*server)(nil)

func RegisterServer(s grpc.ServiceRegistrar) error {
	gen.RegisterJobsServer(s, server{})
	return nil
}

func (s server) ListCompanies(ctx context.Context, request *gen.ListCompaniesRequest) (resp *gen.ListCompaniesResponse, err error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListCompanies not implemented")
}

func (s server) CreateCompany(ctx context.Context, request *gen.CreateCompanyRequest) (resp *gen.Company, err error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateCompany not implemented")
}

func (s server) DeleteCompany(ctx context.Context, request *gen.DeleteCompanyRequest) (*emptypb.Empty, err error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteCompany not implemented")
}

func (s server) ListJobs(ctx context.Context, request *gen.ListJobsRequest) (resp *gen.ListJobsResponse, err error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListJobs not implemented")
}

func (s server) CreateJob(ctx context.Context, request *gen.CreateJobRequest) (resp *gen.Job, err error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateJob not implemented")
}

func (s server) GetJob(ctx context.Context, request *gen.GetJobRequest) (resp *gen.Job, err error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetJob not implemented")
}

func (s server) DeleteJob(ctx context.Context, request *gen.DeleteJobRequest) (emptypb.Empty, err error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteJob not implemented")
}
