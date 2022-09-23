package service

import (
	"FiletoDBMapper/models"
	"FiletoDBMapper/pb/pb"
	"context"

	"google.golang.org/grpc/codes"
)

func (m *MapperServer) UpdateMapDetails(ctx context.Context, req *pb.UpdateMapDetailsRequest) (*pb.UpdateMapDetailsResponse, error) {
	var MapDetail models.MapDetail
	if err:= m.Db.First(&MapDetail, req.GetId()).Error; err != nil{
		return &pb.UpdateMapDetailsResponse{Status: int32(codes.Internal), Message: "Internal Server Error", Data: &pb.MapDetails{}}, err
	}
	if req.GetCategoryName() == MapDetail.CategoryName && req.GetMapDescription() == MapDetail.MapDescription && req.GetMapName() == req.MapName{
		return &pb.UpdateMapDetailsResponse{Status: int32(codes.AlreadyExists), Message: "No details to update", Data: m.preloadModelIntoPb(&MapDetail)}, nil
	}
	if req.GetCategoryName() != ""{
		MapDetail.CategoryName = req.CategoryName
	}
	if req.GetMapDescription() != ""{
		MapDetail.MapDescription = req.MapDescription
	}
	if req.GetMapName() != ""{
		MapDetail.MapName = req.MapName
	}
	if err:= m.Db.Save(&MapDetail).Error; err != nil{
		return &pb.UpdateMapDetailsResponse{Status: int32(codes.Internal), Message: "Internal Server Error", Data: &pb.MapDetails{}}, err
	}
	return &pb.UpdateMapDetailsResponse{Status: int32(codes.OK), Message: "Updated successfully", Data: m.preloadModelIntoPb(&MapDetail)}, nil
}

