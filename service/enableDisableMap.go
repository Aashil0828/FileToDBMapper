package service

import (
	"FiletoDBMapper/models"
	"FiletoDBMapper/pb/pb"
	"context"

	"google.golang.org/grpc/codes"
)

func (m *MapperServer) EnableDisableMap(ctx context.Context, req *pb.EnableDisableMapRequest) (*pb.EnableDisableMapResponse, error) {
	var MapDetail models.MapDetail
	var message string
	if err := m.Db.Unscoped().First(&MapDetail, req.GetId()).Error; err != nil {
		return &pb.EnableDisableMapResponse{Status: int32(codes.Internal), Message: "Cannot perform query", Data: &pb.MapDetails{}}, err
	}
	if !MapDetail.DeletedAt.Valid {
		MapDetail.MapStatus = false
		if err := m.Db.Delete(&MapDetail).Error; err != nil {
			return &pb.EnableDisableMapResponse{Status: int32(codes.Internal), Message: "Cannot perform delete", Data: &pb.MapDetails{}}, err
		}
		message = "Disabled Map Successfully"
	} else {
		MapDetail.MapStatus = true
		m.Db.Unscoped().Model(&MapDetail).Update("deleted_at", nil)
		message = "Enabled Map Successfully"
	}
	return &pb.EnableDisableMapResponse{Status: int32(codes.OK), Message: message, Data: m.preloadModelIntoPb(&MapDetail)}, nil
}
