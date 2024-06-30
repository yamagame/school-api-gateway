package service

import (
	"context"

	"github.com/yamagame/school-api-gateway/proto/school"
)

// School サービスの構造体
type Server struct {
	LaboService LaboInterface
}

// ListLabos 研究室の一覧を返す
func (r *Server) ListLabos(ctx context.Context, in *school.ListLabosRequest) (*school.ListLabosResponse, error) {
	pageSize := int32(5)
	if in.PageSize != nil {
		pageSize = *in.PageSize
	}
	offset := int32(0)
	if in.Offset != nil {
		offset = *in.Offset
	}
	labos, err := r.LaboService.List(ctx, pageSize, offset)
	if err != nil {
		return nil, err
	}
	return &school.ListLabosResponse{Labos: labos, Offset: pageSize + offset}, err
}
