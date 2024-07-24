package service

import (
	"context"
	"errors"
	"fmt"

	"github.com/yamagame/school-api-gateway/infra/infconv"
	"github.com/yamagame/school-api-gateway/infra/model"
	"github.com/yamagame/school-api-gateway/infra/repository"
	"github.com/yamagame/school-api-gateway/irmodel"
	"github.com/yamagame/school-api-gateway/proto/school"
	"github.com/yamagame/school-api-gateway/service/svcconv"
)

var (
	labosConv = svcconv.LabosConv{}
)

type LaboInterface interface {
	Create(ctx context.Context) (int32, error)
	CreateWithMap(ctx context.Context, records []map[string]interface{}) (int32, error)
	UpdateWithMap(ctx context.Context, records []map[string]interface{}) (int32, error)
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
	labos := infconv.Labo.NewSlice()
	labos.Append(&model.Labo{})
	if err := s.laborepo.Create(ctx, labos); err != nil {
		return 0, err
	}
	return labos.First().ID, nil
}

func (s *Labo) CreateWithMap(ctx context.Context, in []map[string]interface{}) (int32, error) {
	zero := int32(0)
	irmodels := irmodel.NewLabo().NewRecords(in)
	labos := infconv.Labos.ToStruct(irmodels)
	if labos.HasError() {
		return zero, labos.Error
	}
	if err := s.laborepo.Create(ctx, labos); err != nil {
		return zero, err
	}
	return labos.First().ID, nil
}

func (s *Labo) UpdateWithMap(ctx context.Context, in []map[string]interface{}) (int32, error) {
	zero := int32(0)
	irmodels := irmodel.NewLabo().NewRecords(in)
	labos := infconv.Labos.ToStruct(irmodels)
	if labos.HasError() {
		return zero, labos.Error
	}
	if err := s.laborepo.Update(ctx, labos); err != nil {
		return zero, err
	}
	return labos.First().ID, nil
}

func (s *Labo) Find(ctx context.Context, id int32) (*school.Labo, error) {
	var zero *school.Labo
	results := s.laborepo.Find(ctx, []int32{id})
	if results.HasError() {
		return zero, results.Error
	}
	labos := labosConv.ToProto(results)
	if labos.HasError() {
		return zero, labos.Error
	}
	if top := labos.First(); top != nil {
		return top, nil
	}
	return zero, errors.Join(ErrNotFound, fmt.Errorf("Find id: %d", id))
}

func (s *Labo) Update(ctx context.Context, in *school.Labo) (int32, error) {
	zero := int32(0)
	labos := labosConv.ToInfra(school.NewLabos(in))
	if labos.HasError() {
		return zero, labos.Error
	}
	if labos.Length() > 0 {
		err := s.laborepo.Update(ctx, labos)
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
	if results.HasError() {
		return zero, results.Error
	}
	in := labosConv.ToProto(results)
	if in.HasError() {
		return zero, in.Error
	}
	if top := in.First(); top != nil {
		top.Id = 0
		labos := labosConv.ToInfra(in)
		if labos.HasError() {
			return zero, labos.Error
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
	if results.HasError() {
		return nil, results.Error
	}
	r := labosConv.ToProto(results)
	if r.HasError() {
		return nil, r.Error
	}
	return r.ShallowCopy()
}
