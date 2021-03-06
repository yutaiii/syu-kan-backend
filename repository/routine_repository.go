package repository

import (
	"context"

	"github.com/yutaiii/syu-kan-backend/domain/entity"
	"github.com/yutaiii/syu-kan-backend/domain/model"
	"github.com/yutaiii/syu-kan-backend/store"

	"gorm.io/gorm"
)

type RoutineRepository struct {
	ctx   context.Context
	store *store.RoutineStore
}

func NewRoutineRepository(ctx context.Context) *RoutineRepository {
	return &RoutineRepository{
		ctx:   ctx,
		store: store.NewRoutineStore(ctx),
	}
}

func (r *RoutineRepository) GetAllRoutines(db *gorm.DB) ([]*model.Routine, error) {
	routines, err := r.store.GetAll(db)
	if err != nil {
		return nil, err
	}
	models := model.NewRoutines(routines)
	return models, nil
}

func (r *RoutineRepository) GetByUserId(db *gorm.DB, m *model.RoutineForGetInput) ([]*model.Routine, error) {
	routines, err := r.store.GetByConditions(db, "user_id = ?", m.UserID)
	if err != nil {
		return nil, err
	}
	models := model.NewRoutines(routines)
	return models, nil
}

func (r *RoutineRepository) CreateRoutines(models []*model.Routine, db *gorm.DB) ([]*model.Routine, error) {
	entities := r.convertModelsToEntity(models)
	result, err := r.store.CreateRoutines(entities, db)
	if err != nil {
		return nil, err
	}
	m := model.NewRoutines(result)
	return m, nil
}

func (r *RoutineRepository) UpdateRoutines(db *gorm.DB, models []*model.Routine) ([]*model.Routine, error) {
	entities := r.convertModelsToEntityWithID(models)
	// create records if not exists
	result, err := r.store.Upsert(db, entities)
	if err != nil {
		return nil, err
	}
	m := model.NewRoutines(result)
	return m, nil
}

func (r *RoutineRepository) DeleteRoutines(db *gorm.DB, models []*model.Routine) error {
	entities := r.convertModelsToEntityWithID(models)
	return r.store.Delete(db, entities)
}

func (r *RoutineRepository) convertModelsToEntity(models []*model.Routine) []*entity.Routine {
	var entities []*entity.Routine
	for _, m := range models {
		e := r.convertModelToEntity(m)
		entities = append(entities, e)
	}
	return entities
}

func (r *RoutineRepository) convertModelToEntity(m *model.Routine) *entity.Routine {
	return &entity.Routine{
		Name:      m.Name,
		UserID:    m.UserID,
		StartedAt: m.StartedAt,
	}
}

func (r *RoutineRepository) convertModelsToEntityWithID(models []*model.Routine) []*entity.Routine {
	var entities []*entity.Routine
	for _, m := range models {
		e := r.convertUpdateModelToEntity(m)
		entities = append(entities, e)
	}
	return entities
}

func (r *RoutineRepository) convertUpdateModelToEntity(m *model.Routine) *entity.Routine {
	return &entity.Routine{
		BaseEntity: entity.BaseEntity{
			ID: m.ID,
		},
		Name:      m.Name,
		UserID:    m.UserID,
		StartedAt: m.StartedAt,
	}
}
