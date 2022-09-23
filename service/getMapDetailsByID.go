package service

import (
	"FiletoDBMapper/models"
	"FiletoDBMapper/pb/pb"
	"context"
	"fmt"
	"os"
	"strconv"

	"google.golang.org/grpc/codes"
)

func (m *MapperServer) GetMapDetailsByID(ctx context.Context, req *pb.GetMapDetailsByIDRequest) (*pb.GetMapDetailsByIdResponse, error) {
	MapDetail := models.MapDetail{}
	TenantID, _ := strconv.Atoi(os.Getenv("TENANT_ID"))
	var tenant_id uint
	m.Db.Model(&models.TenantMasterMapping{}).Select("tenant_id").Where("map_detail_id = ?", req.GetId()).Find(&tenant_id)
	fmt.Println(tenant_id)
	if TenantID != int(tenant_id){
		return &pb.GetMapDetailsByIdResponse{Status: int32(codes.NotFound), Message: "record not found"}, nil
	}
	if err := m.Db.First(&MapDetail, req.GetId()).Error; err != nil {
		return &pb.GetMapDetailsByIdResponse{
			Status:  int32(codes.Internal),
			Message: "Cannot find Map Details By ID",
			Data:    &pb.MapDetails{},
		}, err
	}
	return &pb.GetMapDetailsByIdResponse{Status: int32(codes.OK), Message: "get map details by id successful", Data: m.preloadModelIntoPb(&MapDetail)}, nil
}

func (m *MapperServer) preloadModelIntoPb(MapDetail *models.MapDetail) *pb.MapDetails {
	Pbmapping := []*pb.Mapping{}
	m.Db.Table("mappings").Select("field_mappings.template_field_name, field_mappings.customer_field_name").Joins("join field_mappings on field_mappings.id = mappings.field_mapping_id").Where("mappings.map_detail_id = ? AND mappings.deleted_at IS NULL", MapDetail.ID).Scan(&Pbmapping)
	return &pb.MapDetails{
		Id:             uint32(MapDetail.ID),
		CreatedAt:      MapDetail.CreatedAt.String(),
		UpdatedAt:      MapDetail.UpdatedAt.String(),
		DeletedAt:      MapDetail.DeletedAt.Time.String(),
		MapName:        MapDetail.MapName,
		MapDescription: MapDetail.MapDescription,
		MapStatus:      MapDetail.MapStatus,
		Mappings:       Pbmapping,
		CategoryName:   MapDetail.CategoryName,
		CreatedBy:      MapDetail.CreatedBy,
		UpdatedBy:      MapDetail.UpdatedBy,
	}
}
