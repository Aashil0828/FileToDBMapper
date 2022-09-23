package service

import (
	"FiletoDBMapper/models"
	"FiletoDBMapper/pb/pb"
	"context"
	"fmt"
	"os"
	"strconv"

	"google.golang.org/grpc/codes"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type MapperServer struct {
	Db *gorm.DB
	pb.UnimplementedFileToDBMapperServer
}

func (m *MapperServer) CreateMap(ctx context.Context, req *pb.CreateMapRequest) (*pb.CreateMapResponse, error) {
	var MapName string
	TenantID, _ := strconv.Atoi(os.Getenv("TENANT_ID"))
	m.Db.Table("tenant_master_mappings").Select("map_details.map_name").Joins("join map_details on map_details.id = tenant_master_mappings.map_detail_id").Where("tenant_id = ? AND master_template_id = ?", TenantID, req.TemplateId).Find(&MapName)
	if MapName != req.MapName {
		MapDetail := models.MapDetail{
			MapName:              req.GetMapName(),
			MapDescription:       req.GetMapDescription(),
			CategoryName:         req.GetCategoryName(),
			CreatedBy:            "aashil",
			MapStatus:            true,
			TenantMasterMappings: []models.TenantMasterMapping{{TenantID: uint(TenantID), MasterTemplateID: uint(req.GetTemplateId())}},
		}
		if err := m.Db.Clauses(clause.OnConflict{DoNothing: true}).Save(&MapDetail).Error; err != nil {
			return &pb.CreateMapResponse{Status: int32(codes.Internal), Message: "Internal Server Error"}, nil
		}
		MapDetails := &pb.MapDetails{
			Id:             uint32(MapDetail.ID),
			CreatedAt:      MapDetail.CreatedAt.String(),
			UpdatedAt:      MapDetail.UpdatedAt.String(),
			DeletedAt:      MapDetail.DeletedAt.Time.String(),
			MapName:        MapDetail.MapName,
			MapDescription: MapDetail.MapDescription,
			CategoryName:   MapDetail.CategoryName,
			MapStatus:      MapDetail.MapStatus,
			CreatedBy:      MapDetail.CreatedBy,
			UpdatedBy:      MapDetail.UpdatedBy,
		}
		return &pb.CreateMapResponse{Status: int32(codes.OK), Message: "Map Creation Successful", Data: MapDetails}, nil
	} else {
		return &pb.CreateMapResponse{Status: int32(codes.AlreadyExists), Message: fmt.Sprintf("Map with name %v already exists!", MapName)}, nil
	}
}
