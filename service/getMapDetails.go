package service

import (
	"FiletoDBMapper/models"
	"FiletoDBMapper/pb/pb"
	"context"
	"os"
	"strconv"

	"google.golang.org/grpc/codes"
	"google.golang.org/protobuf/types/known/emptypb"
)

func (m *MapperServer) GetMapDetails(context.Context, *emptypb.Empty) (*pb.GetMapDetailsResponse, error) {
	var MapDetails []models.MapDetail
	TenantID, _ := strconv.Atoi(os.Getenv("TENANT_ID"))
	var pbMapDetails []*pb.MapDetails
	type Result struct {
		MapDetailId uint
	}
	var results []Result
	var map_details_id []uint
	m.Db.Model(&models.TenantMasterMapping{}).Where("tenant_id = ?", TenantID).Find(&results)
	for _, result := range results {
		map_details_id = append(map_details_id, result.MapDetailId)
	}
	if err := m.Db.Where("id IN ?", map_details_id).Find(&MapDetails).Error; err != nil {
		return &pb.GetMapDetailsResponse{
			Status:  int32(codes.Internal),
			Message: "Cannot find Map Details By ID",
			Data:    []*pb.MapDetails{},
		}, err
	}
	//m.Db.Model(&models.MapDetail{}).Preload("Mappings").Find(&MapDetails)
	for _, MapDetail := range MapDetails {
		pbMapDetails = append(pbMapDetails, m.preloadModelIntoPb(&MapDetail))
	}

	return &pb.GetMapDetailsResponse{Status: int32(codes.OK), Message: "Get Map Details Successful!", Data: pbMapDetails}, nil
}
