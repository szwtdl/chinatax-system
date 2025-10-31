package service

import (
	"github.com/szwtdl/chinatax-system/internal/repository"
	"gorm.io/gorm"
)

type BaseService[T any] struct {
	repo repository.BaseRepository[T]
	db   *gorm.DB
}

func (s *BaseService[T]) Db() *gorm.DB {
	return s.db
}

func (s *BaseService[T]) GetByID(id uint) (*T, error) {
	return s.repo.FindOne(s.db, map[string]interface{}{"id": id})
}

func (s *BaseService[T]) Find(where map[string]interface{}) ([]T, error) {
	return s.repo.Find(s.db, where)
}

func (s *BaseService[T]) FindOne(where map[string]interface{}) (*T, error) {
	return s.repo.FindOne(s.db, where)
}

func (s *BaseService[T]) Pagination(where map[string]interface{}, page, pageSize int, order string) ([]T, int64, error) {
	return s.repo.FindWithPagination(s.db, where, page, pageSize, order)
}

func (s *BaseService[T]) Create(entity *T) error {
	return s.repo.Create(s.db, entity)
}

func (s *BaseService[T]) Update(entity *T) error {
	return s.repo.Update(s.db, entity)
}

func (s *BaseService[T]) Delete(where map[string]interface{}) error {
	return s.repo.Delete(s.db, where)
}

func (s *BaseService[T]) DeleteByID(id uint) error {
	return s.repo.DeleteByID(s.db, id)
}

func (s *BaseService[T]) UpdateFields(where map[string]interface{}, fields map[string]interface{}) (int64, error) {
	entity := new(T)
	tx := s.db.Model(entity).Where(where).Updates(fields)
	return tx.RowsAffected, tx.Error
}

func (s *BaseService[T]) Transaction(fn func(tx *gorm.DB) error) error {
	return s.db.Transaction(fn)
}
