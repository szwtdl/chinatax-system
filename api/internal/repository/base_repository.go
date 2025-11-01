package repository

import (
	"gorm.io/gorm"
	"strings"
)

type BaseRepository[T any] interface {
	Find(db *gorm.DB, where map[string]interface{}, preloads ...string) ([]T, error)
	FindOne(db *gorm.DB, where map[string]interface{}, preloads ...string) (*T, error)
	FindWithPagination(db *gorm.DB, where map[string]interface{}, page, pageSize int, order string, preloads ...string) ([]T, int64, error)
	Create(db *gorm.DB, entity *T) error
	Update(db *gorm.DB, entity *T) error
	Delete(db *gorm.DB, where map[string]interface{}) error

	UpdateFields(db *gorm.DB, where map[string]interface{}, fields map[string]interface{}) (int64, error)
	DeleteByID(db *gorm.DB, id uint) error
	CreateWithTx(tx *gorm.DB, entity *T) error
	FindOneTx(tx *gorm.DB, where map[string]interface{}, preloads ...string) (*T, error)
}

type baseRepository[T any] struct{}

func NewBaseRepository[T any]() BaseRepository[T] {
	return &baseRepository[T]{}
}

func (r *baseRepository[T]) Find(db *gorm.DB, where map[string]interface{}, preloads ...string) ([]T, error) {
	var result []T
	tx := db
	for _, p := range preloads {
		tx = tx.Preload(p)
	}
	if err := tx.Where(where).Find(&result).Error; err != nil {
		return nil, err
	}
	return result, nil
}

func (r *baseRepository[T]) FindOne(db *gorm.DB, where map[string]interface{}, preloads ...string) (*T, error) {
	var result T
	tx := db
	for _, p := range preloads {
		tx = tx.Preload(p)
	}
	if err := tx.Where(where).First(&result).Error; err != nil {
		return nil, err
	}
	return &result, nil
}

func (r *baseRepository[T]) FindWithPagination(
	db *gorm.DB,
	where map[string]interface{},
	page, pageSize int,
	order string,
	preloads ...string,
) ([]T, int64, error) {
	var result []T
	var total int64

	tx := db.Model(new(T))
	// 预加载关联
	for _, p := range preloads {
		tx = tx.Preload(p)
	}
	// 处理 where 条件（核心优化）
	if len(where) > 0 {
		for k, v := range where {
			// 判断是否为原始 SQL 条件（如 "username LIKE ?"）
			// 特征：k 包含 SQL 关键字（如 LIKE/IN），且 v 是单一值（非切片）
			if strings.Contains(k, "LIKE") || strings.Contains(k, "IN") || strings.Contains(k, "=") {
				// 按原始 SQL 条件处理：tx.Where("username LIKE ?", "%key%")
				tx = tx.Where(k, v)
			} else {
				// 按普通键值对处理：tx.Where(map[string]interface{}{"id": 1})
				tx = tx.Where(map[string]interface{}{k: v})
			}
		}
	}
	// 统计总数
	if err := tx.Count(&total).Error; err != nil {
		return nil, 0, err
	}
	// 默认排序
	if order == "" {
		order = "id DESC"
	}
	// 分页查询
	if err := tx.Order(order).
		Offset((page - 1) * pageSize).
		Limit(pageSize).
		Find(&result).Error; err != nil {
		return nil, 0, err
	}
	return result, total, nil
}

func (r *baseRepository[T]) Create(db *gorm.DB, entity *T) error {
	return db.Create(entity).Error
}

func (r *baseRepository[T]) Update(db *gorm.DB, entity *T) error {
	return db.Save(entity).Error
}

func (r *baseRepository[T]) Delete(db *gorm.DB, where map[string]interface{}) error {
	return db.Where(where).Delete(new(T)).Error
}

func (r *baseRepository[T]) UpdateFields(db *gorm.DB, where map[string]interface{}, fields map[string]interface{}) (int64, error) {
	entity := new(T)
	tx := db.Model(entity).Where(where).Updates(fields)
	return tx.RowsAffected, tx.Error
}

func (r *baseRepository[T]) DeleteByID(db *gorm.DB, id uint) error {
	return db.Unscoped().Delete(new(T), id).Error
}

func (r *baseRepository[T]) CreateWithTx(tx *gorm.DB, entity *T) error {
	return tx.Create(entity).Error
}

func (r *baseRepository[T]) FindOneTx(tx *gorm.DB, where map[string]interface{}, preloads ...string) (*T, error) {
	var result T
	for _, p := range preloads {
		tx = tx.Preload(p)
	}
	if err := tx.Where(where).First(&result).Error; err != nil {
		return nil, err
	}
	return &result, nil
}
