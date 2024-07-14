package repository

import (
	"context"

	"github.com/yamagame/school-api-gateway/infra/dao/query"
	"github.com/yamagame/school-api-gateway/infra/infconv"
	"github.com/yamagame/school-api-gateway/infra/model"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type LaboInterface interface {
	UpsertInBatches(ctx context.Context, labos *model.Labos, columns []string) error
	Upsert(ctx context.Context, labos *model.Labos) error
	Create(ctx context.Context, labos *model.Labos) error
	Update(ctx context.Context, labos *model.Labos) error
	Find(ctx context.Context, ids []int32) *model.Labos
	List(ctx context.Context, limit, offset int32) *model.Labos
}

type Labo struct {
	db        *gorm.DB
	batchSize int
}

func NewLabo(db *gorm.DB) *Labo {
	return &Labo{
		db:        db,
		batchSize: 100,
	}
}

func (r *Labo) UpsertInBatches(ctx context.Context, labos *model.Labos, columns []string) error {
	q := query.Use(r.db)
	lb := q.Labo
	return lb.WithContext(ctx).
		Clauses(clause.OnConflict{
			DoUpdates: clause.AssignmentColumns(columns),
		}).
		CreateInBatches(labos.ShallowCopy(), r.batchSize)
}

func (r *Labo) Upsert(ctx context.Context, labos *model.Labos) error {
	var err error
	creates := infconv.Labo.NewSlice()
	updates := infconv.Labo.NewSlice()
	it := labos.NewIterator()
	for it.HasNext() {
		labo := it.Next()
		if labo.ID == 0 {
			creates.Append(labo)
		} else {
			updates.Append(labo)
		}
	}
	if err = r.Update(ctx, updates); err != nil {
		return err
	}
	if err = r.Create(ctx, creates); err != nil {
		return err
	}
	return nil
}

func (r *Labo) Create(ctx context.Context, labos *model.Labos) error {
	q := query.Use(r.db)
	lb := q.Labo
	return lb.WithContext(ctx).CreateInBatches(labos.ShallowCopy(), r.batchSize)
}

func (r *Labo) Update(ctx context.Context, labos *model.Labos) error {
	q := query.Use(r.db)
	lb := q.Labo
	it := labos.NewIterator()
	for it.HasNext() {
		labo := it.Next()
		_, err := lb.WithContext(ctx).Where(lb.ID.Eq(labo.ID)).Updates(labo)
		if err != nil {
			return err
		}
	}
	return nil
}

func (r *Labo) Find(ctx context.Context, ids []int32) *model.Labos {
	q := query.Use(r.db)
	lb := q.Labo
	records, err := lb.WithContext(ctx).Where(lb.ID.In(ids...)).Find()
	ret := infconv.Labo.NewSlice()
	ret.Append(records...)
	if err != nil {
		ret.Error = err
	}
	return ret
}

func (r *Labo) List(ctx context.Context, limit, offset int32) *model.Labos {
	q := query.Use(r.db)
	lb := q.Labo
	records, err := lb.WithContext(ctx).
		Joins(lb.Building, lb.Group, lb.Program).
		Limit(int(limit)).
		Offset(int(offset)).
		Order(lb.ID.Asc()).
		Find()
	ret := infconv.Labo.NewSlice()
	ret.Append(records...)
	if err != nil {
		ret.Error = err
	}
	return ret
}
