package service

import (
	"FiletoDBMapper/models"
	"FiletoDBMapper/pb/pb"
	"context"
	"fmt"

	"google.golang.org/grpc/codes"
)

func (m *MapperServer) RemoveMappingFromMap (ctx context.Context, req *pb.RemoveMappingFromMapRequest) (*pb.RemoveMappingFromMapResponse, error){
	var FieldMappingId uint
	if err := m.Db.Model(&models.FieldMapping{}).Select("id").Where(&models.FieldMapping{CustomerFieldName: req.GetMapping().GetCustomerFieldName(), TemplateFieldName: req.GetMapping().GetTemplateFieldName()}).Find(&FieldMappingId).Error; err!=nil{
		return &pb.RemoveMappingFromMapResponse{Status: int32(codes.Internal), Message: "cannot perform find query for field mappings", Data: &pb.MapDetails{}}, err
	}
	if FieldMappingId != 0 {
		var MappingId uint
		if err:= m.Db.Model(&models.Mapping{}).Select("id").Where(&models.Mapping{MapDetailID: uint(req.GetMapId()), FieldMappingID: FieldMappingId}).Find(&MappingId).Error; err!=nil{
			return &pb.RemoveMappingFromMapResponse{Status: int32(codes.Internal), Message: "cannot perform find query for mappings", Data: &pb.MapDetails{}}, err
		}
		if MappingId == 0{
			return &pb.RemoveMappingFromMapResponse{Status: int32(codes.NotFound), Message: fmt.Sprintf("the mapping %v <--> %v has not been added yet to the map id %v", req.GetMapping().GetTemplateFieldName(), req.GetMapping().GetCustomerFieldName(), req.GetMapId())}, nil
		} else {
			if err := m.Db.Delete(&models.Mapping{ID: MappingId}).Error; err!=nil{
				return &pb.RemoveMappingFromMapResponse{Status: int32(codes.Internal), Message: "cannot delete record", Data: &pb.MapDetails{}}, err
			}
			var MapDetail models.MapDetail
			if err := m.Db.First(&MapDetail, req.GetMapId()).Error; err != nil {
				return &pb.RemoveMappingFromMapResponse{
					Status:  int32(codes.Internal),
					Message: "cannot find map details by provided id",
					Data:    &pb.MapDetails{},
				}, err
			}
			return &pb.RemoveMappingFromMapResponse{
				Status:  int32(codes.OK),
				Message: "removed mappings successfully",
				Data:    m.preloadModelIntoPb(&MapDetail),
			}, nil
		}
	} else {
		return &pb.RemoveMappingFromMapResponse{Status: int32(codes.NotFound), Message: "mapping does not exist yet", Data: &pb.MapDetails{}}, nil
	}
}
