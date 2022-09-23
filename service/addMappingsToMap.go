package service

import (
	"FiletoDBMapper/models"
	"FiletoDBMapper/pb/pb"
	"context"

	"google.golang.org/grpc/codes"
)

func (m *MapperServer) AddMappingsToMap(ctx context.Context, req *pb.AddMappingsToMapRequest) (*pb.AddMappingsToMapResponse, error) {
	var MappingId uint
	var MapDetail models.MapDetail
	for _, mapping := range req.Mappings {
		var FieldMappingId uint
		m.Db.Model(&models.FieldMapping{}).Select("id").Where(
			&models.FieldMapping{
				TemplateFieldName: mapping.TemplateFieldName,
				CustomerFieldName: mapping.CustomerFieldName,
			},
		).Find(&FieldMappingId)
		if FieldMappingId != 0 {
			m.Db.Model(&models.Mapping{}).Select("id").Where(
				&models.Mapping{
					MapDetailID:    uint(req.GetMapId()),
					FieldMappingID: FieldMappingId,
				},
			).Find(&MappingId)
			if MappingId == 0 {
				if err := m.Db.Create(
					&models.Mapping{
						FieldMappingID: FieldMappingId,
						MapDetailID:    uint(req.GetMapId()),
					},
				).Error; err != nil {
					return &pb.AddMappingsToMapResponse{
						Status:  int32(codes.Internal),
						Message: "cannot add mapping to the map",
						Data:    &pb.MapDetails{},
					}, err
				}
			}
		} else {
			if err := m.Db.Create(
				&models.FieldMapping{
					TemplateFieldName: mapping.TemplateFieldName,
					CustomerFieldName: mapping.CustomerFieldName,
					Mappings: []models.Mapping{
						{MapDetailID: uint(req.GetMapId())},
					},
				}).Error; err != nil {
				return &pb.AddMappingsToMapResponse{
					Status:  int32(codes.Internal),
					Message: "cannot create new mapping to the map",
					Data:    &pb.MapDetails{},
				}, err
			}
		}
	}
	if err := m.Db.First(&MapDetail, req.GetMapId()).Error; err != nil {
		return &pb.AddMappingsToMapResponse{
			Status:  int32(codes.Internal),
			Message: "cannot find map details by provided id",
			Data:    &pb.MapDetails{},
		}, err
	}
	return &pb.AddMappingsToMapResponse{
		Status:  int32(codes.OK),
		Message: "added mappings successfully",
		Data:    m.preloadModelIntoPb(&MapDetail),
	}, nil
}
