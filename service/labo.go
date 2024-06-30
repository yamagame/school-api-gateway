package service

import (
	"context"

	"github.com/yamagame/school-api-gateway/entity"
	"github.com/yamagame/school-api-gateway/infra/infconv"
	"github.com/yamagame/school-api-gateway/infra/model"
	"github.com/yamagame/school-api-gateway/infra/repository"
	"github.com/yamagame/school-api-gateway/pkg/conv"
	"github.com/yamagame/school-api-gateway/proto/school"
	"github.com/yamagame/school-api-gateway/service/svcconv"
)

type LaboInterface interface {
	Create(ctx context.Context) (int32, error)
	CreateWithMap(ctx context.Context, records []map[string]string) (int32, error)
	Find(ctx context.Context, id int32) (*school.Labo, error)
	Update(ctx context.Context, labo *school.Labo) (int32, error)
	Copy(ctx context.Context, id int32) (int32, error)
	List(ctx context.Context, limit, offset int32) ([]*school.Labo, error)
}

type Labo struct {
	repo repository.LaboInterface
}

func NewLabo(repo repository.LaboInterface) *Labo {
	return &Labo{
		repo: repo,
	}
}

func (s *Labo) Create(ctx context.Context) (int32, error) {
	labos := []*model.Labo{
		{},
	}
	if err := s.repo.Create(ctx, labos); err != nil {
		return 0, err
	}
	return labos[0].ID, nil
}

func (s *Labo) CreateWithMap(ctx context.Context, records []map[string]string) (int32, error) {
	zero := int32(0)
	labos := []*model.Labo{}
	for _, record := range records {
		labo, err := conv.NewRecordWithMap(record, entity.NewLabo)
		if err != nil {
			return zero, err
		}
		l, err := infconv.LaboToInfra(labo)
		if err != nil {
			return zero, err
		}
		labos = append(labos, l)
	}
	if err := s.repo.Create(ctx, labos); err != nil {
		return 0, err
	}
	return labos[0].ID, nil
}

func (s *Labo) Find(ctx context.Context, id int32) (*school.Labo, error) {
	var zero *school.Labo
	results, err := s.repo.Find(ctx, []int32{id})
	if err != nil {
		return zero, err
	}
	labos, err := laboToProto(results)
	if err != nil {
		return zero, err
	}
	if len(labos) > 0 {
		return labos[0], nil
	}
	return zero, ErrNotFound
}

func (s *Labo) Update(ctx context.Context, in *school.Labo) (int32, error) {
	zero := int32(0)
	labos, err := laboToInfra([]*school.Labo{in})
	if err != nil {
		return zero, err
	}
	if len(labos) > 0 {
		err = s.repo.Update(ctx, labos)
		if err != nil {
			return zero, err
		}
		return labos[0].ID, nil
	}
	return zero, ErrNotFound
}

func (s *Labo) Copy(ctx context.Context, id int32) (int32, error) {
	zero := int32(0)
	results, err := s.repo.Find(ctx, []int32{id})
	if err != nil {
		return zero, err
	}
	in, err := laboToProto(results)
	if err != nil {
		return zero, err
	}
	if len(in) > 0 {
		in[0].Id = 0
		labos, err := laboToInfra(in)
		if err != nil {
			return zero, err
		}
		if err := s.repo.Create(ctx, labos); err != nil {
			return 0, err
		}
		return labos[0].ID, nil
	}
	return zero, ErrNotFound
}

func (s *Labo) List(ctx context.Context, limit, offset int32) ([]*school.Labo, error) {
	results, err := s.repo.List(ctx, limit, offset)
	if err != nil {
		return nil, err
	}
	return laboToProto(results)
}

func laboToInfra(labos []*school.Labo) ([]*model.Labo, error) {
	res := []*model.Labo{}
	for _, labo := range labos {
		t, err := svcconv.LaboToEntity(labo)
		if err != nil {
			return nil, err
		}
		l, err := infconv.LaboToInfra(t)
		if err != nil {
			return nil, err
		}
		res = append(res, l)
	}
	return res, nil
}

func laboToProto(labos []*model.Labo) ([]*school.Labo, error) {
	res := []*school.Labo{}
	for _, labo := range labos {
		t, err := infconv.LaboToEntity(labo)
		if err != nil {
			return nil, err
		}
		l, err := svcconv.LaboToProto(t)
		if err != nil {
			return nil, err
		}
		res = append(res, l)
	}
	return res, nil
}
