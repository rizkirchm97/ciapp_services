package repository

import (
	"ciapp/category/model/domain"
	"context"
	"database/sql"
)

type CategoryRepository interface {
	Save(ctx context.Context, tx *sql.Tx, category domain.Category) (domain.Category, error)
	Update(ctx context.Context, tx *sql.Tx, category domain.Category) (domain.Category, error)
	Delete(ctx context.Context, tx *sql.Tx, category domain.Category)
	FindAll(ctx context.Context, tx *sql.Tx) ([]domain.Category, error)
	FindById(ctx context.Context, tx *sql.Tx, categoryId int) (domain.Category, error)
}
