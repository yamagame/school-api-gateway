package service

import (
	"context"

	"github.com/yamagame/school-api-gateway/entity"
	"github.com/yamagame/school-api-gateway/infra/infconv"
	"github.com/yamagame/school-api-gateway/infra/model"
	"github.com/yamagame/school-api-gateway/infra/repository"
	"github.com/yamagame/school-api-gateway/proto/school"
	"github.com/yamagame/school-api-gateway/service/svcconv"
)

type LaboInterface interface {
	Create(ctx context.Context) (int32, error)
	CreateWithMap(ctx context.Context, records []map[string]string) (int32, error)
	UpdateWithMap(ctx context.Context, records []map[string]string) (int32, error)
	Find(ctx context.Context, id int32) (*school.Labo, error)
	Update(ctx context.Context, labo *school.Labo) (int32, error)
	Copy(ctx context.Context, id int32) (int32, error)
	List(ctx context.Context, limit, offset int32) ([]*school.Labo, error)
}

type Labo struct {
	laborepo repository.LaboInterface
}

func NewLabo(repo repository.LaboInterface) *Labo {
	return &Labo{
		laborepo: repo,
	}
}

func (s *Labo) Create(ctx context.Context) (int32, error) {
	labos := model.NewLabos(&model.Labo{})
	if err := s.laborepo.Create(ctx, labos); err != nil {
		return 0, err
	}
	return labos.First().ID, nil
}

func (s *Labo) CreateWithMap(ctx context.Context, records []map[string]string) (int32, error) {
	zero := int32(0)
	labos := &model.Labos{}
	err := infconv.Labos.ToInfraWithMap(records, labos, entity.NewLabo)
	if err != nil {
		return zero, err
	}
	if err := s.laborepo.Create(ctx, labos); err != nil {
		return 0, err
	}
	return labos.First().ID, nil
}

func (s *Labo) UpdateWithMap(ctx context.Context, records []map[string]string) (int32, error) {
	zero := int32(0)
	labos := &model.Labos{}
	err := infconv.Labos.ToInfraWithMap(records, labos, entity.NewLabo)
	if err != nil {
		return zero, err
	}
	if err := s.laborepo.Update(ctx, labos); err != nil {
		return 0, err
	}
	return labos.First().ID, nil
}

func (s *Labo) Find(ctx context.Context, id int32) (*school.Labo, error) {
	var zero *school.Labo
	results := s.laborepo.Find(ctx, []int32{id})
	if results.Error != nil {
		return zero, results.Error
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
	if len(labos.Records) > 0 {
		err = s.laborepo.Update(ctx, labos)
		if err != nil {
			return zero, err
		}
		return labos.First().ID, nil
	}
	return zero, ErrNotFound
}

func (s *Labo) Copy(ctx context.Context, id int32) (int32, error) {
	zero := int32(0)
	results := s.laborepo.Find(ctx, []int32{id})
	if results.Error != nil {
		return zero, results.Error
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
		if err := s.laborepo.Create(ctx, labos); err != nil {
			return 0, err
		}
		return labos.First().ID, nil
	}
	return zero, ErrNotFound
}

func (s *Labo) List(ctx context.Context, limit, offset int32) ([]*school.Labo, error) {
	results := s.laborepo.List(ctx, limit, offset)
	if results.Error != nil {
		return nil, results.Error
	}
	return laboToProto(results)
}

func laboToInfra(labos []*school.Labo) (*model.Labos, error) {
	res := []*model.Labo{}
	for _, labo := range labos {
		t, err := svcconv.Labo.ToEntity(labo)
		if err != nil {
			return nil, err
		}
		l, err := infconv.Labo.ToInfra(t)
		if err != nil {
			return nil, err
		}
		res = append(res, l)
	}
	return model.NewLabos(res...), nil
}

func laboToProto(labos *model.Labos) ([]*school.Labo, error) {
	res := []*school.Labo{}
	it := labos.NewIterator()
	for it.HasNext() {
		labo := it.Next()
		t, err := infconv.Labo.ToEntity(labo)
		if err != nil {
			return nil, err
		}
		l, err := svcconv.Labo.ToProto(t)
		if err != nil {
			return nil, err
		}
		res = append(res, l)
	}
	return res, nil
}
