package service

import (
	"FiletoDBMapper/pb/pb"
	"context"

	"google.golang.org/grpc/codes"
)

func (m *MapperServer) GetPreFilledMappings(ctx context.Context, req *pb.GetPreFilledMappingsRequest) (*pb.GetPreFilledMappingsResponse, error) {
	return &pb.GetPreFilledMappingsResponse{Status: int32(codes.OK), Message: "get prefilled mappings successful!"}, nil
}
